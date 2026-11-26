package api // health.go와 같은 api 패키지 (요청을 받는 계층 = 홀 직원)

import (
	"encoding/json" // JSON ↔ Go struct 변환 (Java의 Jackson/Gson 역할)
	"net/http"      // HTTP 요청/응답

	"jobqueue/internal/job" // 작업과 저장소가 있는 job 패키지 (주방/창고)
)

// JobHandler는 /jobs 요청을 처리하는 홀 직원
// store를 들고 있어서, 손님 요청을 받으면 창고(Store)에 맡기거나 꺼내옴
type JobHandler struct {
	store *job.Store // 작업 창고의 주소 (모든 요청이 같은 창고 하나를 같이 씀)
}

// NewJobHandler는 창고를 받아서 홀 직원을 만들어 돌려줌 (생성자 역할)
func NewJobHandler(store *job.Store) *JobHandler {
	return &JobHandler{store: store} // 받은 창고를 들려서 직원 생성
}

// createRequest는 POST /jobs로 들어오는 주문서 모양 (요청용 DTO)
// 예: {"task": "send-email"}
type createRequest struct {
	Task string `json:"task"` // JSON의 "task" 값이 여기에 들어옴
}

// Create는 POST /jobs 처리: 작업 등록
func (h *JobHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createRequest                                        // 빈 주문서 한 장 준비
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { // 요청 본문(JSON)을 읽어 주문서에 채움
		http.Error(w, "JSON 형식이 잘못됐어요", http.StatusBadRequest) // 못 읽으면 400 (손님이 주문을 잘못 씀)
		return                                                 // 여기서 처리 끝
	}
	if req.Task == "" { // task가 비어 있으면
		http.Error(w, "task는 비어 있으면 안 돼요", http.StatusBadRequest) // 역시 400
		return                                                    // 처리 끝
	}

	j := h.store.Add(req.Task) // 빈칸 1: 창고(h.store)에 req.Task를 등록하고, 번호 붙은 작업을 받기 (힌트: store.go의 Add)

	writeJSON(w, http.StatusCreated, j) // 201 Created + 만든 작업을 JSON으로 응답
}

// Get은 GET /jobs/{id} 처리: 번호로 작업 조회
func (h *JobHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id") // 주소의 {id} 자리 값 꺼내기 (/jobs/3 → "3")

	j, ok := h.store.Get(id) // 빈칸 2: 창고에서 id로 찾기 → j(작업), ok(찾았는지) (힌트: store.go의 Get)
	if !ok {                 // 못 찾았으면
		http.Error(w, "그런 작업 번호는 없어요", http.StatusNotFound) // 빈칸 3: 404 Not Found 상태코드 (힌트: 위의 http.StatusBadRequest 참고)
		return                                              // 처리 끝
	}

	writeJSON(w, http.StatusOK, j) // 200 OK + 찾은 작업을 JSON으로 응답
}

// writeJSON은 "상태코드 + JSON 본문" 응답을 쓰는 공통 도우미 (소문자라 api 패키지 안에서만 씀)
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json") // "본문은 JSON이에요" 표시 (헤더는 WriteHeader보다 먼저)
	w.WriteHeader(status)                              // 상태코드 쓰기 (200, 201 등)
	json.NewEncoder(w).Encode(v)                       // v를 JSON으로 바꿔 응답 본문에 씀
}
