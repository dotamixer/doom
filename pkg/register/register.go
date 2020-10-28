package register

type Register interface {
	Register() error
	Deregister() error
}
