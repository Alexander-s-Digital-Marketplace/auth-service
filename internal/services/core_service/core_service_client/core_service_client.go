package coreserviceclient

import (
	"context"
	"time"

	useraccount "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/models/account_model"
	pb "github.com/Alexander-s-Digital-Marketplace/auth-service/internal/services/core_service/core_service_gen"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ProfileRegister(profile useraccount.ProfileTDO) (int, string) {
	conn, err := grpc.NewClient(
		"localhost:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logrus.Errorln("Error connect:", err)
		return 503, ""
	}
	defer conn.Close()

	client := pb.NewCoreServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.Request{
		UserName:      profile.UserName,
		AccountInfoId: int32(profile.AccountInfoId),
		Wallet:        profile.Wallet,
	}

	res, err := client.ProfileRegister(ctx, req)
	if err != nil {
		logrus.Errorln("Error profile register:", err)
		return 503, ""
	}

	return int(res.Code), res.Message
}
