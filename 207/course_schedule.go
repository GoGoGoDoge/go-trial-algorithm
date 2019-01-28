package course_schedule

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	m := make(map[int][]int, 0)
	items := make([]int, numCourses)
	for i := range items {
		items[i] = i
		m[i] = make([]int, 0)
	}

	for _, pairs := range prerequisites {
		cur, preq := pairs[0], pairs[1]
		m[cur] = append(m[cur], preq)
	}

	perm := make([]bool, numCourses)
	temp := make([]bool, numCourses)

	if _, err := visit(items, m, perm, temp); err != nil {
		return false
	}

	return true
}

func visit(items []int, m map[int][]int, perm []bool, temp []bool) ([]int, error) {
	res := make([]int, 0)
	for _, item := range items {
		if temp[item] {
			return nil, fmt.Errorf("Not a DAG")
		}
		if perm[item] {
			continue
		}
		temp[item] = true
		tmpRes, err := visit(m[item], m, perm, temp)
		if err != nil {
			return nil, err
		}
		res = append(res, tmpRes...)
		res = append(res, item)
		temp[item] = false
		perm[item] = true
	}
	return res, nil
}
