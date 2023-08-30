package snowflake

import (
	"log"

	"github.com/bwmarrin/snowflake"
)

var (
	node *snowflake.Node
	err  error
)

func Initialize(nodeID int64) {
	node, err = snowflake.NewNode(nodeID)
	if err != nil {
		log.Fatalf("snowflake init failed: %s\n", err.Error())
		return
	}
}

// GetID 获得一个唯一ID
func GetID() int64 {
	return node.Generate().Int64()
}
