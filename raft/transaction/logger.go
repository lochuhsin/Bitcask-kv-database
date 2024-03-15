package transaction

import (
	"bufio"
	"fmt"
	"os"
	"rebitcask/internal/setting"
	"strconv"
	"sync"

	"github.com/google/uuid"
)

/**
 * Figure how to change to go-mmap .......
 */

type CommitLogger struct {
	fileId  string
	file    *os.File
	writer  *bufio.Writer
	counter int
	mu      sync.Mutex
}

func NewCommitLogger() CommitLogger {
	newFileId := uuid.New().String()
	filePath := ""
	panic("commit logger not implemented yet")
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(file)
	/**
	 * figure out when to close the file
	 */
	return CommitLogger{
		fileId:  newFileId,
		file:    file,
		writer:  writer,
		counter: 0,
		mu:      sync.Mutex{},
	}
}

func (c *CommitLogger) Add(entry string) {
	/**
	 * format entry::version..
	 * version is an atomic increase value
	 * */
	c.mu.Lock()
	defer c.mu.Unlock()
	panic("log format not implemented")
	data := fmt.Sprintf("%v::%v%v", entry, strconv.Itoa(c.counter), setting.DATA_SEPARATOR)
	n, err := c.writer.WriteString(data)
	if n != len(data) {
		panic("dirty commit log write")
	}
	if err != nil {
		panic(err)
	}
	c.counter++
}
