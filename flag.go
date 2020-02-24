// Package flag implements two types used to create boolean flag variables.
// One of these types can be used by concurrent goroutines.
package flag

import (
	"sync"
)

// Flag defines a type for values that can be used as boolena flags.
// Not safe for concurrent use. See CcFlag for that.
type Flag struct {
	state bool
}

// New returns a FLag
func New() *Flag {
	return &Flag{false}
}

func (f *Flag) Flip() {
	f.state = !f.state
}

func (f *Flag) Set(b bool) {
	f.state = b
}

func (f *Flag) IsTrue() bool {
	return f.state
}

// CcFlag is a boolean flag that is safe to modify concurrently.
type CcFlag struct {
	state bool
	mu    *sync.Mutex
}

// NewCC returns a FLag safe to use in a concurrent setting.
func NewCC() *CcFlag {
	return &CcFlag{false, &sync.Mutex{}}
}

func (f *CcFlag) Flip() {
	f.mu.Lock()
	f.state = !f.state
	f.mu.Unlock()
}

func (f *CcFlag) Set(b bool) {
	f.mu.Lock()
	f.state = b
	f.mu.Unlock()
}

func (f *CcFlag) IsTrue() bool {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.state
}
