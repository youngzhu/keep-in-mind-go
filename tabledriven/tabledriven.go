package tabledriven

// 表驱动开发

var daysOfMonths = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

// DaysOfMonth 返回某个月份（month）的天数
// 方法并不完善，主要是演示表驱动开发方法
func DaysOfMonth(month int) int {
	return daysOfMonths[month-1]
}
