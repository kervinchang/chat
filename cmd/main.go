// main.go
package main

import (
	"github.com/kervinchang/chat/config"
	_ "github.com/kervinchang/chat/internal/model"
	"github.com/kervinchang/chat/internal/router"
)

func main() {
	engin := router.Setup()
	err := engin.Run(config.Addr)
	if err != nil {
		panic(err)
	}
}
