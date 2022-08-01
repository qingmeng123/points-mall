/*******
* @Author:qingmeng
* @Description:
* @File:point
* @Date:2022/7/30
 */

package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"points-mall/service"
	"points-mall/tool"
)

//增加积分
func addPoint(ctx *gin.Context) {
	parsePoint := tool.ParsePoint(ctx)
	if parsePoint.Uid==0||parsePoint.Number<1{
		tool.RespParamsError(ctx)
		return
	}

	ps:=service.PointService{}
	num,err:=ps.AddPoint(parsePoint.Uid,parsePoint.Number)
	if err!=nil{
		log.Println("add point err:",err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithData(ctx,num)


}

//查看积分
func getPointByUid(ctx *gin.Context){
	parsePoint := tool.ParsePoint(ctx)
	if parsePoint.Uid==0{
		tool.RespParamsError(ctx)
		return
	}
	ps:=service.PointService{}
	point, err := ps.GetPointByUid(parsePoint.Uid)
	if err!=nil{
		log.Println("get point err:",err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithData(ctx,point)
	return
}
