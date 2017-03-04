package tree

import (
	"sync"
	"github.com/polaris1119/bitmap"
	"github.com/h4ck3rm1k3/gogccintro/models"
)

type TreeMap struct {
	Bitmap * bitmap.Bitmap
	Nodes map[int] * models.GccTuParserNode
	Mutex sync.RWMutex
}

func NewTreeMap(size int) *TreeMap {
	return &TreeMap{
		Mutex: sync.RWMutex{},
		Bitmap: bitmap.NewBitmapSize(size),
		Nodes : make(map[int] * models.GccTuParserNode),
	}
}

// return true if we are the first to set this bit, false otherwise
func (this* TreeMap) SetBitFirst( pos uint64 ) bool {
	if this.Bitmap.GetBit(pos)==0  {
		this.Mutex.RLock()
		this.Bitmap.SetBit(pos,1)
		this.Mutex.RUnlock()
		return true
	}
	return false
}
