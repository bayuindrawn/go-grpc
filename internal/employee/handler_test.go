package employee_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"go-grpc/internal/employee"
	"go-grpc/internal/employee/mocks"
	pb "go-grpc/proto/employee"
)

type HandlerTestSuite struct {
	suite.Suite
	mockService *mocks.MockService
	handler     *employee.Handler
	ctx         context.Context
}

func (s *HandlerTestSuite) SetupTest() {
	s.ctx = context.Background()
	s.mockService = new(mocks.MockService)
	s.handler = employee.NewHandler(s.mockService)
}

func (s *HandlerTestSuite) TestGetEmployees() {
	s.Run("Success", func() {
		mockEmployees := []*employee.Employee{
			{ID: 1, Name: "Bayu", Position: "Engineer"},
			{ID: 2, Name: "Rina", Position: "Manager"},
		}
		s.mockService.
			On("GetEmployeesWithFilter", mock.Anything, 1, 10, "").
			Return(mockEmployees, int64(2), nil)

		req := &pb.GetEmployeesRequest{Page: 1, Limit: 10, Name: ""}
		resp, err := s.handler.GetEmployees(s.ctx, req)

		s.NoError(err)
		s.NotNil(resp)
		s.Equal(int32(2), resp.Pagination.Total)
		s.Equal(int32(1), resp.Pagination.Page)
		s.Equal(int32(10), resp.Pagination.Limit)
		s.Len(resp.Data, 2)
		s.Equal("Bayu", resp.Data[0].Name)

		s.mockService.AssertExpectations(s.T())
	})

	s.Run("Error", func() {
		s.mockService.
			On("GetEmployeesWithFilter", mock.Anything, 1, 10, "error").
			Return(nil, int64(0), errors.New("failed to fetch data"))

		req := &pb.GetEmployeesRequest{Page: 1, Limit: 10, Name: "error"}
		resp, err := s.handler.GetEmployees(s.ctx, req)

		s.NoError(err)
		s.NotNil(resp)
		s.Equal(int32(500), resp.Status)
		s.Equal("Failed to fetch employees", resp.Message)

		s.mockService.AssertExpectations(s.T())
	})

}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}
