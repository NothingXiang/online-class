[toc]

 登录/注册接口会返回userID和password, 小程序端保存下来，每次请求带上这两个参数



# 用户 /user



## 用户(手机号)注册 POST  /create

### 入参

| 名称      | 类型   | 解释             | 必选 |
| --------- | ------ | ---------------- | ---- |
| name      | string | 用户名           | 是   |
| phone     | string | 手机号           | 是   |
| user_type | int    | 0是学生，1是教师 | 是   |
| password  | string | 密码             | 是   |

示例

```json
{
    "name": "香👨老师",
    "user_type": 1,
    "password": "123456",
    "phone": "13112669929"
}
```



### 出参

| 名称        | 类型   | 解释                         |
| ----------- | ------ | ---------------------------- |
| id          | string | 用户唯一标识                 |
| name        | string | 用户名                       |
| open_id     | string | 用户微信id（现在可能用不上） |
| phone       | string | 手机号                       |
| Email       | string | 邮箱                         |
| user_type   | int    | 0是学生，1是教师             |
| avatar      | string | 头像                         |
| create_time | time   | 创建时间                     |
| update_time | time   | 更新时间                     |

示例

```json
{
    "code": 200,
    "msg": "success",
    "data": {
        "id": "471b8b4d-391e-4498-8ed4-83026403edc9",
        "name": "香👨老师",
        "open_id": "",
        "phone": "13112669929",
        "email": "",
        "password": "123456",
        "user_type": 1,
        "avatar": "",
        "create_time": "2020-03-18T16:21:57.0514736+08:00",
        "update_time": "0001-01-01T00:00:00Z"
    }
}
```



## 用户（手机号）登录 POST /login



## 上传头像 POST /avatar



## 获取头像 GET /avatar





# 班级 /class



## 创建班级 POST /create

### 入参

| 名称      | 类型   | 解释             | 必选 |
| --------- | ------ | ---------------- | ---- |
| name      | string | 班级名           | 是   |
| school    | string | 所属学校         | 是   |
| user_type | int    | 0是学生，1是教师 | 是   |
| password  | string | 密码             | 是   |

示例

```json
{
    "name": "香👨老师",
    "user_type": 1,
    "password": "123456",
    "phone": "13112669929"
}
```



### 出参

| 名称        | 类型   | 解释                         |
| ----------- | ------ | ---------------------------- |
| id          | string | 用户唯一标识                 |
| name        | string | 用户名                       |
| open_id     | string | 用户微信id（现在可能用不上） |
| phone       | string | 手机号                       |
| Email       | string | 邮箱                         |
| user_type   | int    | 0是学生，1是教师             |
| avatar      | string | 头像                         |
| create_time | time   | 创建时间                     |
| update_time | time   | 更新时间                     |

示例

```json
{
    "code": 200,
    "msg": "success",
    "data": {
        "id": "471b8b4d-391e-4498-8ed4-83026403edc9",
        "name": "香👨老师",
        "open_id": "",
        "phone": "13112669929",
        "email": "",
        "password": "123456",
        "user_type": 1,
        "avatar": "",
        "create_time": "2020-03-18T16:21:57.0514736+08:00",
        "update_time": "0001-01-01T00:00:00Z"
    }
}
```



## 获取该用户所在的所有班级 GET /get

### 入参

| 名称      | 类型   | 解释             | 必选 |
| --------- | ------ | ---------------- | ---- |
| name      | string | 用户名           | 是   |
| phone     | string | 手机号           | 是   |
| user_type | int    | 0是学生，1是教师 | 是   |
| password  | string | 密码             | 是   |

示例

```json
{
    "name": "香👨老师",
    "user_type": 1,
    "password": "123456",
    "phone": "13112669929"
}
```



### 出参

| 名称        | 类型   | 解释                         |
| ----------- | ------ | ---------------------------- |
| id          | string | 用户唯一标识                 |
| name        | string | 用户名                       |
| open_id     | string | 用户微信id（现在可能用不上） |
| phone       | string | 手机号                       |
| Email       | string | 邮箱                         |
| user_type   | int    | 0是学生，1是教师             |
| avatar      | string | 头像                         |
| create_time | time   | 创建时间                     |
| update_time | time   | 更新时间                     |

