/*******
* @Author:qingmeng
* @Description:
* @File:task
* @Date:2022/7/31
 */

package service

import (
	"points-mall/dao"
	"points-mall/model"
)

type TaskService struct {
}

func (s *TaskService) CreateTask(task model.Task) error {
	d:=dao.TaskDao{}
	return d.CreateTask(task)
}

//通过用户id获取任务
func (s *TaskService) GetTasksByUid(uid int) ([]model.Task,error) {
	d:=dao.TaskDao{}
	rwLock.RLock()
	defer rwLock.RUnlock()
	tasks,err:= d.GetTaskByUid(uid)
	return tasks, err
}

//通过任务id获取任务
func (s *TaskService) GetTask(id uint) (model.Task, error) {
	d:=dao.TaskDao{}
	rwLock.RLock()
	defer rwLock.RUnlock()
	task,err:= d.GetTask(id)
	return task,err
}

func (s *TaskService) FinishTask(task model.Task)(bool,error) {
	d:=dao.TaskDao{}
	ps:=PointService{}

	//任务已经完成了
	if task.State{
		return false,nil
	}
	rwLock.RLock()
	defer rwLock.RUnlock()
	point, err := ps.GetPointByUid(task.Uid)
	if err!=nil{
		return false,err
	}
	lock.Lock()
	err=d.FinishTask(task,point)
	lock.Unlock()
	if err!=nil{
		return false,err
	}
	return true,nil
}