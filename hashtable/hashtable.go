package hashtable

import (
	"sync"
)

type HTable struct {
	Table map[int]string
	lock  sync.RWMutex
}

func NewHashTable() (hTable HTable) {
	return
}

func hash(val string) int {
	var h byte
	leng := len(val)
	for i := 0; i < leng; i++ {
		h += val[i] + byte(i)
	}
	separator := (val[0] + val[leng/2] + val[leng-1])
	h += separator
	return int(h)
}

func (ht *HTable) Put(val string) {
	ht.lock.Lock()
	defer ht.lock.Unlock()
	hash := hash(val)
	if ht.Table == nil {
		ht.Table = make(map[int]string)
	}
	ht.Table[hash] = val
}

func (ht *HTable) Remove(val string) {
	ht.lock.Lock()
	defer ht.lock.Unlock()
	hash := hash(val)
	delete(ht.Table, hash)
}

func (ht *HTable) Get(val string) string {
	hash := hash(val)
	return ht.Table[hash]
}

func (ht *HTable) Size() int {
	return len(ht.Table)
}
