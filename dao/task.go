/*******
* @Author:qingmeng
* @Description:
* @File:task
* @Date:2022/7/31
 */

package dao

import (
	"gorm.io/gorm"
	"points-mall/model"
)

type TaskDao struct {

}

func (d *TaskDao) CreateTask(task model.Task) error {
	err:= db.Create(&task).Error
	return err
}

func (d *TaskDao) GetTaskByUid(uid int) (tasks []model.Task,err error) {
	tx:= db.Where("uid=?",uid).Find(&tasks)
	return tasks,tx.Error
}

func (d *TaskDao) GetTask(id uint) (task model.Task,err error) {
	tx:=db.First(&task,id)
	return task,tx.Error
}

//事务处理，完成任务，增加积分
func (d *TaskDao) FinishTask(task model.Task, point model.Point) error {
	return db.Transaction(func(tx *gorm.DB) error {
		task.State=true
		if err:=tx.Save(&task).Error;err!=nil{
			return err
		}
		point.Number+=task.Point
		if err:=tx.Save(&point).Error;err!=nil{
			return err
		}
		return nil
	})
}
