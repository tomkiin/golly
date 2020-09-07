package main

import (
	"fmt"
	"golly/api"
	"golly/api/model"
	"golly/cache"
	"golly/cron"
)

func main() {
	fmt.Println("golly")

	cache.InitCache()
	cron.InitCron()
	model.InitDB()
	api.InitAPI()
}