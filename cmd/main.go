/*******
* @Author:qingmeng
* @Description:
* @File:main
* @Date:2022/7/30
 */

package main

import (
	"points-mall/api"
	"points-mall/conf"
)

func main() {
	conf.Init()
	api.InitEngine()
}
