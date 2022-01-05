package algorithms

// Topological sort
func canFinish(numCourses int, prerequisites [][]int) bool {
	prereqDegrees := map[int]map[int]bool{}
	prereqCount := make([]int, numCourses)

	// count all out degrees
	for _, prereq := range prerequisites {
		course := prereq[0]
		coursePrereq := prereq[1]
		if prereqDegrees[coursePrereq] == nil {
			prereqDegrees[coursePrereq] = make(map[int]bool)
		}
		prereqDegrees[coursePrereq][course] = true
		prereqCount[course]++
	}
	queue := make([]int, 0)
	for course, count := range prereqCount {
		if count == 0 {
			queue = append(queue, course)
		}
	}

	noOfCourses := 0
	for len(queue) != 0 {
		currentCourse := queue[0]
		queue = queue[1:]
		noOfCourses++
		currentCourseDeps := prereqDegrees[currentCourse]
		for course := range currentCourseDeps {
			prereqCount[course]--
			if prereqCount[course] == 0 {
				queue = append(queue, course)
			}
		}
	}
	return numCourses == noOfCourses
}
func canFinish1(numCourses int, prerequisites [][]int) bool {
	outs := make(map[int][]int)
	cycleExists := make(map[int]bool)
	for _, preq := range prerequisites {
		in := preq[0]
		out := preq[1]
		outs[out] = append(outs[out], in)
	}
	for course := range outs {
		courseMap := map[int]bool{}
		recursive := dfs(course, &courseMap, outs, cycleExists)
		if recursive {
			return false
		}
	}
	return true
}

func dfs(course int, courseMap *map[int]bool, outs map[int][]int, cycleExists map[int]bool) bool {
	if len(outs[course]) == 0 {
		return false
	}
	if (*courseMap)[course] {
		return true
	}
	if isCycle, e := cycleExists[course]; e {
		return isCycle
	}
	(*courseMap)[course] = true
	for _, out := range outs[course] {
		if dfs(out, courseMap, outs, cycleExists) {
			return true
		}
	}
	(*courseMap)[course] = false
	cycleExists[course] = false
	return false
}
