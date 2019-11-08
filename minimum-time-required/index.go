package algorithms

func minTime(machines []int64, goal int64) int64 {
	var max_day int64 = 0
	for _, x := range machines {
		if x > max_day {
			max_day = x
		}
	}
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
