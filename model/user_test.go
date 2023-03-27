package model

import (
	"testing"
	"time"
)

func TestUser_AgeStage(t *testing.T) {
	timeVal1, _ := time.ParseInLocation("2006-01-02 15:04:05", "2012-01-22 09:30:00", time.Local)
	timeVal2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-01-22 09:30:00", time.Local)
	type fields struct {
		UserID     uint64
		Name       string
		Password   string
		Nick       string
		Birthday   time.Time
		Phone      int
		CreateTime time.Time
		UpdateTime time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   AgeStage
	}{
		{
			name: "超过6岁",
			fields: fields{
				Birthday: timeVal1,
			},
			want: YearMoreThan6,
		},
		{
			name: "5-6岁",
			fields: fields{
				Birthday: timeVal2,
			},
			want: Year5To6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := User{
				UserID:     tt.fields.UserID,
				Name:       tt.fields.Name,
				Password:   tt.fields.Password,
				Nick:       tt.fields.Nick,
				Birthday:   tt.fields.Birthday,
				Phone:      tt.fields.Phone,
				CreateTime: tt.fields.CreateTime,
				UpdateTime: tt.fields.UpdateTime,
			}
			if got := w.AgeStage(); got != tt.want {
				t.Errorf("AgeStage() = %v, want %v", got, tt.want)
			}
		})
	}
}
