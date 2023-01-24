package model

// AgeState 年龄段
type AgeState uint8

const (
	// Month0To6 0-6个月
	Month0To6 AgeState = 1
	// Month7To9 7-9个月
	Month7To9 AgeState = 2
	// Month10To12 10-12个月
	Month10To12 AgeState = 3
	// Year1To2 1-2岁
	Year1To2 AgeState = 4
	// Year2To3 2-3岁
	Year2To3 AgeState = 5
	// Year3To4 3-4岁
	Year3To4 AgeState = 6
	// Year4To5 4-5岁
	Year4To5 AgeState = 7
	// Year5To6 5-6岁
	Year5To6 AgeState = 8
)
