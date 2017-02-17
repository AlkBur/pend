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

/*
========================================
	BENCHMARKS!!
========================================
	# go test -bench=.
*/

func BenchmarkWithPend(b *testing.B) {
	p := New()
	for i := 0; i < b.N; i++ {
		p.Go(TestingFunc)
	}
	err := p.Wait()
	if err != nil {
		b.Fatal(err)
	}
	//b.Log("OK")
}

func BenchmarkWithoutPend(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		err = TestingFunc()
		if err != nil {
			b.Fatal(err)
		}
	}
	//b.Log("OK")
}

func TestingFunc() (err error) {
	time.Sleep(250 * time.Millisecond)
	return
}
