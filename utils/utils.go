package utils

import (
	"fmt"
	"time"
)

func ParseDateStr(str string) (time.Time, error) {
	loc, _ := time.LoadLocation("Asia/Tokyo")

	now := time.Now()
	year := now.Year()

	const format = "2006/1/2"
	dateStr := fmt.Sprintf("%d/%s", year, str)
	t, err := time.ParseInLocation(format, dateStr, loc)
	if err != nil {
		return time.Time{}, err
	}

	// 今年ではもう過ぎている日の場合、来年に設定
	if now.After(t) {
		year++

		dateStr := fmt.Sprintf("%d/%s", year, str)
		return time.ParseInLocation(format, dateStr, loc)
	}

	return t, err
}
