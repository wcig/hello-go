package test

import (
	"fmt"
	"log"
	"testing"
	"time"
)

// 1.获取当时时间 + 年月日时分秒 + 时间戳 (1s=10^3ms(毫秒)=10^6μs(微秒)=10^9ns(纳秒))
func TestTime1(t *testing.T) {
	now := time.Now()
	fmt.Printf("current time:%v\n", now)

	year := now.Year()             // 年
	month := now.Month()           // 月
	day := now.Day()               // 日
	hour := now.Hour()             // 小时
	minute := now.Minute()         // 分钟
	second := now.Second()         // 秒
	nanosecond := now.Nanosecond() // 纳秒
	fmt.Printf("%04d-%02d-%02d %02d:%02d:%02d:%09d\n", year, month, day, hour, minute, second, nanosecond)

	dayOfWeek := now.Weekday().String()
	fmt.Printf("day of week:%s\n", dayOfWeek)

	timestamp1 := now.Unix()           // 秒时间戳
	timestamp2 := now.UnixNano()       // 纳秒时间戳
	timestamp3 := now.UnixNano() / 1e6 // 毫秒时间戳
	fmt.Printf("现在的时间戳：%v\n", timestamp1)
	fmt.Printf("现在的纳秒时间戳：%v\n", timestamp2)
	fmt.Printf("现在的毫秒时间戳：%v\n", timestamp3)
}

// current time:2020-01-15 17:47:00.835769665 +0800 CST m=+0.000262435
// 2020-01-15 17:47:00:835769665
// day of week:Wednesday
// 现在的时间戳：1579081620
// 现在的纳秒时间戳：1579081620835769665
// 现在的毫秒时间戳：1579081620835

// 2.时间戳 -> 时间
func TestTime2(t *testing.T) {
	ts1 := time.Now().Unix()
	ts2 := time.Now().UnixNano()
	ts3 := time.Now().UnixNano() / 1e6

	timeObj1 := time.Unix(ts1, 0)
	fmt.Printf("time:%v\n", timeObj1)

	timeObj2 := time.Unix(0, ts2)
	fmt.Printf("time:%v\n", timeObj2)

	timeObj3 := time.Unix(0, ts3*1e6)
	fmt.Printf("time:%v\n", timeObj3)
}

// time:2020-01-15 17:46:49 +0800 CST
// time:2020-01-15 17:46:49.577604025 +0800 CST
// time:2020-01-15 17:46:49.577 +0800 CST

// 3.字符串格式时间 -> 时间
func TestTime3(t *testing.T) {
	dateTimeStr := "2020-01-15 17:35:37.561"
	timeObj1, err := time.Parse("2006-01-02 15:04:05.000", dateTimeStr)
	if err != nil {
		log.Fatal("parse datetime str err:", err)
	}
	fmt.Printf("time:%v\n", timeObj1)

	timeObj2, _ := time.ParseInLocation("2006-01-02 15:04:05.000", dateTimeStr, time.Local)
	fmt.Printf("time:%v\n", timeObj2)
	fmt.Println(timeObj2.Unix())

	timeObj3, _ := time.ParseInLocation("2006-01-02 15:04:05.000", dateTimeStr, time.UTC)
	fmt.Printf("time:%v\n", timeObj3)
	fmt.Println(timeObj3.Unix())
}

// time:2020-01-15 17:35:37.561 +0000 UTC
// time:2020-01-15 17:35:37.561 +0800 CST
// 1579080937
// time:2020-01-15 17:35:37.561 +0000 UTC
// 1579109737

// 4.时间戳 -> 字符串格式时间
func TestTime4(t *testing.T) {
	var ts int64 = 1579080937
	timeObj := time.Unix(ts, 0)

	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(timeObj.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(timeObj.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(timeObj.Format("2006/01/02 15:04"))
	fmt.Println(timeObj.Format("15:04 2006/01/02"))
	fmt.Println(timeObj.Format("2006/01/02"))

	timeObj2 := time.Unix(0, ts*1e9)
	fmt.Println(timeObj2.Format("2006/01/02 15:04:05.000"))
}

// 2020-01-15 17:35:37.000 Wed Jan
// 2020-01-15 05:35:37.000 PM Wed Jan
// 2020/01/15 17:35
// 17:35 2020/01/15
// 2020/01/15
// 2020/01/15 17:35:37.000

// 5.时间计算
func TestTime5(t *testing.T) {
	now := time.Now()
	fmt.Println(now, "==>> now")

	addOneHour := now.Add(time.Hour)
	fmt.Println(addOneHour, "==>> after one hour")

	addDateTimeObj := now.AddDate(1, 1, 1)
	fmt.Println(addDateTimeObj, "==>> after one year one month one day")

	timeDuration := now.Sub(addOneHour)
	fmt.Println("time duration:", int64(timeDuration.Seconds()))

	isEqual := now.Equal(addOneHour)
	fmt.Println("is equal:", isEqual)

	isBefore := now.Before(addOneHour)
	fmt.Println("is before:", isBefore)

	isAfter := now.After(addOneHour)
	fmt.Println("is after:", isAfter)
}

// 2020-01-15 17:59:28.065065837 +0800 CST m=+0.000216584 ==>> now
// 2020-01-15 18:59:28.065065837 +0800 CST m=+3600.000216584 ==>> after one hour
// 2021-02-16 17:59:28.065065837 +0800 CST ==>> after one year one month one day
// time duration: -3600
// is equal: false
// is before: true
// is after: false

// 6.时区
func TestTime6(t *testing.T) {
	now := time.Now()
	local1, err1 := time.LoadLocation("") // UTC
	if err1 != nil {
		fmt.Println(err1)
	}
	local2, err2 := time.LoadLocation("Local") // 服务器设置时区
	if err2 != nil {
		fmt.Println(err2)
	}
	local3, err3 := time.LoadLocation("America/Los_Angeles") // 美国洛杉矶时区
	// local3, err3 := time.LoadLocation("Asia/Shanghai")
	if err3 != nil {
		fmt.Println(err3)
	}

	fmt.Println(now.In(local1))
	fmt.Println(now.In(local2))
	fmt.Println(now.In(local3))

	fmt.Println(now.In(local1).Unix())
	fmt.Println(now.In(local2).Unix())
	fmt.Println(now.In(local3).Unix())
}

// 2020-01-15 10:04:30.617381925 +0000 UTC
// 2020-01-15 18:04:30.617381925 +0800 CST
// 2020-01-15 02:04:30.617381925 -0800 PST

// 7.定时器
func TestTime7(t *testing.T) {
	ticker := time.Tick(time.Second) // 定义一个1秒间隔的定时器
	for i := range ticker {
		fmt.Println(i) // 每秒都会执行的任务
	}
}
