package pkg

import (
	"sync"
)

// StatCounter is an interface describing what a counter should look like
type StatCounter interface {
	// Add a "hit" to a given request
	Add(req Request) error
	// GetMostRequested will return the request who had the most "hits"
	GetMostRequested() (req Request, hits uint, err error)
}

type statCounter struct {
	sync.RWMutex
	m map[string]uint
}

// Add a "hit" to a given request
func (stat *statCounter) Add(r Request) error {
	key, err := r.Key()
	if err != nil {
		return err
	}
	stat.Lock()
	stat.m[string(key)]++
	stat.Unlock()
	return nil
}

// GetMostRequested will return the most requested fizzbuzz combination
func (stat *statCounter) GetMostRequested() (req Request, hits uint, err error) {
	var r request
	stat.RLock()
	defer stat.RUnlock()
	for key, hitNum := range stat.m {
		if hitNum > hits {
			req, err = r.FromKey([]byte(key))
			if err != nil {
				return
			}
			hits = hitNum
		}
	}
	return
}

// GetMostRequested will return the most requested fizzbuzz combination
func GetMostRequested() (req Request, hits uint, err error) {
	return counter.GetMostRequested()
}

// The global stat counter. Can be set using SetStatCounter
var counter StatCounter

// SetStatCounter will set the global stat counter.
// If not set, will use the default, map based counter
func SetStatCounter(stat StatCounter) {
	counter = stat
}

// Initialize default stat counter
func init() {
	countR := new(statCounter)
	countR.m = make(map[string]uint)
	counter = countR
}
