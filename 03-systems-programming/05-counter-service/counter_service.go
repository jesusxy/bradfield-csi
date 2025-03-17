// Author: Patch Neranartkomol

package counterservice

import (
	"sync"
	"sync/atomic"
)

type CounterService interface {
	// Returns values in ascending order; it should be safe to call
	// getNext() concurrently from multiple goroutines without any
	// additional synchronization on the caller's side.
	getNext() uint64
}

type UnsynchronizedCounterService struct {
	/* Please implement this struct and its getNext method */
	counter uint64
}

// getNext() - This one can be UNSAFE
func (usync *UnsynchronizedCounterService) getNext() uint64 {
	usync.counter++
	return usync.counter
}

type AtomicCounterService struct {
	/* Please implement this struct and its getNext method */
	counter uint64
}

// getNext() with sync/atomic
func (as *AtomicCounterService) getNext() uint64 {
	return atomic.AddUint64(&as.counter, 1)
}

type MutexCounterService struct {
	/* Please implement this struct and its getNext method */
	mut     sync.Mutex
	counter uint64
}

// getNext() with sync/Mutex
func (ms *MutexCounterService) getNext() uint64 {
	ms.mut.Lock()
	defer ms.mut.Unlock()
	ms.counter++
	return ms.counter
}

type ChannelCounterService struct {
	/* Please implement this struct and its getNext method */
	requestChan  chan struct{}
	responseChan chan uint64
}

// A constructor for ChannelCounterService
func newChannelCounterService() *ChannelCounterService {
	cs := &ChannelCounterService{
		requestChan:  make(chan struct{}),
		responseChan: make(chan uint64),
	}
	go cs.run()
	return cs
}

func (cs *ChannelCounterService) run() {
	var counter uint64
	for range cs.requestChan {
		counter++
		cs.responseChan <- counter
	}
}

// getNext() with goroutines and channels
func (cs *ChannelCounterService) getNext() uint64 {
	cs.requestChan <- struct{}{}
	return <-cs.responseChan
}
