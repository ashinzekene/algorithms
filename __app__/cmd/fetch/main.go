package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var (
	leetcodeURLRe   = regexp.MustCompile(`https?://(?:www\.)?leetcode\.com/problems/([a-z0-9-]+)`)
	hackerrankURLRe = regexp.MustCompile(`https?://(?:www\.)?hackerrank\.com/challenges/([a-z0-9-]+)`)
	difficultyRe    = regexp.MustCompile(`(?i)\b(easy|medium|hard)\b`)
	titleRe         = regexp.MustCompile(`(?m)^#{1,2}\s+\S`)
	tagRe           = regexp.MustCompile(`<[^>]+>`)
	preRe           = regexp.MustCompile(`(?is)<pre[^>]*>\s*(?:<code[^>]*>)?(.*?)(?:</code>)?\s*</pre>`)
	multiNL         = regexp.MustCompile(`\n{3,}`)
	skipDirs        = map[string]bool{
		"__app__": true, ".github": true, ".vscode": true,
		"utils": true, "algorithms-course": true, "sorting": true,
		"node_modules": true, ".git": true,
	}
)

type problem struct {
	Title      string
	Difficulty string
	Content    string
	URL        string
	Markdown   bool // true if Content is already markdown (skip HTML conversion)
}

func main() {
	dryRun := flag.Bool("dry-run", false, "Print what would change without writing")
	rootArg := flag.String("root", "", "Repo root (defaults to autodetect)")
	only := flag.String("only", "", "Only fetch this single slug (for previewing)")
	verbose := flag.Bool("verbose", false, "Print full formatted README to stdout")
	flag.Parse()

	root := *rootArg
	if root == "" {
		root = findRepoRoot()
	}
	if root == "" {
		fmt.Fprintln(os.Stderr, "Could not find repo root")
		os.Exit(1)
	}
	fmt.Printf("Scanning: %s\n\n", root)

	entries, err := os.ReadDir(root)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading repo: %v\n", err)
		os.Exit(1)
	}

	var fixed, skipped, failed int
	for _, e := range entries {
		if !e.IsDir() || skipDirs[e.Name()] || strings.HasPrefix(e.Name(), ".") || strings.HasPrefix(e.Name(), "_") {
			continue
		}
		if *only != "" && e.Name() != *only {
			continue
		}
		readmePath := filepath.Join(root, e.Name(), "README.md")
		data, err := os.ReadFile(readmePath)
		if err != nil {
			continue
		}

		needsFix := !titleRe.Match(data) || !difficultyRe.Match(data)
		if !needsFix {
			continue
		}

		var (
			p      *problem
			source string
		)
		if m := leetcodeURLRe.FindStringSubmatch(string(data)); m != nil {
			source = "leetcode"
			fmt.Printf("[fetching] %s (leetcode: %s)\n", e.Name(), m[1])
			p, err = fetchLeetcode(m[1], m[0])
		} else if m := hackerrankURLRe.FindStringSubmatch(string(data)); m != nil {
			source = "hackerrank"
			fmt.Printf("[fetching] %s (hackerrank: %s)\n", e.Name(), m[1])
			p, err = fetchHackerrank(m[1], m[0])
		} else {
			fmt.Printf("[no-url]   %s\n", e.Name())
			skipped++
			continue
		}

		if err != nil {
			fmt.Printf("[failed]   %s (%s): %v\n", e.Name(), source, err)
			failed++
			continue
		}

		newReadme := formatReadme(p)
		if *verbose {
			fmt.Printf("\n--- %s ---\n%s\n--- end %s ---\n", e.Name(), newReadme, e.Name())
		}
		if *dryRun {
			fmt.Printf("[would write] %s — %s (%s)\n", e.Name(), p.Title, p.Difficulty)
		} else {
			if err := os.WriteFile(readmePath, []byte(newReadme), 0644); err != nil {
				fmt.Printf("[failed]   %s: write: %v\n", e.Name(), err)
				failed++
				continue
			}
			fmt.Printf("[wrote]    %s — %s (%s)\n", e.Name(), p.Title, p.Difficulty)
		}
		fixed++
		time.Sleep(600 * time.Millisecond)
	}

	mode := "wrote"
	if *dryRun {
		mode = "would write"
	}
	fmt.Printf("\nDone: %d %s, %d skipped (no URL), %d failed\n", fixed, mode, skipped, failed)
}

func findRepoRoot() string {
	wd, _ := os.Getwd()
	for _, dir := range []string{wd, filepath.Dir(wd), filepath.Dir(filepath.Dir(wd)), filepath.Dir(filepath.Dir(filepath.Dir(wd)))} {
		if _, err := os.Stat(filepath.Join(dir, "climbing-stairs", "README.md")); err == nil {
			abs, _ := filepath.Abs(dir)
			return abs
		}
	}
	return ""
}

const userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"

