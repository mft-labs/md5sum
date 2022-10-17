package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func GetMD5CheckSum(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := md5.New()
	_, err = io.Copy(hasher, file)
	if err != nil {
		return "", err
	}

	sum := hasher.Sum(nil)
	//fmt.Printf("Md5 checksum: %x\n", sum)
	return fmt.Sprintf("%x", sum), nil
}
