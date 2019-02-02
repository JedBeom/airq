package airq

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
)

var (
	serviceKey string
)

// SetKey 함수는 서비스 키를 직접 인자로 받습니다.
func SetKey(key string) (err error) {
	if key == "" {
		err = errors.New("Expected content, but no content from env")
		return
	}
	serviceKey = key
	return
}

// GetKeyEnv 함수는 환경변수에서 값을 가져옵니다.
func GetKeyEnv(env string) (err error) {
	key := os.Getenv(env)
	if key == "" {
		err = errors.New("Expected content, but no content from env")
		return
	}
	serviceKey = key
	return
}

// GetKeyFile 함수는 filename 파일을 열어서 그 안에 있는 서비스 키를 그대로 입력받습니다.
func GetKeyFile(filename string) error {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	if len(content) < 10 {
		err = errors.New("Service Key is too short")
		return err
	}
	serviceKey = string(content[:len(content)-1])
	return err
}
