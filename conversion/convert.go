package conversion

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strconv"
)

func HexToBytes(hex string) ([]byte, error) {
	ret := []byte{}

	if len(hex)%2 == 1 {
		return ret, fmt.Errorf("malformed hex string %s", hex)
	}

	for len(hex) > 0 {
		str := hex[0:2]
		hex = hex[2:]
		num, err := strconv.ParseInt(str, 16, 0)
		if err != nil {
			return []byte{}, err
		}
		ret = append(ret, byte(num))
	}

	return ret, nil
}

func BytesToHex(bytes []byte) string {
	ret := ""
	for _, b := range bytes {
		hex := fmt.Sprintf("%02x", b)
		ret += hex
	}
	return ret
}

func BytesToBase64(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}

func Base64ToBytes(b64 string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(b64)
}

func ReadBase64File(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return []byte{}, err
	}
	defer file.Close()

	decoder := base64.NewDecoder(base64.StdEncoding, file)

	bytes := []byte{}
	buf := make([]byte, 256)
	for {
		n, err := decoder.Read(buf)
		if err != io.EOF {
			bytes = append(bytes, buf[0:n]...)
		} else {
			break
		}
	}

	return bytes, nil
}
