package dict

import (
	"fmt"
	"os"
)

type DictList map[string]uint64

func SaveDict(filename string, dictList DictList) {
	f, err := os.OpenFile(filename,os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)
	if err !=nil {
		panic(err)
	}

	for  k,_:=range dictList {
		f.WriteString(k+"\n")
		f.WriteString(fmt.Sprintf("%d",dictList[k])+"\n")
	}

}
