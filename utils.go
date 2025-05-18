package dotenv

func getFilenamesOrDefault(filenames []string) []string {
	if len(filenames) == 0 {
		return []string{"./.env"}
	}
	return filenames
}
