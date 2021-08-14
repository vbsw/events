/*
 *        Copyright 2019, 2021 Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package events provides events and event queues.
//
// Version 2.0.0.
package events

import "sync"

const (
	Generic = iota
	KeyDown
	KeyUp
	LeftButtonDown
	LeftButtonUp
	RightButtonDown
	RightButtonUp
	MiddleButtonDown
	MiddleButtonUp
	X1ButtonDown
	X1ButtonUp
	X2ButtonDown
	X2ButtonUp
	MouseWheel
	MouseMove
	FocusLoose
	FocusGain
	WindowResize
	WindowMove
	Destroy
)

// Event is posted to and retrieved from an event queue.
type Event interface {
	TypeID() int
}

// Queue holds the events.
type Queue interface {
	Close()
	NextEvent() Event
	PostEvent(event Event)
}

// GenericEvent serves as a template for other events.
type GenericEvent struct {
	typeId int
	Time   uint64
}

// KeyDownEvent is an event of pressed key.
type KeyDownEvent struct {
	GenericEvent
	KeyCode int
	Repeat  int
}

// KeyUpEvent is an event of released key.
type KeyUpEvent struct {
	GenericEvent
	KeyCode int
}

// ButtonDownEvent is an event of pressed mouse button.
type ButtonDownEvent struct {
	GenericEvent
	DoubleClick bool
}

// ButtonUpEvent is an event of released mouse button.
type ButtonUpEvent struct {
	GenericEvent
}

// MouseWheelEvent is an event of spun mouse wheel.
type MouseWheelEvent struct {
	GenericEvent
	Delta int
}

// MouseMoveEvent is an event of moved mouse.
type MouseMoveEvent struct {
	GenericEvent
	X int
	Y int
}

// WindowMoveEvent is an event of moved window.
type WindowMoveEvent struct {
	GenericEvent
	WindowX int
	WindowY int
	ClientX int
	ClientY int
}

// WindowResizeEvent is an event of resized window.
type WindowResizeEvent struct {
	GenericEvent
	WindowWidth  int
	WindowHeight int
	ClientWidth  int
	ClientHeight int
}

// DestroyEvent is an event of resized window.
type DestroyEvent struct {
	GenericEvent
	Confirmed bool
}

// DefaultQueue serves as a template for other event queues.
type DefaultQueue struct {
	events    []Event
	currIndex int
	nextIndex int
	full      bool
}

// DefaultThreadSafeQueue serves as a template for other event queues.
type DefaultThreadSafeQueue struct {
	DefaultQueue
	mutex sync.Mutex
}

// NewGenericEvent returns a new instance of GenericEvent.
func NewGenericEvent(typeId int, time uint64) *GenericEvent {
	event := new(GenericEvent)
	event.typeId = typeId
	event.Time = time
	return event
}

// NewKeyDownEvent returns a new instance of KeyDownEvent.
func NewKeyDownEvent(time uint64, keyCode, repeat int) *KeyDownEvent {
	event := new(KeyDownEvent)
	event.typeId = KeyDown
	event.Time = time
	event.KeyCode = keyCode
	event.Repeat = repeat
	return event
}

// NewKeyUpEvent returns a new instance of KeyUpEvent.
func NewKeyUpEvent(time uint64, keyCode int) *KeyUpEvent {
	event := new(KeyUpEvent)
	event.typeId = KeyUp
	event.Time = time
	event.KeyCode = keyCode
	return event
}

// NewButtonDownEvent returns a new instance of ButtonDownEvent.
func NewButtonDownEvent(typeId int, time uint64, doubleClick bool) *ButtonDownEvent {
	event := new(ButtonDownEvent)
	event.typeId = typeId
	event.Time = time
	event.DoubleClick = doubleClick
	return event
}

// NewLeftButtonDownEvent returns a new instance of ButtonDownEvent.
func NewLeftButtonDownEvent(time uint64, doubleClick bool) *ButtonDownEvent {
	return NewButtonDownEvent(LeftButtonDown, time, doubleClick)
}

// NewRightButtonDownEvent returns a new instance of ButtonDownEvent.
func NewRightButtonDownEvent(time uint64, doubleClick bool) *ButtonDownEvent {
	return NewButtonDownEvent(RightButtonDown, time, doubleClick)
}

// NewMiddleButtonDownEvent returns a new instance of ButtonDownEvent.
func NewMiddleButtonDownEvent(time uint64, doubleClick bool) *ButtonDownEvent {
	return NewButtonDownEvent(MiddleButtonDown, time, doubleClick)
}

// NewX1ButtonDownEvent returns a new instance of ButtonDownEvent.
func NewX1ButtonDownEvent(time uint64, doubleClick bool) *ButtonDownEvent {
	return NewButtonDownEvent(X1ButtonDown, time, doubleClick)
}

// NewX2ButtonDownEvent returns a new instance of ButtonDownEvent.
func NewX2ButtonDownEvent(time uint64, doubleClick bool) *ButtonDownEvent {
	return NewButtonDownEvent(X2ButtonDown, time, doubleClick)
}

// NewButtonUpEvent returns a new instance of ButtonUpEvent.
func NewButtonUpEvent(typeId int, time uint64) *ButtonUpEvent {
	event := new(ButtonUpEvent)
	event.typeId = typeId
	event.Time = time
	return event
}

// NewLeftButtonUpEvent returns a new instance of ButtonUpEvent.
func NewLeftButtonUpEvent(time uint64) *ButtonUpEvent {
	return NewButtonUpEvent(LeftButtonUp, time)
}

// NewRightButtonUpEvent returns a new instance of ButtonUpEvent.
func NewRightButtonUpEvent(time uint64) *ButtonUpEvent {
	return NewButtonUpEvent(RightButtonUp, time)
}

// NewMiddleButtonUpEvent returns a new instance of ButtonUpEvent.
func NewMiddleButtonUpEvent(time uint64) *ButtonUpEvent {
	return NewButtonUpEvent(MiddleButtonUp, time)
}

// NewX1ButtonUpEvent returns a new instance of ButtonUpEvent.
func NewX1ButtonUpEvent(time uint64) *ButtonUpEvent {
	return NewButtonUpEvent(X1ButtonUp, time)
}

// NewX2ButtonUpEvent returns a new instance of ButtonUpEvent.
func NewX2ButtonUpEvent(time uint64) *ButtonUpEvent {
	return NewButtonUpEvent(X2ButtonUp, time)
}

// NewMouseWheelEvent returns a new instance of MouseWheelEvent.
func NewMouseWheelEvent(time uint64, delta int) *MouseWheelEvent {
	event := new(MouseWheelEvent)
	event.typeId = MouseWheel
	event.Time = time
	event.Delta = delta
	return event
}

// NewMouseMoveEvent returns a new instance of MouseMoveEvent.
func NewMouseMoveEvent(time uint64, x, y int) *MouseMoveEvent {
	event := new(MouseMoveEvent)
	event.typeId = MouseMove
	event.Time = time
	event.X = x
	event.Y = y
	return event
}

// NewFocusLooseEvent returns a new instance of GenericEvent.
func NewFocusLooseEvent(time uint64) *GenericEvent {
	event := new(GenericEvent)
	event.typeId = FocusLoose
	event.Time = time
	return event
}

// NewFocusGainEvent returns a new instance of GenericEvent.
func NewFocusGainEvent(time uint64) *GenericEvent {
	event := new(GenericEvent)
	event.typeId = FocusGain
	event.Time = time
	return event
}

// NewWindowResizeEvent returns a new instance of WindowResizeEvent.
func NewWindowResizeEvent(time uint64, windowWidth, windowHeight, clientWidth, clientHeight int) *WindowResizeEvent {
	event := new(WindowResizeEvent)
	event.typeId = WindowResize
	event.Time = time
	event.WindowWidth = windowWidth
	event.WindowHeight = windowHeight
	event.ClientWidth = clientWidth
	event.ClientHeight = clientHeight
	return event
}

// NewWindowMoveEvent returns a new instance of WindowMoveEvent.
func NewWindowMoveEvent(time uint64, windowX, windowY, clientX, clientY int) *WindowMoveEvent {
	event := new(WindowMoveEvent)
	event.typeId = WindowMove
	event.Time = time
	event.WindowX = windowX
	event.WindowY = windowY
	event.ClientX = clientX
	event.ClientY = clientY
	return event
}

// NewDestroyEvent returns a new instance of DestroyEvent.
func NewDestroyEvent(time uint64, confirmed bool) *DestroyEvent {
	event := new(DestroyEvent)
	event.typeId = Destroy
	event.Time = time
	event.Confirmed = confirmed
	return event
}

// NewQueue returns an event queue. Default capacity is 6.
// Queues' capacity is increased, if events number exceeds the capacity.
func NewQueue(initialCapacity ...int) Queue {
	queue := new(DefaultQueue)
	if len(initialCapacity) > 0 && initialCapacity[0] > 0 {
		queue.events = make([]Event, initialCapacity[0])
	} else {
		queue.events = make([]Event, 6)
	}
	return queue
}

// NewThreadSafeQueue returns a thread safe event queue. Default capacity is 6.
// Queues' capacity is increased, if events number exceeds the capacity.
func NewThreadSafeQueue(initialCapacity ...int) Queue {
	queue := new(DefaultThreadSafeQueue)
	if len(initialCapacity) > 0 && initialCapacity[0] > 0 {
		queue.events = make([]Event, initialCapacity[0])
	} else {
		queue.events = make([]Event, 6)
	}
	return queue
}

// TypeID returns the typeId of the event type.
func (event *GenericEvent) TypeID() int {
	return event.typeId
}

// Close removes all events from queue. Further posted events are ignored. Next call of NextEvent returns nil.
func (queue *DefaultQueue) Close() {
	if queue.nextIndex >= 0 {
		queue.events = nil
		queue.currIndex = -1
		queue.nextIndex = -1
		queue.full = false
	}
}

// NextEvent returns the next event in the queue if available, otherwise nil.
func (queue *DefaultQueue) NextEvent() Event {
	if queue.currIndex != queue.nextIndex || queue.full {
		event := queue.events[queue.currIndex]
		queue.events[queue.currIndex] = nil
		queue.currIndex = queue.subsequentIndex(queue.currIndex)
		queue.full = false
		return event
	}
	return nil
}

// PostEvent puts an event into queue.
func (queue *DefaultQueue) PostEvent(event Event) {
	if queue.nextIndex >= 0 {
		queue.ensureCapacity()
		queue.events[queue.nextIndex] = event
		queue.nextIndex = queue.subsequentIndex(queue.nextIndex)
		queue.full = bool(queue.nextIndex == queue.currIndex)
	}
}

// Close removes all events from queue. Further posted events are ignored. Next call of NextEvent returns nil.
func (queue *DefaultThreadSafeQueue) Close() {
	queue.mutex.Lock()
	queue.DefaultQueue.Close()
	queue.mutex.Unlock()
}

// NextEvent returns the next event in the queue if available, otherwise nil.
func (queue *DefaultThreadSafeQueue) NextEvent() Event {
	queue.mutex.Lock()
	event := queue.DefaultQueue.NextEvent()
	queue.mutex.Unlock()
	return event
}

// PostEvent puts an event in queue.
func (queue *DefaultThreadSafeQueue) PostEvent(event Event) {
	queue.mutex.Lock()
	queue.DefaultQueue.PostEvent(event)
	queue.mutex.Unlock()
}

func (queue *DefaultQueue) subsequentIndex(index int) int {
	nextIndex := index + 1
	if nextIndex < len(queue.events) {
		return nextIndex
	}
	return 0
}

func (queue *DefaultQueue) ensureCapacity() {
	if queue.full {
		events := make([]Event, (len(queue.events)+1)*2)
		copy(events, queue.events[queue.currIndex:])
		copy(events[len(queue.events)-queue.currIndex:], queue.events[:queue.currIndex])
		queue.currIndex = 0
		queue.nextIndex = len(queue.events)
		queue.events = events
		queue.full = false
	}
}
