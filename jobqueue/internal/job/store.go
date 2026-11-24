package job // job.go와 같은 패키지 (같은 폴더면 같은 패키지)

import (
	"strconv" // 숫자 → 문자열 변환 (Java의 Integer.toString)
	"sync"    // Mutex(자물쇠) 같은 동시성 도구
)

// Store는 작업들을 메모리에 보관하는 창고 (Repository 역할)
type Store struct {
	mu     sync.Mutex     // 자물쇠: 한 번에 goroutine 하나만 아래 데이터를 만지게 함
	jobs   map[string]Job // 작업 번호 → 작업 (Java의 HashMap<String, Job>)
	nextID int            // 마지막으로 발급한 번호 (번호표 발급기)
}

// NewStore는 빈 창고를 만들어 돌려줌 (Java의 생성자 new Store() 역할)
func NewStore() *Store {
	return &Store{jobs: make(map[string]Job)} // map은 make로 만들어야 쓸 수 있음, &는 "만든 것의 주소"
}

// Add는 새 작업을 등록하고, 번호가 붙은 작업을 돌려줌
// (s *Store)는 "이 함수는 Store에 붙은 메서드"라는 뜻 (s는 Java의 this)
func (s *Store) Add(task string) Job {
	s.mu.Lock()         // 빈칸 1: 자물쇠 잠그기
	defer s.mu.Unlock() // defer: 함수가 끝날 때 자동 실행 → 자물쇠 풀기 (Java의 finally)

	s.nextID++                                                              // 번호 1 증가
	j := Job{ID: strconv.Itoa(s.nextID), Task: task, Status: StatusPending} // 새 작업 만들기
	s.jobs[j.ID] = j                                                        // 빈칸 2: map에 j를 번호(j.ID)로 저장
	return j                                                                // 만든 작업을 돌려줌
}

// Get은 번호로 작업을 찾음. 반환값이 2개: 작업, 찾았는지 여부(bool)
func (s *Store) Get(id string) (Job, bool) {
	s.mu.Lock()         // 읽을 때도 잠금 (누가 쓰는 중에 읽으면 깨진 값을 볼 수 있음)
	defer s.mu.Unlock() // 함수가 끝나면 풀기

	j, ok := s.jobs[id] // 빈칸 3: map에서 id로 꺼내기 → j(작업), ok(있으면 true)
	return j, ok        // 찾은 결과 돌려줌
}
