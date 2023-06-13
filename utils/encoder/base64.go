package encoder

import "encoding/base64"

func Base64Encode(content []byte) string {
	return base64.StdEncoding.EncodeToString(content)
}

func Base64Decode(content string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(content)
}
