package message

const (
	PACKAGE_INIT_LENGTH      int = 64
	PACKAGE_INIT_HALF_LENGTH int = 64 / 2
)

type NetBuffer struct {
	Buffer []byte
	Length int
}

func (this *NetBuffer) Init(size int) {
	this.Buffer = make([]byte, size)
	this.Length = 0
}

func (this *NetBuffer) Resize(size int) {
	if size <= len(this.Buffer) {
		return
	}
	next := len(this.Buffer)
	if len(this.Buffer) == 0 {
		next = PACKAGE_INIT_HALF_LENGTH
	}
	for {
		next *= 2
		if next >= size {
			break
		}
	}
	newBuffer := make([]byte, next)
	if this.Length != 0 {
		copy(newBuffer, this.Buffer)
	}
	this.Buffer = newBuffer
}
