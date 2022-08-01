/*******
* @Author:qingmeng
* @Description:
* @File:goods
* @Date:2022/7/30
 */

package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"points-mall/cache"
	"points-mall/model"
	"points-mall/service"
	"points-mall/tool"
	"sync"
)

//商品主页
func viewGoods(ctx *gin.Context) {
	parseGoods := tool.ParseGoods(ctx)
	gs := service.GoodsService{}
	//商品id查找
	if parseGoods.ID != 0 {
		goods, err := gs.GetGoodsById(parseGoods.ID)
		if err != nil {
			log.Println("get goods by id err:", err)
			tool.RespInternalError(ctx)
			return
		}
		if goods.GoodsName == "" {
			tool.RespSuccessfulWithData(ctx, "暂无该商品")
			return
		}
		tool.RespSuccessfulWithData(ctx, goods)
		return
	}

	//查看所有商品或按名称模糊查找
	goodses, err := gs.GetGoodsesByName(parseGoods.GoodsName)
	if err != nil {
		log.Println("get goodses by name err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if goodses == nil {
		tool.RespSuccessfulWithData(ctx, "暂无商品")
		return
	}
	tool.RespSuccessfulWithData(ctx, goodses)
	return

}

//更新商品
func updateGoods(ctx *gin.Context) {
	parseGoods := tool.ParseGoods(ctx)

	goods := model.Goods{}
	gs := service.GoodsService{}

	if parseGoods.ID==0{
		tool.RespParamsError(ctx)
		return
	}
	//取出商品
	flag, err := gs.IsExistGoodsId(parseGoods.ID)
	if err != nil {
		fmt.Println("judge exist goods id err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if !flag {
		tool.RespParamsError(ctx)
		return
	}
	goods, err = gs.GetGoodsById(parseGoods.ID)

	//取出有效参数,起个协程帮哈忙
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if parseGoods.GoodsIntro!=""&&!tool.CheckIfSensitive(parseGoods.GoodsIntro){
			goods.GoodsIntro=parseGoods.GoodsIntro
		}

		if parseGoods.GoodsName!=""&&!tool.CheckIfSensitive(parseGoods.GoodsName){
			goods.GoodsName=parseGoods.GoodsName
		}
		wg.Done()
	}()

	if parseGoods.Number>0{
		goods.Number=parseGoods.Number
	}

	if parseGoods.Price>0{
		goods.Price=parseGoods.Price
	}

	if parseGoods.Picture!=""{
		goods.Picture=parseGoods.Picture
	}

	wg.Wait()
	err = gs.UpdateGoods(goods)
	if err != nil {
		log.Println("update goods err:", err)
		tool.RespInternalError(ctx)
		return
	}
	//返回修改后的结果
	selectGoodsById, err := gs.GetGoodsById(goods.ID)
	if err!=nil{
		log.Println("get goods by id err:",err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithData(ctx,selectGoodsById)
}

//添加商品
func addGoods(ctx *gin.Context) {
	parseGoods := tool.ParseGoods(ctx)

	if parseGoods.GoodsName == "" || parseGoods.Price < 0 || parseGoods.Number < 0  {
		tool.RespParamsError(ctx)
		return
	}

	//验证goodsName
	flag := tool.CheckIfSensitive(parseGoods.GoodsName)
	if flag {
		tool.RespSensitiveError(ctx)
		return
	}

	//验证goodsIntro
	flag = tool.CheckIfSensitive(parseGoods.GoodsIntro)
	if flag {
		tool.RespSensitiveError(ctx)
		return
	}

	goods := model.Goods{
		GoodsName:  parseGoods.GoodsName,
		Picture:    parseGoods.Picture,
		Price:      parseGoods.Price,
		GoodsIntro: parseGoods.GoodsIntro,
		Number:     parseGoods.Number,
	}
	gs := service.GoodsService{}
	_,err:= gs.AddGoods(goods)
	if err != nil {
		log.Println("add goods err:", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}

//下架商品
func delGoods(ctx *gin.Context)  {
	parseGoods := tool.ParseGoods(ctx)
	if parseGoods.ID==0{
		tool.RespParamsError(ctx)
		return
	}

	gs:=service.GoodsService{}
	ok,err:=gs.IsExistGoodsId(parseGoods.ID)
	if err!=nil{
		log.Println("judge goods id err:",err)
		tool.RespInternalError(ctx)
		return
	}
	if !ok{
		tool.RespParamsError(ctx)
		return
	}

	err=gs.DelGoods(parseGoods.ID)
	if err!=nil{
		log.Println("del goods err:",err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}


//添加秒杀商品
func addSicKillGoods(ctx *gin.Context) {
	parseGoods := tool.ParseGoods(ctx)

	if parseGoods.GoodsName == "" || parseGoods.Price < 0 || parseGoods.Number < 0  {
		tool.RespParamsError(ctx)
		return
	}

	//验证goodsName
	flag := tool.CheckIfSensitive(parseGoods.GoodsName)
	if flag {
		tool.RespSensitiveError(ctx)
		return
	}

	//验证goodsIntro
	flag = tool.CheckIfSensitive(parseGoods.GoodsIntro)
	if flag {
		tool.RespSensitiveError(ctx)
		return
	}

	goods := model.Goods{
		GoodsName:  parseGoods.GoodsName,
		Picture:    parseGoods.Picture,
		Price:      parseGoods.Price,
		GoodsIntro: parseGoods.GoodsIntro,
		Number:     parseGoods.Number,
	}
	gs := service.GoodsService{}
	createAt,err:= gs.AddGoods(goods)
	if err != nil {
		log.Println("add goods err:", err)
		tool.RespInternalError(ctx)
		return
	}
	goods,err=gs.GetGoodsByNameAndTime(goods.GoodsName,createAt)
	if err!=nil{
		log.Println("get goods err:",err)
		tool.RespInternalError(ctx)
		return
	}

	err=cache.SetCacheGoods(goods)
	if err!=nil{
		log.Println("cache goods err:",err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithData(ctx,goods.ID)
}

//查看秒杀商品数量
func getSicKillGoodsNum(ctx *gin.Context) {
	parseGoods := tool.ParseGoods(ctx)
	if parseGoods.ID<1{
		tool.RespParamsError(ctx)
		return
	}

	num, err:=cache.GetCacheGoods(parseGoods.ID)
	if err!=nil{
		log.Println("cache goods err:",err)
		tool.RespParamsError(ctx)
		return
	}

	tool.RespSuccessfulWithData(ctx,num)
}
