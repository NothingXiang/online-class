package class

import (
	"github.com/NothingXiang/online-class/middle"
	"github.com/gin-gonic/gin"
)

func RegisterRouters(e *gin.Engine) {
	class := e.Group("/class")
	{
		// 创建班级
		class.POST("/create", middle.ValidateTeacher, CreateClass)

		// 通过userID获取班级
		class.GET("/get", GetClassesByUser)

		// 添加学生
		class.POST("/add/student", AddStudent)

		// 添加教师
		class.POST("/add/teacher", AddTeacher)

		student := class.Group("/student")
		{
			// 获取学生列表
			student.GET("/lists", ListStudentPageable)

			// 获取学生
			//student.GET("/get", GetStudents)

		}

		teacher := class.Group("/teacher")
		{
			// 获取教师列表
			teacher.GET("/lists", ListTeacher)

			// 获取教师信息
			//teacher.GET("/get", GetTeacher)

			//	增加该名老师教授的科目
			teacher.POST("/subject", AddSubject)

			// 减少该名老师教授的科目
			teacher.DELETE("/subject", DeleteSubject)

		}
	}

}
