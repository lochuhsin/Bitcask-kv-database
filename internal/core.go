package internal

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func toDisk(memory *memoryMap, currSeg *SegmentMap, segContainer *SegmentContainer) error {
	filepath := fmt.Sprintf("%v%v/%v.log", LOGFOLDER, SEGMENTFOLDER, currSeg.CurrentSegmentNo)
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}

	byteHeadPosition := currSeg.byteFileLength
	for k, v := range memory.keyvalue {

		byteValue := v
		bytes, err := file.Write(byteValue)

		if err != nil {
			panic("Something went wrong while writing to disk")
		}

		currSeg.byteLengthMap[k] = bytes
		currSeg.bytePositionMap[k] = byteHeadPosition

		byteHeadPosition += bytes

		if byteHeadPosition >= FILEBYTELIMIT {
			// close file
			file.Close()

			// store segment
			segContainer.memo = append(segContainer.memo, *currSeg)

			// create new segment
			newSegmentNo := currSeg.CurrentSegmentNo + 1
			file, *currSeg = createNewSegment(newSegmentNo)

			// new bytehead
			byteHeadPosition = 0
		}
	}
	currSeg.byteFileLength = byteHeadPosition
	file.Close()
	return nil
}

func isKeyInSegment(k string, segment *SegmentMap) (v []byte, status bool) {

	if _, ok := segment.bytePositionMap[k]; !ok {
		return []byte(""), false
	}

	filepath := fmt.Sprintf("%v%v/%v.log", LOGFOLDER, SEGMENTFOLDER, segment.CurrentSegmentNo)
	bytePos, _ := segment.bytePositionMap[k]
	byteLen, _ := segment.byteLengthMap[k]

	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		return []byte("Something went wrong while opening file"), false
	}

	_, err = file.Seek(int64(bytePos), io.SeekStart)
	if err != nil {
		panic("Something went wrong while seeking file")
	}

	readByte := make([]byte, byteLen)

	_, err = file.Read(readByte)
	if err != nil {
		panic("Something went wrong while reading file")
	}

	return readByte, true
}

type keyPosPair struct {
	key string
	pos int
}

func compressSegments(segments []SegmentMap) (newSegments []SegmentMap) {
	keyValue := make(map[string][]byte)

	// reading reverse order
	for segIndex := len(segments) - 1; segIndex >= 0; segIndex-- {
		segment := segments[segIndex]

		// Sort segment hashmap by byte position
		// since we need to read everything reverse, so the order is extremely important
		bytePositionMap := segment.bytePositionMap

		// using struct is way more efficient than other datatype
		keyPosPairArr := make([]keyPosPair, len(bytePositionMap))
		i := 0
		for key, pos := range bytePositionMap {
			keyPosPairArr[i] = keyPosPair{key: key, pos: pos}
		}

		sort.Slice(keyPosPairArr, func(i int, j int) bool {
			return keyPosPairArr[i].pos > keyPosPairArr[j].pos
		})

		// Since the key, pos is sort with reverse order,
		// we can loop everything and write it in to temprary memory
		filepath := fmt.Sprintf("%v%v/%v.log", LOGFOLDER, SEGMENTFOLDER, segment.CurrentSegmentNo)
		file, _ := os.Open(filepath)
		for _, pair := range keyPosPairArr {
			pos := pair.pos
			key := pair.key

			if _, ok := keyValue[key]; ok {
				continue
			}

			_, err := file.Seek(int64(pos), io.SeekStart)
			if err != nil {
				panic("Something went wrong while seeking file")
			}

			readByte := make([]byte, segment.byteLengthMap[key])

			_, err = file.Read(readByte)
			if err != nil {
				panic("Something went wrong while seeking file")
			}

			keyValue[key] = readByte
		}
		//TODO: Remove segment folder after finish reading file
	}

	//keyvalue contains all values, create a start looping and create a new segment
	memoMap := memoryMap{keyvalue: keyValue}
	newSegContainer := SegmentContainer{
		memo: []SegmentMap{},
	}
	tempSegment := SegmentMap{
		bytePositionMap:  make(map[string]int),
		byteLengthMap:    make(map[string]int),
		byteFileLength:   0,
		CurrentSegmentNo: 0,
	}
	// Since we are compressing
	// the maximum segment number will be less than currentSegNo (for worst case)
	// the key in each segment will be unique, so the order of writing
	// to disk is irrelevant

	err := toDisk(&memoMap, &tempSegment, &newSegContainer)
	if err != nil {
		panic("something went wrong while compressing")
	}
	// add the temp Segment to newSegContainer if length is not zero
	// we cannot ensure there are no segment left.
	if tempSegment.byteFileLength != 0 {
		newSegContainer.memo = append(newSegContainer.memo, tempSegment)
	}
	return newSegContainer.memo
}
