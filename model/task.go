/*******
* @Author:qingmeng
* @Description:
* @File:task
* @Date:2022/7/30
 */

package model

import "gorm.io/gorm"

//任务
type Task struct {
	gorm.Model
	Uid int `gorm:"default:0" json:"uid" form:"uid"`			//用户id
	TaskName string `gorm:"default:'null'" json:"task_name" form:"task_name"` //任务名
	Intro   string   `gorm:"default:'null'" json:"intro" form:"intro"` //任务内容介绍
	Point 	float32   `gorm:"default:0" json:"point" form:"point"`   //任务积分
	State    bool `gorm:"default:0" json:"state" form:"state"`	//任务状态(true为已完成)
	Type    int   `gorm:"default:0" json:"type" form:"type"` 	//任务类别（0为日常任务。1为限定一次的任务,待加）
}
