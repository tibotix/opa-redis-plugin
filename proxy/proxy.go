package proxy

import (
	"errors"
	"sync"
)

type Proxy[T any] struct {
	mu    sync.Mutex
	obj   T
	isSet bool
}

func (p *Proxy[T]) Set(obj T) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.obj = obj
	p.isSet = true
}

func (p *Proxy[T]) Get() (T, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.isSet {
		return p.obj, nil
	}
	return p.obj, errors.New("Object not set yet")
}

func (p *Proxy[T]) IsSet() bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.isSet
}

func (p *Proxy[T]) Unset() {
	p.isSet = false
}
