package http

import (
	"net/http"
	"test-architecture/usecases"
)

type Server struct {
	UserController    *UserController
	PaymentController *PaymentController
}

func NewServer(createUser usecases.CreateUser, listUsers usecases.ListUsers, processPayment usecases.ProcessPayment, getPayment usecases.GetPayment) *Server {
	userController := NewUserController(createUser, listUsers)
	paymentController := NewPaymentController(processPayment, getPayment)

	return &Server{
		UserController:    userController,
		PaymentController: paymentController,
	}
}

func (s *Server) SetupRoutes() {
	http.HandleFunc("/user/create", s.UserController.CreateUserHandler)
	http.HandleFunc("/user/list", s.UserController.ListUsersHandler)
	http.HandleFunc("/payment/process", s.PaymentController.ProcessPaymentHandler)
	http.HandleFunc("/payment/get", s.PaymentController.GetPaymentHandler)
}

func (s *Server) Run(port string) {
	http.ListenAndServe(":"+port, nil)
}
