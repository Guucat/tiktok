package service

import (
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
)

func GetStoreId() (int64, error) {
	// 单机版机器id固定
	node, err := snowflake.NewNode(0)
	if err != nil {
		return -1, err
	}
	id := node.Generate()
	return id.Int64(), nil
}

func StoreFileWithId(c *gin.Context, id int64) error {
	return nil
}
