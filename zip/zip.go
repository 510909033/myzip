package zip

import (
	"baotian0506.com/myzip/dict"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func Zip(rawFilename string, dictList dict.DictList) string  {
	rawFilenamePtr, err := os.Open(rawFilename)

	if err !=nil {
		panic(err)
	}
	defer rawFilenamePtr.Close()

	zipFilename := rawFilename + "_zip"
	zipFilenamePtr, err1 := os.OpenFile(zipFilename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err1 != nil {
		panic(err)
	}
	defer  zipFilenamePtr.Close()

	for {
		b := make([]byte, 64)
		n,err:= rawFilenamePtr.Read(b)
		_=n
		switch err {
			case nil:
				str:=string(b)
				if _,ok:= dictList[str];ok{
					//命中字典
					buf:=make([]byte,108)
					writeLen:=binary.PutUvarint(buf, dictList[str])
					_=writeLen
					//fmt.Println("writeLen:", writeLen)
					newWriteLen,err:= zipFilenamePtr.Write(buf[:8])//?
					if err!=nil{
						panic(err)
					}
					_=newWriteLen
					//fmt.Println("newWriteLen:", newWriteLen)
				} else {
					//未命中字典
					fmt.Println("未命中字典")
					panic("未命中字典")
				}
			case io.EOF:
			goto forBreak
		default:
			panic(err)
		}
		//fmt.Println(n)
	}
forBreak:

	dict.SaveDict("e:/tmp1/dict/dict.txt",dictList)
	fmt.Println("新压缩文件：", zipFilename)
	fmt.Println("字典文件：", "")
	fmt.Println("over")
	return zipFilename
}

func Unzip(zipFilename string, dictList dict.DictList, destFilename string) {
	var reserveDickList = make(map[uint64]string)
	for k,_:=range dictList{
		reserveDickList[dictList[k]] = k
	}

	zipFilenamePtr, err := os.Open(zipFilename)
	if err != nil {
		panic(err)
	}
	defer  zipFilenamePtr.Close()

	destFilenamePtr, err := os.OpenFile(destFilename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err !=nil {
		panic(err)
	}
	defer destFilenamePtr.Close()


	for {
		b := make([]byte, 8)
		n,err:= zipFilenamePtr.Read(b)
		switch err {
		case nil:
			index,uvarintLen:= binary.Uvarint(b)
			fmt.Println("index=", index," ,uvarintLen=", uvarintLen)
			destFilenamePtr.WriteString(reserveDickList[index])
		case io.EOF:
			goto forBreak
		default:
			panic(err)
		}
		fmt.Println(n,b[:n])
	}
forBreak:

}