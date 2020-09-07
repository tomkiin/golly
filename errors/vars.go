package errors

var (
	OK             = newHttpError(0, "ok")
	InternalError  = newHttpError(1, "内部错误")
	ParamError     = newHttpError(2, "参数错误")
	TaskRunError   = newHttpError(3, "当前任务正在运行")
	TaskAbortError = newHttpError(4, "当前任务已经结束")
	TaskNotFound   = newHttpError(5, "该任务不存在")
	TaskRuleError  = newHttpError(6, "请先配置规则")
)
