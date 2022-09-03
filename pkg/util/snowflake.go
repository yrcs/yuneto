package util

import "github.com/bwmarrin/snowflake"

func ID() uint64 {
	node, _ := snowflake.NewNode(1)
	return uint64(node.Generate().Int64())
}
