package employee

import (
	"context"
	pb "go-grpc/proto/employee"
)

type Handler struct {
	pb.UnimplementedEmployeeServiceServer
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetEmployees(ctx context.Context, _ *pb.Empty) (*pb.EmployeeList, error) {
	emps, err := h.service.GetAllEmployees()
	if err != nil {
		return nil, err
	}

	var res []*pb.Employee
	for _, emp := range emps {
		res = append(res, &pb.Employee{
			Id:       int32(emp.ID),
			Name:     emp.Name,
			Position: emp.Position,
		})
	}

	return &pb.EmployeeList{Employees: res}, nil
}
