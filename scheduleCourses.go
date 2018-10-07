package leetcode

import "sort"

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool { return courses[i][1] < courses[j][1] })
	day := 1
	num := 0
	for _, course := range courses {
		if day+course[0]-1 <= course[1] {
			num++
			day += course[0]
		}
	}
	return num
}
