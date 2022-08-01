
[生成订单]

##### 简要描述

- multipart/form-data
- 生成订单接口

##### 请求URL

- ` http://101.43.160.254：8080/mall/user/order`

##### 请求方式

- PUT

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |
|goods_id |是  |int |商品id   |


##### 返回示例

``` 

{
    "data": "成功",
    "status": true
}
```


[查看订单]

##### 简要描述

- multipart/form-data
- 查看订单接口

##### 请求URL

- ` http://101.43.160.254：8080/mall/user/order`

##### 请求方式

- GET

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |
|order_state |否  |bool |是否查看已完成的订单   |


##### 返回示例

``` 
{
    "data": [
        {
            "ID": 4,
            "CreatedAt": "2022-07-31T04:36:53.459Z",
            "UpdatedAt": "2022-08-01T08:38:27.124Z",
            "DeletedAt": null,
            "uid": 7,
            "goods_id": 2,
            "point": 3,
            "is_paid": true,
            "order_state": true
        }
    ],
    "status": true
}
```

[删除订单]

##### 简要描述

- multipart/form-data
- 删除订单接口

##### 请求URL

- ` http://101.43.160.254：8080/mall/user/order`

##### 请求方式

- DELETE

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |
|ID|是  |int |订单id   |


##### 返回示例

``` 
{
    "data": "成功",
    "status": true
}
```

[更新订单]

##### 简要描述

- multipart/form-data
- 更新订单接口

##### 请求URL

- ` http://101.43.160.254：8080/mall/user/order`

##### 请求方式

- POST

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |
|ID|是  |int |订单id   |
|is_paid |否 |bool |支付状态   |
|order_state |否 |bool |订单完成与否状态   |


##### 返回示例

``` 
{
    "data": "成功",
    "status": true
}
```

[生成秒杀订单]

##### 简要描述

- multipart/form-data
- 生成秒杀订单接口

##### 请求URL

- ` http://101.43.160.254：8080/mall/user/order/spike`

##### 请求方式

- PUT

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |
|goods_id|是  |int |订单商品id   |



##### 返回示例

``` 
{
    "data": "用户试图重复抢购",
    "status": false
}
```


[支付秒杀订单]

##### 简要描述

- multipart/form-data
- 支付秒杀订单接口

##### 请求URL

- ` http://101.43.160.254：8080/mall/user/order/spike`

##### 请求方式

- POST

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |
|goods_id|是  |int |订单商品id   |
|is_paid |是 |bool |支付状态   |


##### 返回示例

``` 
{
    "data": "成功",
    "status": true
}
```


[查看秒杀订单]

##### 简要描述

- multipart/form-data
- 查看秒杀订单接口，支付成功后请通过一般订单接口查看

##### 请求URL

- ` http://101.43.160.254：8080/mall/user/order/spike`

##### 请求方式

- GET

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |
|goods_id|是  |int |订单商品id   |



##### 返回示例

``` 
{
    "data": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "uid": 7,
        "goods_id": 7,
        "point": 0,
        "is_paid": false,
        "order_state": false
    },
    "status": true
}
```


[删除秒杀订单]

##### 简要描述

- multipart/form-data
- 删除秒杀订单接口

##### 请求URL

- ` http://101.43.160.254：8080/mall/user/order/spike`

##### 请求方式

- DELETE

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|token |是  |String |用户token   |
|goods_id|是  |int |订单商品id   |



##### 返回示例

``` 
{
    "data": "成功",
    "status": true
}
```