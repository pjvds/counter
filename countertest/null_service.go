package countertest

import (
	"github.com/pjvds/counter"
)

type nullService struct {
}

func NewNullService() counter.CountService {
	return &nullService{}
}

func (s *nullService) Increase(name counter.Name) error {
	return nil
}

func (s *nullService) Get(name counter.Name) (int, error) {
	return 0, nil
}
