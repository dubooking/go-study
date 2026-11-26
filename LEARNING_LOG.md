# LEARNING_LOG

## day2 (2026-11-20)
- 한 것: CLAUDE.md 작성, WSL에 Git credential manager 설정, 첫 commit/push
- 막힌 것: 명령어가 두 번 붙여져 "no action specified" 에러 / commit과 push 차이가 헷갈림
- 배운 것: commit은 내 컴퓨터에 저장, push는 GitHub에 업로드, pull은 다른 기기에서 받아오기

## day3 (2026-11-24)
- 한 것: jobqueue 폴더 시작, net/http로 8080 포트 서버를 띄우고 GET /health → "OK" 응답 확인 (브라우저) / 2-1: Job struct와 메모리 저장소(map + Mutex) 만들고 go test로 확인
- 막힌 것: 브라우저에 localhost:8080만 쳐서 404 (등록한 경로 /health를 붙여야 함), Windows 폴더와 Ubuntu 폴더를 헷갈림, VS Code를 WSL에 연결하는 법, 읽기(Get)에도 자물쇠를 거는 이유와 defer가 필요한 이유가 바로 이해되지 않음
- 배운 것: Go는 Tomcat 없이 프로그램이 직접 서버가 됨, ServeMux는 web.xml의 servlet-mapping 역할, main은 조립만 하고 처리는 internal/api에 둠, 여러 goroutine이 map을 동시에 만지면 가끔 서버가 죽음(경쟁 상태) → Mutex로 한 번에 하나만, defer로 반드시 풀기

## day4 (2026-11-26)
- 한 것: jobqueue 2-2: POST /jobs(등록), GET /jobs/{id}(조회) API를 만들고 빈칸 3개를 직접 채움, curl로 201 / 200 / 404 확인
- 막힌 것: `Task`만 쓰면 "누구의 Task?"라서 `req.Task`로 써야 함, 400(주문서 잘못 씀)과 404(그런 번호 없음)를 헷갈림, 화면에서 고친 파일을 저장하지 않아 디스크에는 반영이 안 됨
- 배운 것: main은 창고를 하나만 만들어 직원(Handler)에게 들려주는 조립 담당, Handler는 요청 해석과 응답만, Store는 보관만 / map은 메모리라 서버를 끄면 데이터가 사라짐
