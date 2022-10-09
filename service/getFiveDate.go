package service

import (
	"fmt"
	"time"
)

func NumToHanzi(num int) string {
	switch num {
	case 1:
		return "一"
	case 2:
		return "二"
	case 3:
		return "三"
	case 4:
		return "四"
	case 5:
		return "五"
	case 6:
		return "六"
	case 7:
		return "日"
	}
	return ""
}

func GetFiveDate() [5]string {
	var date [5]string
	now := time.Now()
	date[0] = fmt.Sprintf("今天%d月%d日", now.Month(), now.Day())
	tomorrow := now.AddDate(0, 0, 1)
	date[1] = fmt.Sprintf("明天%d月%d日", tomorrow.Month(), tomorrow.Day())
	afterTomorrow := tomorrow.AddDate(0, 0, 1)
	date[2] = fmt.Sprintf("后天%d月%d日", afterTomorrow.Month(), afterTomorrow.Day())
	other1 := afterTomorrow.AddDate(0, 0, 1)
	date[3] = fmt.Sprintf("周%s%d月%d日", NumToHanzi(int(other1.Weekday())), other1.Month(), other1.Day())
	other2 := other1.AddDate(0, 0, 1)
	date[4] = fmt.Sprintf("周%s%d月%d日", NumToHanzi(int(other2.Weekday())), other2.Month(), other2.Day())
	return date
}
