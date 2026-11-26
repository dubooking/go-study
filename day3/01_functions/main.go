package main // 실행 가능한 프로그램은 main 패키지

import "fmt" // 화면 출력용 표준 라이브러리

// 예제: 두 수를 받아 더한 값을 돌려주는 함수
// Java로 쓰면: public static int add(int a, int b) { return a + b; }
func add(a int, b int) int { // func 이름(매개변수 타입) 반환타입
	return a + b // 계산 결과를 호출한 쪽으로 돌려줌
}

// 연습: 아래 ___ 를 알맞은 반환 타입으로 바꾸기
func remainTries(maxFail int, failed int) int { // 남은 로그인 시도 횟수를 돌려줌
	return maxFail - failed // 최대 허용 횟수 - 실패한 횟수
}

func main() { // 프로그램 시작점
	fmt.Println("add(3, 4) =", add(3, 4))     // 7이 출력됨
	fmt.Println("남은 시도 =", remainTries(5, 2)) // 3이 나오면 정답
}
