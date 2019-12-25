package engine

import (
	"sync"
)

type Handler interface {
	Post(cmd Command)
}

type Command interface {
	Execute(handler Handler)
}

type finishCommand func (handler Handler)

func (c finishCommand) Execute(handler Handler) {
	c(handler)
}

type messageQueue struct {
	sync.Mutex
	data []Command
	waiting bool
	receiveSignal chan struct{}
}

func (mq *messageQueue) push(cmd Command) {
	mq.Lock()
	defer mq.Unlock()
	mq.data = append(mq.data, cmd)
	if mq.waiting {
		mq.waiting = false
		mq.receiveSignal <- struct{}{}
	}
}

func (mq *messageQueue) pull() Command {
	mq.Lock()
	defer mq.Unlock()
	if len(mq.data) == 0 {
		mq.waiting = true;
		mq.Unlock()
		<- mq.receiveSignal
		mq.Lock()
	}
	res := mq.data[0]
	mq.data[0] = nil
	mq.data = mq.data[1:]
	return res
}

func (mq *messageQueue) size() int {
	return len(mq.data)
}

type EventLoop struct {
	queue *messageQueue
	terminateReceived bool
	stopSignal chan struct{}
}

func (el *EventLoop) Start() {
	el.queue = new(messageQueue)
	el.queue.receiveSignal = make(chan struct{})
	el.stopSignal = make(chan struct{})
	go func() {
		for (!el.terminateReceived) || (el.queue.size() != 0) {
			cmd := el.queue.pull()
			cmd.Execute(el)
		}
		el.stopSignal <- struct{}{}
	}()
}

func (el *EventLoop) AwaitFinish() {
	el.Post(finishCommand(func (h Handler) {
		h.(*EventLoop).terminateReceived = true
	}))
	<- el.stopSignal
}

func (el * EventLoop) Post(cmd Command) {
	el.queue.push(cmd)
}
