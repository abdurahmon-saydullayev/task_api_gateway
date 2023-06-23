package client

import (
	"task_go_api_gateway/config"
	"task_go_api_gateway/genproto/blogpost_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	BookService() blogpost_service.BookServiceClient
	AuthorService() blogpost_service.AuthorServiceClient
}

type grpcClients struct {
	bookService   blogpost_service.BookServiceClient
	authorService blogpost_service.AuthorServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	connBookService, err := grpc.Dial(
		cfg.BlogpostServiceHost+cfg.BlogpostServicePort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	connAuthorService, err := grpc.Dial(
		cfg.BlogpostServiceHost+cfg.BlogpostServicePort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		bookService:   blogpost_service.NewBookServiceClient(connBookService),
		authorService: blogpost_service.NewAuthorServiceClient(connAuthorService),
	}, nil
}

func (g *grpcClients) BookService() blogpost_service.BookServiceClient {
	return g.bookService
}

func (g *grpcClients) AuthorService() blogpost_service.AuthorServiceClient {
	return g.authorService
}
