package taskManager

import "testing"

func TestObserver_handleEvent1(t *testing.T) {
	var x int
	observer := NewObserver()
	observer.AddListener(EventCreate, func(i interface{}) {
		x = 1
	})
	observer.handleEvent(EventCreate)

	if x != 1 {
		t.Errorf("got %d; want %d", x, 1)
	}
}

func TestObserver_handleEvent2(t *testing.T) {
	var x int
	observer := NewObserver()
	observer.AddListener(EventCreate, func(i interface{}) {
		x = 1
	})
	observer.handleEvent(EventComplete)

	if x != 0 {
		t.Errorf("got %d; want %d", x, 0)
	}
}

func TestObserver_Close(t *testing.T) {
	observer := NewObserver()
	err := observer.StartObserve()
	if err != nil {
		t.Errorf("got %s; want nil", err)
	}
	observer.Close()
	if _, ok :=<-observer.quit; ok {
		t.Errorf("channel quit nol closed")
	}
	if _, ok :=<-observer.events; ok {
		t.Errorf("channel events nol closed")
	}

}