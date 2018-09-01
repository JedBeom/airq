package airq

import (
	"io/ioutil"
)

var (
	serviceKey string
)

// SetServiceKey 함수는 서비스 키를 직접 인자로 받습니다. 추천하지 않습니다.
func SetServiceKey(key string) {
	serviceKey = key
}

// LoadServiceKey 함수는 filename 파일을 열어서 그 안에 있는 서비스 키를 그대로 입력받습니다.
func LoadServiceKey(filename string) error {
	content, err := ioutil.ReadFile(filename)
	serviceKey = string(content[:len(content)-1])

	return err
}
