package buffer

import (
	"encoding/binary"
	"github.com/dobyte/due/v2/errors"
	"math"
)

type Reader struct {
	buf []byte
	off int
}

func NewReader(data []byte) *Reader {
	return &Reader{buf: data}
}

// Reset 重置
func (r *Reader) Reset() {
	r.off = 0
}

// ReadBool 读取bool值
func (r *Reader) ReadBool() (bool, error) {
	buf, err := r.slice(1)
	if err != nil {
		return false, err
	}

	return buf[0] == 1, nil
}

// ReadInt8 读取int8值
func (r *Reader) ReadInt8() (int8, error) {
	buf, err := r.slice(1)
	if err != nil {
		return 0, err
	}

	return int8(buf[0]), nil
}

// ReadUint8 读取uint8值
func (r *Reader) ReadUint8() (uint8, error) {
	buf, err := r.slice(1)
	if err != nil {
		return 0, err
	}

	return buf[0], nil
}

// ReadInt16 读取int16值
func (r *Reader) ReadInt16(order binary.ByteOrder) (int16, error) {
	buf, err := r.slice(2)
	if err != nil {
		return 0, err
	}

	return int16(order.Uint16(buf)), nil
}

// ReadUint16 读取uint16值
func (r *Reader) ReadUint16(order binary.ByteOrder) (uint16, error) {
	buf, err := r.slice(2)
	if err != nil {
		return 0, err
	}

	return order.Uint16(buf), nil
}

// ReadInt32 读取int32值
func (r *Reader) ReadInt32(order binary.ByteOrder) (int32, error) {
	buf, err := r.slice(4)
	if err != nil {
		return 0, err
	}

	return int32(order.Uint32(buf)), nil
}

// ReadUint32 读取uint32值
func (r *Reader) ReadUint32(order binary.ByteOrder) (uint32, error) {
	buf, err := r.slice(4)
	if err != nil {
		return 0, err
	}

	return order.Uint32(buf), nil
}

// ReadInt64 读取int64值
func (r *Reader) ReadInt64(order binary.ByteOrder) (int64, error) {
	buf, err := r.slice(8)
	if err != nil {
		return 0, err
	}

	return int64(order.Uint64(buf)), nil
}

// ReadUint64 读取uint64值
func (r *Reader) ReadUint64(order binary.ByteOrder) (uint64, error) {
	buf, err := r.slice(8)
	if err != nil {
		return 0, err
	}

	return order.Uint64(buf), nil
}

// ReadFloat32 读取float32值
func (r *Reader) ReadFloat32(order binary.ByteOrder) (float32, error) {
	buf, err := r.slice(4)
	if err != nil {
		return 0, err
	}

	return math.Float32frombits(order.Uint32(buf)), nil
}

// ReadFloat64 读取float64值
func (r *Reader) ReadFloat64(order binary.ByteOrder) (float64, error) {
	buf, err := r.slice(8)
	if err != nil {
		return 0, err
	}

	return math.Float64frombits(order.Uint64(buf)), nil
}

// ReadRune 读取rune值
func (r *Reader) ReadRune(order binary.ByteOrder) (rune, error) {
	return r.ReadInt32(order)
}

// ReadByte 读取byte值
func (r *Reader) ReadByte() (byte, error) {
	return r.ReadUint8()
}

// ReadString 读取string值
func (r *Reader) ReadString(n int) (string, error) {
	buf, err := r.slice(n)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

func (r *Reader) slice(n int) ([]byte, error) {
	if r.off+n > len(r.buf) {
		return nil, errors.ErrUnexpectedEOF
	}

	buf := r.buf[r.off : r.off+n]
	r.off += n

	return buf, nil
}
