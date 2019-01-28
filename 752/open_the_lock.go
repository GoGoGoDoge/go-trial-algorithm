package openthelock

func openLock(deadends []string, target string) int {
	dMap := make(map[string]bool, len(deadends))
	usedMap := make(map[string]bool)

	for _, x := range deadends {
		dMap[x] = true
	}

	if dMap[target] {
		return -1
	}

	queue := []string{"0000"}
	step := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[i]
			if cur == target {
				return step
			}

			ns := neighbors(cur, usedMap)

			for _, s := range ns {
				if dMap[cur] {
					continue
				}
				usedMap[s] = true
				queue = append(queue, s)
			}
		}
		step++
		queue = queue[size:]
	}

	return -1
}

func neighbors(cur string, usedMap map[string]bool) []string {
	rs := []byte(cur)
	res := []string{}
	for i, r := range rs {
		rs[i] = (r-'0'+1)%10 + '0'
		if !usedMap[string(rs)] {
			res = append(res, string(rs))
		}
		rs[i] = (r-'0'+9)%10 + '0'
		if !usedMap[string(rs)] {
			res = append(res, string(rs))
		}
		rs[i] = r
	}
	return res
}
