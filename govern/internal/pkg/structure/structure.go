package structure

import (
	"github.com/jinzhu/copier"
)

// Temporary:

func Copy(to, from interface{}) error {
	return copier.Copy(to, from)
}
