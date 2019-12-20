package deliverer

import (
	"ck/models"
	"testing"

	q "ck/models/pq"

	"ck/.gen/mocks"

	"github.com/stretchr/testify/suite"
)

type TestDeliverer struct {
	suite.Suite
	deliverer *Deliverer
}

func (s *TestDeliverer) SetupTest() {
	s.deliverer = NewDeliverer(new(mocks.PQ))
}

func (s *TestDeliverer) TestProcess() {
	pq := s.deliverer.pq.(*mocks.PQ)
	order := &models.Order{ID: 0, PrepareTime: 8}
	item := &q.Item{
		Value:    order,
		Priority: 8,
	}
	pq.On("Pop").Return(item)
	s.deliverer.Process()
}

func Test(t *testing.T) {
	suite.Run(t, new(TestDeliverer))
}
