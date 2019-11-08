def minTime(machines, goal):
    max_days = max(machines) * goal / len(machines)
    start = 0
    end = max_days

    while end > start+1:
        curr_day = int((end  + start)/2)
        total_production = 0
        today_production = 0
        for machine in machines:
            if curr_day % machine == 0:
                today_production += 1
            total_production += int(curr_day / machine)
        prev_day_production = total_production - today_production
        
        if prev_day_production < goal and total_production >= goal:
            return curr_day
        elif prev_day_production >= goal:
            end = curr_day - 1
        elif total_production < goal:
            start = curr_day
    return end