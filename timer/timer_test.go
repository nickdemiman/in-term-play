package timer

// import (
// 	"testing"
// 	"time"
// )

// type (
// 	testingSender struct{}
// )

// func (t *testingSender) NotifyTimer() {

// }

// func TestUnregisterFunc(t *testing.T) {
// 	SetInterval(time.Microsecond)
// 	_timer := GetTimer()

// 	defer func() {
// 		go _timer.Stop()
// 	}()

// 	go _timer.Run()

// 	sender := testingSender{}

// 	_timer.Register(&sender)
// 	err := _timer.Unregister(&sender)

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	if _timer.searchSender(&sender) != -1 {
// 		t.Errorf("sender not removed")
// 	}
// }
