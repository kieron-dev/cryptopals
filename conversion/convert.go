package conversion

import (
	"encoding/base64"
	"encoding/hex"
	"io/ioutil"
	"os"
)

func HexToBytes(in string) ([]byte, error) {
	return hex.DecodeString(in)
}

func BytesToHex(bytes []byte) string {
	return hex.EncodeToString(bytes)
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
	return ioutil.ReadAll(decoder)
}
