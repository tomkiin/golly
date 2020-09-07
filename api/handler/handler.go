package handler

import (
	"github.com/gin-gonic/gin"
	"golly/api/model"
	"golly/errors"
	"net/http"
)

type Request struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func render(c *gin.Context, err error, data interface{}) {
	code, msg := errors.ParseError(err)
	c.JSON(http.StatusOK, &Request{Code: code, Message: msg, Data: data})
}

func renderPagination(c *gin.Context, data interface{}, p *model.Pagination) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    data,
		"total":   p.Total,
		"page":    p.Page,
		"size":    p.Size,
	})
}
