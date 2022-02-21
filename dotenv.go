package dotenv

import (
	"bufio"
	"errors"
	"os"
	"strings"
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
		parts := strings.Split(scanner.Text(), "=")
		if len(parts) == 1 {
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
