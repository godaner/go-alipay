package randomutil

import (
	"math/rand"
	"sync"
	"github.com/bwmarrin/snowflake"
)

var (
	node *snowflake.Node
	once sync.Once
)
func RandomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(RandInt(65, 90))
	}
	return string(bytes)
}
func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

