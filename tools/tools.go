package tools

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"time"
)

type State string

const (
	queueSize       = 1024
	waitTimeForData = time.Second
)

var (
	ID        State = "ID"
	DATA_SIZE State = "DATA_SIZE"
	DATA      State = "DATA"
)

type Result struct {
	Element  *ElementData
	DataSize int64
	Data     []byte
	Err      error
}

// TODO: Accept channel to close/stop the parsing
func Parse(r io.Reader) chan Result {
	results := make(chan Result, queueSize)

	go func() {
		bfr := bufio.NewReader(r)
		state := ID
		var dataSize int64
		var el *ElementData
		var found bool
		var tag Result
		for {
			switch state {
			case ID:
				tag = Result{}
				_, length, err := ReadVintS(bfr)
				if err == io.EOF || err == io.ErrUnexpectedEOF {
					// Not enough bytes to read the VInt
					time.Sleep(waitTimeForData)
					continue
				} else if err != nil {
					// TODO: Send the error on the channel
					results <- Result{
						Err: fmt.Errorf("could not read tag ID: %s", err),
					}
					return
				}
				tagBytes, err := bfr.Peek(length)
				if err == io.EOF || err == io.ErrUnexpectedEOF {
					time.Sleep(waitTimeForData)
					continue
				} else if err != nil {
					results <- Result{
						Err: fmt.Errorf("could not read tag ID: %s", err),
					}
					return
				}
				discarded, err := bfr.Discard(length)
				if discarded < length {
					results <- Result{
						Err: fmt.Errorf("discarded only %d of %d bytes: %s", discarded, length, err),
					}
					return
				}
				idHex := ReadTagHex(tagBytes, 0, int64(length))

				state = DATA_SIZE
				el, found = Schema[idHex]
				if !found {
					fmt.Printf("Element %s not found\n", idHex)
				}
				tag.Element = el

			case DATA_SIZE:
				size, length, err := ReadVintS(bfr)
				if err == io.EOF || err == io.ErrUnexpectedEOF {
					time.Sleep(waitTimeForData)
					continue
				} else if err != nil {
					return
				}
				discarded, err := bfr.Discard(length)
				if discarded < length {
					results <- Result{
						Err: fmt.Errorf("discarded only %d of %d bytes: %s", discarded, length, err),
					}
					return
				}
				dataSize = int64(size)
				tag.DataSize = dataSize
				state = DATA

			case DATA:
				if found {
					switch el.Typ {
					case Master:
						tag.Element = el
						state = ID
					case String:
						fallthrough
					case UInteger:
						fallthrough
					case Binary:
						var read int = 0
						var n int
						var err error
						contentBytes := make([]byte, dataSize)
						n, err = io.ReadFull(bfr, contentBytes)
						for err == io.ErrUnexpectedEOF {
							read += n
							n, err = io.ReadFull(bfr, contentBytes[read:])
						}
						if err != nil {
							results <- Result{
								Err: fmt.Errorf("unexpected error :%s", err),
							}
							return
						}

						tag.Data = contentBytes
						state = ID
					}
					results <- tag
				} else {
					var read int64 = 0
					for read < dataSize {
						if _, err := bfr.ReadByte(); err != nil {
							results <- Result{
								Err: fmt.Errorf("unexpected error: %s", err),
							}
							return
						}
						read++
					}
					state = ID
				}
			}
		}
	}()

	return results
}

func ParseWhole(buffer []byte) {
	var offset int64 = 0
	state := ID
	var dataSize int64
	var el *ElementData
	var found bool
	for {
		switch state {
		case ID:
			_, len := ReadVint(buffer, offset)
			idHex := ReadTagHex(buffer, offset, offset+int64(len))
			fmt.Printf("ID=%s\n", idHex)
			offset += int64(len)
			state = DATA_SIZE
			el, found = Schema[idHex]

		case DATA_SIZE:
			size, len := ReadVint(buffer, offset)
			fmt.Printf("DATA_SIZE=%d\n", size)
			offset += int64(len)
			dataSize = int64(size)
			state = DATA

		case DATA:
			if found {
				switch el.Typ {
				case String:
					result := string(buffer[offset : offset+dataSize])
					fmt.Printf("DATA=[%s]\n", result)
					offset += dataSize
					state = ID
				case Master:
					fmt.Printf("%s[%s]\n", el.Name, el.Hex)
					state = ID
				case UInteger:
					data := ToUint64(buffer[offset : offset+dataSize])
					fmt.Printf("%s[%s]=%d\n", el.Name, el.Hex, data)
					offset += dataSize
					state = ID
				case Binary:
					// result := buffer[offset : offset+dataSize]
					// fmt.Printf("%+v\n", result)
					offset += dataSize
					state = ID
				}
			} else {
				offset += dataSize
				state = ID
			}
		}
	}
}

func ReadVint(b []byte, offset int64) (uint64, int) {
	length := LeadingZeros(b[offset])
	remaining := make([]byte, length)
	copy(remaining[0:], b[offset:])
	remaining[0] = remaining[0] & ((1 << (8 - length)) - 1)
	return ToUint64(remaining), length
}

func ReadVintS(r *bufio.Reader) (uint64, int, error) {
	firstByte, err := r.Peek(1)
	if err != nil {
		return 0, 0, err
	}
	length := LeadingZeros(firstByte[0])
	remaining, err := r.Peek(length)
	if err != nil {
		return 0, 0, err
	}
	result := make([]byte, length)
	copy(result, remaining)
	result[0] = result[0] & ((1 << (8 - length)) - 1)
	return ToUint64(result), length, nil
}

func ReadTagHex(buffer []byte, start, end int64) string {
	data := ToUint64(buffer[start:end])
	return ToHex(data)
}

func ToHex(data uint64) string {
	return fmt.Sprintf("%x", data)
}

func LeadingZeros(b byte) int {
	r := 8 - math.Floor(math.Log2(float64(b)))
	return int(r)
}

func ToUint64(bytes []byte) uint64 {
	// binary.BigEndian.Uint32()
	// return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
	var result uint64 = 0
	offset := 0
	for i := len(bytes) - 1; i >= 0; i-- {
		result |= (uint64(bytes[i]) << offset)
		offset += 8
	}
	return result
}

func ToInt64(bytes []byte) int64 {
	// return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
	var result int64 = 0
	offset := 0
	for i := len(bytes) - 1; i >= 0; i-- {
		result |= int64(bytes[i]) << offset
		offset += 8
	}
	return result
}

func UInt64ToByte(n uint64) []byte {
	result := make([]byte, 0)
	offset := 0
	// i := 0
	for {
		x := byte(n >> offset)
		if x == 0 {
			break
		}
		result = append(result, 255&x)
		offset += 8
	}
	if len(result) == 0 {
		result = []byte{0}
	}
	reverseArr(result)
	return result
}

func reverseArr(arr []byte) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
