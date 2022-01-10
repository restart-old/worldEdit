package worldedit

import (
	"sync"

	"github.com/df-mc/dragonfly/server/item"
)

var wandItem item.Stack
var wandMu sync.RWMutex

func SetWand(i item.Stack) {
	wandMu.Lock()
	defer wandMu.Unlock()
	wandItem = i
}

func Wand() item.Stack {
	wandMu.RLock()
	defer wandMu.RUnlock()
	return wandItem
}
