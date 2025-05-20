package service

import "hexademo/internal/domain"

// GreetingService is the default implementation of the Greeter port.
type GreetingService struct{}

func NewGreetingService() *GreetingService {
	return &GreetingService{}
}

func (s *GreetingService) Greet(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello, " + name + "!"
}

var _ domain.Greeter = (*GreetingService)(nil)
