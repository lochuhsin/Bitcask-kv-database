package internal

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"os"
	"rebitcask/internal/models"
	"strconv"
	"strings"
	"time"
)

/*
To be noticed, file.Sync() doesn't actually sync
the file on MACOS system. This needs to be tested on linux environment.
If still, doesn't work, needs to implement a buffer to handle this situation...etc
*/
func toSegment(memory *models.BinarySearchTree, segContainer *SegmentContainer) error {

	segID := uuid.New().String()

	filepath := fmt.Sprintf("%v%v/%v.log", ENVVAR.logFolder, ENVVAR.segmentFolder, segID)
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	writer := bufio.NewWriter(file)

	if err != nil {
		return err
	}

	// create new Segment
	kvPairs := memory.GetAll()
	currentSeg := SegmentMap{segID: segID}
	currentSeg.segHead = kvPairs[0].Key // TODO: save seg end and store it in tree, performance will way better

	for _, pair := range kvPairs {
		key := pair.Key
		item := pair.Val

		// file structure for every line is key,val
		line := fmt.Sprintf("%v,%v", key, item.Val)
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error while writing to segment !!!!!")
			panic(1)
		}
	}

	// this will cause performance issue,
	// change all memory model to generic type or at least using interface
	segContainer.memo.Set(currentSeg.segHead, models.Item{
		Val:        []byte(segID),
		CreateTime: strconv.FormatInt(time.Now().UnixNano(), 2),
	})
	_ = writer.Flush()
	file.Sync()
	file.Close()
	return nil
}

// search all sub tree in memory tree
func isKeyInSegments(k *string, segContainer *SegmentContainer) (v []byte, status bool) {
	kvPairs := segContainer.memo.GetAllValueUnder(k)

	for _, pair := range kvPairs {
		segID := string(pair.Val.Val)
		filepath := fmt.Sprintf("%v%v/%v.log", ENVVAR.logFolder, ENVVAR.segmentFolder, segID)
		file, err := os.Open(filepath)

		if err != nil {
			panic("Something wrong with reading file")
		}

		sc := bufio.NewScanner(file)

		for sc.Scan() {
			line := sc.Text()
			stringList := strings.Split(line, ",")
			key, v := stringList[0], stringList[1]

			// TODO: optimize this without parinsg string to []byte
			if key == *k {
				return []byte(v), true
			}
		}
	}
	return []byte(""), false
}

type keyPosPair struct {
	key string
	pos int
}

// currently, compress function will compress the entire history
// this can be optimized
// Redesign tree entry to save, key: value, create_time, additional_attribute: map
//func compressSegments(segContainer *SegmentContainer) (newSegContainer SegmentContainer) {
//	allKVPair := segContainer.memo.GetAll()
//	return newSegContainer
//}
