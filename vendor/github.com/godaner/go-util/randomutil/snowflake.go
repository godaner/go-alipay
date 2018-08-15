package randomutil

import (
	"github.com/bwmarrin/snowflake"
	"fmt"
)

//diff workMachineId will produce never same snow flake id
func GetSnowFlakeId(workMachineId int64) int64 {
	once.Do(func() {
		getSnowFlake(workMachineId)
	})
	id := node.Generate()
	return id.Int64()
}

func GetSnowFlakeIdStr(workMachineId int64) string {

	return fmt.Sprintf("%d", GetSnowFlakeId(workMachineId))
}

func getSnowFlake(workMachineId int64) {
	node1, err := snowflake.NewNode(workMachineId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		node = node1
	}
}