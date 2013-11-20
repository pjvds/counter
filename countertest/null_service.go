package countertest

import (
	"github.com/pjvds/counter"
)

type CountService struct {
	IncreaseFunc func(name counter.Name) error
	GetFunc      func(name counter.Name) (int, error)
}

func NewCountService() *CountService {
	return &CountService{}
}

func (s *CountService) Increase(name counter.Name) error {
	if s.IncreaseFunc != nil {
		return s.IncreaseFunc(name)
	}

	return nil
}

func (s *CountService) Get(name counter.Name) (int, error) {
	if s.GetFunc != nil {
		return s.GetFunc(name)
	}

	return 0, nil
}
