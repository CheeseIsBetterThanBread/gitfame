package stats

import (
	"sync"
	"sync/atomic"
)

type safeMapBool struct {
	mutex sync.RWMutex
	data  map[string]bool
}

func (s *safeMapBool) Found(key string) {
	s.mutex.RLock()

	if _, ok := s.data[key]; ok {
		s.mutex.RUnlock()
		return
	}
	s.mutex.RUnlock()

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.data[key] = true
}

func (s *safeMapBool) Len() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return len(s.data)
}

func newSafeMapBool() *safeMapBool {
	return &safeMapBool{data: make(map[string]bool)}
}

type safeMapHelper struct {
	mutex sync.RWMutex
	data  map[string]*authorStatsHelper
}

func newSafeMapHelper() *safeMapHelper {
	return &safeMapHelper{data: make(map[string]*authorStatsHelper)}
}

func (s *safeMapHelper) Create(author string) {
	s.mutex.RLock()

	if _, ok := s.data[author]; ok {
		s.mutex.RUnlock()
		return
	}
	s.mutex.RUnlock()

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.data[author] = &authorStatsHelper{Name: author, Commits: newSafeMapBool()}
}

func (s *safeMapHelper) IncreaseLines(author string, lines int) {
	atomic.AddInt32(&s.data[author].Lines, int32(lines))
}

func (s *safeMapHelper) IncreaseFiles(author string) {
	atomic.AddInt32(&s.data[author].Files, 1)
}

func (s *safeMapHelper) LogCommits(author, commit string) {
	s.data[author].Commits.Found(commit)
}
