package resp

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type Resp struct {
	reader *bufio.Reader
}

func NewResp(rd io.Reader) *Resp {
	return &Resp{reader: bufio.NewReader(rd)}
}

func (r *Resp) readLine() (line []byte, n int, err error) {
	for {
		b, err := r.reader.ReadByte()
		if err != nil {
			return nil, 0, err
		}
		n += 1
		line = append(line, b)
		if len(line) >= 2 && line[len(line)-2] == '\r' {
			break
		}
	}

	return line[:len(line)-2], n, nil
}

func (r *Resp) readInt() (x int, n int, err error) {
	line, n, err := r.readLine()
	if err != nil {
		return 0, 0, err
	}

	i64, err := strconv.ParseInt(string(line), 10, 64)
	if err != nil {
		return 0, n, err
	}
	return int(i64), n, nil
}

func (r *Resp) Read() (Val, error) {
	_type, err := r.reader.ReadByte()
	if err != nil {
		return Val{}, err
	}

	switch _type {
	case ARR:
		return r.readArr()
	case BULK:
		return r.readBulk()
	default:
		fmt.Printf("🔴 Unknown type: %v", string(_type))
		return Val{}, nil
	}
}

func (r *Resp) readArr() (Val, error) {
	v := Val{}
	v.Typ = "Array"

	length, _, err := r.readInt()
	if err != nil {
		return Val{}, err
	}

	v.Array = make([]Val, length)
	for i := range length {
		val, err := r.Read()
		if err != nil {
			return v, err
		}

		v.Array[i] = val
	}

	return v, nil
}

func (r *Resp) readBulk() (Val, error) {
	v := Val{}
	v.Typ = "Bulk"

	len, _, err := r.readInt()
	if err != nil {
		return v, err
	}

	bulk := make([]byte, len)
	_, err = r.reader.Read(bulk)
	if err != nil {
		fmt.Println("🔴 Failed to read with reader")
	}
	v.Bulk = string(bulk)

	_, _, err = r.readLine()
	if err != nil {
		fmt.Println("🔴 Failed to read line")
	}

	return v, nil
}
