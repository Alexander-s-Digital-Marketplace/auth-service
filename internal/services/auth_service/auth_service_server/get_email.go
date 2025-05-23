package authserviceserver

import (
	"context"
	"errors"
	"net/http"

	accountmodel "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/models/account_model"
	pb "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/services/auth_service/auth_service_gen"
)

func (s *Server) GetEmailByAccountId(ctx context.Context, req *pb.EmailRequest) (*pb.EmailResponse, error) {
	user := accountmodel.UserAccount{
		Id: int(req.Id),
	}

	code := user.GetFromTable()
	if code != 200 {
		return &pb.EmailResponse{
			Code:  int32(code),
			Email: "",
		}, errors.New("error get user from table")
	}

	return &pb.EmailResponse{
		Code:  int32(http.StatusOK),
		Email: user.Email,
	}, nil
}
