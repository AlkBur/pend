package main

import (
	"sync"
)

type Pend struct {
	wg  sync.WaitGroup
	err []error

	errCh chan error
	done  chan interface{}
}

func New() *Pend {

	p := &Pend{
		err:   make([]error, 0),
		errCh: make(chan error, 1),
		done:  make(chan interface{}),
	}
	go p.Run()

	return p
}

func (p *Pend) Run() {
	for {
		select {
		case <-p.done:
			//Println("Done Pend")
			return
		case err := <-p.errCh:
			p.err = append(p.err, err)
		}
	}
}

func (p *Pend) Go(h func() error) {
	p.wg.Add(1)
	go func() {
		err := h()
		p.errCh <- err
		p.Done()
	}()
}

func (p *Pend) Wait() error {
	p.wg.Wait()
	p.done <- struct{}{}
	close(p.errCh)
	close(p.done)

	if len(p.err) > 0 {
		return p.err[0]
	}
	return nil
}

func (p *Pend) Done() {
	p.wg.Done()
}
