/*******
* @Author:qingmeng
* @Description:
* @File:point
* @Date:2022/7/30
 */

package model

import "gorm.io/gorm"

//用户积分
type Point struct {
	gorm.Model
	Uid int `gorm:"default:0; NOT NULL;primary_key" json:"uid" form:"uid"`
	Number  float32 `gorm:"default:0" json:"number" form:"number"` //积分数量
}

