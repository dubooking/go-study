package main // 실행 가능한 프로그램은 항상 main 패키지로 시작

import "fmt" // 화면 출력용 표준 라이브러리 (Java의 System.out 역할)

func main() { // 프로그램 시작점 (Java의 public static void main)
	// 1) var 키워드로 선언: "var 이름 타입 = 값" (타입이 이름 뒤에 옴)
	var userName string = "eunsang"

	// 2) 타입 생략: 값을 보고 Go가 타입을 추론 (int로 추론됨)
	var port = 8080

	// 3) 짧은 선언 := (함수 안에서만 사용 가능, 가장 많이 씀)
	sessionTimeout := 30

	// 4) 값을 안 주면 "제로값"이 자동으로 들어감 (Java처럼 null이 아님)
	var count int       // 0
	var token string    // "" (빈 문자열)
	var isLoggedIn bool // false

	// 5) 상수: 한 번 정하면 못 바꿈 (Java의 static final)
	const issuer = "https://sso.example.com"

	// 6) 값 바꾸기: 이미 선언된 변수에는 = 사용 (:= 아님)
	sessionTimeout = 60

	// %v는 값을 그대로, %q는 따옴표를 붙여서, %T는 타입을 출력
	fmt.Printf("userName=%v (타입 %T)\n", userName, userName)
	fmt.Printf("port=%v (타입 %T)\n", port, port)
	fmt.Printf("sessionTimeout=%v\n", sessionTimeout)
	fmt.Printf("제로값 → count=%v, token=%q, isLoggedIn=%v\n", count, token, isLoggedIn)
	fmt.Println("issuer =", issuer)

	maxLoginFail := 5
	fmt.Printf("maxLoginFail=%d (타입 %T)\n", maxLoginFail, maxLoginFail)

}
