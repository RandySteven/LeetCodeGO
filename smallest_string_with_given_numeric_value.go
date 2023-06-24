package main

func getSmallestString(n int, k int) string {
	result := make([]byte, n)
	for i := n - 1; i >= 0; i-- {
		if k-i > 26 {
			k -= 26
			result[i] = byte('z')
		} else if k-i <= 26 {
			result[i] = byte('a') + byte(k-i-1)
			k -= k - i
		}
	}
	return string(result)
}
