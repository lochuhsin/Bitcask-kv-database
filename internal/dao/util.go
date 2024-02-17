package dao

import (
	"errors"
	"rebitcask/internal/settings"
	"rebitcask/internal/util"
	"strconv"
	"strings"
)

// CRC::TimeStamp::KeyDataType::KeyLen::Key::ValueDataType::ValueLen::Value
func Serialize(k Entry) (string, error) {
	var builder strings.Builder
	builder.WriteString("CRC::")
	builder.WriteString(strconv.Itoa(int(k.CreateTime)))
	builder.WriteString("::")
	builder.WriteString(string(String))
	builder.WriteString("::")
	builder.WriteString(strconv.Itoa(len(k.Key)))
	builder.WriteString("::")
	builder.Write(k.Key)
	builder.WriteString("::")
	builder.WriteString(util.BytesToString(k.Val.Format()))
	return builder.String(), nil
}

func DeSerialize(line string) (Entry, error) {
	// CRC::TimeStamp::KeyDataType::KeyLen::Key::ValueDataType::ValueLen::Value
	strList := strings.Split(line, "::")
	// crc = strList[0] for validation usage
	timestamp := strList[1]
	// keyDataType := strList[2]
	// KeyLen := strList[3] for validation usage
	key := strList[4]
	valueDataType := strList[5]
	// ValueLen := strList[6] for validation usage
	val := strList[7]

	valData, err := toBaseType(valueDataType, val)
	if err != nil {
		panic(err)
	}
	ts, err := strconv.Atoi(timestamp)
	if err != nil {
		panic(err)
	}
	return Entry{
		Key:        util.StringToBytes(key),
		Val:        valData,
		CreateTime: int64(ts),
	}, nil

}

func toBaseType(valType, val string) (Base, error) {
	var data Base

	isNil := false
	if val == settings.Config.NIL_DATA_REP {
		isNil = true
	}

	switch valType {
	case string(String):
		data = NilString{
			IsNil: isNil, Val: util.StringToBytes(val),
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
