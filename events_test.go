/*
 *          Copyright 2019, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package events

import "testing"

func TestEnsureCapacity(t *testing.T) {
	queue := NewEventQueue(1).(*DefaultEventQueue)
	queue.PostEvent(NewEvent(0, 0))

	if queue.currIndex != 0 {
		t.Error(queue.currIndex)
	}
	if queue.nextIndex != 0 {
		t.Error(queue.nextIndex)
	}
	if len(queue.events) != 1 {
		t.Error(len(queue.events))
	}
	if queue.events[0].EventTypeID() != 0 {
		t.Error(queue.events[0].EventTypeID())
	}

	queue.PostEvent(NewEvent(1, 0))
	if queue.currIndex != 0 {
		t.Error(queue.currIndex)
	}
	if queue.nextIndex != 2 {
		t.Error(queue.nextIndex)
	}
	if len(queue.events) <= 1 {
		t.Error(len(queue.events))
	}
	if queue.events[0].EventTypeID() != 0 {
		t.Error(queue.events[0].EventTypeID())
	}
	if queue.events[1].EventTypeID() != 1 {
		t.Error(queue.events[1].EventTypeID())
	}
}

func TestNextEvent(t *testing.T) {
	queue := NewEventQueue(10).(*DefaultEventQueue)
	for i := 0; i < 10; i++ {
		queue.PostEvent(NewEvent(i, 0))
	}
	if queue.currIndex != 0 {
		t.Error(queue.currIndex)
	}
	if queue.nextIndex != 0 {
		t.Error(queue.nextIndex)
	}
	if queue.NextEvent().EventTypeID() != 0 {
		t.Error(queue.NextEvent().EventTypeID())
	}
	if queue.currIndex != 1 {
		t.Error(queue.currIndex)
	}
	for i := 1; i < 5; i++ {
		if queue.NextEvent().EventTypeID() != i {
			t.Error(i, queue.NextEvent().EventTypeID())
		}
	}
	queue.PostEvent(NewEvent(10, 0))
	if queue.currIndex != 5 {
		t.Error(queue.currIndex)
	}
	if queue.nextIndex != 1 {
		t.Error(queue.nextIndex)
	}
	if len(queue.events) != 10 {
		t.Error(len(queue.events))
	}
	if queue.events[queue.currIndex].EventTypeID() != 5 {
		t.Error(queue.events[0].EventTypeID())
	}
	if queue.events[0].EventTypeID() != 10 {
		t.Error(queue.events[0].EventTypeID())
	}
	for i := 1; i < 5; i++ {
		queue.PostEvent(NewEvent(10+i, 0))
		if queue.nextIndex != i+1 {
			t.Error(i+1, queue.nextIndex)
		}
		if queue.events[5].EventTypeID() != 5 {
			t.Error(i, 5, queue.events[5].EventTypeID())
		}
		if queue.events[i].EventTypeID() != 10+i {
			t.Error(10+i, queue.events[i].EventTypeID())
		}
	}
	if queue.currIndex != 5 {
		t.Error(queue.currIndex)
	}
	if queue.nextIndex != 5 {
		t.Error(queue.nextIndex)
	}
	if len(queue.events) != 10 {
		t.Error(len(queue.events))
	}
	if queue.events[queue.currIndex].EventTypeID() != 5 {
		t.Error(queue.events[queue.currIndex].EventTypeID())
	}
	for i := 0; i < 5; i++ {
		if queue.events[i].EventTypeID() != i+10 {
			t.Error(i+10, queue.events[i].EventTypeID())
		}
		if queue.events[i+5].EventTypeID() != i+5 {
			t.Error(i+5, queue.events[i+5].EventTypeID())
		}
	}
	queue.PostEvent(NewEvent(15, 0))
	if len(queue.events) == 10 {
		t.Error(len(queue.events))
	}
	if queue.currIndex != 0 {
		t.Error(queue.currIndex)
	}
	if queue.nextIndex != 11 {
		t.Error(queue.nextIndex)
	}
	for i := 0; i < 10; i++ {
		if queue.events[i].EventTypeID() != i+5 {
			t.Error(i+5, queue.events[i].EventTypeID())
		}
	}
}
