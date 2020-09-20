package program

import (
	"fmt"
	"os"
)

type Program struct {
}

func NewProgram() *Program {
	return &Program{}
}

func (p *Program) Run() {
	rootCmd := p.createRootCmd()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		p.Exit(1)
	}

	p.Exit(0)
}

func (p *Program) Exit(code int) {
	os.Exit(code)
}
