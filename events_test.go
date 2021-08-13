/*
 *        Copyright 2019, 2021 Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package events

import "testing"

func TestNewQueue(t *testing.T) {
	queue1 := NewQueue(1).(*DefaultQueue)
	queue2 := NewQueue(2).(*DefaultQueue)

	if len(queue1.events) != 1 {
		t.Error(len(queue1.events))
	}
	if len(queue2.events) != 2 {
		t.Error(len(queue2.events))
	}
}

func TestPostEvent(t *testing.T) {
	queue1 := NewQueue(1).(*DefaultQueue)
	queue2 := NewQueue(2).(*DefaultQueue)

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
	if queue1.events[0].ID() != 0 {
		t.Error(queue1.events[0].ID())
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
	if queue2.events[1].ID() != 1 {
		t.Error(queue2.events[1].ID())
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
	if queue1.events[0].ID() != 0 {
		t.Error(queue1.events[0].ID())
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
	if queue2.events[1].ID() != 1 {
		t.Error(queue2.events[1].ID())
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
	if event.ID() != 0 {
		t.Error(event.ID())
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
	if event.ID() != 0 {
		t.Error(event.ID())
	}
	if queue2.events[0] != nil {
		t.Error(queue2.events[0])
	}
	if queue2.events[1] == nil {
		t.Error(queue2.events[0])
	}
	if queue2.events[1].ID() != 1 {
		t.Error(queue2.events[1].ID())
	}

	event = queue2.NextEvent()
	if queue2.currIndex != 0 {
		t.Error(queue2.currIndex)
	}
	if queue2.nextIndex != 0 {
		t.Error(queue2.nextIndex)
	}
	if event.ID() != 1 {
		t.Error(event.ID())
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
	if queue1.events[0].ID() != 0 {
		t.Error(queue1.events[0].ID())
	}
	if queue1.events[1].ID() != 1 {
		t.Error(queue1.events[1].ID())
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
	if queue2.events[0].ID() != 1 {
		t.Error(queue2.events[0].ID())
	}
	if queue2.events[1].ID() != 2 {
		t.Error(queue2.events[1].ID())
	}
	if queue2.events[2].ID() != 3 {
		t.Error(queue2.events[2].ID())
	}
}

func TestClose(t *testing.T) {
	queue2 := newFilledDefaultQueue(2)

	queue2.Close()
	if len(queue2.events) != 0 {
		t.Error(len(queue2.events))
	}
	if event := queue2.NextEvent(); event != nil {
		t.Error(event)
	}
}

func newFilledDefaultQueue(capacity int) *DefaultQueue {
	queue := NewQueue(capacity).(*DefaultQueue)
	for i := 0; i < capacity; i++ {
		queue.PostEvent(NewEvent(i, 0))
	}
	return queue
}
