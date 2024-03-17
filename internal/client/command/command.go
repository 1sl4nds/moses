package command

type Command interface {
	Init(args []string) error
	Name() string
	Run() error
}
