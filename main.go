package main

import (
	"flag"
	"github.com/mft-labs/md5sum/utils"
	"log"
)

func main() {
	var filename string
	var md5sum string
	var err error
	flag.StringVar(&filename, "f", "", "filename to check")
	flag.Parse()

	md5sum, err = utils.GetMD5CheckSum(filename)
	if err != nil {
		log.Printf("Error occurred:%v", err)
	} else {
		log.Printf("MD5 checksum: %s", md5sum)
	}
}
