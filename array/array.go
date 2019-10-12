package array

import (
	"strings"
)
func ArraySplit(str string) []string {
	arrSplit:=[]string{}
	strArr:=strings.Split(str,",")
	aStr:=""
	for k,v:=range strArr{
		key:=k+1
		aStr += v+","
		if key%100==0 {
			aStr=strings.TrimRight(aStr,",")
			aStr="`"+aStr+"`"
			arrSplit=append(arrSplit, aStr)
			aStr=""
		}
	}
	if aStr!="" {
		aStr=strings.TrimRight(aStr,",")
		aStr="`"+aStr+"`"
		arrSplit=append(arrSplit, aStr)
	}
	return arrSplit
}