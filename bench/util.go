package test

import (
	"fmt"
	"os"
	"rebitcask/internal/setting"
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
	os.RemoveAll(setting.Config.DATA_FOLDER_PATH)
}
