package main

import (
	"fmt"
	"testing"
	"time"
)

func TestPend(t *testing.T) {
	p := New()
	p.Go(func() error {
		fmt.Println("1")
		return nil
	})
	p.Go(func() error {
		fmt.Println("2")
		time.Sleep(250 * time.Millisecond)
		return nil
	})
	p.Go(func() error {
		fmt.Println("3")
		return nil
	})

	err := p.Wait()
	if err != nil {
		t.Fatal(err)
	}
}
