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

func (h *Handler) GetEmployees(ctx context.Context, req *pb.GetEmployeesRequest) (*pb.GetEmployeesResponse, error) {
	page := int(req.GetPage())
	limit := int(req.GetLimit())
	name := req.GetName()

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	emps, total, err := h.service.GetEmployeesWithFilter(ctx, page, limit, name)
	if err != nil {
		return &pb.GetEmployeesResponse{
			Status:  500,
			Message: "Failed to fetch employees",
			Data:    nil,
			Pagination: &pb.Pagination{
				Total: 0,
				Page:  int32(page),
				Limit: int32(limit),
			},
		}, nil
	}

	var res []*pb.Employee
	for _, emp := range emps {
		res = append(res, &pb.Employee{
			Id:       int32(emp.ID),
			Name:     emp.Name,
			Position: emp.Position,
		})
	}

	return &pb.GetEmployeesResponse{
		Status:  200,
		Message: "Employees fetched successfully",
		Data:    res,
		Pagination: &pb.Pagination{
			Total: int32(total),
			Page:  int32(page),
			Limit: int32(limit),
		},
	}, nil
}
