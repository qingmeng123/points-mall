/*******
* @Author:qingmeng
* @Description:
* @File:mq
* @Date:2022/8/1
 */

package api

import (
	"log"
	"points-mall/model"
	"points-mall/service"
)

const maxOrdersNum = 30000
var SecKillChannel = make(chan model.Order, maxOrdersNum)

//从channel中读取信息，更新数据库
func secKillConsumer() {
	for {
		order := <-SecKillChannel
		log.Printf("start:deal order uid:%d,gid:%d\n",order.Uid,order.GoodsId)
		os:=service.OrderService{}
		//先验证通道中的订单是否已确认支付
		if !order.IsPaid{
			log.Printf("err order state uid:%d,gid:%d\n",order.Uid,order.GoodsId)
			continue
		}
		err:=os.DealSicKillOrder(order.Uid,order.GoodsId)
		if err!=nil{
			log.Printf("err:deal order uid:%d,gid:%d,%v,\n",order.Uid,order.GoodsId,err)
		}else {
			log.Printf("success:deal order uid:%d,gid:%d\n",order.Uid,order.GoodsId)
		}

	}

}
