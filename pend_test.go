package main

import (
	"fmt"
	"testing"
	"time"
)

func TestPend(t *testing.T) {
	pend := NewPend()
	pend.Go(func() error {
		fmt.Println("1")
		return nil
	})
	pend.Go(func() error {
		fmt.Println("2")
		time.Sleep(250 * time.Millisecond)
		return nil
	})
	pend.Go(func() error {
		fmt.Println("3")
		return nil
	})

	err := pend.Wait()
	if err != nil {
		t.Fatal(err)
	}
}
