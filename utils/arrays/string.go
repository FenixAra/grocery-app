package arrays

func Contains(sa []string, s string) bool {
	for _, v := range sa {
		if v == s {
			return true
		}
	}
	return false
}

func Diff(sa, s []string) []string {
	m := make(map[string]bool)
	for _, v := range s {
		m[v] = true
	}

	var val []string
	for _, v := range sa {
		if _, ok := m[v]; !ok {
			val = append(val, v)
		}
	}
	return val
}

func ContainsAny(sa, s []string) bool {
	m := make(map[string]bool)
	for _, v := range s {
		m[v] = true
	}

	for _, v := range sa {
		if _, ok := m[v]; ok {
			return true
		}
	}
	return false
}

func RemoveDuplicates(sa []string) []string {
	m := make(map[string]bool)
	var res []string
	for _, v := range sa {
		if _, ok := m[v]; !ok {
			res = append(res, v)
			m[v] = true
		}
	}

	return res
}

func IsEqual(src []string, s []string) bool {
	m := make(map[string]bool)
	for _, v := range src {
		m[v] = true
	}

	eq := true
	for _, v := range s {
		if _, ok := m[v]; !ok {
			eq = false
		}
	}
	return eq
}

func AppendWithoutDuplicates(src []string, s []string) []string {
	m := make(map[string]bool)
	for _, v := range src {
		m[v] = true
	}

	for _, v := range s {
		if _, ok := m[v]; !ok {
			src = append(src, v)
		}
	}
	return src
}

func RemoveFromArray(src []string, s []string) []string {
	m := make(map[string]bool)
	for _, v := range s {
		m[v] = true
	}

	var newArray []string

	for _, v := range src {
		if _, ok := m[v]; !ok {
			newArray = append(newArray, v)
		}
	}
	return newArray
}

func RemoveFirstElement(src []string) []string {
	return RemoveNthElement(src, 1)
}

func RemoveNthElement(src []string, n int) []string {
	var a []string
	for i, _ := range src {
		if i != (n - 1) {
			a = append(a, src[i])
		}
	}
	return a
}
