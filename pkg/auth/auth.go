package auth

import (
	"context"
	"errors"
	"fmt"
	"gitlab.com/tego-partner/kardiachain/kai-auth/generated/auth"
	"gitlab.com/tego-partner/kardiachain/kai-auth/pkg/crypto"
	"gitlab.com/tego-partner/kardiachain/kai-auth/third_party/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	secretKey []byte
}

// User model
type User struct {
	Domain  string
	AuthKey string
	Hash    string
	Role	interface{}
}

func NewAuthServer(key string) *Server {
	return &Server{ secretKey: []byte(key) }
}

func (s *Server) Run(port string) error {
	srv := grpc.NewServer()
	auth.RegisterAuthenticationServer(srv, s)

	lis, err := net.Listen("tcp", ":" + port)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	return srv.Serve(lis)
}

func (s *Server) Register(ctx context.Context, req *auth.RegisterMsg) (*auth.Token, error) {
	cipheredText, err := crypto.Encrypt([]byte(req.SecretKey), s.secretKey)
	if err != nil {
		log.Fatal(err)
	}

	newUser := &User{
		Domain:  req.Domain,
		AuthKey: req.AuthKey,
		Hash:    string(cipheredText),
	}

	createdUser, _ := mongodb.GetDB().Collection("users").InsertOne(ctx, newUser)
	createdUserID, _ := createdUser.InsertedID.(primitive.ObjectID)

	return genToken(createdUserID.String()), nil
}

func (s *Server) Login(ctx context.Context, req *auth.LoginMsg) (*auth.Token, error) {
	var user User
	err := mongodb.GetDB().Collection("users").FindOne(ctx, bson.M{
		"authkey": req.AuthKey,
		"domain" : req.Domain,
	}).Decode(&user)
	if err != nil {
		return nil, errors.New("not found")
	}

	plaintext, err := crypto.Decrypt([]byte(user.Hash), s.secretKey)
	if err != nil {
		return nil, errors.New("unauthorized")
	} else if string(plaintext) == req.SecretKey {
		return &auth.Token{AccessToken: req.GetAuthKey() + req.GetSecretKey()}, nil
	}

	return nil, errors.New("unauthorized")
}

func (s *Server) Authorization(ctx context.Context, req *auth.Token) (*auth.Permission, error) {
	return &auth.Permission{Permission: true}, nil
}

func genToken(userId string) *auth.Token {
	return &auth.Token{AccessToken: userId}
}