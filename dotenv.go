package dotenv

import (
	"bufio"
	"errors"
	"os"
)

//
func Load(filepaths ...string) error {
	for _, path := range filepaths {
		err := loadFromFile(path)
		if err != nil {
			return err
		}
	}
	return nil
}

func loadFromFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var parts [2]string
		currPart := 0

		for _, c := range scanner.Text() {
			if c == '=' && currPart == 0 {
				currPart = 1
				continue
			}
			parts[currPart] += string(c)
		}

		if len(parts) != 2 {
			return errors.New("wrong format for " + scanner.Text())
		}

		key := parts[0]
		value := parts[1]

		err = os.Setenv(key, value)
		if err != nil {
			return err
		}
	}

	return scanner.Err()
}
