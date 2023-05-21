package cli

type SubCommand interface {
	Run(*Context) error
}
