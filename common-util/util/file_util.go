package util

import (
	"io/ioutil"
	"os"
	log "github.com/sirupsen/logrus"
	"fmt"
	"github.com/pkg/errors"
	"github.com/dispatchlabs/disgo/commons/utils"
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

func DeleteFile(file string) error {
	err := os.Remove(file)
	if err != nil {
		utils.Error(err)
		return err
	}
	return nil
}

func DeleteSubDirs(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			err := os.Remove(file.Name())
			if err != nil {
				utils.Error(err)
				return err
			}
		}
	}
	return nil
}

func DeleteDir(dir string) error {
	err := os.RemoveAll(dir)
	if err != nil {
		utils.Error(err)
		return err
	}
	return nil
}

func GetCurrentWorkingDir() string {
	dir, err := os.Getwd()
	if err != nil {
		utils.Error(err)
	}
	return dir
}