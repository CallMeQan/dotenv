package dotenv

func getFilenamesOrDefault(filenames []string) []string {
	if len(filenames) == 0 {
		return []string{"./.env"}
	}
	return filenames
}
func hasQuotePrefix(str string) bool {
	if len(str) == 0 {
		return false
	}
	switch str[0] {
	case '"', '\'':
		return true
	default:
		return false
	}
}
