package service

import (
	"golly/api/model"
	"golly/app"
	"golly/cache"
	"golly/errors"
)

func RunTask(id int) error {
	var task model.Task
	task.ID = id
	if err := task.One(); err != nil {
		return err
	}
	if len(task.Rules) == 0 {
		return errors.TaskRuleError
	}
	link := initTaskLink(task)
	if task.Status == "运行" {
		return errors.TaskRunError
	}
	task.Status = "运行"
	if err := task.Update(); err != nil {
		return err
	}
	link.Run(task.Seed)
	return nil
}

func AbortTask(id int) error {
	var task model.Task
	task.ID = id
	if err := task.One(); err != nil {
		return err
	}
	link := cache.TaskMap.Get(id)
	if link == nil {
		return errors.TaskNotFound
	}
	if task.Status != "运行" {
		return errors.TaskAbortError
	}
	task.Status = "结束"
	if err := task.Update(); err != nil {
		return err
	}
	link.Abort()
	return nil
}

func LogTask(id int) ([]*app.LogContainer, error) {
	var task model.Task
	task.ID = id
	if err := task.One(); err != nil {
		return []*app.LogContainer{}, err
	}
	link := cache.TaskMap.Get(id)
	if link == nil {
		return []*app.LogContainer{}, errors.TaskNotFound
	}
	res := link.GetLog()
	return res, nil
}

func initTaskLink(task model.Task) *app.Link {
	link := app.NewLink(task.ID)
	for _, rule := range task.Rules {
		curRule := rule // 深拷贝
		c := app.InitCollector(link, curRule.Name, curRule.Runtime, link.StopCh())
		c.SetParse(curRule.Pattern, func(e *app.HtmlElement) string {
			if curRule.Find == "attr" {
				return task.BaseURL + e.Attr(curRule.Attr)
			} else if curRule.Find == "text" {
				return e.Text()
			}
			return ""
		})
		link.Add(c)
	}
	cache.TaskMap.Put(link)
	return link
}
