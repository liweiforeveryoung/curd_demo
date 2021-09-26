package main

import (
	"math/rand"
	"time"

	"curd_demo/apis"
	"curd_demo/config"
)

func main() {
	config.Initialize()
	rand.Seed(time.Now().Unix())
	apis.StartHttpService()
}
