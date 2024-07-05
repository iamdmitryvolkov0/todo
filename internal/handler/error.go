package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Error struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, msg string) {
	logrus.Error(msg)
	c.AbortWithStatusJSON(statusCode, Error{msg})
}
