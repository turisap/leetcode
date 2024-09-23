package first_substr

/*
mississippi

	issip
*/
func firstSbuStrNaive(haystack, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}

	if needle == haystack {
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		substr := haystack[i : i+len(needle)]

		if substr == needle {
			return i
		}
	}

	return -1
}
