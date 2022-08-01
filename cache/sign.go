/*******
* @Author:qingmeng
* @Description:
* @File:sign
* @Date:2022/7/31
 */

package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

const day =  time.Hour*24

type Sign struct {

}

//用户当日签到,返回签到排名和error
func (s Sign) DoSign(uid int) (int,error) {
	//计算offset,从0开始，即每月当前天数减一
	var offset  = time.Now().Local().Day() - 1
	var keys  = s.buildSignKey(uid)
	//将指定key的offset下标设置为1,即为签到
	_,err := RedisClient.SetBit(keys,int64(offset),1).Result()
	if err != nil {
		return 0, err
	}

	//获取当日日期
	dayStr:=time.Now().Format("2006-01-02 ")
	result, _ := RedisClient.Get(dayStr).Result()
	//第一个签到的生成dayStr
	if result==""{
		err = RedisClient.Set(dayStr, 1,day).Err()
		if err != nil {
			return 0, err
		}
		num:=1
		//将该用户签到排名存入缓存中
		RedisClient.Set(s.buildRankKey(uid), num, day)
		return num, nil
	}
	//当日已签到的人数加1
	num, err := RedisClient.Incr(dayStr).Result()
	if err!=nil{
		return 0,err
	}
	RedisClient.Set(s.buildRankKey(uid), num, day)
	return int(num), nil
}

//获取用户当日签到排名
func (s Sign) GetRank(uid int) (int,error) {
	result, err := RedisClient.Get(s.buildRankKey(uid)).Result()
	if err!=nil{
		return 0,err
	}
	num,err:=strconv.Atoi(result)
	if err!=nil{
		return 0,err
	}
	return num,nil
}

//判断用户今日是否签到（0未签，1已签）
func (s Sign) CheckSign(uid int)(int64,error) {
	var keys  = s.buildSignKey(uid)
	var offset  = time.Now().Local().Day() - 1

	return RedisClient.GetBit(keys,int64(offset)).Result()
}

//获取用户本月签到的次数
func (s Sign) GetSignCount(uid int)(int64,error)  {
	var keys  = s.buildSignKey(uid)
	//未到的日期都为0
	count:=redis.BitCount{Start: 0,End: 30}
	//获得字符串类型键中值是1的二进制位个数
	return RedisClient.BitCount(keys,&count).Result()
}

//连续签到天数
func (s Sign) ContinueSignDay(uid int)(int,error)  {
	keys:=s.buildSignKey(uid)
	offset:=time.Now().Local().Day() - 1
	count:=0
	for{
		//从今天开始算，往前倒数有几个1
		result, err := RedisClient.GetBit(keys, int64(offset)).Result()
		if err!=nil{
			return 0, err
		}
		if result==1{
			count++
		}
		//退出循环条件
		if result!=1||offset<=0{
			return count,nil
		}
		offset--

	}
}

//用户签到的key
func (s Sign) buildSignKey(uid int) string {
	var nowDate = s.formatDate()
	return fmt.Sprintf("u:sign:%d:%s",uid,nowDate)
}

//用户每日签到排名的key
func (s Sign) buildRankKey(uid int) string {
	var nowDate = s.formatDate()
	return fmt.Sprintf("u:rank:%d:%s",uid,nowDate)
}


//获取当前的月份
func (s Sign) formatDate() string {
	return time.Now().Format("2006-01")
}
