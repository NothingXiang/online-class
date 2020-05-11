package class

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouters(e *gin.Engine) {
	class := e.Group("/class")
	{
		// 创建班级
		class.POST("/create", CreateClass)

		// 通过userID获取班级
		class.GET("/get", GetClassesByUser)

		class.GET("/get/id", GetClassById)

		// 添加学生
		//class.POST("/add/student", AddStudent)

		// 添加教师
		//class.POST("/add/teacher", AddTeacher)

		student := class.Group("/student")
		{
			// 获取学生列表
			student.GET("/lists", ListStudentPageable)

			student.POST("/add", AddStudent)
			// 获取学生
			//student.GET("/get", GetStudents)

		}

		teacher := class.Group("/teacher")
		{
			teacher.POST("/add", AddTeacher)

			// 获取教师列表
			teacher.GET("/lists", ListTeacher)

			// 获取教师信息
			teacher.GET("/get", GetTeacher)

			//	增加该名老师教授的科目
			teacher.POST("/subject", AddSubject)

			// 减少该名老师教授的科目
			teacher.DELETE("/subject", DeleteSubject)

			// 更新所教授的科目
			teacher.POST("/subject/update", UpdateSubject)

		}
	}

}
