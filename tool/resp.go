/*******
* @Author:qingmeng
* @Description:
* @File:resp
* @Date2021/12/10
 */

package tool

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	ERRGOODSNUMBER =errors.New("商品库存不足")
	ERRUSERPOINT=errors.New("用户余额不足")
	ERRREAPETSIGN=errors.New("用户试图重复签到")
	ERRREAPETSICKILL=errors.New("用户试图重复抢购")
	ERRREAPETPAY=errors.New("用户试图重复支付???")
	ERRHASPAID=errors.New("sorry啦，您已经确定支付咯")
)

func RespErrorWithData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": false,
		"data":   data,
	})
}

func RespParamsError(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": false,
		"data":   "参数错误",
	})
}


func RespInternalError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"status": false,
		"data":   "服务器错误",
	})
}

func RespSuccessful(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   "成功",
	})
}

func RespSuccessfulWithData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   data,
	})
}

func RespSensitiveError(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": false,
		"data":   "含有非法词汇",
	})
}
