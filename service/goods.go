/*******
* @Author:qingmeng
* @Description:
* @File:goods
* @Date:2022/7/30
 */

package service

import (
	"gorm.io/gorm"
	"points-mall/dao"
	"points-mall/model"
	"sync"
	"time"
)
var lock sync.Mutex
var rwLock sync.RWMutex
type GoodsService struct {
}

func (s *GoodsService) GetGoodsesByName(name string) ([]model.Goods, error) {
	d := dao.GoodsDao{}
	rwLock.RLock()
	defer rwLock.RUnlock()
	return d.GetGoodsesByName(name)
}

func (s *GoodsService) GetGoodsById(id uint) (model.Goods, error) {
	d:=dao.GoodsDao{}
	rwLock.RLock()
	defer rwLock.RUnlock()
	return d.GetGoodsById(int(id))
}

func (s *GoodsService) IsExistGoodsId(gid uint) (bool, error) {
	//防止并发导致脏读使id是否存在错误判断
	rwLock.RLock()
	_, err := s.GetGoodsById(gid)
	rwLock.RUnlock()
	if err!=nil{
		if err==gorm.ErrRecordNotFound{
			return false,nil
		}
		return false,err
	}
	return true,nil
}

func (s *GoodsService) UpdateGoods(goods model.Goods) error {
	d:=dao.GoodsDao{}
	lock.Lock()
	err:= d.SaveGoods(goods)
	lock.Unlock()
	return err
}

func (s *GoodsService) AddGoods(goods model.Goods) (time.Time,error) {
	d:=dao.GoodsDao{}
	lock.Lock()
	createAt,err:= d.CreateGoods(goods)
	lock.Unlock()
	return createAt,err
}

func (s *GoodsService) DelGoods(id uint) error {
	d:=dao.GoodsDao{}
	return d.DelGoods(id)
}

func (s *GoodsService) GetGoodsByNameAndTime(name string, createAt time.Time) (model.Goods,error) {
	d:=dao.GoodsDao{}
	rwLock.RLock()
	defer rwLock.RUnlock()
	return d.GetGoodsByNameAndTime(name,createAt)
}
