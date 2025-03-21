package frameworks

import "github.com/urfave/cli/v2"

type framework struct {
	Stander interface {
		CreatStanderNode()
	}
	Advanced interface{
		CreatAdvancedNode()
	}

	Commands interface {
		NodeFunc() *cli.Command
	}

}
