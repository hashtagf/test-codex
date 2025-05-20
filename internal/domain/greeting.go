package domain

// Greeter defines the primary port for greeting someone.
type Greeter interface {
	Greet(name string) string
}
