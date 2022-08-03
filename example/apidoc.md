# API Document
This [API document](https://github.com/WGrape/apimock) can be automatically generated and quickly released in the CI stage without manual editing.

#### Version
Last update time at ```2022-08-03 20:35:43```
#### Content
- [1. 获取用户接口](#1-获取用户接口)

## 1. 获取用户接口

### (1) Description
获取用户接口

### (2) URL
/api/get_user

### (3) Request
| Field | Type | Required | Mark |
| :----: | :----: | :----: | :----: |
```json
{
    "uid": 0
}
```

### (4) Response
| Field | Type | Mark |
| :----: | :----: | :----: |
|data.uid|int|用户ID|
|data.name|string|用户昵称|
|data.age|int|用户年龄|
|data.gender|int|用户性别, 1女, 2男|
|data.city|string|用户所在城市|
```json
{
    "code": 0,
    "data": {
        "uid": 88328876,
        "name": "Peter",
        "age": 45,
        "gender": 1,
        "city": "New York"
    },
    "is_mock": false
}
```
