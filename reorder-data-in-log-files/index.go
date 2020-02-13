package algorithms

import (
	"sort"
	"strconv"
	"strings"
)

func reorderLogFiles(logs []string) []string {
	letterLogs := make([]string, 0)
	numLogs := make([]int, 0)
	for i, log := range logs {
		if isLetterLog(log) {
			letterLogs = append(letterLogs, log)
		} else {
			numLogs = append(numLogs, i)
		}
	}
	sort.SliceStable(letterLogs, func(i, j int) bool {
		a := strings.Index(letterLogs[i], " ")
		b := strings.Index(letterLogs[j], " ")
		if letterLogs[i][a+1:] == letterLogs[j][b+1:] {
			return letterLogs[i][:a] < letterLogs[j][:b]
		}
		return letterLogs[i][a+1:] < letterLogs[j][b+1:]
	})
	for _, i := range numLogs {
		letterLogs = append(letterLogs, logs[i])
	}
	return letterLogs
}

func isLetterLog(log string) bool {
	for i, ch := range log {
		if string(ch) == " " {
			_, err := strconv.Atoi(string(log[i+1]))
			return err != nil
		}
	}
	return false
}
