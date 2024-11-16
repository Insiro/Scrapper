package util

import (
	"Scrapper/internal/entity/enum"
	"errors"
	"sync"
)

type OnScrap struct {
	Key    string
	Source enum.PageType
}

type OnScrapSet struct {
	items sync.Map
}

func NewScrapSet(values ...OnScrap) *OnScrapSet {
	set := &OnScrapSet{}
	for _, value := range values {
		set.Add(value.Source, value.Key)
	}
	return set
}

func (s *OnScrapSet) Add(page enum.PageType, key string) error {
	if s.Contains(page, key) {
		return errors.New("exist")
	}
	s.items.Store(OnScrap{key, page}, nil)
	return nil
}

func (s *OnScrapSet) Remove(page enum.PageType, key string) {
	s.items.Delete(OnScrap{key, page})
}

func (s *OnScrapSet) Contains(page enum.PageType, key string) bool {
	_, exist := s.items.Load(OnScrap{key, page})
	return exist
}

func (s *OnScrapSet) Size() int {
	count := 0
	s.items.Range(func(_, _ interface{}) bool {
		count++
		return true
	})
	return count
}

func (s *OnScrapSet) Slice() []OnScrap {
	var keys []OnScrap
	s.items.Range(func(key, _ interface{}) bool {
		keys = append(keys, key.(OnScrap))
		return true
	})
	return keys
}
