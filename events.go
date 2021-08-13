/*
 *        Copyright 2019, 2021 Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package events provides generic event queues and generic events.
//
// Version 2.0.0.
package events

import "sync"

// Event is posted to and retrieved from an event queue.
type Event interface {
	ID() int
	Time() uint64
}

// Queue holds the events.
type Queue interface {
	Close()
	NextEvent() Event
	PostEvent(event Event)
}

// DefaultEvent serves as a template for other events.
type DefaultEvent struct {
	id   int
	time uint64
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

// NewEvent returns a new event.
func NewEvent(id int, time uint64) Event {
	event := new(DefaultEvent)
	event.id = id
	event.time = time
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

// ID returns the id of the event type.
func (event *DefaultEvent) ID() int {
	return event.id
}

// Time returns the creation time of the event.
func (event *DefaultEvent) Time() uint64 {
	return event.time
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
