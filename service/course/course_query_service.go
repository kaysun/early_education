package course

import (
	"github.com/circle/early_education/infra"
	"github.com/circle/early_education/model"
)

// CourseQueryService 课程查询服务
type CourseQueryService struct {
	// courseDAO 课程访问层
	courseDAO infra.CourseDAO
}

// GetAllCourses 获取从现在开始，往后的开放中的课程
func (c CourseQueryService) GetAllCourses(user model.User) ([]model.Course, error) {
	return c.courseDAO.GetAllCourses(user)
}