func fetchLeetcode(slug, url string) (*problem, error) {
	body, _ := json.Marshal(map[string]interface{}{
		"query":         `query questionData($titleSlug: String!) { question(titleSlug: $titleSlug) { title difficulty content } }`,
		"variables":     map[string]string{"titleSlug": slug},
		"operationName": "questionData",
	})

	req, _ := http.NewRequest("POST", "https://leetcode.com/graphql/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Referer", "https://leetcode.com/problems/"+slug+"/")
	req.Header.Set("User-Agent", userAgent)

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		rb, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("status %d: %.200s", resp.StatusCode, string(rb))
	}

	var out struct {
		Data struct {
			Question *struct {
				Title      string `json:"title"`
				Difficulty string `json:"difficulty"`
				Content    string `json:"content"`
			} `json:"question"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	if out.Data.Question == nil || out.Data.Question.Title == "" {
		return nil, fmt.Errorf("question not found")
	}
	return &problem{
		Title:      out.Data.Question.Title,
		Difficulty: out.Data.Question.Difficulty,
		Content:    out.Data.Question.Content,
		URL:        url,
	}, nil
}

func fetchHackerrank(slug, url string) (*problem, error) {
	apiURL := "https://www.hackerrank.com/rest/contests/master/challenges/" + slug
	req, _ := http.NewRequest("GET", apiURL, nil)
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		rb, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("status %d: %.200s", resp.StatusCode, string(rb))
	}

	var out struct {
		Model *struct {
			Name             string `json:"name"`
			DifficultyName   string `json:"difficulty_name"`
			ProblemStatement string `json:"problem_statement"`
			InputFormat      string `json:"input_format"`
			OutputFormat     string `json:"output_format"`
			Constraints      string `json:"constraints"`
			SampleInput      string `json:"sample_input"`
			SampleOutput     string `json:"sample_output"`
		} `json:"model"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	if out.Model == nil || out.Model.Name == "" {
		return nil, fmt.Errorf("challenge not found")
	}
	m := out.Model

	var b strings.Builder
	b.WriteString(strings.TrimSpace(m.ProblemStatement))
	if s := strings.TrimSpace(m.InputFormat); s != "" {
		b.WriteString("\n\n**Input Format**\n\n")
		b.WriteString(s)
	}
	if s := strings.TrimSpace(m.OutputFormat); s != "" {
		b.WriteString("\n\n**Output Format**\n\n")
		b.WriteString(s)
	}
	if s := strings.TrimSpace(m.Constraints); s != "" {
		b.WriteString("\n\n**Constraints**\n\n")
		b.WriteString(s)
	}
	if s := strings.TrimSpace(m.SampleInput); s != "" {
		b.WriteString("\n\n**Sample Input**\n\n```\n")
		b.WriteString(s)
		b.WriteString("\n```")
	}
	if s := strings.TrimSpace(m.SampleOutput); s != "" {
		b.WriteString("\n\n**Sample Output**\n\n```\n")
		b.WriteString(s)
		b.WriteString("\n```")
	}

	return &problem{
		Title:      m.Name,
		Difficulty: m.DifficultyName,
		Content:    b.String(),
		URL:        url,
		Markdown:   true,
	}, nil
}

func formatReadme(p *problem) string {
	body := p.Content
	if !p.Markdown {
		body = htmlToMarkdown(body)
	}
	return fmt.Sprintf("# %s\n**%s**\n\n%s\n\n%s\n", p.Title, p.Difficulty, p.URL, body)
}

var (
	styleRe = regexp.MustCompile(`(?is)<style[^>]*>.*?</style>`)
	// MathJax stores the original LaTeX in a <script type="math/tex"> element — extract it before stripping scripts
	mathTexRe   = regexp.MustCompile(`(?is)<script[^>]*type=["']math/tex[^"']*["'][^>]*>(.*?)</script>`)
	scriptRe    = regexp.MustCompile(`(?is)<script[^>]*>.*?</script>`)
	mathJaxSpan = regexp.MustCompile(`(?is)<span[^>]*class=["']?MathJax[^"'>]*["']?[^>]*>.*?</span>`)
)

func htmlToMarkdown(html string) string {
	h := html
	// Step 0: extract MathJax LaTeX before stripping script/style/spans
	h = mathTexRe.ReplaceAllString(h, "`$1`")
	// Drop the rendered MathJax SVG/HTML wrappers (already extracted the source above)
	h = mathJaxSpan.ReplaceAllString(h, "")
	h = styleRe.ReplaceAllString(h, "")
	h = scriptRe.ReplaceAllString(h, "")
	// Step 1: convert HTML structure to markdown while content still has entities
	for from, to := range map[string]string{
		"<br>": "\n", "<br/>": "\n", "<br />": "\n", "&nbsp;": " ",
	} {
		h = strings.ReplaceAll(h, from, to)
	}
	h = preRe.ReplaceAllString(h, "\n```\n$1\n```\n")
	for _, sub := range []struct{ re, repl string }{
		{`(?is)<strong[^>]*>(.*?)</strong>`, "**$1**"},
		{`(?is)<b[^>]*>(.*?)</b>`, "**$1**"},
		{`(?is)<em[^>]*>(.*?)</em>`, "*$1*"},
		{`(?is)<i[^>]*>(.*?)</i>`, "*$1*"},
		{`(?is)<code[^>]*>(.*?)</code>`, "`$1`"},
		{`(?is)<li[^>]*>(.*?)</li>`, "- $1\n"},
		{`(?is)<sup[^>]*>(.*?)</sup>`, "^$1"},
		{`(?is)<sub[^>]*>(.*?)</sub>`, "_$1"},
		{`(?is)<p[^>]*>`, ""},
		{`</p>`, "\n\n"},
	} {
		h = regexp.MustCompile(sub.re).ReplaceAllString(h, sub.repl)
	}
	// Step 2: strip any remaining tags (entities still encoded, so < > inside text are safe)
	h = tagRe.ReplaceAllString(h, "")
	// Step 3: decode entities last, so decoded < > don't get treated as tags
	for from, to := range map[string]string{
		"&lt;": "<", "&gt;": ">", "&amp;": "&",
		"&quot;": "\"", "&#39;": "'", "&apos;": "'",
	} {
		h = strings.ReplaceAll(h, from, to)
	}
	h = multiNL.ReplaceAllString(h, "\n\n")
	return strings.TrimSpace(h)
}
