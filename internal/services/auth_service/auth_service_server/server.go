package authserviceserver

import (
	pb "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/services/auth_service/auth_service_gen"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
}
