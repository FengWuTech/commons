package app

import (
	"github.com/FengWuTech/commons/logger"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int, error) {
	err := c.ShouldBind(form)
	if err != nil {
		errMsg := err.Error()
		logger.Warn(errMsg)
		return http.StatusOK, INVALID_PARAMS, err
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		logger.Warn(err.Error())
		return http.StatusOK, ERROR, err
	}
	if !check {
		logger.Warn("valid check failed")
		return http.StatusOK, INVALID_PARAMS, valid.Errors[0]
	}

	return http.StatusOK, SUCCESS, nil
}

