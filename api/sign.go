/*******
* @Author:qingmeng
* @Description:
* @File:sign
* @Date:2022/7/31
 */

package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"points-mall/cache"
	"points-mall/tool"
)

//用户签到
func sign(ctx *gin.Context) {
	ids,_:=ctx.Get("uid")
	id,_:=ids.(int)
	s:=cache.Sign{}
	//是否签到
	check, err := s.CheckSign(id)
	if err!=nil{
		log.Println("check sign err:",err)
		tool.RespInternalError(ctx)
		return
	}
	if check==1{	//已签到
		tool.RespErrorWithData(ctx,tool.ERRREAPETSIGN.Error())
		return
	}

	//签到
	num,err := s.DoSign(id)
	if err!=nil{
		log.Println("sign err:",err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithData(ctx,num)
}

//获取用户连续签到天数
func getContinueSignDay(ctx *gin.Context)  {
	ids,_:=ctx.Get("uid")
	id,_:=ids.(int)
	s:=cache.Sign{}
	count, err := s.ContinueSignDay(id)
	if err !=nil{
		log.Println("count err:", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithData(ctx,count)
}

//获取用户今日签到排名
func getSignRank(ctx *gin.Context)  {
	ids,_:=ctx.Get("uid")
	id,_:=ids.(int)
	s:=cache.Sign{}
	count,err:=s.GetRank(id)
	if err !=nil{
		log.Println("get rank err:", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithData(ctx,count)

}