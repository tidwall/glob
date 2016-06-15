package glob

func Match(pattern, str string) bool {
	if pattern == "*" {
		return true
	}
	return match(pattern, str, false)
}

func Parse(pattern string, desc bool) (min, max, key string, ok bool) {
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
				if desc {
					if c == 0x00 {
						max = min + string(0xFF)
					} else {
						max = min[:len(min)-1] + string(min[len(min)-1]-1)
					}
					min, max = max, min
					// extend the max range by one
					c = max[len(max)-1]
					if c == 0xFF {
						max = max + string(0)
					} else {
						max = max[:len(max)-1] + string(max[len(max)-1]+1)
					}
				} else {
					if c == 0xFF {
						max = min + string(0)
					} else {
						max = min[:len(min)-1] + string(min[len(min)-1]+1)
					}
				}
			}
			return min, max, pattern, true
		}
	}
	if len(skips) == 0 {
		return "", "", pattern, false
	}
	key = escape(pattern, skips)
	return "", "", key, false
}

func escape(pattern string, skips []int) string {
	key := pattern
	for i, j := range skips {
		key = key[:j-i] + key[j-i+1:]
	}
	return key
}
