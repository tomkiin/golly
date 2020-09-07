package handler

import (
	"github.com/gin-gonic/gin"
	"golly/api/model"
	"golly/api/service"
	"golly/errors"
	"strconv"
)

func ListTask(c *gin.Context) {
	pageQuery := c.DefaultQuery("page", "0")
	sizeQuery := c.DefaultQuery("size", "10")
	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		render(c, errors.ParamError, nil)
		return
	}
	size, err := strconv.Atoi(sizeQuery)
	if err != nil {
		render(c, errors.ParamError, nil)
		return
	}
	p := &model.Pagination{Page: page, Size: size}
	var task model.Task
	list, err := task.All(p)
	if err != nil {
		render(c, err, nil)
		return
	}
	renderPagination(c, list, p)
}

func GetTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		render(c, errors.ParamError, nil)
		return
	}
	var task model.Task
	task.ID = id
	if err = task.One(); err != nil {
		render(c, err, nil)
		return
	}
	render(c, nil, task)
}

func CreateTask(c *gin.Context) {
	var task model.Task
	err := c.BindJSON(&task)
	if err != nil {
		render(c, errors.ParamError, nil)
		return
	}
	if err = task.Add(); err != nil {
		render(c, err, nil)
		return
	}
	render(c, nil, nil)
}

func UpdateTask(c *gin.Context) {
	var task model.Task
	err := c.BindJSON(&task)
	if err != nil {
		render(c, errors.ParamError, nil)
		return
	}
	if err = task.Update(); err != nil {
		render(c, err, nil)
		return
	}
	render(c, nil, nil)
}

func DeleteTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		render(c, errors.ParamError, nil)
		return
	}
	var task model.Task
	task.ID = id
	if err = task.Delete(); err != nil {
		render(c, err, nil)
		return
	}
	render(c, nil, nil)
}

func RunTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		render(c, errors.ParamError, nil)
		return
	}
	if err := service.RunTask(id); err != nil {
		render(c, err, nil)
		return
	}
	render(c, nil, nil)
}

func AbortTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		render(c, errors.ParamError, nil)
		return
	}
	if err := service.AbortTask(id); err != nil {
		render(c, err, nil)
		return
	}
	render(c, nil, nil)
}

func LogTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		render(c, errors.ParamError, nil)
		return
	}
	logs, err := service.LogTask(id)
	if err != nil {
		render(c, err, nil)
		return
	}
	render(c, nil, logs)
}
