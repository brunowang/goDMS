package shared

import (
	"utils"
)

type Identity struct {
	prefix string
	name   string
	hash   uint
}

func (this *Identity) Hash() uint {
	return utils.Hash(this.prefix + this.name)
}
