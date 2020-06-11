package common

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
