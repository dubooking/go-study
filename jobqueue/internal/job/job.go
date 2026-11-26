package job // 작업(Job)과 작업 저장소를 담는 패키지

// 작업 상태값 (const = 바꿀 수 없는 상수, Java의 static final)
const (
	StatusPending = "pending" // 대기 중: 등록만 되고 아직 처리 전
	StatusDone    = "done"    // 완료: 워커가 처리를 끝냄 (3단계에서 사용)
)

// Job은 작업 하나의 정보를 담는 상자 (DTO 역할)
// 필드 이름이 대문자로 시작해야 다른 패키지(api)에서 읽을 수 있음
// `json:"id"` 같은 꼬리표(태그): JSON으로 바꿀 때 쓸 이름 (없으면 "ID"처럼 대문자로 나감)
type Job struct {
	ID     string `json:"id"`     // 작업 번호 (예: "1")
	Task   string `json:"task"`   // 무슨 작업인지 (예: "send-email")
	Status string `json:"status"` // 현재 상태 (StatusPending 또는 StatusDone)
}
