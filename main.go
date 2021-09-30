package main

import (
	"math/rand"
	"time"

	"curd_demo/api"
	"curd_demo/config"
	"curd_demo/dep"
)

func main() {
	config.Initialize()
	dep.Prepare()
	rand.Seed(time.Now().Unix())
	api.StartHttpService()
}
