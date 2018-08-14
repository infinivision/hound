package proxy

import (
	"time"
)

// Option option
type Option func(*options)

type options struct {
	timeout time.Duration
}

func defaultOptions() *options {
	return &options{
		timeout: time.Millisecond * 200,
	}
}

// Proxy is proxy for versiondb
type Proxy interface {
	UpdateWithIds(db uint64, extXb []float32, extXids []int64) error
	AddWithIds(newXb []float32, newXids []int64) error
	Search(xq []float32) ([]float32, []int64, error)
}