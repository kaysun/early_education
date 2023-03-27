package model

// AgeStage 年龄段
type AgeStage uint8

const (
	// Month0To6 0-6个月
	Month0To6 AgeStage = 1
	// Month7To9 7-9个月
	Month7To9 AgeStage = 2
	// Month10To12 10-12个月
	Month10To12 AgeStage = 3
	// Year1To2 1-2岁
	Year1To2 AgeStage = 4
	// Year2To3 2-3岁
	Year2To3 AgeStage = 5
	// Year3To4 3-4岁
	Year3To4 AgeStage = 6
	// Year4To5 4-5岁
	Year4To5 AgeStage = 7
	// Year5To6 5-6岁
	Year5To6 AgeStage = 8
	// YearMoreThan6 超过6岁
	YearMoreThan6 AgeStage = 9
)
