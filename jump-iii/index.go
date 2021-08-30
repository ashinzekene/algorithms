package algorithms

func canReach(arr []int, start int) bool {
	visited := make([]bool, len(arr))
	return checkIfZero(arr, visited, start)
}

func checkIfZero(arr []int, visited []bool, current int) bool {
	if visited[current] {
		return false
	}
	if arr[current] == 0 {
		return true
	}

	val := arr[current]
	back := current - val
	front := current + val
	visited[current] = true

	if back >= 0 && checkIfZero(arr, visited, back) {
		return true
	}
	if front < len(arr) && checkIfZero(arr, visited, front) {
		return true
	}
	return false
}
