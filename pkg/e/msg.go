/*
 * @Author: Jecosine
 * @Date: 2021-01-02 06:03:50
 * @LastEditTime: 2021-01-02 06:56:51
 * @LastEditors: Jecosine
 * @Description: Error code message
 */

package e

// MsgDict Dictionary of message
var MsgDict = map[int]string{
	SUCCESS: "OK",
	ERROR:   "Internal Error",
}

// GetMsg get message by error code
func GetMsg(code int) string {
	msg, flag := MsgDict[code]
	if flag {
		return msg
	}
	return MsgDict[ERROR]
}
