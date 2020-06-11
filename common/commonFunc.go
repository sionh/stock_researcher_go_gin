package common

import (
	"fmt"
	"github.com/dustin/go-humanize"
	strconv "strconv"
)

// GetQuarterFromNum 期コード→文字列
func GetQuarterFromNum(quarter int) string {
	qstr := ""

	switch quarter {
	case 1:
		qstr = "春"
	case 2:
		qstr = "夏"
	case 3:
		qstr = "秋"
	case 4:
		qstr = "冬"
	}

	return qstr
}

// Comma 数字を桁区切りにして返します
func Comma(num int) string {
	return humanize.Comma(int64(num))
}

func RoundFloatStrIfParsable(floatStr string) string {
	floatNum, err := strconv.ParseFloat(floatStr, 64)
	if err == nil {
		return fmt.Sprintf("%.1f", floatNum)
	}
	return floatStr
}
