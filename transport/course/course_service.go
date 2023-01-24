package course

import (
	"github.com/circle/early_education/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// List 课程列表
func List(context *gin.Context) {
	//timeVal1, _ := time.ParseInLocation("2006-01-02 15:04:05", "2023-01-22 09:30:00", time.Local)
	//timeVal2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2023-01-22 11:00:00", time.Local)

	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": []model.Course{
			{
				CourseID:  1,
				Mechanism: "金宝贝",
				ClassTime: "2023-01-22 09:30:00",
				ClassName: "SEL1",
				Teacher: model.Teacher{
					Name:      "jasmine",
					Age:       25,
					Area:      "北京",
					CourseIDs: []uint32{1},
				},
				OpenState: 0,
			},
			{
				CourseID:  2,
				Mechanism: "金宝贝",
				ClassTime: "2023-01-22 11:00:00",
				ClassName: "SEL2",
				Teacher: model.Teacher{
					Name:      "lucky",
					Age:       23,
					Area:      "北京",
					CourseIDs: []uint32{2},
				},
				OpenState: 0,
			},
		},
	})
}
