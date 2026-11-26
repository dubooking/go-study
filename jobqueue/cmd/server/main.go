package main // 실행 프로그램은 main 패키지에서 시작

import ( // 여러 개를 import할 때는 괄호로 묶음
	"log"      // 로그 출력용 (시간이 같이 찍힘)
	"net/http" // HTTP 서버 표준 라이브러리

	"jobqueue/internal/api" // 우리가 만든 api 패키지 (모듈 이름/폴더 경로)
	"jobqueue/internal/job" // 작업 저장소가 있는 job 패키지
)

func main() { // 프로그램 시작점
	store := job.NewStore()          // 창고를 딱 하나 만듦 (모든 요청이 이 창고를 같이 씀)
	jobs := api.NewJobHandler(store) // 창고를 들려서 홀 직원 생성 (조립)
	mux := http.NewServeMux()        // 요청 주소 → 처리 함수를 연결하는 표 (web.xml의 servlet-mapping 역할)

	mux.HandleFunc("GET /health", api.Health)  // 서버 살아 있는지 확인
	mux.HandleFunc("POST /jobs", jobs.Create)  // 작업 등록 → Create
	mux.HandleFunc("GET /jobs/{id}", jobs.Get) // 작업 조회 → Get ({id}는 주소 속 빈칸)

	log.Println("서버 시작: http://localhost:8080")  // 시작 메시지 출력
	log.Fatal(http.ListenAndServe(":8080", mux)) // 8080 포트로 서버 띄우기 (실패하면 에러 찍고 종료)
}
