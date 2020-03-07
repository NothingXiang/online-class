[toc]

# User

```go

type Users struct {

	// 用户唯一标识（感觉作用不大，用wx_code说不定更好）
	ID string

	// 用户名
	Name string 

	// 微信接口返回的用户唯一标识
	OpenId string 

	// 用户手机号
	Phone string 

	// 用户类型: 暂定为:1教师 2学生
	UserType int 

	// 用户头像路径
	Avatar string 

	// 创建时间
	CreateTime time.Time

	// 最后一次修改时间
	UpdateTime time.Time 
}

```





# Teacher

```go
type Teacher struct {
	// 唯一id
	ID int64

	// 教授科目
	Subject Subject
}
```



**注**：

## 科目对照表

```

	Chinese 	0
	Maths   	1 
	English		2
	Science		3
	Physics		4
	Chemistry	5
	Biology		6
	Politics	7
	Moral		8
	History		9
	Geography	10
	Natural		11
	Sports		12
	Technology	13
	Art			14
	Music		15
	Other		16
```



