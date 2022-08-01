/*******
* @Author:qingmeng
* @Description:
* @File:point
* @Date:2022/7/30
 */

package dao

import (
	"points-mall/model"
)

type PointDao struct {
}

func (d *PointDao) GetPointByUid(uid int) (point model.Point,err error) {
	tx:=db.Where("uid=?",uid).First(&point)
	return point,tx.Error
}

func (d *PointDao) CreatePoint(uid int) error {
	point:=model.Point{Uid: uid}
	tx:=db.Create(&point)
	return tx.Error
}

func (d *PointDao) UpdatePoint(uid int, number float32) error {
	point:=model.Point{Uid: uid,Number: number}
	return db.Model(&point).Where("uid=?",uid).Update("number",number).Error
}


