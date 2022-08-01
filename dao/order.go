/*******
* @Author:qingmeng
* @Description:
* @File:order
* @Date:2022/7/31
 */

package dao

import (
	"gorm.io/gorm"
	"points-mall/model"
	"points-mall/tool"
)

type OrderDao struct {

}

func (d *OrderDao) CreateOrder(order model.Order) error {
	tx:=db.Create(&order)
	return tx.Error
}

func (d *OrderDao) GetOrdersByUid(uid int) (orders[]model.Order,err error) {
	tx:=db.Where("uid=?",uid).Find(&orders)
	return orders,tx.Error
}

func (d *OrderDao) GetOrderById(id uint) (model.Order, error) {
	order:=model.Order{}
	order.ID=id
	tx:= db.First(&order)
	return order,tx.Error
}

//创建事务处理，支付订单
func (d *OrderDao) PayOrder(order model.Order) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var point model.Point
		var goods model.Goods
		if err:=tx.Where("uid=?",order.Uid).First(&point).Error;err!=nil{
			return err
		}
		if err:=tx.First(&goods,order.GoodsId).Error;err!=nil{
			return err
		}
		//验证参数
		if goods.Number < 1 {
			return tool.ERRGOODSNUMBER
		}
		if point.Number < order.Point {
			return tool.ERRUSERPOINT
		}
		//减少goods数量
		goods.Number-=1
		if err := tx.Save(&goods).Error;err!=nil{
			return err
		}

		//扣除账户余额
		point.Number-=order.Point
		if err:=tx.Save(&point).Error;err!=nil{
			return err
		}

		//确认订单
		order.IsPaid=true
		if err:=tx.Save(&order).Error;err!=nil{
			return err
		}
		return nil
	})

}

func (d *OrderDao) ConfirmOrder(order model.Order) error {
	return db.Model(&order).Update("order_state",order.OrderState).Error
}

func (d *OrderDao) DelOrder(order model.Order) error {
	return db.Delete(&order).Error
}
