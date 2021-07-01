package Proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

//沾包解决
func Encode(msg string) ([]byte, error) {
	//获取消息体的长度 转换成int32
	var length = int32(len(msg))
	var pkg = new(bytes.Buffer)

	//写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	//写入消息体
	err = binary.Write(pkg, binary.LittleEndian, []byte(string(msg)))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

func Decode(reader *bufio.Reader) (string, error) {
	//读取消息的长度
	lengthByte, _ := reader.Peek(4)
	lengthBuff := bytes.NewBuffer(lengthByte)

	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}

	// buffered 返回缓冲中先有的刻度字数
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	//读取真正的字符串
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack[:])
	if err != nil {
		return "", err
	}

	return string(pack[4:]), err
}

func main() {

}
