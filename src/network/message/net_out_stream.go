package message

import (
	"utils"
)

type NetOutStream struct {
	buffer *NetBuffer
	offset int
}

func (this *NetOutStream) Init(buffer *NetBuffer) {
	this.buffer = buffer
}

func (this *NetOutStream) WriteBool(val bool) {
	if val == true {
		this.WriteByte(1)
	} else {
		this.WriteByte(0)
	}
}

func (this *NetOutStream) WriteByte(val byte) {
	this.buffer.Resize(this.offset + 1)
	this.buffer.Buffer[this.offset] = val
	this.offset += 1
	this.buffer.Length += 1
}

func (this *NetOutStream) WriteInt16(val int16) {
	arr := utils.Int16ToByte(val)
	this.WriteBytes(arr[:])
}

func (this *NetOutStream) WriteInt(val int) {
	arr := utils.IntToByte(val)
	this.WriteBytes(arr[:])
}

func (this *NetOutStream) WriteInt64(val int64) {
	arr := utils.Int64ToByte(val)
	this.WriteBytes(arr[:])
}

func (this *NetOutStream) WriteFloat(val float32) {
	arr := utils.FloatToByte(val)
	this.WriteBytes(arr[:])
}

func (this *NetOutStream) WriteFloat64(val float64) {
	arr := utils.Float64ToByte(val)
	this.WriteBytes(arr[:])
}

func (this *NetOutStream) WriteBytes(arr []byte) {
	this.buffer.Resize(this.offset + len(arr))
	this.buffer.Buffer = append(this.buffer.Buffer, arr...)
	this.offset += len(arr)
	this.buffer.Length += len(arr)
}
