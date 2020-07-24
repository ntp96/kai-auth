package main

import (
	"context"
	"log"

	"gitlab.com/tego-partner/kardiachain/kai-auth/generated/auth"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":3333", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := auth.NewAuthenticationClient(conn)
	req := &auth.RegisterMsg{
		Domain:    "dev.tego.com",
		AuthKey:   "dev@tego.com",
		SecretKey: "password",
	}
	//response, err := c.Login(context.Background(), req)
	//if err != nil {
	//	log.Fatalf("Error when calling SayHello: %s", err)
	//}
	//log.Printf("Response from server: %s", response)

	RegisterRes, _ := c.Register(context.Background(), req)
	log.Printf("Response from server: %s", RegisterRes)

	// authorized, _ := c.Authorization(context.Background(), &auth.Token{AccessToken: "123-123"})
	// log.Printf("Response from server: %s", authorized)
}
