package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Problem struct {
	Slug       string
	Title      string
	Difficulty string
}

type ProblemDetail struct {
	Problem
	Readme    string
	Solutions []Solution
}

type Solution struct {
	Language string
	Filename string
	Code     string
}

var baseDir string

func dirExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

var skipDirs = map[string]bool{
	"__app__":           true,
	".github":           true,
	".vscode":           true,
	"utils":             true,
	"algorithms-course": true,
	"sorting":           true,
	"node_modules":      true,
	".git":              true,
}

func loadProblems() []Problem {
	entries, err := os.ReadDir(baseDir)
	if err != nil {
		log.Printf("Error reading directory %s: %v", baseDir, err)
		return nil
	}

	var problems []Problem
	for _, entry := range entries {
		if !entry.IsDir() || skipDirs[entry.Name()] || strings.HasPrefix(entry.Name(), ".") || strings.HasPrefix(entry.Name(), "_") {
			continue
		}

		dir := filepath.Join(baseDir, entry.Name())
		readmePath := filepath.Join(dir, "README.md")
		if _, err := os.Stat(readmePath); err != nil {
			continue
		}

		title, difficulty := parseReadmeHeader(readmePath)
		if title == "" {
			title = slugToTitle(entry.Name())
		}
		problems = append(problems, Problem{
			Slug:       entry.Name(),
			Title:      title,
			Difficulty: difficulty,
		})
	}

	sort.Slice(problems, func(i, j int) bool {
		return strings.ToLower(problems[i].Title) < strings.ToLower(problems[j].Title)
	})

	return problems
}

func slugToTitle(slug string) string {
	parts := strings.Split(slug, "-")
	for i, p := range parts {
		if p == "" {
			continue
		}
		parts[i] = strings.ToUpper(p[:1]) + p[1:]
	}
	return strings.Join(parts, " ")
}

func parseReadmeHeader(path string) (title, difficulty string) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", ""
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if title == "" && strings.HasPrefix(line, "#") {
			title = strings.TrimSpace(strings.TrimLeft(line, "#"))
		}
		lower := strings.ToLower(line)
		if difficulty == "" && !strings.HasPrefix(line, "#") {
			if strings.Contains(lower, "**easy**") || lower == "easy" {
				difficulty = "Easy"
			} else if strings.Contains(lower, "**medium**") || lower == "medium" {
				difficulty = "Medium"
			} else if strings.Contains(lower, "**hard**") || lower == "hard" {
				difficulty = "Hard"
			}
		}
		if title != "" && difficulty != "" {
			break
		}
	}
	return title, difficulty
}

