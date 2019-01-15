/*
 *          Copyright 2019, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package events provides an event queue and events.
//
// Version 0.1.0.
package events


// Event is posted into and retrieved from the event queue.
type Event interface {
	EventTypeID() int
	TimeStamp() uint64
}

// EventQueue holds the events.
type EventQueue interface {
	NextEvent() Event
	PostEvent(event Event)
}

// DefaultEvent serves as a template for other events
type DefaultEvent struct {
	eventTypeID int
	timeStamp uint64
}

// DefaultEventQueue serves as a template for other event queues
type DefaultEventQueue struct {
	events []Event
	currIndex int
	nextIndex int
}

// NewEvent returns a simple event, that may serve as a stub.
func NewEvent(eventTypeID int, timeStamp uint64) Event {
	event := new(DefaultEvent)
	event.eventTypeID = eventTypeID
	event.timeStamp = timeStamp
	return event
}

// NewEventQueue returns a simple unsynchronized event queue. Default capacity is 6.
func NewEventQueue(initialCapacity ...int) EventQueue {
	eventQueue := new(DefaultEventQueue)
	if len(initialCapacity) > 0 {
		eventQueue.events = make([]Event, initialCapacity[0])
	} else {
		eventQueue.events = make([]Event, 6)
	}
	return eventQueue
}

// EventTypeID returns a value representing the type of the event.
func (event DefaultEvent) EventTypeID() int {
	return event.eventTypeID
}

// TimeStamp returns a value representing the creation time of the event.
func (event DefaultEvent) TimeStamp() uint64 {
	return event.timeStamp
}

// NextEvent returns the next event in the queue. If no event is available, nil is returned.
func (queue DefaultEventQueue) NextEvent() Event {
	event := queue.events[queue.currIndex]
	queue.events[queue.currIndex] = nil
	if queue.currIndex != queue.nextIndex {
		queue.currIndex = queue.previewNextIndex(queue.currIndex)
	}
	return event
}

// PostEvent puts an event into queue.
func (queue DefaultEventQueue) PostEvent(event Event) {
	queue.ensureCapacity()
	queue.events[queue.nextIndex] = event
	queue.nextIndex = queue.previewNextIndex(queue.nextIndex)
}

func (queue DefaultEventQueue) previewNextIndex(index int) int {
	nextIndex := index + 1
	if nextIndex < len(queue.events) {
		return nextIndex
	} else {
		return 0
	}
}

func (queue DefaultEventQueue) ensureCapacity() {
	if queue.events[queue.nextIndex] != nil {
		events := make([]Event, (len(queue.events) + 1) * 2)
		copy(events, queue.events[queue.currIndex:])
		copy(events, queue.events[:queue.currIndex])
		queue.currIndex = 0
		queue.nextIndex = len(queue.events)
		queue.events = events
	}
}
