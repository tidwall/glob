package glob

func Match(pattern, str string) bool {
	return match(pattern, str, false)
}

func Parse(pattern string) (min, max, key string, ok bool) {
	var skips []int
	for i := 0; i < len(pattern); i++ {
		switch pattern[i] {
		case '\\':
			skips = append(skips, i)
			i += 1
		case '[', '*', '?':
			min = escape(pattern[:i], skips)
			if len(min) > 0 {
				c := min[len(min)-1]
				if c == 0xFF {
					max = min + string(0)
				} else {
					max = min[:len(min)-1] + string(min[len(min)-1]+1)
				}
			}
			return min, max, pattern, true
		}
	}
	if len(skips) == 0 {
		return "", "", pattern, false
	}
	return "", "", escape(pattern, skips), false
}
func escape(pattern string, skips []int) string {
	key := pattern
	for i, j := range skips {
		key = key[:j-i] + key[j-i+1:]
	}
	return key
}
