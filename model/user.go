/*******
* @Author:qingmeng
* @Description:
* @File:user
* @Date2022/7/14
 */

package model

type User struct {
	ID       int   `json:"id" form:"id"`
	Token   string  `json:"token" form:"token"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Gender   int    `json:"gender" form:"gender"` //0为男，1为女
	Name     string `json:"name" form:"name"`              //昵称
	Phone    string `json:"phone" form:"phone"`
	Email    string `json:"email" form:"email"`
	State    int    `json:"state" form:"state"`       //(0为有效用户，1为无效)
	GroupId  int    `json:"group_id" form:"group_id"` //成员组id,1为超级管理员，0为普通用户
}
