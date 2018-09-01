# airq
[![GoDoc](https://godoc.org/github.com/JedBeom/airq?status.svg)](https://godoc.org/github.com/JedBeom/airq)

```shell
$ go get -u github.com/JedBeom/airq
```

[공공 데이터 포털](https://www.data.go.kr/dataset/15000581/openapi.do)을 이용해 대기 오염 정보를 불러오는 개발 상태의 패키지입니다. 기능이 갑작스레 추가되거나 삭제될 수 있습니다.

현재 지원하는 기능은 *측정소별 실시간 측정정보 조회*이며, 공공 데이터 포털에서 서비스 키를 발급받아야 사용 가능합니다.

## Example

```go
package main

import (
    "fmt"
    "github.com/JedBeom/airq"
)

func main() {
    err := airq.LoadServiceKey("key.txt")
    if err != nil {
        panic(err)
    }

    quality, err := airq.GetAirqOfNowByStation("종로구")
    if err != nil {
        panic(err)
    }

    fmt.Println(quality)
}
```

key.txt
```
yourservicekeyhere
```

