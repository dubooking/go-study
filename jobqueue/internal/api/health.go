package api // 이 파일은 api 패키지 소속 (폴더 이름과 맞춤)

import "net/http" // HTTP 요청/응답을 다루는 표준 라이브러리

// Health는 서버가 살아 있는지 알려주는 핸들러
// Java Servlet의 doGet(HttpServletRequest req, HttpServletResponse resp)과 같은 역할
// 이름이 대문자 H로 시작해야 다른 패키지(main)에서 부를 수 있음
func Health(w http.ResponseWriter, r *http.Request) { // w: 응답을 쓰는 곳, r: 들어온 요청
	w.Write([]byte("OK")) // 빈칸 1: 응답 본문에 "OK"를 쓰기
}
