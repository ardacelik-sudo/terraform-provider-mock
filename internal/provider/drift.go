package provider

import (
	"sync"
	"time"
)

// DriftSimulator records simulated external out-of-band modifications
type DriftSimulator struct {
	mu        sync.RWMutex
	modified  map[string]time.Time
}

func NewDriftSimulator() *DriftSimulator {
	return &DriftSimulator{
		modified: make(map[string]time.Time),
	}
}

func (s *DriftSimulator) MarkDrifted(resourceID string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.modified[resourceID] = time.Now()
}

func (s *DriftSimulator) HasDrifted(resourceID string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, exists := s.modified[resourceID]
	return exists
}
