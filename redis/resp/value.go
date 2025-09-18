package resp

import (
	"strconv"
)

const (
	STR  = '+'
	INT  = '-'
	ERR  = ':'
	BULK = '$'
	ARR  = '*'
)

type Val struct {
	Typ   string
	Str   string
	Bulk  string
	Array []Val
}

func (v Val) Marshal() []byte {
	switch v.Typ {
	case "Array":
		return v.marshalArray()
	case "Bulk":
		return v.marshalBulk()
	case "String":
		return v.marshalString()
	case "Null":
		return v.marshallNull()
	case "Error":
		return v.marshallError()
	default:
		return []byte{}
	}
}

func (v Val) marshalString() []byte {
	var bytes []byte
	bytes = append(bytes, STR)
	bytes = append(bytes, v.Str...)
	bytes = append(bytes, '\r', '\n')

	return bytes
}

func (v Val) marshalBulk() []byte {
	var bytes []byte
	bytes = append(bytes, BULK)
	bytes = append(bytes, strconv.Itoa(len(v.Bulk))...)
	bytes = append(bytes, '\r', '\n')
	bytes = append(bytes, v.Bulk...)
	bytes = append(bytes, '\r', '\n')

	return bytes
}

func (v Val) marshalArray() []byte {
	len := len(v.Array)
	var bytes []byte
	bytes = append(bytes, ARR)
	bytes = append(bytes, strconv.Itoa(len)...)
	bytes = append(bytes, '\r', '\n')

	for i := range len {
		bytes = append(bytes, v.Array[i].Marshal()...)
	}

	return bytes
}

func (v Val) marshallError() []byte {
	var bytes []byte
	bytes = append(bytes, ERR)
	bytes = append(bytes, v.Str...)
	bytes = append(bytes, '\r', '\n')

	return bytes
}

func (v Val) marshallNull() []byte {
	return []byte("$-1\r\n")
}