示例

```json
{
    "code": 200,
    "msg": "success",
    "data": {
        "id": "471b8b4d-391e-4498-8ed4-83026403edc9",
        "name": "香👨老师",
        "open_id": "",
        "phone": "13112669929",
        "email": "",
        "password": "123456",
        "user_type": 1,
        "avatar": "",
        "create_time": "2020-03-18T16:21:57.0514736+08:00",
        "update_time": "0001-01-01T00:00:00Z"
    }
}
```



## 加入学生 POST /add/student 

### 入参

| 名称      | 类型   | 解释             | 必选 |
| --------- | ------ | ---------------- | ---- |
| name      | string | 用户名           | 是   |
| phone     | string | 手机号           | 是   |
| user_type | int    | 0是学生，1是教师 | 是   |
| password  | string | 密码             | 是   |

示例

```json
{
    "name": "香👨老师",
    "user_type": 1,
    "password": "123456",
    "phone": "13112669929"
}
```



### 出参

| 名称        | 类型   | 解释                         |
| ----------- | ------ | ---------------------------- |
| id          | string | 用户唯一标识                 |
| name        | string | 用户名                       |
| open_id     | string | 用户微信id（现在可能用不上） |
| phone       | string | 手机号                       |
| Email       | string | 邮箱                         |
| user_type   | int    | 0是学生，1是教师             |
| avatar      | string | 头像                         |
| create_time | time   | 创建时间                     |
| update_time | time   | 更新时间                     |

示例

```json
{
    "code": 200,
    "msg": "success",
    "data": {
        "id": "471b8b4d-391e-4498-8ed4-83026403edc9",
        "name": "香👨老师",
        "open_id": "",
        "phone": "13112669929",
        "email": "",
        "password": "123456",
        "user_type": 1,
        "avatar": "",
        "create_time": "2020-03-18T16:21:57.0514736+08:00",
        "update_time": "0001-01-01T00:00:00Z"
    }
}
```



## 加入教师 POST /add/teacher

### 入参

| 名称      | 类型   | 解释             | 必选 |
| --------- | ------ | ---------------- | ---- |
| name      | string | 用户名           | 是   |
| phone     | string | 手机号           | 是   |
| user_type | int    | 0是学生，1是教师 | 是   |
| password  | string | 密码             | 是   |

示例

```json
{
    "name": "香👨老师",
    "user_type": 1,
    "password": "123456",
    "phone": "13112669929"
}
```



### 出参

| 名称        | 类型   | 解释                         |
| ----------- | ------ | ---------------------------- |
| id          | string | 用户唯一标识                 |
| name        | string | 用户名                       |
| open_id     | string | 用户微信id（现在可能用不上） |
| phone       | string | 手机号                       |
| Email       | string | 邮箱                         |
| user_type   | int    | 0是学生，1是教师             |
| avatar      | string | 头像                         |
| create_time | time   | 创建时间                     |
| update_time | time   | 更新时间                     |

示例

```json
{
    "code": 200,
    "msg": "success",
    "data": {
        "id": "471b8b4d-391e-4498-8ed4-83026403edc9",
        "name": "香👨老师",
        "open_id": "",
        "phone": "13112669929",
        "email": "",
        "password": "123456",
        "user_type": 1,
        "avatar": "",
        "create_time": "2020-03-18T16:21:57.0514736+08:00",
        "update_time": "0001-01-01T00:00:00Z"
    }
}
```



## 获取学生列表 GET  /student/lists

### 入参

| 名称      | 类型   | 解释             | 必选 |
| --------- | ------ | ---------------- | ---- |
| name      | string | 用户名           | 是   |
| phone     | string | 手机号           | 是   |
| user_type | int    | 0是学生，1是教师 | 是   |
| password  | string | 密码             | 是   |

示例

```json
{
    "name": "香👨老师",
    "user_type": 1,
    "password": "123456",
    "phone": "13112669929"
}
```



