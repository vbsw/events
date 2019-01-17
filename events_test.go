/*
 *          Copyright 2019, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package events

import "testing"

func TestNewEventQueue(t *testing.T) {
	queue1 := NewEventQueue(1).(*DefaultEventQueue)
	queue2 := NewEventQueue(2).(*DefaultEventQueue)

	if len(queue1.events) != 1 {
		t.Error(len(queue1.events))
	}
	if len(queue2.events) != 2 {
		t.Error(len(queue2.events))
	}
}

func TestPostEvent(t *testing.T) {
	queue1 := NewEventQueue(1).(*DefaultEventQueue)
	queue2 := NewEventQueue(2).(*DefaultEventQueue)

	queue1.PostEvent(NewEvent(0, 0))
	if queue1.currIndex != 0 {
		t.Error(queue1.currIndex)
	}
	if queue1.nextIndex != 0 {
		t.Error(queue1.nextIndex)
	}
	if len(queue1.events) != 1 {
		t.Error(len(queue1.events))
	}
	if queue1.events[0].EventTypeID() != 0 {
		t.Error(queue1.events[0].EventTypeID())
	}

	queue2.PostEvent(NewEvent(0, 0))
	if queue2.currIndex != 0 {
		t.Error(queue2.currIndex)
	}
	if queue2.nextIndex != 1 {
		t.Error(queue2.nextIndex)
	}
	queue2.PostEvent(NewEvent(1, 0))
	if queue2.currIndex != 0 {
		t.Error(queue2.currIndex)
	}
	if queue2.nextIndex != 0 {
		t.Error(queue2.nextIndex)
	}
	if len(queue2.events) != 2 {
		t.Error(len(queue2.events))
	}
	if queue2.events[1].EventTypeID() != 1 {
		t.Error(queue2.events[1].EventTypeID())
	}
}

func TestFillDefaultQueue(t *testing.T) {
	queue1 := newFilledDefaultQueue(1)
	queue2 := newFilledDefaultQueue(2)

	if queue1.currIndex != 0 {
		t.Error(queue1.currIndex)
	}
	if queue1.nextIndex != 0 {
		t.Error(queue1.nextIndex)
	}
	if len(queue1.events) != 1 {
		t.Error(len(queue1.events))
	}
	if queue1.events[0].EventTypeID() != 0 {
		t.Error(queue1.events[0].EventTypeID())
	}

	if queue2.currIndex != 0 {
		t.Error(queue2.currIndex)
	}
	if queue2.nextIndex != 0 {
		t.Error(queue2.nextIndex)
	}
	if len(queue2.events) != 2 {
		t.Error(len(queue2.events))
	}
	if queue2.events[1].EventTypeID() != 1 {
		t.Error(queue2.events[1].EventTypeID())
	}
}

func TestNextEvent1(t *testing.T) {
	queue1 := newFilledDefaultQueue(1)

	event := queue1.NextEvent()
	if queue1.currIndex != 0 {
		t.Error(queue1.currIndex)
	}
	if queue1.nextIndex != 0 {
		t.Error(queue1.nextIndex)
	}
	if event.EventTypeID() != 0 {
		t.Error(event.EventTypeID())
	}
	if queue1.events[0] != nil {
		t.Error(queue1.events[0])
	}
}

func TestNextEvent2(t *testing.T) {
	queue2 := newFilledDefaultQueue(2)

	event := queue2.NextEvent()
	if queue2.currIndex != 1 {
		t.Error(queue2.currIndex)
	}
	if queue2.nextIndex != 0 {
		t.Error(queue2.nextIndex)
	}
	if event.EventTypeID() != 0 {
		t.Error(event.EventTypeID())
	}
	if queue2.events[0] != nil {
		t.Error(queue2.events[0])
	}
	if queue2.events[1] == nil {
		t.Error(queue2.events[0])
	}
	if queue2.events[1].EventTypeID() != 1 {
		t.Error(queue2.events[1].EventTypeID())
	}

	event = queue2.NextEvent()
	if queue2.currIndex != 0 {
		t.Error(queue2.currIndex)
	}
	if queue2.nextIndex != 0 {
		t.Error(queue2.nextIndex)
	}
	if event.EventTypeID() != 1 {
		t.Error(event.EventTypeID())
	}
	if queue2.events[0] != nil {
		t.Error(queue2.events[0])
	}
	if queue2.events[1] != nil {
		t.Error(queue2.events[0])
	}

	event = queue2.NextEvent()
	if queue2.currIndex != 0 {
		t.Error(queue2.currIndex)
	}
	if queue2.nextIndex != 0 {
		t.Error(queue2.nextIndex)
	}
	if event != nil {
		t.Error(event)
	}
}

func TestEnsureCapacity1(t *testing.T) {
	queue1 := newFilledDefaultQueue(1)

	queue1.PostEvent(NewEvent(1, 0))
	if queue1.currIndex != 0 {
		t.Error(queue1.currIndex)
	}
	if queue1.nextIndex != 2 {
		t.Error(queue1.nextIndex)
	}
	if len(queue1.events) <= 1 {
		t.Error(len(queue1.events))
	}
	if queue1.events[0].EventTypeID() != 0 {
		t.Error(queue1.events[0].EventTypeID())
	}
	if queue1.events[1].EventTypeID() != 1 {
		t.Error(queue1.events[1].EventTypeID())
	}
}

func TestEnsureCapacity2(t *testing.T) {
	queue2 := newFilledDefaultQueue(2)

	queue2.NextEvent()
	queue2.PostEvent(NewEvent(2, 0))
	queue2.PostEvent(NewEvent(3, 0))
	if queue2.currIndex != 0 {
		t.Error(queue2.currIndex)
	}
	if queue2.nextIndex != 3 {
		t.Error(queue2.nextIndex)
	}
	if len(queue2.events) <= 2 {
		t.Error(len(queue2.events))
	}
	if queue2.events[0] == nil {
		t.Error(queue2.events[0])
	}
	if queue2.events[1] == nil {
		t.Error(queue2.events[1])
	}
	if queue2.events[2] == nil {
		t.Error(queue2.events[1])
	}
	if queue2.events[3] != nil {
		t.Error(queue2.events[1])
	}
	if queue2.events[0].EventTypeID() != 1 {
		t.Error(queue2.events[0].EventTypeID())
	}
	if queue2.events[1].EventTypeID() != 2 {
		t.Error(queue2.events[1].EventTypeID())
	}
	if queue2.events[2].EventTypeID() != 3 {
		t.Error(queue2.events[2].EventTypeID())
	}
}

func newFilledDefaultQueue(capacity int) *DefaultEventQueue {
	queue := NewEventQueue(capacity).(*DefaultEventQueue)
	for i := 0; i < capacity; i++ {
		queue.PostEvent(NewEvent(i, 0))
	}
	return queue
}
