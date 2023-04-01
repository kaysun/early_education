package infra

import "github.com/circle/early_education/model"

type CourseDAO struct {
}

// GetAllCourses 获取从现在开始，往后的开放中的课程
func (CourseDAO) GetAllCourses(user model.User) ([]model.Course, error) {
	return nil, nil
}
