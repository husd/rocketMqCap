package trace

import (
	"encoding/binary"
)

/**
 *
 * @author hushengdong
 */
func bytesToInt(buf []byte) int {
	return int(binary.BigEndian.Uint32(buf))
}

//这里的buf 一定是需要3个字节
func bytes3ToInt(b []byte) int {

	temp := make([]byte, 4, 4)
	temp[0] = 0
	temp[1] = b[0]
	temp[2] = b[1]
	temp[3] = b[2]
	return bytesToInt(temp)
}
