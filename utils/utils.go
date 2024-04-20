package utils

func SplitStringByLength(input string, length int) []string {
	var parts []string
	for i := 0; i < len(input); i += length {
		end := i + length
		if end > len(input) {
			end = len(input)
		}
		parts = append(parts, input[i:end])
	}

	return parts
}
