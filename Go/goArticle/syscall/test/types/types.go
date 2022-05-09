/*
 * File: types.go
 * Created Date: 2022-05-07 04:35:00
 * Author: ysj
 * Description:
 */

package types

const (
	IPC_CREATE = 00001000
	KEY        = 0x00000001
	SIZE       = 512
)

type Data struct {
	F1 string
	F2 int
	F3 []string
	F4 map[int]int
}
