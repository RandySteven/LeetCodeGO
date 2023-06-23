package main

func countAndSay(n int) string {
	if n == 0 {
		return ""
	}
	if n == 1 {
		return "1"
	}

	prev := countAndSay(n - 1)
	lenPrev := len(prev)

	var count byte
	curr := make([]byte, 0, 64)

	for i := 0; i < lenPrev; i++ {
		if i == lenPrev-1 {
			curr = append(curr, '1'+count)
			curr = append(curr, prev[i])
			break
		}
		if prev[i] == prev[i+1] {
			count++
		} else {
			curr = append(curr, '1'+count)
			curr = append(curr, prev[i])
			count = 0
		}
	}
	return string(curr)
}
