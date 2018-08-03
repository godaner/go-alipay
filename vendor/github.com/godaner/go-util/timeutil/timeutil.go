package timeutil

import (
	"strconv"
	"time"
)

func UnixStr()(string){
	t := time.Now()
	return strconv.FormatInt(t.UTC().UnixNano(), 10)[:10]
}
func Unix() (int64){
	t := time.Now()
	return t.Unix()
}
