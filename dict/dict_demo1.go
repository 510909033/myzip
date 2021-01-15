package dict

import (
	"io"
	"os"
)

func CreateDictFileAndContent(filename string) DictList{
	var dictList = make(DictList)
	var index uint64 = 0
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	for {
		b := make([]byte, 64)
		n,err:= f.Read(b)
		_=n
		switch err {
		case nil:
			str:=string(b)
			if _,ok:=dictList[str];!ok{
				dictList[str] = index
				index++
			}
		case io.EOF:
			goto forBreak
		default:
			panic(err)
		}

	}
forBreak:

	return dictList
}
