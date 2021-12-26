package algorithms

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	if version1 == version2 {
		return 0
	}
	v1Revisions := strings.Split(version1, ".")
	v2Revisions := strings.Split(version2, ".")
	return compareRevisions(v1Revisions, v2Revisions)
}

func compareRevisions(v1Revisions, v2Revisions []string) int {
	if len(v1Revisions) == 0 && len(v2Revisions) == 0 {
		return 0
	}
	if len(v1Revisions) == 0 {
		if x, _ := strconv.Atoi(v2Revisions[0]); x == 0 {
			return compareRevisions(v1Revisions, v2Revisions[1:])
		} else {
			return -1
		}
	}
	if len(v2Revisions) == 0 {
		if x, _ := strconv.Atoi(v1Revisions[0]); x == 0 {
			return compareRevisions(v1Revisions[1:], v2Revisions)
		} else {
			return 1
		}
	}
	v1Revision, _ := strconv.Atoi(v1Revisions[0])
	v2Revision, _ := strconv.Atoi(v2Revisions[0])
	if v1Revision == v2Revision {
		return compareRevisions(v1Revisions[1:], v2Revisions[1:])
	} else if v1Revision > v2Revision {
		return 1
	} else {
		return -1
	}
}
