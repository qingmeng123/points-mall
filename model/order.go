/*******
* @Author:qingmeng
* @Description:
* @File:order
* @Date:2022/7/30
 */

package model

import (
	"gorm.io/gorm"
)

//订单详情
type Order struct {
	gorm.Model
	Uid       		int           `gorm:"default:0"  json:"uid" form:"uid"`		//订单所有者
	GoodsId			int 			`gorm:"default:0" json:"goods_id" form:"goods_id"`	//订单商品id
	Point          float32        `gorm:"default:0" json:"point"` 	//该订单所需积分
	IsPaid			bool	`gorm:"default:0" json:"is_paid" form:"is_paid"`		//false未支付。true已支付
	OrderState     bool           `gorm:"default:0" json:"order_state" form:"order_state"`     //订单状态，true则完成订单
}
