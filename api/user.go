/*******
* @Author:qingmeng
* @Description:
* @File:user
* @Date:2022/7/30
 */

package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"points-mall/api/auth"
	"points-mall/cache"
	"points-mall/conf"
	"points-mall/model"
	"points-mall/pbfile/pb"
	"points-mall/service"
	"points-mall/tool"
	"strconv"
	"sync"
	"time"
)

const day = 24 * 60 * 60 //一天时间

//连接用户中心
func LinkUserServer()  (*grpc.ClientConn,pb.UserServiceClient) {
	conn, err := grpc.Dial(conf.UserTcpPort, grpc.WithTransportCredentials(auth.GetCreds()),grpc.WithUnaryInterceptor(auth.UnaryInterceptor))
	if err!=nil{
		log.Fatalln("did not connect:",err)
	}
	c:=pb.NewUserServiceClient(conn)
	return conn,c
}

//注册
func register(ctx *gin.Context)  {
	conn, c := LinkUserServer()
	defer conn.Close()
	parseUser := tool.ParseUser(ctx)

	user:=&pb.User{
		Username: parseUser.Username,
		Password: parseUser.Password,
	}
	resp, err := c.Register(context.Background(), &pb.RegisterReq{User: user})
	if err!=nil{
		log.Println("register err:",resp.Resp)
		tool.RespInternalError(ctx)
		return
	}
	if !resp.Resp.Status{
		tool.RespErrorWithData(ctx,resp.Resp.Data)
		return
	}

	tool.RespSuccessful(ctx)

	return
}

//登录
func login(ctx *gin.Context)  {
	conn, client := LinkUserServer()
	defer conn.Close()
	parseUser:= tool.ParseUser(ctx)

	//起个协程帮忙登录。。。
	wg:=sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		resp, err := client.Login(context.Background(), &pb.LoginReq{
			Username: parseUser.Username,
			Password: parseUser.Password,
		})
		if err!=nil{
			log.Println("login err:",resp.Resp)
			tool.RespInternalError(ctx)
			return
		}

		if !resp.Resp.Status{
			tool.RespErrorWithData(ctx,resp.Resp.Data)
			return
		}

	}()

	//获取用户
	UserResp, err:= client.GetUserByUserName(context.Background(), &pb.GetUserByUserNameReq{UserName: parseUser.Username})
	if err!=nil{
		log.Println("get user err：",UserResp.Resp)
		tool.RespInternalError(ctx)
		return
	}

	if !UserResp.Resp.Status{
		tool.RespErrorWithData(ctx,UserResp.Resp.Data)
		return
	}

	user:=model.User{
		ID: int(UserResp.User.ID),
		Username: UserResp.User.Username,
		Password: UserResp.User.Password,
		State:    int(UserResp.User.State),
		GroupId:  int(UserResp.User.GroupId),
	}

	//生成token
	token, err := service.CreateToken(user, day)
	if err!=nil{
		log.Println("create token err:",err)
		tool.RespInternalError(ctx)
		return
	}
	//将token存入缓存，每次登陆刷新该token
	_, err = cache.RedisClient.Set("token"+strconv.Itoa(int(user.ID)), token, day*time.Second).Result()
	if err != nil {
		log.Println("set token to redis err:", err)
		tool.RespInternalError(ctx)
		return
	}
	user.Token=token
	wg.Wait()
	tool.RespSuccessfulWithData(ctx,user)
	return
}

//修改密码
func changePassword(ctx *gin.Context)  {
	id,_:=ctx.Get("uid")
	oldPassword := ctx.PostForm("old_password")
	newPassword := ctx.PostForm("new_password")

	conn, client := LinkUserServer()
	defer conn.Close()
	resp, err := client.ChangePwd(context.Background(), &pb.ChangePwdReq{
		ID:          int32(id.(int)),
		OldPassword: oldPassword,
		NewPassword: newPassword,
	})
	if err!=nil{
		log.Println("change password err:",err)
		tool.RespInternalError(ctx)
		return
	}

	if !resp.Resp.Status{
		tool.RespErrorWithData(ctx,resp.Resp.Data)
		return
	}
	tool.RespSuccessfulWithData(ctx,resp.Resp.Data)
}

//获取用户
func getUser(ctx *gin.Context) {
	id,_:=ctx.Get("uid")
	conn, client := LinkUserServer()
	defer conn.Close()
	resp, err := client.GetUserById(context.Background(), &pb.GetUserByIdReq{ID: int32(id.(int))})
	if err!=nil{
		log.Println("get user by id err:",err)
		tool.RespInternalError(ctx)
		return
	}
	if !resp.Resp.Status{
		tool.RespErrorWithData(ctx,resp.Resp)
		return
	}
	tool.RespSuccessfulWithData(ctx,resp.User)
}

//更新用户
func updateUser(ctx *gin.Context) {
	conn, client := LinkUserServer()
	defer conn.Close()
	id,_:=ctx.Get("uid")
	parseUser := tool.ParseUser(ctx)
	pbUser:=pb.User{
		ID:       int32(id.(int)),
		Username: parseUser.Username,
		Gender:   0,
		Name:     parseUser.Name,
		Phone:    parseUser.Phone,
		Email:    parseUser.Email,
		State:    0,
		GroupId:  0,
	}
	if parseUser.Gender==1{
		pbUser.Gender=1
	}
	if parseUser.State==1{
		pbUser.State=1
	}
	if parseUser.GroupId==1{
		pbUser.GroupId=1
	}

	resp, err := client.UpdateUser(context.Background(),&pb.UpdateUserReq{User: &pbUser})
	if err!=nil{
		log.Println("update user err:",err)
		tool.RespInternalError(ctx)
		return
	}
	if !resp.Resp.Status{
		tool.RespErrorWithData(ctx,resp.Resp)
		return
	}
	tool.RespSuccessfulWithData(ctx,resp.User)
}

