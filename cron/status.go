package cron

import (
	"golly/api/model"
	"golly/cache"
	"time"
)

func CheckTaskRuntimeStatus() {
	for {
		check()
		time.Sleep(10 * time.Second)
	}
}

func check() {
	for id, link := range cache.TaskMap.Cache() {
		if link.CheckOver() {
			var task model.Task
			task.ID = id
			if err := task.One(); err != nil {
				continue
			}
			task.Status = "结束"
			if err := task.Update(); err != nil {
				continue
			}
			cache.TaskMap.Delete(id)  // clear stop task link
		}
	}
}
