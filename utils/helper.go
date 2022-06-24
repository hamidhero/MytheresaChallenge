package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type Error struct {
	ErrorCode int
	ErrorMsg  string
}

type Output struct {
	Timestamp  time.Time   `json:"TimeStamp"`
	Status     int         `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Error      []Error     `json:"error"`
}

//NewOutput produces a basic response
func NewOutput() Output {
	output := Output{}
	output.Timestamp = time.Now()
	output.Status = http.StatusOK
	output.Message = "operation successful!"
	return output
}

//SetError returns appropriate error output
func SetError(e interface{}, c *gin.Context, output Output, internalErrorCode int, externalErrorCode int) {
	var err Error
	var errorMsg string
	switch e.(type) {
	case error:
		errorMsg = e.(error).Error()
	case string:
		errorMsg = e.(string)
	}

	err.ErrorCode = internalErrorCode
	err.ErrorMsg = errorMsg
	output.Error = append(output.Error, err)

	output.Status = externalErrorCode
	output.Message = errorMsg
	c.JSON(output.Status, output)
	return
}

//accepts price and discount percent and return price after percent applied
func ApplyDiscount(price int64, percent int) int64 {
	percentFloat := float64(100 - percent) / 100
	s := fmt.Sprintf("%.0f", float64(price) * percentFloat)
	p, _ := strconv.ParseInt(s, 10, 64)
	return p
}