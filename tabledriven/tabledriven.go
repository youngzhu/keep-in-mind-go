package tabledriven

// 表驱动开发

var daysOfMonths = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

// DaysOfMonth 返回某个月份（month）的天数
// 方法并不完善，主要是演示表驱动开发方法
func DaysOfMonth(month int) int {
	return daysOfMonths[month-1]
}

/*
评分系统
>=90.0	A
<90.0	B
<75.0	C
<65.0	D
<50.0	F
*/

var (
	rangeUpper = []float32{50.0, 65.0, 75.0, 90.0, 100.0}
	grades     = []rune{'F', 'D', 'C', 'B', 'A'}
)

func GetGrade(score float32) (grade rune) {
	idx := 0
	grade = 'A'

	for grade == 'A' && idx < len(grades) {
		if score < rangeUpper[idx] {
			grade = grades[idx]
		}
		idx++
	}

	return
}
