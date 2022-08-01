##### 返回参数统一说明

##### 一般参数
|status|data|说明|
|:-----  |:-----|-----    |
|false |"参数错误"   |参数错误  |
|false |"含有非法词汇"   |参数含有敏感词  |
|false |"服务器错误"   |服务端错误  |
|true |"成功！"   |参数合法  |
|true |"data{}"   |参数合法,返回数据  |

##### 特殊参数

|status|data|
|:-----  |:-----|
|false |"商品库存不足"|
|false |"用户余额不足"|
|false |"用户试图重复签到"|
|false |"用户试图重复抢购"|
|false |"用户试图重复支付？？？"|
|false |"sorry啦，您已经确定支付咯"|

[用户注册]

##### 简要描述

- multipart/form-data
- 用户注册接口

##### 请求URL

- ` http://101.43.160.254：8080/mall/user`

##### 请求方式

- POST

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|username |是  |string |用户名   |
|password |是  |string | 密码    |
|name     |否  |string | 昵称    |

##### 返回示例

``` 
  {
    "status": true,
    "data": "注册成功!"
    "username": "12154545",
    "name": "清梦",
    "groupId": 0
    
  }
```

[用户登陆]

##### 简要描述

- multipart/form-data
- 用户登陆接口，检查用户名及密码

##### 请求URL

- ` http://101.43.160.254：8080/mall/user`

##### 请求方式

- GET

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|username |是  |string |用户名   |
|password |是  |string | 密码    |

##### 返回示例

``` 
  {
    "data": {
        "id": 7,
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjcsIlVzZXJOYW1lIjoiIiwiU3RhdGUiOjAsIkdyb3VwSWQiOjAsIlRpbWUiOiIwMDAxLTAxLTAxVDAwOjAwOjAwWiIsImV4cCI6MTY1OTM3NDE3Mn0.G1l_KyWfp6WE9e09RQYmMMOvRZ7C4NkUdJTe5Qk-BY0",
        "username": "小明11",
        "password": "",
        "gender": 0,
        "name": "",
        "phone": "",
        "email": "",
        "state": 0,     //是否封号
        "group_id": 0   //用户组，1.商家，0.普通用户
    },
    "status": true
}
```


[查看用户个人信息]

##### 简要描述

- multipart/form-data
- 用户个人信息接口

##### 请求URL

- ` http://101.43.160.254：8080/mall/user/info`

##### 请求方式

- GET

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |

##### 返回示例

``` 

{
    "data": {
        "ID": 7,
        "Username": "小明11",
        "Name": "null"
    },
    "status": true
}

```

##### 返回参数说明

|status|data|说明|
|:-----  |:-----|-----                           |
|false |"token无效"   | |
|false |"旧密码错误"   | |
|false |"新密码无效"   | |
|true |"修改成功！"   |参数合法  |


[用户个人信息修改]

##### 简要描述

- multipart/form-data
- 用户个人信息修改接口

##### 请求URL

- ` http://101.43.160.254：8080/mall/user/info`

##### 请求方式

- POST

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |
|username |否 |String |用户名修改   |
|name |否 |String |用户昵称修改   |
|gender |否 |String |用户性别修改 (0为男，1为女)  |
|phone |否  |String |用户电话修改   |

##### 返回示例

``` 

{
    "data": {
        "ID": 7,
        "Username": "小明111",
        "Name": "null"
    },
    "status": true
}

```



[登录后修改密码]

##### 简要描述

- multipart/form-data
- 登录后修改密码接口

##### 请求URL

- ` http://101.43.160.254：8080/mall/user/password`

##### 请求方式

- POST

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |
|username |否 |String |用户名修改   |
|name |否 |String |用户昵称修改   |


##### 返回示例

``` 

{
    "data": "成功",
    "status": true
}

```

[通过用户id查看任务]

##### 简要描述

- multipart/form-data
- 通过用户id查看任务接口

##### 请求URL

- ` http://101.43.160.254：8080/mall/user/task`

##### 请求方式

- GET

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |
|state |否 |String |筛选是否完成任务   |


##### 返回示例

