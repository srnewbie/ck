package preparer

import (
	"ck/models"
	"testing"

	"ck/.gen/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestPreparer struct {
	suite.Suite
	preparer *Preparer
}

func (s *TestPreparer) SetupTest() {
	s.preparer = NewPreparer(new(mocks.PQ), new(mocks.Queue))
}

func (s *TestPreparer) TestProcess() {
	oq := s.preparer.oq.(*mocks.Queue)
	pq := s.preparer.pq.(*mocks.PQ)
	order := &models.Order{ID: 0, PrepareTime: 8}
	oq.On("Pop").Return(order)
	pq.On("Len").Return(1)
	pq.On("Push", mock.Anything).Return(nil)
	s.preparer.Process()
}

func (s *TestPreparer) TestProcessNoOrder() {
	oq := s.preparer.oq.(*mocks.Queue)
	oq.On("Pop").Return(nil)
	s.preparer.Process()
}

func Test(t *testing.T) {
	suite.Run(t, new(TestPreparer))
}
