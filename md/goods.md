[查看或搜索商品]

##### 简要描述

- multipart/form-data
- 查看或搜索商品接口

##### 请求URL

- ` http://101.43.160.254：8080/mall/goods`

##### 请求方式

- GET

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|ID |否  |int |商品ID |
|goods_name |否  |string |商品名模糊查找|


##### 返回示例

``` 
{
    "data": [
        {
            "ID": 3,
            "CreatedAt": "2022-08-01T05:40:07.248Z",
            "UpdatedAt": "2022-08-01T05:40:07.248Z",
            "DeletedAt": null,
            "goods_name": "飞机",
            "picture": "null",
            "price": 5,
            "goods_intro": "null",
            "number": 111
        },
        {
            "ID": 4,
            "CreatedAt": "2022-08-01T05:54:10.013Z",
            "UpdatedAt": "2022-08-01T05:54:10.013Z",
            "DeletedAt": null,
            "goods_name": "飞机1",
            "picture": "null",
            "price": 5,
            "goods_intro": "null",
            "number": 111
        },
        {
            "ID": 5,
            "CreatedAt": "2022-08-01T06:01:35.893Z",
            "UpdatedAt": "2022-08-01T06:15:23.503Z",
            "DeletedAt": null,
            "goods_name": "飞机11",
            "picture": "null",
            "price": 5,
            "goods_intro": "null",
            "number": 110
        }
    ],
    "status": true
}
```

[增加商品]

##### 简要描述

- multipart/form-data
- 增加商品接口

##### 请求URL

- ` http://101.43.160.254：8080/mall/goods`

##### 请求方式

- PUT

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|goods_name |是  |string |商品名 |
|price |是  |int |商品价格 |
|number |是  |int |商品数量 |
|goods_intro |否  |string |商品介绍 |
|picture |否  |int |商品图片 |


##### 返回示例

``` 
{
    "data": "成功",
    "status": true
}
```

[更新商品]

##### 简要描述

- multipart/form-data
- 更新商品接口

##### 请求URL

- ` http://101.43.160.254：8080/mall/goods`

##### 请求方式

- POST

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|goods_name |否  |string |商品名 |
|price |否  |int |商品价格 |
|number |否  |int |商品数量 |
|goods_intro |否  |string |商品介绍 |
|picture |否  |int |商品图片 |


##### 返回示例

``` 
{
    "data": {
        "ID": 4,
        "CreatedAt": "2022-08-01T05:54:10.013Z",
        "UpdatedAt": "2022-08-01T08:03:32.808Z",
        "DeletedAt": null,
        "goods_name": "飞机aaa",
        "picture": "null",
        "price": 5,
        "goods_intro": "null",
        "number": 111
    },
    "status": true
}
```

[删除商品]

##### 简要描述

- multipart/form-data
- 删除商品接口

##### 请求URL

- ` http://101.43.160.254：8080/mall/goods`

##### 请求方式

- DELETE

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|ID |是  |int |商品id |


##### 返回示例

``` 
{
    "data": "成功",
    "status": true
}
```


[添加秒杀商品]

##### 简要描述

- multipart/form-data
- 添加秒杀商品接口

##### 请求URL

- ` http://101.43.160.254：8080/mall/goods/spike`

##### 请求方式

- PUT

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|goods_name |是  |string |商品名 |
|price |是  |int |商品价格 |
|number |是  |int |商品数量 |
|goods_intro |否  |string |商品介绍 |
|picture |否  |int |商品图片 |


##### 返回示例

``` 
{
    "data": 7,  //商品id
    "status": true
}
```

[查看秒杀商品数量]

##### 简要描述

- multipart/form-data
- 查看秒杀商品数量接口

##### 请求URL

- ` http://101.43.160.254：8080/mall/goods/spike`

##### 请求方式

- GET

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|ID |是  |int |商品id |



##### 返回示例

``` 
{
    "data": 111,    //还未抢购的秒杀商品数量
    "status": true
}
```