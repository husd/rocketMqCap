package trace

import (
	"fmt"
	"time"
)

/**
 *
 * @author hushengdong
 */
func now() string {
	now := time.Now()      //获取当前时间
	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d ", year, month, day, hour, minute, second)
}
