package main

// Go언어는 사용하지 않는 모듈이 있으면 오류가 나온다.
import (
	"fmt"
)

func main() {
	// 어떤 모듈의 export되는 function은 무조건 대문자로 시작한다
	fmt.Println("hi")

}
