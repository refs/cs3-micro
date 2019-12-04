package main

import (
	"context"
	"fmt"
	"log"

	demoAuth "github.com/cs3org/reva/pkg/auth/manager/demo"
	userManager "github.com/cs3org/reva/pkg/user/manager/demo"
	"github.com/micro/go-micro"
	providerv1beta1 "github.com/refs/cs3-micro/pkg/cs3/auth/provider/v1beta1"
	userv1beta1 "github.com/refs/cs3-micro/pkg/cs3/identity/user/v1beta1"
)

type authHandler struct{}

func (h *authHandler) Authenticate(ctx context.Context, req *providerv1beta1.AuthenticateRequest, res *providerv1beta1.AuthenticateResponse) error {
	demoManager, _ := demoAuth.New(nil)
	// forward request to Demo manager
	// userid, err := demoManager.Authenticate(context.TODO(), req.GetClientId(), req.GetClientSecret())
	userid, err := demoManager.Authenticate(context.TODO(), "einstein", "relativity")
	if err != nil {
		log.Fatal(err)
	}

	// we have a user id, how do we get the user from the user id?
	um, err := userManager.New(nil)
	if err != nil {
		log.Fatal(err)
	}

	user, err := um.GetUser(context.TODO(), userid)
	if err != nil {
		log.Fatal(err)
	}

	res.User = &userv1beta1.User{
		Username:    user.GetUsername(),
		Mail:        user.GetMail(),
		DisplayName: user.GetDisplayName(),
	}

	return nil
}

func main() {
	// behind a gateway
	service := micro.NewService(
		micro.Name("go.micro.api.authprovider"),
		micro.Version("latest"),
	)

	service.Init()

	// handler registration happens only after service initialization...obviously
	err := providerv1beta1.RegisterProviderAPIHandler(service.Server(), new(authHandler))
	if err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
