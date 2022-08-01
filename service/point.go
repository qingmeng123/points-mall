/*******
* @Author:qingmeng
* @Description:
* @File:potin
* @Date:2022/7/30
 */

package service

import (
	"gorm.io/gorm"
	"points-mall/dao"
	"points-mall/model"
)

type PointService struct {

}

//增加积分
func (s *PointService) AddPoint(uid int, number float32)(float32 ,error) {
	ok, err := s.IsExistUid(uid)
	if err!=nil{
		return 0,err
	}
	if !ok{
		err= s.CreatePoint(uid)
		if err!=nil{
			return 0, err
		}
	}
	pd:=dao.PointDao{}
	lock.Lock()
	point,err:=pd.GetPointByUid(uid)
	lock.Unlock()
	if err!=nil{
		return 0,err
	}

	return point.Number+number, pd.UpdatePoint(uid,number)

}

//生成积分表
func (s *PointService) CreatePoint(uid int) error {
	ok,err:=s.IsExistUid(uid)
	if err!=nil{
		return err
	}
	if ok{
		return nil
	}
	pd:=dao.PointDao{}
	lock.Lock()
	err=pd.CreatePoint(uid)
	lock.Unlock()
	if err!=nil{
		return err
	}
	return nil
}

//是否存在该uid的point表
func (s *PointService) IsExistUid(uid int)(bool, error) {
	pd:=dao.PointDao{}
	rwLock.RLock()
	defer rwLock.RUnlock()
	_,err:=pd.GetPointByUid(uid)

	if err!=nil{
		if err==gorm.ErrRecordNotFound{
			return false,nil
		}
		return false,err
	}
	return true,nil
}

func (s *PointService) GetPointByUid(uid int) (model.Point,error) {
	pd:=dao.PointDao{}
	rwLock.RLock()
	defer rwLock.RUnlock()
	err := s.CreatePoint(uid)
	if err!=nil{
		return model.Point{}, err
	}
	return pd.GetPointByUid(uid)
}

func (s *PointService) UpdatePoint(point model.Point) error {
	d:=dao.PointDao{}
	return d.UpdatePoint(point.Uid,point.Number)
}


