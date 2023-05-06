package main

import (
	"fmt"
	"tiktok/server/route/handler"
)

func main() {
	h := handler.SetupRouter()
	err := h.Run(":7070")
	if err != nil {
		fmt.Println(err)
	}
}
