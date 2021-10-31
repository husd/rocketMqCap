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
