package cron

func InitCron() {
	go CheckTaskRuntimeStatus()
}
