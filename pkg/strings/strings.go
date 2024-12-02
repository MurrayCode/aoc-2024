package strings

func RemoveNonNumeric(s string) string {
	var res string
	for _, c := range s {
		if c >= '0' && c <= '9' {
			res += string(c)
		}
	}
	return res
}
