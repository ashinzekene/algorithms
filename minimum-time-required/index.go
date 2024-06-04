package algorithms

import . "github.com/ashinzekene/algorithms/utils"

func minTime(machines []int64, goal int64) int64 {
	max_day := Max(machines...)
	var start int64 = 0
	end := max_day * goal / int64(len(machines))

	for end > start+1 {
		curr_day := (start + end) / 2
		var today_production int64 = 0
		var total_production int64 = 0
		for _, x := range machines {
			if curr_day%x == 0 {
				today_production++
			}
			total_production += curr_day / x
		}
		prev_day_production := total_production - today_production

		if prev_day_production < goal && total_production >= goal {
			return curr_day
		} else if prev_day_production >= goal {
			end = curr_day - 1
		} else if total_production < goal {
			start = curr_day
		}
	}
	return end
}

func minTime1(machines []int64, goal int64) int64 {
	var start int64 = 0
	maxDay := Max(machines...)
	var end int64 = maxDay * goal / int64(len(machines))
	for start <= end {
		m := (start + end) / 2
		prod := totalProductionAt(m, machines)
		if prod < goal {
			start = m + 1
		} else {
			end = m - 1
		}
	}
	return start
}

func totalProductionAt(days int64, machines []int64) int64 {
	var total int64 = 0
	for _, machine := range machines {
		total += (days / machine)
	}
	return total
}
