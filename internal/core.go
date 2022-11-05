package internal

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"os"
	"rebitcask/internal/models"
	"sort"
	"strconv"
	"strings"
	"time"
)

/*
To be noticed, file.Sync() doesn't actually sync
the file on MACOS system. This needs to be tested on linux environment.
If still, doesn't work, needs to implement a buffer to handle this situation...etc
*/
func toSegment(memory models.MemoryModel, segContainer *SegmentContainer) error {

	segID := uuid.New().String()

	filepath := fmt.Sprintf("%v%v/%v.log", ENVVAR.logFolder, ENVVAR.segmentFolder, segID)
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	writer := bufio.NewWriter(file)

	if err != nil {
		return err
	}

	// create new Segment
	kvPairs := memory.GetAll()
	segHead := kvPairs[0].Key // TODO: save seg end and store it in tree, performance will way better

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
	segContainer.memo.Set(segHead, models.Item{
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

// currently, compress function will compress the entire history
// this can be optimized
// Redesign tree entry to save, key: value, create_time, additional_attribute: map
func compressSegments(segContainer *SegmentContainer) (newSegContainer SegmentContainer) {
	allKVPair := segContainer.memo.GetAll()

	// initialize newSegContainer
	newSegContainer.Init()

	// Sort all KV pair by CreateTime, with backwords
	sort.SliceStable(allKVPair, func(i, j int) bool {
		return allKVPair[i].Val.CreateTime >= allKVPair[j].Val.CreateTime
	})

	// this can be merged using merge sort like method
	// and using multi threading TODO: Optimize this entire algorithm for performance
	keyvalue := make(map[string]string)
	keyArr := make([]string, 0)
	for _, kvPair := range allKVPair {
		segID := string(kvPair.Val.Val)

		filepath := fmt.Sprintf("%v%v/%v.log", ENVVAR.logFolder, ENVVAR.segmentFolder, segID)
		file, _ := os.Open(filepath)

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()

			strArr := strings.Split(line, ",")
			key, val := strArr[0], strArr[1]

			if _, ok := keyvalue[key]; !ok {
				keyvalue[key] = val
			}

			keyArr = append(keyArr, key)
		}

		file.Close()
		_ = os.RemoveAll(filepath)
	}

	// sort keys
	sort.Strings(keyArr)

	segID := uuid.New().String()
	segFilePath := fmt.Sprintf("%v%v/%v.log", ENVVAR.logFolder, ENVVAR.segmentFolder, segID)
	segFile, _ := os.Create(segFilePath)
	writer := bufio.NewWriter(segFile)

	segHead := keyArr[0]

	for i, key := range keyArr {
		val, _ := keyvalue[key]
		// file structure for every line is key,val
		line := fmt.Sprintf("%v,%v", key, val)
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println(line)
			fmt.Println("Error while writing to segment !!!!!", err)
			panic(1)
		}

		if i%ENVVAR.fileLineLimit == 0 && i != 0 {

			// close file
			_ = writer.Flush()
			segFile.Sync()
			segFile.Close()

			// store new segment
			newSegContainer.memo.Set(segHead, models.Item{
				Val:        []byte(segID),
				CreateTime: strconv.FormatInt(time.Now().UnixNano(), 2),
			})
			if i < len(keyArr)-1 {
				segHead = keyArr[i+1]
			}

			// Update new file
			segFilePath = fmt.Sprintf("%v%v/%v.log", ENVVAR.logFolder, ENVVAR.segmentFolder, uuid.New().String())
			segFile, _ = os.Create(segFilePath)
			writer = bufio.NewWriter(segFile)
		}
	}

	return newSegContainer
}
