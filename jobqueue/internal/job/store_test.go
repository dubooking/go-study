package job // 테스트도 같은 패키지 (파일 이름이 _test.go로 끝나면 테스트 파일)

import (
	"sync"    // WaitGroup: goroutine 여러 개가 다 끝날 때까지 기다리는 도구
	"testing" // Go 기본 테스트 도구 (Java의 JUnit)
)

// 등록한 작업을 번호로 다시 찾을 수 있는지 확인
func TestAddAndGet(t *testing.T) { // 함수 이름이 Test로 시작하면 go test가 자동 실행
	s := NewStore()              // 빈 창고 만들기
	added := s.Add("send-email") // 작업 하나 등록

	got, ok := s.Get(added.ID)           // 방금 받은 번호로 찾기
	if !ok || got.Task != "send-email" { // 못 찾았거나 내용이 다르면
		t.Errorf("Get(%q) = %v, %v", added.ID, got, ok) // 테스트 실패로 표시
	}

	if _, ok := s.Get("999"); ok { // 없는 번호는 ok가 false여야 함 (_는 "안 쓰는 값")
		t.Error("없는 번호인데 찾았다고 나옴") // 테스트 실패로 표시
	}
}

// 100개 작업을 동시에 등록해도 100개가 다 저장되는지 확인 (Mutex 실험용)
func TestAddConcurrent(t *testing.T) {
	s := NewStore() // 빈 창고 만들기

	var wg sync.WaitGroup      // 끝난 goroutine 수를 세는 카운터
	for i := 0; i < 100; i++ { // 100번 반복
		wg.Add(1)   // "일꾼 1명 출발" 표시
		go func() { // go: 새 goroutine을 띄움 → 100개가 동시에 실행됨
			defer wg.Done()     // 끝나면 "1명 끝" 표시
			s.Add("send-email") // 작업 등록
		}()
	}
	wg.Wait() // 100명이 다 끝날 때까지 기다림

	if len(s.jobs) != 100 { // 저장된 개수가 100이 아니면
		t.Errorf("저장된 작업 = %d개, 기대값 100개", len(s.jobs)) // 테스트 실패로 표시
	}
}
