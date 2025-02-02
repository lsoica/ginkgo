package interrupthandler

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type InterruptHandler struct {
	interruptCount int
	lock           *sync.Mutex
	C              chan bool
}

func NewInterruptHandler() *InterruptHandler {
	h := &InterruptHandler{
		lock: &sync.Mutex{},
		C:    make(chan bool),
	}

	go h.handleInterrupt()
	SwallowSigQuit()

	return h
}

func (h *InterruptHandler) WasInterrupted() bool {
	h.lock.Lock()
	defer h.lock.Unlock()

	return h.interruptCount > 0
}

func (h *InterruptHandler) handleInterrupt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	<-c
	signal.Stop(c)

	h.lock.Lock()
	h.interruptCount++
	fmt.Printf("\nInterrupt %d out of 100\n", h.interruptCount)
	if h.interruptCount == 1 {
		close(h.C)
	} else if h.interruptCount > 100 {
		os.Exit(1)
	}
	h.lock.Unlock()

	go h.handleInterrupt()
}
