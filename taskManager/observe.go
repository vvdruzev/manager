package taskManager

type Listener func(interface{})

type Observer struct {
	quit      chan bool
	events    chan string
	listeners map[string][]Listener
}

func NewObserver() *Observer {
	var o Observer

	// Create the observer channels.
	o.quit = make(chan bool)
	o.events = make(chan string)
	o.listeners = make(map[string][]Listener)
	return &o
}

func (o *Observer) StartObserve() error {
	return o.eventLoop()
}

func (o *Observer) Close() {
	if o.events != nil {
		// Send a quit signal.
		o.quit <- true

		// Close channels.
		close(o.quit)
		close(o.events)
	}

}

func (o *Observer) eventLoop() error {
	go func() {
		for {
			select {
			case event := <-o.events:
				o.handleEvent(event)
			case <-o.quit:
				return
			}
		}
	}()
	return nil
}

func (o *Observer) handleEvent(event string) {
	if _, ok := o.listeners[event]; !ok {
		return
	}
	for _, listener := range o.listeners[event] {
		listener("ok")
	}
}

func (o *Observer) AddListener(event string, l Listener) {
	o.listeners[event] = append(o.listeners[event], l)
}

func (o *Observer) Emit(event string) {
	o.events <- event
}
