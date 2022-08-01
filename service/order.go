/*******
* @Author:qingmeng
* @Description:
* @File:order
* @Date:2022/7/31
 */

package service

import (
	"gorm.io/gorm"
	"points-mall/dao"
	"points-mall/model"
	"points-mall/tool"
)
type OrderService struct {

}

//处理秒杀的订单
func (s *OrderService) DealSicKillOrder(uid int,gid int)error  {
	//完善订单信息
	gs:=GoodsService{}
	rwLock.RLock()
	goods, err := gs.GetGoodsById(uint(gid))
	rwLock.RUnlock()
	if err!=nil{
		return err
	}
	order:=model.Order{
		Uid:        uid,
		GoodsId:    gid,
		Point:      goods.Price,
		IsPaid:     false,	//现在还未扣除用户积分
		OrderState: false,
	}
	err = s.CreateOrder(order)
	if err!=nil{
		return err
	}
	_, err= s.PayOrder(order)
	return err
}

func (s *OrderService) CreateOrder(order model.Order) error {
	d:=dao.OrderDao{}
	rwLock.RLock()
	defer rwLock.RUnlock()
	return d.CreateOrder(order)
}

func (s *OrderService) GetOrdersByUid(uid int) ([]model.Order, error) {
	d:=dao.OrderDao{}
	rwLock.RLock()
	defer rwLock.RUnlock()
	return d.GetOrdersByUid(uid)

}

func (s *OrderService) GetOrdersById(id uint) (model.Order, error) {
	d:=dao.OrderDao{}
	rwLock.RLock()
	defer rwLock.RUnlock()
	return d.GetOrderById(id)
}

func (s *OrderService) IsExistOrder(id uint) (bool, error) {
	rwLock.RLock()
	defer rwLock.RUnlock()
	_,err:=s.GetOrdersById(id)

	if err!=nil{
		if err==gorm.ErrRecordNotFound{
			return false,nil
		}
		return false,err
	}
	return true,nil
}

//系统错误返回false,err
func (s *OrderService) PayOrder(order model.Order)(bool, error) {
	d:=dao.OrderDao{}
	lock.Lock()
	err:=d.PayOrder(order)
	lock.Unlock()
	if err!=nil{
		if err==tool.ERRUSERPOINT||err==tool.ERRGOODSNUMBER{
			return true,err
		}
		return false,err
	}

	return true,nil
}

func (s *OrderService) ConfirmOrder(order model.Order) error {
	lock.Lock()
	defer lock.Unlock()
	order.OrderState=true
	d:=dao.OrderDao{}
	return d.ConfirmOrder(order)
}

func (s *OrderService)DelOrder(order model.Order) error {
	d:=dao.OrderDao{}
	return d.DelOrder(order)
}
