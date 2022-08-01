/*******
* @Author:qingmeng
* @Description:
* @File:goods
* @Date:2022/7/30
 */

package model

import "gorm.io/gorm"

type Goods struct {
	gorm.Model
	GoodsName     string    `gorm:"default:'null'" json:"goods_name" form:"goods_name"`     //商品名字
	Picture       string    `gorm:"default:'null'" json:"picture" form:"picture"`        //商品图片
	Price         float32   `gorm:"default:0" json:"price" form:"price"`          		//积分价格
	GoodsIntro    string    `gorm:"default:'null'" json:"goods_intro" form:"goods_intro"`    //商品介绍
	Number        int       `gorm:"default:0" json:"number" form:"number"`         //商品数量
}