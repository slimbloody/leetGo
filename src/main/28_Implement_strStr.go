package main

type P28 struct {
}

func strStr(haystack string, needle string) int {
	var p P28
	return p.sol1(haystack, needle)
}

func (p P28)sol1(haystack string, needle string) int {
}

func (p P28)sol2Optimize(haystack string, needle string) int {
	for i := 0;; i++ {
		for j := 0;; j++ {
			// If at this point we have navigated through the entire length
			// of needle, we have found a solution, haystack[i].
			if j == len(needle) {
				return i
			}
			// This happens when we run out of elements in haystack,
			// but there are still elements in needle.
			if i + j == len(haystack) {
				return -1
			}
			if haystack[i + j] != needle[j] {
				break
			}
		}
	}
}

func (p P28)sol2(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	for i := 0;; i++ {
		if i + len(needle) > len(haystack) {
			break
		}

		for j := 0;; j++ {
			if haystack[i + j] != needle[j] {
				break
			} else {
				if j == len(needle) - 1 {
					return j
				}
			}
		}
	}

	return -1
}

func main() {
}
