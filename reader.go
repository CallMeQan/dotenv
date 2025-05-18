package dotenv

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func loadFile(filename string, forceLoad bool) error {
	mapper, err := readFile(filename)
	if err != nil {
		return err
	}

	for k, v := range mapper {
		if !forceLoad {
			if _, exist := os.LookupEnv(k); exist {
				continue
			}
		}
		os.Setenv(k, v)
	}
	return nil
}

func readFile(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Start process of reading file
	scanner := bufio.NewScanner(file)
	res := make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()

		// Skipped if empty line or comment line
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		idx := strings.Index(line, "=")
		if idx == -1 {
			return nil, errors.New(filename + " has unknown line: " + line)
		}

		// Prevented from whitespace bug
		key := strings.TrimSpace(line[:idx])
		val := strings.TrimSpace(line[idx+1:])
		res[key] = val
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return res, nil
}
