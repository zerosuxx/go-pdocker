package pkg

import "os"

func IsFileExists(file string) bool {
	_, statError := os.Stat(file)

	return statError == nil
}
