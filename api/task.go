/*******
* @Author:qingmeng
* @Description:
* @File:task
* @Date:2022/7/31
 */

package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"points-mall/model"
	"points-mall/service"
	"points-mall/tool"
	"sync"
)

//分发任务给指定用户
func addTask(ctx *gin.Context)  {
	parseTask := tool.ParseTask(ctx)
	wg:=sync.WaitGroup{}
	wg.Add(1)
	//帮忙查参
	go func() {
		defer wg.Done()
		if parseTask.TaskName==""||parseTask.Point<=0{
			tool.RespParamsError(ctx)
			return
		}

		if tool.CheckIfSensitive(parseTask.TaskName)||tool.CheckIfSensitive(parseTask.Intro){
			tool.RespSensitiveError(ctx)
			return
		}
	}()

	//验证用户id
	ps:=service.PointService{}
	if parseTask.Uid<=0{
		tool.RespParamsError(ctx)
		return
	}
	ok, err := ps.IsExistUid(parseTask.Uid)
	if err!=nil{
		log.Println("judge exist uid err:",err)
		tool.RespInternalError(ctx)
		return
	}
	if !ok{
		tool.RespParamsError(ctx)
		return
	}
	wg.Wait()
	task:=model.Task{
		Uid:      parseTask.Uid,
		TaskName: parseTask.TaskName,
		Intro:    parseTask.Intro,
		Point:    parseTask.Point,
		State:    false,
		Type:     0,
	}
	if parseTask.Type==1{
		task.Type=1
	}
	ts:=service.TaskService{}
	err=ts.CreateTask(task)
	if err!=nil{
		log.Println("create task err:",err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithData(ctx,task)
}

//根据任务id获取任务
func getTask(ctx *gin.Context)  {
	parseTask := tool.ParseTask(ctx)
	if parseTask.ID<=0 {
		tool.RespParamsError(ctx)
		return
	}
	ts:=service.TaskService{}
	task, err := ts.GetTask(parseTask.ID)
	if err!=nil{
		log.Println("get task err",err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithData(ctx,task)
}

//根据用户id获取任务，查询用户已完成的任务则为积分获取记录
func getTaskByUid(ctx *gin.Context){
	id,_:=ctx.Get("uid")
	ts:=service.TaskService{}
	parseTask := tool.ParseTask(ctx)
	tasks, err := ts.GetTasksByUid(id.(int))
	if err!=nil{
		log.Println("get tasks by uid err:",err)
		tool.RespInternalError(ctx)
		return
	}

	//已完成的任务
	var finishedTasks []model.Task
	if parseTask.State{
		for _, task := range tasks {
			if task.State{
				finishedTasks=append(finishedTasks,task)
			}
		}
		tool.RespSuccessfulWithData(ctx,finishedTasks)
		return
	}
	tool.RespSuccessfulWithData(ctx,tasks)
}

//完成任务
func finishTask(ctx *gin.Context) {
	id,_:=ctx.Get("uid")	//token获取用户id
	parseTask := tool.ParseTask(ctx)
	if !parseTask.State{
		tool.RespParamsError(ctx)
		return
	}
	ts:=service.TaskService{}
	task,err:=ts.GetTask(parseTask.ID)
	if err!=nil{
		log.Println("get task err:",err)
		tool.RespInternalError(ctx)
		return
	}

	//验证用户
	if task.Uid!=id.(int){
		tool.RespParamsError(ctx)
		return
	}

	ok,err:=ts.FinishTask(task)
	if err!=nil{
		log.Println("finish task err",err)
		tool.RespInternalError(ctx)
		return
	}
	if !ok{
		tool.RespErrorWithData(ctx,"完成失败")
		return
	}
	tool.RespSuccessful(ctx)
}
