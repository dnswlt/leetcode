package leetcode

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	m := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	return letterCombs(digits, m)
}

func letterCombs(digits string, m map[byte]string) []string {
	if len(digits) == 0 {
		return []string{""}
	}
	cs := m[digits[0]]
	tail := letterCombs(digits[1:], m)
	r := []string{}
	for _, c := range cs {
		for _, t := range tail {
			r = append(r, string(c)+t)
		}
	}
	return r
}
