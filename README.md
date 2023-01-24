# early_education 基于golang的早教系统

## 一个终端执行进程
ling.sun@FVFHP1SFQ05P early_education % go run main.go

## 另开一个终端，curl
ling.sun@192 ~ % curl "http://127.0.0.1:8080/course/list"
{"code":0,"data":[{"course_id":1,"mechanism":"金宝贝","class_time":"2023-01-22 09:30:00","class_name":"SEL1","teacher":{"name":"jasmine","age":25,"area":"北京"},"open_state":0},{"course_id":2,"mechanism":"金宝贝","class_time":"2023-01-22 11:00:00","class_name":"SEL2","teacher":{"name":"lucky","age":23,"area":"北京"},"open_state":0}],"msg":"success"}