``` 

{
    "data": [
        {
            "ID": 1,
            "CreatedAt": "2022-07-31T03:23:02.898Z",
            "UpdatedAt": "2022-07-31T03:25:26.559Z",
            "DeletedAt": null,
            "uid": 7,
            "task_name": "run",
            "intro": "null",
            "point": 5,
            "state": true,
            "type": 0
        }
    ],
    "status": true
}

```

[完成任务]

##### 简要描述

- multipart/form-data
- 完成任务接口

##### 请求URL

- ` http://101.43.160.254：8080/mall/user/task`

##### 请求方式

- POST

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |
|state |是 |String |确定是否完成任务   |


##### 返回示例

``` 

{
    "data": "完成失败",
    "status": false
}
```

[用户签到]

##### 简要描述

- multipart/form-data
- 用户签到接口

##### 请求URL

- ` http://101.43.160.254：8080/mall/user/sign`

##### 请求方式

- POST

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |


##### 返回示例

``` 

{
    "data": 1,  //连续签到的天数
    "status": true
}

{
    "data": "用户试图重复签到",
    "status": false
}
```

[获取用户连续签到的天数]

##### 简要描述

- multipart/form-data
- 获取用户连续签到的天数接口

##### 请求URL

- ` http://101.43.160.254：8080/mall/user/sign`

##### 请求方式

- GET

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |


##### 返回示例

``` 

{
    "data": 1,  //连续签到的天数
    "status": true
}

```

[获取用户签到排名]

##### 简要描述

- multipart/form-data
- 获取用户签到排名接口

##### 请求URL

- ` http://101.43.160.254：8080/mall/user/sign/rank`

##### 请求方式

- GET

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |


##### 返回示例

``` 

{
    "data": 2, 
    "status": true
}

```

[查看用户积分]

##### 简要描述

- multipart/form-data
- 查看用户积分接口

##### 请求URL

- ` http://101.43.160.254：8080/mall/point`

##### 请求方式

- GET

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uid |是  |int |用户id |


##### 返回示例

``` 

{
    "data": {
        "ID": 1,
        "CreatedAt": "2022-07-30T09:03:46.375Z",
        "UpdatedAt": "2022-08-01T06:15:23.505Z",
        "DeletedAt": null,
        "uid": 7,
        "number": 106
    },
    "status": true
}

```
##### 以下三个接口其实是管理员功能，方便测试没做管理员验证

[给用户增加积分]

##### 简要描述

- multipart/form-data
- 给用户增加积分接口

##### 请求URL

- ` http://101.43.160.254：8080/mall/point`

##### 请求方式

- POST

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uid |是  |int |用户id |
|number |是  |int |增加积分数 |

##### 返回示例

``` 

{
    "data": 217,
    "status": true
}
```


[给用户增加任务]

##### 简要描述

- multipart/form-data
- 给用户增加任务接口

##### 请求URL

- ` http://101.43.160.254：8080/mall/task`

##### 请求方式

- PUT

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|uid |是  |int |用户id |
|task_name |是  |string |任务名 |
|point |是  |int |任务积分数 |
|intro |否  |string |任务介绍 |
|type |否  |int |任务类别（0为日常任务。1为限定一次的任务,待加） |

##### 返回示例

``` 
{
    "data": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "uid": 7,
        "task_name": "起飞",
        "intro": "",
        "point": 5,
        "state": false,
        "type": 0
    },
    "status": true
}
```

[通过任务id查看任务]

##### 简要描述

- multipart/form-data
- 通过任务id查看任务接口

##### 请求URL

- ` http://101.43.160.254：8080/mall/task`

##### 请求方式

- GET

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|ID |是  |int |任务ID |


##### 返回示例

``` 
{
    "data": {
        "ID": 4,
        "CreatedAt": "2022-08-01T07:50:06.511Z",
        "UpdatedAt": "2022-08-01T07:50:06.511Z",
        "DeletedAt": null,
        "uid": 7,
        "task_name": "起飞",
        "intro": "null",
        "point": 5,
        "state": false,
        "type": 0
    },
    "status": true
}
```
