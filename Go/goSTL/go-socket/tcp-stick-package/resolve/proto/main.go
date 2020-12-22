/**
出现”粘包”的关键在于接收方不确定将要传输的数据包的大小，因此我们可以对数据包进行封包和拆包的操作。
封包就是给一段数据加上包头，这样一来数据包就分为包头和包体两部分了。
包头部分的长度是固定的，它存储了包体的长度
*/

package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

// Encode 发送消息进行编码封包
func Encode(msg string) ([]byte, error) {
	// 数据长度,4个字节
	length := int32(len(msg))
	// 数据缓冲
	pkg := new(bytes.Buffer)
	// 写入数据长度
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	// 写入数据主体
	err = binary.Write(pkg, binary.LittleEndian, []byte(msg))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

// Decode 接收消息进行解码拆包
func Decode(reader *bufio.Reader) (string, error) {
	// 读取数据长度,4个字节
	lengthByte, _ := reader.Peek(4)
	lengthBuf := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuf, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}

	if int32(reader.Buffered()) < (length + 4) {
		return "", err
	}

	//读取消息主体
	pack := make([]byte, int(4+length))
	n, err := reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:n]), nil
}
