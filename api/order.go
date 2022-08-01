/*******
* @Author:qingmeng
* @Description:
* @File:order
* @Date:2022/7/31
 */

package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"points-mall/cache"
	"points-mall/model"
	"points-mall/service"
	"points-mall/tool"
)

//生成订单
func createOrder(ctx *gin.Context) {
	id, _ := ctx.Get("uid")
	parseOrder := tool.ParseOrder(ctx)
	if parseOrder.GoodsId < 1 {
		tool.RespParamsError(ctx)
		return
	}
	gs := service.GoodsService{}
	os := service.OrderService{}

	//验证商品
	ok, err := gs.IsExistGoodsId(uint(parseOrder.GoodsId))
	if err != nil {
		log.Println("judge exist goods err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if !ok {
		tool.RespParamsError(ctx)
		return
	}

	goods, err := gs.GetGoodsById(uint(parseOrder.GoodsId))
	if err != nil {
		log.Println("get goods by id err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if goods.Number < 1 {
		tool.RespErrorWithData(ctx, tool.ERRGOODSNUMBER.Error())
		return
	}

	//生成订单
	order := model.Order{
		Uid:        id.(int),
		GoodsId:    int(goods.ID),
		Point:      goods.Price,
		OrderState: false,
	}

	err = os.CreateOrder(order)
	if err != nil {
		log.Println("create order err:", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
}

//通过用户token获取订单,查看已完成的订单可用于计入积分消费流水
func getOrders(ctx *gin.Context) {
	id, _ := ctx.Get("uid")
	os := service.OrderService{}
	parseOrder := tool.ParseOrder(ctx)
	orders, err := os.GetOrdersByUid(id.(int))
	if err != nil {
		log.Println("get orders by uid err:", err)
		tool.RespInternalError(ctx)
		return
	}

	//已完成的订单
	var finishedOrders []model.Order
	if parseOrder.OrderState {
		for _, order := range orders {
			if order.OrderState {
				finishedOrders = append(finishedOrders, order)
			}
		}
		tool.RespSuccessfulWithData(ctx, finishedOrders)
		return
	}

	tool.RespSuccessfulWithData(ctx, orders)
}

//更新订单
func updateOrder(ctx *gin.Context) {
	id, _ := ctx.Get("uid")
	os := service.OrderService{}
	parseOrder := tool.ParseOrder(ctx)
	//验证参数
	if !parseOrder.OrderState && !parseOrder.IsPaid {
		tool.RespParamsError(ctx)
		return
	}
	if parseOrder.OrderState && !parseOrder.IsPaid {
		tool.RespParamsError(ctx)
		return
	}
	//验证订单id
	if parseOrder.ID < 1 {
		tool.RespParamsError(ctx)
		return
	}
	ok, err := os.IsExistOrder(parseOrder.ID)
	if err != nil {
		log.Println("judge order by id err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if !ok {
		tool.RespParamsError(ctx)
		return
	}
	//获取订单
	order, err := os.GetOrdersById(parseOrder.ID)
	if err != nil {
		log.Println("get order by id err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if order.Uid != id.(int) {
		tool.RespParamsError(ctx)
		return
	}

	if parseOrder.IsPaid {
		ok, err = os.PayOrder(order)
		if !ok {
			log.Println("pay order err:", err)
			tool.RespInternalError(ctx)
			return
		} else {
			if err != nil {
				tool.RespErrorWithData(ctx, err.Error())
				return
			}
		}
	}

	//已完成
	if parseOrder.OrderState {

		err = os.ConfirmOrder(order)
		if err != nil {
			log.Println("confirm order err:", err)
			tool.RespInternalError(ctx)
			return
		}
	}

	tool.RespSuccessful(ctx)

}

//删除订单
func delOrder(ctx *gin.Context) {
	id, _ := ctx.Get("uid")
	os := service.OrderService{}
	parseOrder := tool.ParseOrder(ctx)
	//验证订单id
	if parseOrder.ID < 1 {
		tool.RespParamsError(ctx)
		return
	}
	ok, err := os.IsExistOrder(parseOrder.ID)
	if err != nil {
		log.Println("judge order by id err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if !ok {
		tool.RespParamsError(ctx)
		return
	}
	order, err := os.GetOrdersById(parseOrder.ID)
	if err != nil {
		log.Println("get order by id err:", err)
		tool.RespInternalError(ctx)
		return
	}
	//验证是否为本人的订单
	if order.Uid != id.(int) {
		tool.RespParamsError(ctx)
		return
	}
	err = os.DelOrder(order)
	if err != nil {
		log.Println("del order err:", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}

//秒杀下单
func createSicKillOrder(ctx *gin.Context) {
	id, _ := ctx.Get("uid")
	parseOrder := tool.ParseOrder(ctx)
	if parseOrder.GoodsId < 1 {
		tool.RespParamsError(ctx)
		return
	}

	//缓存中验证商品
	num, err := cache.GetCacheGoods(uint(parseOrder.GoodsId))
	if num < 1 {
		tool.RespErrorWithData(ctx, tool.ERRGOODSNUMBER.Error())
		return
	}
	if err != nil {
		log.Println("get cache goods err:", err)
		tool.RespInternalError(ctx)
		return
	}

	//生成订单存入缓存
	order := model.Order{
		Uid:     id.(int),
		GoodsId: parseOrder.GoodsId,
		IsPaid:  false,
	}
	//验证是否重复抢购
	ok, _ := cache.GetCacheOrder(order.Uid, order.GoodsId)
	if ok {
		tool.RespErrorWithData(ctx, tool.ERRREAPETSICKILL.Error())
		return
	}

	//添加订单
	err = cache.SetCacheOrder(order)
	if err != nil {
		log.Println("create order err:", err)
		tool.RespInternalError(ctx)
		return
	}

	//将缓存中的商品数量减一,抢购嘛，抢到手了再说
	err = cache.DecrCacheGoods(uint(order.GoodsId))
	if err != nil {
		log.Println("decr cache goods err:", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
}

//查看秒杀订单
func getSicKillOrder(ctx *gin.Context) {
	id, _ := ctx.Get("uid")
	parseOrder := tool.ParseOrder(ctx)

	if parseOrder.ID < 1 {
		tool.RespParamsError(ctx)
		return
	}
	ok, order := cache.GetCacheOrder(id.(int), parseOrder.GoodsId)
	if !ok {
		tool.RespParamsError(ctx)
		return
	}
	tool.RespSuccessfulWithData(ctx, order)
}

//秒杀支付
func paySicKillOrder(ctx *gin.Context) {
	id, _ := ctx.Get("uid")
	parseOrder := tool.ParseOrder(ctx)
	if parseOrder.GoodsId < 1 {
		tool.RespParamsError(ctx)
		return
	}
	if !parseOrder.IsPaid {
		tool.RespParamsError(ctx)
		return
	}

	//验证订单
	ok, order := cache.GetCacheOrder(id.(int), parseOrder.GoodsId)
	if !ok {
		tool.RespParamsError(ctx)
		return
	}

	if order.IsPaid {
		tool.RespErrorWithData(ctx, tool.ERRREAPETPAY.Error())
		return
	}

	//支付订单
	order.IsPaid = true
	err := cache.SetCacheOrder(order)
	if err!=nil{
		log.Println("update cache order err:",err)
		tool.RespInternalError(ctx)
		return
	}
	SecKillChannel <- order
	tool.RespSuccessful(ctx)
}

//取消秒杀订单
func delSickOrder(ctx *gin.Context) {
	id, _ := ctx.Get("uid")
	parseOrder := tool.ParseOrder(ctx)
	//验证订单
	if parseOrder.ID < 1 {
		tool.RespParamsError(ctx)
		return
	}
	ok, order := cache.GetCacheOrder(id.(int), parseOrder.GoodsId)
	if !ok {
		tool.RespParamsError(ctx)
		return
	}

	if order.OrderState {
		tool.RespErrorWithData(ctx, tool.ERRHASPAID.Error())
		return
	}

	//删除缓存中订单
	err := cache.DelCacheOrder(order)
	if err != nil {
		log.Println("del cache order err:", err)
		tool.RespInternalError(ctx)
		return
	}

	//缓存中商品数量加一
	err = cache.IncrCacheGoods(uint(order.GoodsId))
	if err != nil {
		log.Println("incr cache goods err:", err)
	}

	tool.RespSuccessful(ctx)
}
