package main

import (
	"math/rand"
	"time"

	"curd_demo/api"
	"curd_demo/config"
)

func main() {
	config.Initialize()
	rand.Seed(time.Now().Unix())
	api.StartHttpService()
}
