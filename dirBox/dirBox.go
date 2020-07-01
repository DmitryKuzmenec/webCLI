package dirBox

import (
	"sync"

	"github.com/gobuffalo/packr"
)

const dir = "../common"

var box packr.Box
var once sync.Once

func NewBox() packr.Box {
	once.Do(func() { box = packr.NewBox(dir) })
	return box
}
