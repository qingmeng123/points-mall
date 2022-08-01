/*******
* @Author:qingmeng
* @Description:
* @File:parse
* @Date2022/6/17
 */

package tool

import (
	"github.com/gin-gonic/gin"
	"log"
	"points-mall/model"
)

func ParseUser(c *gin.Context)*model.User  {
	m:=new(model.User)
	if err:=c.Bind(m);err!=nil{
		log.Println(err.Error())
	}
	return m
}

func ParsePoint(c *gin.Context)*model.Point  {
	m:=new(model.Point)
	if err:=c.Bind(m);err!=nil{
		log.Println(err.Error())
	}
	return m
}

func ParseGoods(c *gin.Context)*model.Goods  {
	m:=new(model.Goods)
	if err:=c.Bind(m);err!=nil{
		log.Println(err.Error())
	}
	return m
}

func ParseOrder(ctx *gin.Context)*model.Order {
	m:=new(model.Order)
	if err:=ctx.Bind(m);err!=nil{
		log.Println(err.Error())
	}
	return m
}


func ParseTask(c *gin.Context)*model.Task  {
	m:=new(model.Task)
	if err:=c.Bind(m);err!=nil{
		log.Println(err.Error())
	}
	return m
}