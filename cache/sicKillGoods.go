/*******
* @Author:qingmeng
* @Description:
* @File:sickillProject
* @Date:2022/8/1
 */

package cache

import (
	"fmt"
	"points-mall/model"
	"strconv"
)

func buildGoodsId(id uint) string {
	return fmt.Sprintf("goods-%d", id)
}

func buildOrderUidGid(uid int,gid int) string {
	return fmt.Sprintf("user-%d-gid-%d", uid,gid)
}


//缓存中添加订单
func SetCacheOrder(order model.Order)error  {
	key:= buildOrderUidGid(order.Uid,order.GoodsId)
	fields := map[string]interface{}{
		"uid": order.Uid,
		"gid":order.GoodsId,
		"is-paid":order.IsPaid,
	}

	return RedisClient.HMSet(key,fields).Err()
}

//删除缓存中的订单
func DelCacheOrder(order model.Order)error  {
	key:=buildOrderUidGid(order.Uid,order.GoodsId)
	_, err := RedisClient.HDel(key, "uid", "gid", "is-paid").Result()
	return err
}

//获取缓存中某用户的订单
func GetCacheOrder(uid int,gid int)(ok bool,order model.Order) {
	key:= buildOrderUidGid(uid,gid)
	ids,_:=RedisClient.HGet(key,"uid").Result()
	if ids==""{
		return false,order
	}

	order.Uid,_=strconv.Atoi(ids)
	gids,_:=RedisClient.HGet(key,"gid").Result()
	if gids==""{
		return false,order
	}
	order.GoodsId,_=strconv.Atoi(gids)

	isPaid,_:=RedisClient.HGet(key,"is-paid").Result()
	is,_:=strconv.Atoi(isPaid)
	if is==1{
		order.OrderState=true
	}else {
		order.IsPaid=false
	}
	return true,order
}

//商家上架商品的信息
func SetCacheGoods(goods model.Goods) error {
	key := buildGoodsId(goods.ID)
	_, err := RedisClient.Set(key,goods.Number,day).Result()
	return  err
}

//获取缓存中的商品数量
func GetCacheGoods(gid uint) (int,error ){
	key := buildGoodsId(gid)
	val, err := RedisClient.Get(key).Result()
	num,err:=strconv.Atoi(val)
	if err!=nil{
		return 0,err
	}
	return num,err
}

//缓存中商品数量减一
func DecrCacheGoods(gid uint) error {
	key:=buildGoodsId(gid)
	_, err := RedisClient.Decr(key).Result()
	return err
}

//缓存中商品数量加一
func IncrCacheGoods(gid uint) error {
	key:=buildGoodsId(gid)
	_, err := RedisClient.Incr(key).Result()
	return err
}



