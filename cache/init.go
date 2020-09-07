package cache

var (
	TaskMap *TaskCache
)

func InitCache() {
	TaskMap = NewTaskCache()
}
