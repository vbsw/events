/*
 *          Copyright 2019, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package events provides generic event queues and generic events.
//
// Version 1.0.0.
package events

import "sync"

// Event is posted into and retrieved from the event queue.
type Event interface {
	EventTypeID() int
	TimeStamp() uint64
}

// Queue holds the events.
type Queue interface {
	Close()
	NextEvent() Event
	PostEvent(event Event)
}

// DefaultEvent serves as a template for other events.
type DefaultEvent struct {
	eventTypeID int
	timeStamp   uint64
}

// DefaultQueue serves as a template for other event queues.
type DefaultQueue struct {
	events    []Event
	currIndex int
	nextIndex int
}

// DefaultSynchronizedQueue serves as a template for other event queues.
type DefaultSynchronizedQueue struct {
	DefaultQueue
	mutex sync.Mutex
}

// NewEvent returns a simple event, that may serve as a stub.
func NewEvent(eventTypeID int, timeStamp uint64) Event {
	event := new(DefaultEvent)
	event.eventTypeID = eventTypeID
	event.timeStamp = timeStamp
	return event
}

// NewQueue returns a simple unsynchronized event queue. Default capacity is 6.
func NewQueue(initialCapacity ...int) Queue {
	queue := new(DefaultQueue)
	if len(initialCapacity) > 0 {
		queue.events = make([]Event, initialCapacity[0])
	} else {
		queue.events = make([]Event, 6)
	}
	return queue
}

// NewSynchronizedQueue returns a simple synchronized event queue. Default capacity is 6.
func NewSynchronizedQueue(initialCapacity ...int) Queue {
	queue := new(DefaultSynchronizedQueue)
	if len(initialCapacity) > 0 {
		queue.events = make([]Event, initialCapacity[0])
	} else {
		queue.events = make([]Event, 6)
	}
	return queue
}

// EventTypeID returns a value representing the type of the event.
func (event *DefaultEvent) EventTypeID() int {
	return event.eventTypeID
}

// TimeStamp returns a value representing the creation time of the event.
func (event *DefaultEvent) TimeStamp() uint64 {
	return event.timeStamp
}

// Close removes all events from queue. Further posted events are ignored. Next call of NextEvent returns nil.
func (queue *DefaultQueue) Close() {
	if queue.nextIndex >= 0 {
		if queue.currIndex < queue.nextIndex {
			setNil(queue.events[queue.currIndex:queue.nextIndex])
		} else if queue.nextIndex < queue.currIndex {
			setNil(queue.events[queue.currIndex:])
			setNil(queue.events[:queue.nextIndex])
		} else if queue.events[queue.currIndex] != nil {
			setNil(queue.events)
		}
		queue.nextIndex = -1
	}
}

// NextEvent returns the next event in the queue if available, otherwise nil.
func (queue *DefaultQueue) NextEvent() Event {
	event := queue.events[queue.currIndex]
	queue.events[queue.currIndex] = nil
	if event != nil {
		queue.currIndex = queue.previewNextIndex(queue.currIndex)
	}
	return event
}

// PostEvent puts an event into queue.
func (queue *DefaultQueue) PostEvent(event Event) {
	if queue.nextIndex >= 0 {
		queue.ensureCapacity()
		queue.events[queue.nextIndex] = event
		queue.nextIndex = queue.previewNextIndex(queue.nextIndex)
	}
}

// Close removes all events from queue. Further posted events are ignored. Next call of NextEvent returns nil.
func (queue *DefaultSynchronizedQueue) Close() {
	queue.mutex.Lock()
	queue.DefaultQueue.Close()
	queue.mutex.Unlock()
}

// NextEvent returns the next event in the queue if available, otherwise nil.
func (queue *DefaultSynchronizedQueue) NextEvent() Event {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	return queue.DefaultQueue.NextEvent()
}

// PostEvent puts an event into queue.
func (queue *DefaultSynchronizedQueue) PostEvent(event Event) {
	queue.mutex.Lock()
	queue.DefaultQueue.PostEvent(event)
	queue.mutex.Unlock()
}

func (queue *DefaultQueue) previewNextIndex(index int) int {
	nextIndex := index + 1
	if nextIndex < len(queue.events) {
		return nextIndex
	}
	return 0
}

func (queue *DefaultQueue) ensureCapacity() {
	if queue.events[queue.nextIndex] != nil {
		events := make([]Event, (len(queue.events)+1)*2)
		copy(events, queue.events[queue.currIndex:])
		copy(events[len(queue.events)-queue.currIndex:], queue.events[:queue.currIndex])
		queue.currIndex = 0
		queue.nextIndex = len(queue.events)
		queue.events = events
	}
}

func setNil(events []Event) {
	for i := range events {
		events[i] = nil
	}
}