func loadProblemDetail(slug string) *ProblemDetail {
	dir := filepath.Join(baseDir, slug)

	if _, err := os.Stat(dir); err != nil {
		return nil
	}

	readmePath := filepath.Join(dir, "README.md")
	title, difficulty := parseReadmeHeader(readmePath)
	if title == "" {
		title = slugToTitle(slug)
	}
	readme, _ := os.ReadFile(readmePath)

	solutionFiles := []struct {
		name string
		lang string
	}{
		{"index.go", "Go"},
		{"index.py", "Python"},
		{"index.js", "JavaScript"},
	}

	var solutions []Solution
	for _, sf := range solutionFiles {
		code, err := os.ReadFile(filepath.Join(dir, sf.name))
		if err != nil {
			continue
		}
		solutions = append(solutions, Solution{
			Language: sf.lang,
			Filename: sf.name,
			Code:     string(code),
		})
	}

	return &ProblemDetail{
		Problem: Problem{
			Slug:       slug,
			Title:      title,
			Difficulty: difficulty,
		},
		Readme:    string(readme),
		Solutions: solutions,
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	problems := loadProblems()
	if err := indexTmpl.Execute(w, problems); err != nil {
		log.Printf("template error: %v", err)
	}
}

func handleProblem(w http.ResponseWriter, r *http.Request) {
	slug := strings.TrimPrefix(r.URL.Path, "/problem/")
	if slug == "" || strings.Contains(slug, "/") || strings.Contains(slug, "..") {
		http.NotFound(w, r)
		return
	}

	detail := loadProblemDetail(slug)
	if detail == nil {
		http.NotFound(w, r)
		return
	}

	if err := problemTmpl.Execute(w, detail); err != nil {
		log.Printf("template error: %v", err)
	}
}

func main() {
	if len(os.Args) > 1 {
		baseDir = os.Args[1]
	} else {
		wd, _ := os.Getwd()
		baseDir = filepath.Dir(wd)
	}

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/problem/", handleProblem)

	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	fmt.Printf("Algorithms viewer running at http://localhost:%s\n", port)
	fmt.Printf("Scanning problems from: %s\n", baseDir)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

var funcMap = template.FuncMap{
	"lower": strings.ToLower,
}

const sharedStyles = `
  * { margin: 0; padding: 0; box-sizing: border-box; }
  body { font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif; background: #0d1117; color: #c9d1d9; line-height: 1.5; }
  .container { max-width: 920px; margin: 0 auto; padding: 32px 24px; }
  a { color: #58a6ff; text-decoration: none; }
  a:hover { text-decoration: underline; }
  .badge { font-size: 11px; padding: 3px 10px; border-radius: 999px; font-weight: 600; letter-spacing: 0.3px; text-transform: uppercase; display: inline-block; }
  .badge-easy { background: rgba(63, 185, 80, 0.15); color: #3fb950; }
  .badge-medium { background: rgba(210, 153, 34, 0.15); color: #d29922; }
  .badge-hard { background: rgba(248, 81, 73, 0.15); color: #f85149; }
`

var indexTmpl = template.Must(template.New("index").Funcs(funcMap).Parse(`<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Algorithms</title>
<style>` + sharedStyles + `
  header { margin-bottom: 24px; }
  h1 { font-size: 32px; color: #f0f6fc; letter-spacing: -0.5px; }
  .count { color: #8b949e; font-size: 14px; margin-top: 4px; }
  .controls { display: flex; flex-wrap: wrap; gap: 12px; align-items: center; margin-bottom: 20px; }
  .search { flex: 1; min-width: 220px; padding: 10px 14px; border: 1px solid #30363d; border-radius: 8px; background: #161b22; color: #c9d1d9; font-size: 15px; outline: none; transition: border-color 0.15s; }
  .search:focus { border-color: #58a6ff; }
  .sort { padding: 10px 14px; border: 1px solid #30363d; border-radius: 8px; background: #161b22; color: #c9d1d9; font-size: 14px; outline: none; cursor: pointer; }
  .filters { display: flex; gap: 8px; flex-wrap: wrap; margin-bottom: 16px; }
  .chip { padding: 6px 14px; border: 1px solid #30363d; border-radius: 999px; background: #161b22; color: #c9d1d9; font-size: 13px; font-weight: 500; cursor: pointer; transition: all 0.15s; }
  .chip:hover { border-color: #484f58; }
  .chip.active { background: #1f6feb; border-color: #1f6feb; color: white; }
  .chip[data-filter="easy"].active { background: #1b3a2a; border-color: #3fb950; color: #3fb950; }
  .chip[data-filter="medium"].active { background: #3b2e1a; border-color: #d29922; color: #d29922; }
  .chip[data-filter="hard"].active { background: #3d1a1a; border-color: #f85149; color: #f85149; }
  .list { list-style: none; border: 1px solid #21262d; border-radius: 8px; overflow: hidden; }
  .list li { border-bottom: 1px solid #21262d; }
  .list li:last-child { border-bottom: none; }
  .list a { display: flex; align-items: center; justify-content: space-between; padding: 14px 16px; color: #c9d1d9; text-decoration: none; transition: background 0.15s; }
  .list a:hover { background: #161b22; text-decoration: none; }
  .list a:hover .title { color: #58a6ff; }
  .title { font-size: 15px; transition: color 0.15s; }
  .visible-count { color: #8b949e; font-size: 13px; padding: 12px 0; }
  .no-results { text-align: center; color: #8b949e; padding: 40px; }
</style>
</head>
<body>
<div class="container">
  <header>
    <h1>Algorithms</h1>
    <p class="count">{{len .}} problems solved</p>
  </header>

  <div class="controls">
    <input class="search" type="text" placeholder="Search problems..." id="search" autofocus>
    <select class="sort" id="sort">
      <option value="alpha-asc">A &rarr; Z</option>
      <option value="alpha-desc">Z &rarr; A</option>
      <option value="difficulty-asc">Easy first</option>
      <option value="difficulty-desc">Hard first</option>
    </select>
  </div>

  <div class="filters">
    <button class="chip active" data-filter="all">All</button>
    <button class="chip" data-filter="easy">Easy</button>
    <button class="chip" data-filter="medium">Medium</button>
    <button class="chip" data-filter="hard">Hard</button>
  </div>

  <div class="visible-count" id="visible-count"></div>

  <ul class="list" id="list">
    {{range .}}
    <li data-title="{{.Title}}" data-difficulty="{{.Difficulty | lower}}">
      <a href="/problem/{{.Slug}}">
        <span class="title">{{.Title}}</span>
        {{if .Difficulty}}<span class="badge badge-{{.Difficulty | lower}}">{{.Difficulty}}</span>{{end}}
      </a>
    </li>
    {{end}}
  </ul>
  <p class="no-results" id="no-results" style="display:none">No matching problems found.</p>
</div>
<script>
(function() {
  var list = document.getElementById('list');
  var items = Array.from(list.querySelectorAll('li'));
  var search = document.getElementById('search');
  var sort = document.getElementById('sort');
  var chips = document.querySelectorAll('.chip');
  var visibleCount = document.getElementById('visible-count');
  var noResults = document.getElementById('no-results');
  var activeFilter = 'all';

  var difficultyOrder = { 'easy': 1, 'medium': 2, 'hard': 3, '': 4 };

  function applyFilters() {
    var q = search.value.toLowerCase().trim();
    var visible = 0;
    items.forEach(function(item) {
      var title = item.getAttribute('data-title').toLowerCase();
      var diff = item.getAttribute('data-difficulty');
      var matchesSearch = !q || title.includes(q);
      var matchesFilter = activeFilter === 'all' || diff === activeFilter;
      var show = matchesSearch && matchesFilter;
      item.style.display = show ? '' : 'none';
      if (show) visible++;
    });
    visibleCount.textContent = visible + (visible === 1 ? ' problem' : ' problems');
    noResults.style.display = visible === 0 ? '' : 'none';
  }

  function applySort() {
    var mode = sort.value;
    var sorted = items.slice().sort(function(a, b) {
      var ta = a.getAttribute('data-title').toLowerCase();
      var tb = b.getAttribute('data-title').toLowerCase();
      var da = difficultyOrder[a.getAttribute('data-difficulty')] || 4;
      var db = difficultyOrder[b.getAttribute('data-difficulty')] || 4;
      switch (mode) {
        case 'alpha-asc':       return ta < tb ? -1 : ta > tb ? 1 : 0;
        case 'alpha-desc':      return ta < tb ? 1 : ta > tb ? -1 : 0;
        case 'difficulty-asc':  return (da - db) || (ta < tb ? -1 : 1);
        case 'difficulty-desc': return (db - da) || (ta < tb ? -1 : 1);
      }
      return 0;
    });
    sorted.forEach(function(item) { list.appendChild(item); });
  }

  search.addEventListener('input', applyFilters);
  sort.addEventListener('change', applySort);
  chips.forEach(function(chip) {
    chip.addEventListener('click', function() {
      chips.forEach(function(c) { c.classList.remove('active'); });
      chip.classList.add('active');
      activeFilter = chip.getAttribute('data-filter');
      applyFilters();
    });
  });

  applyFilters();
})();
</script>
</body>
</html>`))

var problemTmpl = template.Must(template.New("problem").Funcs(funcMap).Parse(`<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>{{.Title}}</title>
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/highlight.js@11.9.0/styles/github-dark.min.css">
<style>` + sharedStyles + `
  .back { font-size: 14px; display: inline-flex; align-items: center; gap: 6px; margin-bottom: 24px; color: #8b949e; }
  .back:hover { color: #58a6ff; text-decoration: none; }
  .header { margin-bottom: 32px; padding-bottom: 20px; border-bottom: 1px solid #21262d; }
  h1 { font-size: 28px; color: #f0f6fc; margin-bottom: 10px; letter-spacing: -0.3px; }
  .card { background: #0d1117; border: 1px solid #21262d; border-radius: 8px; margin-bottom: 16px; overflow: hidden; }
  .card-header { display: flex; align-items: center; justify-content: space-between; padding: 12px 16px; background: #161b22; border-bottom: 1px solid #21262d; cursor: pointer; user-select: none; }
  .card-header:hover { background: #1c2129; }
  .card-title { font-size: 15px; font-weight: 600; color: #f0f6fc; display: flex; align-items: center; gap: 10px; }
  .card-title::before { content: '\25BE'; font-size: 12px; color: #8b949e; transition: transform 0.2s; display: inline-block; }
  .card.collapsed .card-title::before { transform: rotate(-90deg); }
  .lang-badge { font-size: 11px; padding: 3px 8px; border-radius: 4px; background: #21262d; color: #8b949e; font-weight: 500; font-family: 'SF Mono', 'Fira Code', Menlo, monospace; }
  .card-body { padding: 20px 24px; }
  .card.collapsed .card-body { display: none; }

  /* Solution tabs */
  .tabs { display: flex; gap: 4px; padding: 8px 16px 0; background: #161b22; border-bottom: 1px solid #21262d; }
  .tab { padding: 8px 14px; font-size: 13px; font-weight: 500; color: #8b949e; background: transparent; border: none; border-radius: 6px 6px 0 0; cursor: pointer; transition: all 0.15s; }
  .tab:hover { color: #c9d1d9; }
  .tab.active { color: #f0f6fc; background: #0d1117; border: 1px solid #21262d; border-bottom: 1px solid #0d1117; margin-bottom: -1px; }
  .tab-content { display: none; }
  .tab-content.active { display: block; }
  .tab-content pre { margin: 0; }

  /* Markdown styles */
  .markdown { font-size: 15px; color: #c9d1d9; }
  .markdown h1, .markdown h2, .markdown h3, .markdown h4 { color: #f0f6fc; margin-top: 24px; margin-bottom: 12px; line-height: 1.3; }
  .markdown h1 { font-size: 22px; padding-bottom: 8px; border-bottom: 1px solid #21262d; }
  .markdown h2 { font-size: 19px; padding-bottom: 6px; border-bottom: 1px solid #21262d; }
  .markdown h3 { font-size: 17px; }
  .markdown h4 { font-size: 15px; }
  .markdown h1:first-child, .markdown h2:first-child, .markdown h3:first-child { margin-top: 0; }
  .markdown p { margin: 12px 0; }
  .markdown ul, .markdown ol { margin: 12px 0; padding-left: 28px; }
  .markdown li { margin: 4px 0; }
  .markdown blockquote { border-left: 3px solid #30363d; padding: 4px 16px; color: #8b949e; margin: 12px 0; }
  .markdown a { color: #58a6ff; }
  .markdown strong { color: #f0f6fc; font-weight: 600; }
  .markdown table { border-collapse: collapse; margin: 16px 0; }
  .markdown th, .markdown td { border: 1px solid #30363d; padding: 6px 12px; }
  .markdown th { background: #161b22; }
  .markdown img { max-width: 100%; }
  .markdown code:not(pre code) { background: #161b22; border: 1px solid #21262d; border-radius: 4px; padding: 1px 6px; font-size: 0.9em; font-family: 'SF Mono', 'Fira Code', Menlo, monospace; }
  .markdown hr { border: none; border-top: 1px solid #21262d; margin: 16px 0; }

  pre { background: #161b22; border-radius: 6px; padding: 16px; overflow-x: auto; font-size: 13px; line-height: 1.5; }
  pre code { font-family: 'SF Mono', 'Fira Code', 'Fira Mono', Menlo, monospace; background: transparent !important; padding: 0 !important; border: none !important; }
  .markdown pre { border: 1px solid #21262d; }
</style>
</head>
<body>
<div class="container">
  <a href="/" class="back">&larr; All Problems</a>

  <div class="header">
    <h1>{{.Title}}</h1>
    {{if .Difficulty}}<span class="badge badge-{{.Difficulty | lower}}">{{.Difficulty}}</span>{{end}}
  </div>

  <div class="card" id="question-card">
    <div class="card-header" onclick="toggleCard('question-card')">
      <span class="card-title">Question</span>
    </div>
    <div class="card-body">
      <div class="markdown" id="readme"></div>
      <textarea id="readme-source" style="display:none">{{.Readme}}</textarea>
    </div>
  </div>

  {{if .Solutions}}
  <div class="card" id="solution-card">
    <div class="card-header" onclick="toggleCard('solution-card')">
      <span class="card-title">Solution</span>
    </div>
    {{if gt (len .Solutions) 1}}
    <div class="tabs" id="solution-tabs">
      {{range $i, $s := .Solutions}}
      <button class="tab{{if eq $i 0}} active{{end}}" data-tab="sol-{{$i}}">{{$s.Language}}</button>
      {{end}}
    </div>
    {{end}}
    <div class="card-body">
      {{range $i, $s := .Solutions}}
      <div class="tab-content{{if eq $i 0}} active{{end}}" id="sol-{{$i}}">
        <div style="display:flex; justify-content:space-between; align-items:center; margin-bottom:8px; font-size:12px; color:#8b949e;">
          <span class="lang-badge">{{$s.Filename}}</span>
        </div>
        <pre><code class="language-{{$s.Language | lower}}">{{$s.Code}}</code></pre>
      </div>
      {{end}}
    </div>
  </div>
  {{end}}
</div>

<script src="https://cdn.jsdelivr.net/npm/marked@12.0.0/marked.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/highlight.min.js"></script>
<script>
function toggleCard(id) {
  document.getElementById(id).classList.toggle('collapsed');
}

(function() {
  // Render markdown
  var src = document.getElementById('readme-source').value;
  var target = document.getElementById('readme');
  if (window.marked) {
    marked.setOptions({ gfm: true, breaks: false });
    target.innerHTML = marked.parse(src);
  } else {
    target.textContent = src;
  }

  // Syntax highlighting
  if (window.hljs) {
    document.querySelectorAll('pre code').forEach(function(block) {
      hljs.highlightElement(block);
    });
  }

  // Solution tabs
  var tabs = document.querySelectorAll('#solution-tabs .tab');
  tabs.forEach(function(tab) {
    tab.addEventListener('click', function() {
      tabs.forEach(function(t) { t.classList.remove('active'); });
      tab.classList.add('active');
      document.querySelectorAll('.tab-content').forEach(function(c) { c.classList.remove('active'); });
      document.getElementById(tab.getAttribute('data-tab')).classList.add('active');
    });
  });
})();
</script>
</body>
</html>`))
