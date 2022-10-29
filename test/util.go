package test

import (
	"fmt"
	"testing"
	"time"
)

func testTimer(fn func(t *testing.T), t *testing.T) {
	s := time.Now()
	fmt.Println(s)

	fn(t)

	timeLength := time.Since(s)
	fmt.Println("test finished")
	fmt.Printf("Cost: %v", timeLength)
}
