package weekly

func repeatedCharacter(s string) byte {
	var m = make(map[byte]bool)
	for i, _ := range s {
		if m[s[i]] {
			return s[i]
		}
		m[s[i]] = true
	}
	return s[0]
}

