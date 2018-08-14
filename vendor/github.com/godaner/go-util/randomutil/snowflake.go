package randomutil

import (
	"github.com/bwmarrin/snowflake"
	"fmt"
)

//diff randomId will produce never same snow flake id
func GetSnowFlakeId(randomId int64) int64 {
	once.Do(func() {
		getSnowFlake(randomId)
	})
	id := node.Generate()
	return id.Int64()
}

func GetSnowFlakeIdStr(uniqueId int64) string {

	return fmt.Sprintf("%d", GetSnowFlakeId(uniqueId))
}

func getSnowFlake(uniqueId int64) {
	node1, err := snowflake.NewNode(int64(uniqueId))
	if err != nil {
		fmt.Println(err)
		return
	} else {
		node = node1
	}
}