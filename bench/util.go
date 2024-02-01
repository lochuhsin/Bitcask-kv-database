package test

import (
	"fmt"
	"os"
	"rebitcask/internal/settings"
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

func removeSegment() {
	os.RemoveAll(settings.Config.DATA_FOLDER_PATH)
}
