package main

import (
	"baotian0506.com/myzip/dict"
	"baotian0506.com/myzip/zip"
	"fmt"
	"time"
)

func main() {

	fmt.Println("1.0")
	filename:= "e:/tmp1/dict/20190924-1.csv"
	//filename= "e:/tmp1/dict/x.tar.gz"

	createZipByFilename := "e:/tmp1/dict/x.tar.gz"

	dictList:=dict.CreateDictFileAndContent(createZipByFilename)
	zipFilename:=zip.Zip(filename,dictList)

	//zipFilename := "e:/tmp1/dict/20190924-1.csv_zip"
	unzipFilename:= "e:/tmp1/dict/" + time.Now().Format("2006-01-02_15_04_05") + ".csv"
	zip.Unzip(zipFilename, dictList, unzipFilename)
}
