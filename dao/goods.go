/*******
* @Author:qingmeng
* @Description:
* @File:goods
* @Date:2022/7/30
 */

package dao

import (
	"points-mall/model"
	"time"
)

type GoodsDao struct {
}

//模糊查询
func (d *GoodsDao) GetGoodsesByName(name string) ([]model.Goods, error) {
	var goodses []model.Goods
	name = "%" + name + "%"
	tx := db.Where("goods_name LIKE ?", name).Find(&goodses)
	return goodses, tx.Error
}

//准确查询
func (d *GoodsDao) GetGoodsById(id int) (model.Goods, error) {
	var goods model.Goods
	goods.ID = uint(id)
	tx := db.First(&goods)
	return goods, tx.Error
}

func (d *GoodsDao) SaveGoods(goods model.Goods) error {
	tx:=db.Save(&goods)
	return tx.Error
}

func (d *GoodsDao) CreateGoods(goods model.Goods)(time.Time, error) {
	err:= db.Create(&goods).Error
	return goods.CreatedAt, err
}

func (d *GoodsDao) DelGoods(id uint) error {
	return db.Delete(&model.Goods{},id).Error

}

func (d *GoodsDao) GetGoodsByNameAndTime(name string, at time.Time) (goods model.Goods,err error) {
	tx:=db.Where("goods_name=? and created_at=?",name,at).First(&goods)
	return goods,tx.Error
}