### 出参

| 名称        | 类型   | 解释                         |
| ----------- | ------ | ---------------------------- |
| id          | string | 用户唯一标识                 |
| name        | string | 用户名                       |
| open_id     | string | 用户微信id（现在可能用不上） |
| phone       | string | 手机号                       |
| Email       | string | 邮箱                         |
| user_type   | int    | 0是学生，1是教师             |
| avatar      | string | 头像                         |
| create_time | time   | 创建时间                     |
| update_time | time   | 更新时间                     |

示例

```json
{
    "code": 200,
    "msg": "success",
    "data": {
        "id": "471b8b4d-391e-4498-8ed4-83026403edc9",
        "name": "香👨老师",
        "open_id": "",
        "phone": "13112669929",
        "email": "",
        "password": "123456",
        "user_type": 1,
        "avatar": "",
        "create_time": "2020-03-18T16:21:57.0514736+08:00",
        "update_time": "0001-01-01T00:00:00Z"
    }
}
```



## 教师列表 GET /teacher/list

### 入参

| 名称      | 类型   | 解释             | 必选 |
| --------- | ------ | ---------------- | ---- |
| name      | string | 用户名           | 是   |
| phone     | string | 手机号           | 是   |
| user_type | int    | 0是学生，1是教师 | 是   |
| password  | string | 密码             | 是   |

示例

```json
{
    "name": "香👨老师",
    "user_type": 1,
    "password": "123456",
    "phone": "13112669929"
}
```



### 出参

| 名称        | 类型   | 解释                         |
| ----------- | ------ | ---------------------------- |
| id          | string | 用户唯一标识                 |
| name        | string | 用户名                       |
| open_id     | string | 用户微信id（现在可能用不上） |
| phone       | string | 手机号                       |
| Email       | string | 邮箱                         |
| user_type   | int    | 0是学生，1是教师             |
| avatar      | string | 头像                         |
| create_time | time   | 创建时间                     |
| update_time | time   | 更新时间                     |

示例

```json
{
    "code": 200,
    "msg": "success",
    "data": {
        "id": "471b8b4d-391e-4498-8ed4-83026403edc9",
        "name": "香👨老师",
        "open_id": "",
        "phone": "13112669929",
        "email": "",
        "password": "123456",
        "user_type": 1,
        "avatar": "",
        "create_time": "2020-03-18T16:21:57.0514736+08:00",
        "update_time": "0001-01-01T00:00:00Z"
    }
}
```



## 增加某班级某教师教授的科目 POST /teacher/subject

### 入参

| 名称      | 类型   | 解释             | 必选 |
| --------- | ------ | ---------------- | ---- |
| name      | string | 用户名           | 是   |
| phone     | string | 手机号           | 是   |
| user_type | int    | 0是学生，1是教师 | 是   |
| password  | string | 密码             | 是   |

示例

```json
{
    "name": "香👨老师",
    "user_type": 1,
    "password": "123456",
    "phone": "13112669929"
}
```



### 出参

| 名称        | 类型   | 解释                         |
| ----------- | ------ | ---------------------------- |
| id          | string | 用户唯一标识                 |
| name        | string | 用户名                       |
| open_id     | string | 用户微信id（现在可能用不上） |
| phone       | string | 手机号                       |
| Email       | string | 邮箱                         |
| user_type   | int    | 0是学生，1是教师             |
| avatar      | string | 头像                         |
| create_time | time   | 创建时间                     |
| update_time | time   | 更新时间                     |

示例

```json
{
    "code": 200,
    "msg": "success",
    "data": {
        "id": "471b8b4d-391e-4498-8ed4-83026403edc9",
        "name": "香👨老师",
        "open_id": "",
        "phone": "13112669929",
        "email": "",
        "password": "123456",
        "user_type": 1,
        "avatar": "",
        "create_time": "2020-03-18T16:21:57.0514736+08:00",
        "update_time": "0001-01-01T00:00:00Z"
    }
}
```



## 减少某教师教授的科目 DELETE /teacher/subject





