## 积分商城

### 架构设计
│  go.mod
│  go.sum
│  README.md
│                          
├─api
│  │  goods.go
│  │  middleware.go
│  │  mq.go
│  │  order.go
│  │  point.go
│  │  router.go
│  │  sign.go
│  │  task.go
│  │  user.go
│  │  
│  └─auth
│          interceptor.go
│          issue.go
│          
├─cache
│      redis.go
│      sicKillGoods.go
│      sign.go
│      
├─cmd
│      main.go
│      
├─conf
│      conf.go
│      config.ini
│      
├─dao
│      dao.go
│      goods.go
│      order.go
│      point.go
│      task.go
│      
├─md
│      user.md
│      
├─model
│      goods.go
│      order.go
│      point.go
│      task.go
│      token.go
│      user.go
│      
├─pbfile
│  ├─cert
│  │      ca.crt
│  │      ca.key
│  │      client.csr
│  │      client.key
│  │      client.pem
│  │      openssl.cnf
│  │      
│  ├─pb
│  │      user.pb.go
│  │      user_grpc.pb.go
│  │      
│  └─proto
│          user.proto
│          
├─service
│      goods.go
│      order.go
│      point.go
│      task.go
│      token.go
│      
└─tool	//工具类
check.go
parse.go
resp.go
trie.go




### 功能

#### 基础（有一说一

后台接口

- [x] 加积分
- [x] 商品CURD
- [x] 获取订单信息
- [x] 用户兑换商品（将订单标记为已完成）
- [x] 添加任务（分日常任务和只能完成一次的任务）

客户端接口

- [x] 完成任务获取商品信息
- [x] 返回用户相关信息（如有多少积分，签到排名，不同人的签到排名不同，连续签到天数等等
- [x] 签到
- [x] 买商品（将订单标记为已支付
- [x] 生成订单
- [x] 获取积分纪录（通过后端用户完成的任务接口获取积分增加记录，完成的订单接口获取积分消费记录）

### 进阶（大概不难

- [x] 具有秒杀商品功能，即一段时间内商品具有一个较大折扣，这是考虑大量流量处理等等
- （处理的比较简单。。。就先全用缓存来处理大流量的请求，同时协程通过管道慢慢异步处理数据库）

- [x] (部分...)规避掉并发带来的问题 (关键地方采用的加锁)

### 接口介绍

见md文件夹



### 其他

- rpc用户中心(证书、拦截器...没啥用)...字典树敏感词检查，简单过滤防sql注入...
- 为了方便测试，没加管理员验证