package infra

import "github.com/circle/early_education/model"

type CourseDAO struct {
}

func (CourseDAO) GetCourses(pageNo, pageSize int) ([]model.Course, error) {
	// 从mysql取数据
	return nil, nil
}
