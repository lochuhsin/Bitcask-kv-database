package internal

import (
	"bufio"
	"fmt"
	"os"
	"rebitcask/internal/models"
	"sort"
	"strconv"
)

/*
To be noticed, file.Sync() doesn't actually sync
the file on MACOS system. This needs to be tested on linux environment.
If still, doesn't work, needs to implement a buffer to handle this situation...etc
*/
func toSegment(memory *models.BinarySearchTree, segContainer *SegmentContainer) error {
	filepath := fmt.Sprintf("%v%v/%v.log", ENVVAR.logFolder, ENVVAR.segmentFolder)
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	writer := bufio.NewWriter(file)

	if err != nil {
		return err
	}

	kvPairs := memory.GetAll()

	// create new Segment
	currentSeg := SegmentMap{segID: segContainer.segCount}
	currentSeg.segHead = kvPairs[0].Key
	currentSeg.segEnd = kvPairs[len(kvPairs)-1].Key

	for _, pair := range memory.GetAll() {
		key := pair.Key
		val := pair.Val

		// file structure for every line is key,val
		line := fmt.Sprintf("%v,%v", key, val)
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error while writing to segment !!!!!")
			panic(1)
		}
	}

	segContainer.memo.Set(currentSeg.segHead, []byte(strconv.Itoa(currentSeg.segID)))
	writer.Flush()
	file.Sync()
	file.Close()
	return nil
}

func isKeyInSegment(k string, segment *SegmentMap) (v []byte, status bool) {
}

type keyPosPair struct {
	key string
	pos int
}

// currently, compress function will compress the entire history
// this can be optimized
func compressSegments(segContainer *SegmentContainer) (newSegContainer SegmentContainer) {
	keyValue := make(map[string][]byte)

	// reading in reverse order, since the larger the later
	for segIndex := segContainer.segCount - 1; segIndex >= 0; segIndex-- {
		segment := segContainer.memo[segIndex]

		// Sort segment hashmap by byte position
		// since we need to read everything reverse, so the order is extremely important
		bytePositionMap := segment.bytePositionMap

		// using struct is way more efficient than other datatype
		keyPosPairArr := make([]keyPosPair, len(bytePositionMap))
		i := 0
		for key, pos := range bytePositionMap {
			keyPosPairArr[i] = keyPosPair{key: key, pos: pos}
			i += 1
		}

		sort.Slice(keyPosPairArr, func(i int, j int) bool {
			return keyPosPairArr[i].pos > keyPosPairArr[j].pos
		})

		filepath := fmt.Sprintf("%v%v/%v.log", ENVVAR.logFolder, ENVVAR.segmentFolder, segment.CurrentSegmentNo)
		file, _ := os.Open(filepath)
		for _, pair := range keyPosPairArr {
			pos, key := pair.pos, pair.key
			if _, ok := keyValue[key]; ok {
				continue
			}
			keyValue[key] = seekFile(file, pos, segment.byteLengthMap[key])
		}
		os.Remove(filepath)
	}

	//keyvalue contains all values, create a start looping and create a new segment
	var newMemoMap models.Hash
	newMemoMap.Init()
	newMemoMap.SetMemory(keyValue)

	newSegContainer = SegmentContainer{
		memo:     []SegmentMap{},
		segCount: 0,
	}
	tempSegment := SegmentMap{
		bytePositionMap:  make(map[string]int),
		byteLengthMap:    make(map[string]int),
		byteFileLength:   0,
		CurrentSegmentNo: 0,
	}
	// Since we are compressing
	// the maximum segment number will be less than currentSegNo (even in worst case)
	// the key in each segment will be unique, so the order of writing
	// to disk is irrelevant
	err := toSegment(&newMemoMap, &tempSegment, &newSegContainer)
	if err != nil {
		panic("something went wrong while compressing")
	}

	// we cannot ensure there are no segment left.
	if tempSegment.byteFileLength != 0 {
		newSegContainer.memo = append(newSegContainer.memo, tempSegment)
		newSegContainer.segCount++
	}
	return newSegContainer
}
