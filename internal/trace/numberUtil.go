package trace

import (
	"fmt"
	"strconv"
)

/**
 *
 * @author hushengdong
 */
func string2Int(num string) int {

	res, err := strconv.Atoi(num)
	if err != nil {
		panic(fmt.Sprintf("数字转换错误:%s", num))
	}
	return res
}

func string2Uint16(num string) uint16 {

	res, err := strconv.ParseInt(num, 16, 16)
	if err != nil {
		panic(fmt.Sprintf("数字转换错误:%s", num))
	}
	return uint16(res)
}
