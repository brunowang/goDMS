package message

const PACKET_HEADER_LENGTH = 14

type NetPacketHeader struct {
	BodyLength int16
	MsgId      int16
	TimeStamp  int
	CheckSum   int
	Seq        int16
}

func (this *NetPacketHeader) Load(s *NetInStream) {
	this.BodyLength = s.ReadInt16()
	this.MsgId = s.ReadInt16()
	this.TimeStamp = s.ReadInt()
	this.CheckSum = s.ReadInt()
	this.Seq = s.ReadInt16()
}

func (this *NetPacketHeader) Save(s *NetOutStream) {
	s.WriteInt16(this.BodyLength)
	s.WriteInt16(this.MsgId)
	s.WriteInt(this.TimeStamp)
	s.WriteInt(this.CheckSum)
	s.WriteInt16(this.Seq)
}
