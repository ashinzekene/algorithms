package algorithms

func canVisitAllRooms(rooms [][]int) bool {
	canReach := make([]bool, len(rooms))

    queue := []int{0}
    canReach[0] = true
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        for _, child := range rooms[node] {
            if !canReach[child] {
                canReach[child] = true
                queue = append(queue, child)
            }
        }
    }

    for _, v := range canReach {
        if !v {
            return false
        }
    }

    return true
}