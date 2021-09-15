package pkg

import (
	"encoding/json"
	"sync"
)

// StatCount is the StatCounter response
type StatCount struct {
	Request Request
	Hits    uint
}

// MarshalJSON is used to satisfy the json.Marshaler interface
func (statCount StatCount) MarshalJSON() ([]byte, error) {
	js, err := statCount.Request.JSON()
	if err != nil {
		return nil, err
	}
	reqJson := json.RawMessage(js)
	res := struct {
		Request *json.RawMessage `json:"request"`
		Hits    uint             `json:"hits"`
	}{
		Request: &reqJson,
		Hits:    statCount.Hits,
	}
	b, err := json.MarshalIndent(&res, "", "\t")
	if err != nil {
		return nil, err
	}
	return b, nil
}

// StatCounter is an interface describing what a counter should look like
type StatCounter interface {
	// Add a "hit" to a given request
	Add(req Request) error
	// GetMostRequested will return the request who had the most "hits"
	GetMostRequested() (statCount StatCount, err error)
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
func (stat *statCounter) GetMostRequested() (statCount StatCount, err error) {
	var (
		r   request
		req Request
	)
	stat.RLock()
	defer stat.RUnlock()
	for key, hitNum := range stat.m {
		if hitNum > statCount.Hits {
			req, err = r.FromKey([]byte(key))
			if err != nil {
				return
			}
			statCount.Request = req
			statCount.Hits = hitNum
		}
	}
	return
}

// GetMostRequested will return the most requested fizzbuzz combination
func GetMostRequested() (statCount StatCount, err error) {
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
