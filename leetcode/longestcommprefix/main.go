package main

func main() {
	println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return string(strs[0])
	}

	minPrefixLen := len(strs[0])

	for i := 1; i < len(strs); i++ {
		if minPrefixLen > len(strs[i]) {
			minPrefixLen = len(strs[i])
		}
	}

	commonPrefix := []byte{}
	for j := 0; j < minPrefixLen; j++ {
		commonPrefix = append(commonPrefix, strs[0][j])
	}

	for i := 0; i < len(commonPrefix); i++ {
		prefixChar := commonPrefix[i]
		for _, str := range strs {
			if str[i] != prefixChar {
				return string(commonPrefix[:i])
			}
		}
	}

	return string(commonPrefix)
}
