package message

import (
	"utils"
)

type NetInStream struct {
	buffer *NetBuffer
	offset int
}

func (this *NetInStream) Init(buffer *NetBuffer) {
	this.buffer = buffer
}

func (this *NetInStream) ReadBool() bool {
	this.Skip(1)
	return this.buffer.Buffer[this.offset-1] == 1
}

func (this *NetInStream) ReadByte() byte {
	this.Skip(1)
	return this.buffer.Buffer[this.offset-1]
}

func (this *NetInStream) ReadInt16() int16 {
	this.Skip(2)
	return utils.ByteToInt16(this.buffer.Buffer, this.offset-2)
}

func (this *NetInStream) ReadInt() int {
	this.Skip(4)
	return utils.ByteToInt(this.buffer.Buffer, this.offset-4)
}

func (this *NetInStream) ReadInt64() int64 {
	this.Skip(8)
	return utils.ByteToInt64(this.buffer.Buffer, this.offset-8)
}

func (this *NetInStream) ReadFloat() float32 {
	this.Skip(4)
	return utils.ByteToFloat(this.buffer.Buffer, this.offset-4)
}

func (this *NetInStream) ReadFloat64() float64 {
	this.Skip(8)
	return utils.ByteToFloat64(this.buffer.Buffer, this.offset-8)
}

func (this *NetInStream) ReadBytes(size int) []byte {
	arr := this.buffer.Buffer[this.offset : this.offset+size]
	this.Skip(size)
	return arr
}

func (this *NetInStream) Skip(size int) {
	if this.offset+size > this.buffer.Length {
		panic("stream out of range")
	}
	this.offset += size
}
