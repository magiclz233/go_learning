package time_test

import (
	"fmt"
	"testing"
	"time"
)

// 时间类型
func TestDemo1(t *testing.T) {
	now := time.Now()
	fmt.Printf("当前时间： %v\n", now)
	var year int = now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

// 时间戳
/*
时间戳是自1970年1月1日（08:00:00GMT）至当前时间的总毫秒数。它也被称为Unix时间戳（UnixTimestamp）。
*/
func TestTimestamp(t *testing.T) {
	now := time.Now()
	timestamp1 := now.Unix()     //毫秒时间戳
	timestamp2 := now.UnixNano() // 纳秒时间戳
	fmt.Printf("当前时间时间戳(毫秒): %v\n", timestamp1)
	fmt.Printf("当前时间时间戳(纳秒): %v\n", timestamp2)
}

// 时间格式与时间戳的转换
func TestTimeAndTimestamp(t *testing.T) {
	timeAndTimestamp(time.Now().Unix())
}

func timeAndTimestamp(timestamp int64) {
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式 第二个参数为纳秒
	fmt.Println(timestamp, timeObj)
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

// 时间间隔 time.Duration
/*
time.Duration是time包定义的一个类型，它代表两个时间点之间经过的时间，以纳秒为单位。time.Duration表示一段时间间隔，可表示的最长时间段大约290年。
*/
func TestDuration(t *testing.T) {
	now := time.Now()
	fmt.Println(now)
	later := now.Add(time.Hour) // 当前时间
	fmt.Println(later)

	du := now.Sub(later) // 两个时间的差值
	fmt.Printf("%v\n", du)

	isEq := now.Equal(now) // 判断两个时间是否一致 时区有影响
	fmt.Printf("%v\n", isEq)

	isBe := now.Before(later) //now在later之前 true
	fmt.Printf("%v\n", isBe)

	isAf := now.After(later) //now 在 later之后 true
	fmt.Printf("%v\n", isAf)
}

// 定时器
func TestTicker(t *testing.T) {
	ticker := time.Tick(1e9)
	// 每隔一秒打印一次当前时间 定时器range返回的i为当前时间
	for i := range ticker {
		fmt.Println(i)
	}
}

// 时间格式化
/*
	时间类型有一个自带的方法Format进行格式化，需要注意的是Go语言中格式化时间模板不是常见的Y-m-d H:M:S
	而是使用Go的诞生时间2006年1月2号15点04分（记忆口诀为2006 1 2 3 4）。
*/
func TestTimeFormat(t *testing.T) {
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 03:04:05"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}

// 字符串转时间
func TestStrAndTime(t *testing.T) {
	now := time.Now()
	fmt.Println(now)

	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("err", err)
		return
	}

	timeStr := "2020-07-23 17:37:33"
	// 按照指定时区和指定格式解析字符串时间
	//注意: layout 里面 hour 标准格式 如果传入的是03,那么就是12小时制,传入15则为24小时制
	timeObj, err1 := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	if err1 != nil {
		fmt.Println("err1", err1)
		return
	}
	fmt.Println(timeObj)
	fmt.Printf("%T\n", timeObj)
	fmt.Println(timeObj.Sub(now))
}
