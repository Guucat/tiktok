package main

import (
	"tiktok/dao/mysql"
	"tiktok/mid/oss"
	"tiktok/mid/validate"
	"tiktok/router"
)

func init() {
	mysql.Init()
	validate.Init()
	oss.Init()
}

func main() {
	// Register Route
	r := router.SetupRouter()
	r.Run(":7080")
}
