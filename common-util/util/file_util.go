package util

import (
	"io/ioutil"
	"os"
	log "github.com/sirupsen/logrus"
	"fmt"
	"github.com/pkg/errors"
)

func ReadFile(path string) ([]byte, error) {
	log.Debug("ReadFile()")

	fileBytes, err := ioutil.ReadFile(path)

	if err != nil {
		log.Error(err)
		return nil, err
	}
	if len(fileBytes) == 0 {
		err = errors.New("File bytes are length zero")
		return nil, err
	}
	return fileBytes, nil
}

func WriteFile(dir, fileName, content string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	fmt.Fprintf(file, content)
	defer file.Close()
	return nil
}

