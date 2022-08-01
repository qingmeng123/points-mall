/*******
* @Author:qingmeng
* @Description:
* @File:router
* @Date:2022/7/30
 */

package api

import "github.com/gin-gonic/gin"

func InitEngine() {
	engine:=gin.Default()
	engine.Use(Cors())
	apiGroup := engine.Group("/mall")
	{
		apiGroup.POST("/user",register)				//注册
		apiGroup.GET("/user", login)                 //登陆

		userGroup := apiGroup.Group("/user", jwtAuth)
		{
			userGroup.GET("/info", getUser)       //查看信息
			userGroup.POST("/info", updateUser) //修改信息
			userGroup.POST("/password",changePassword)	//登陆后修改密码

			userGroup.GET("/task",getTaskByUid)	//通过用户id查看任务
			userGroup.POST("/task",finishTask)	//完成任务

			orderGroup:=userGroup.Group("/order")
			{
				orderGroup.PUT("",createOrder)   //	生成订单
				orderGroup.GET("",getOrders)     //查看订单
				orderGroup.DELETE("",delOrder)   //删除订单
				orderGroup.POST("", updateOrder) //更新订单
				orderGroup.PUT("/spike",createSicKillOrder)	//生成秒杀订单
				orderGroup.POST("/spike",paySicKillOrder)		//支付秒杀订单
				orderGroup.GET("/spike",getSicKillOrder)		//查看秒杀订单
				orderGroup.DELETE("/spike",delSickOrder)		//删除秒杀订单
			}


			signGroup:=userGroup.Group("/sign")
			{
				signGroup.GET("", getContinueSignDay)	//获取用户连续签到的天数
				signGroup.POST("",sign)				//用户签到
				signGroup.GET("/rank",getSignRank)	//获取用户签到排名
			}
		}

		pointGroup:=apiGroup.Group("/point")
		{
			pointGroup.GET("", getPointByUid) //查看用户积分
			pointGroup.POST("",addPoint)		//给用户增加积分
		}

		goodsGroup:=apiGroup.Group("/goods")
		{
			goodsGroup.GET("",viewGoods)                 //查看或搜索商品
			goodsGroup.PUT("",addGoods)                  //增加商品
			goodsGroup.POST("",updateGoods)              //更新商品
			goodsGroup.DELETE("",delGoods)               //删除商品
			goodsGroup.PUT("/spike",addSicKillGoods)     //添加秒杀商品
			goodsGroup.GET("/spike", getSicKillGoodsNum) //查看秒杀商品数量
		}

		taskGroup:=apiGroup.Group("/task")
		{
			taskGroup.PUT("",addTask)		//给用户增加任务
			taskGroup.GET("",getTask)		//通过任务id查看任务
		}
	}

	go secKillConsumer()

	engine.Run(":8080")
}

