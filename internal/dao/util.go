package dao

import (
	"errors"
	"fmt"
	"rebitcask/internal/settings"
	"strconv"
	"strings"
)

func Serialize(k Entry) (string, error) {
	return fmt.Sprintf("CRC::%v::%v::%v", k.CreateTime, k.Key.Format(), k.Val.Format()), nil
}

func DeSerialize(line string) (Entry, error) {
	// CRC::TimeStamp::KeyDataType::KeyLen::Key::ValueDataType::ValueLen::Value
	strList := strings.Split(line, "::")
	// crc = strList[0] for validation usage
	timestamp := strList[1]
	keyDataType := strList[2]
	// KeyLen := strList[3] for validation usage
	key := strList[4]
	valueDataType := strList[5]
	// ValueLen := strList[6] for validation usage
	val := strList[7]

	keyData, err := toBaseType(keyDataType, key)
	if err != nil {
		panic(err)
	}

	valData, err := toBaseType(valueDataType, val)
	if err != nil {
		panic(err)
	}
	ts, err := strconv.Atoi(timestamp)
	if err != nil {
		panic(err)
	}
	return Entry{
		Key:        keyData.(NilString),
		Val:        valData,
		CreateTime: int64(ts),
	}, nil

}

func toBaseType(valtype, val string) (Base, error) {
	var data Base

	isNil := false
	if val == settings.ENV.NilData {
		isNil = true
	}

	switch valtype {
	case string(String):
		data = NilString{
			IsNil: isNil, Val: val,
		}
	case string(Int):
		i, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		data = NilInt{
			IsNil: isNil, Val: i,
		}
	case string(Byte):
		data = NilByte{
			IsNil: isNil, Val: byte(val[0]),
		}
	case string(Float):
		f, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return nil, err
		}
		data = NilFloat{
			IsNil: isNil, Val: f,
		}
	case string(Bool):
		b, err := strconv.ParseBool(val)
		if err != nil {
			return nil, err
		}
		data = NilBool{
			IsNil: isNil, Val: b,
		}
	case string(Tombstone):
		data = NilTomb{}
	default:
		return nil, errors.New("unsupported data type")
	}
	return data, nil
}
