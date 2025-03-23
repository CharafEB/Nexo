package frameworks

import (
	"github.com/Nexo/cmd/midel"
	"github.com/urfave/cli/v2"
)

type framework struct {
	Stander interface {
		CreatStanderNode()
	}
	Advanced interface {
		CreatAdvancedNode()
	}

	Commands interface {
		NodeFunc() *cli.Command
	}

	Json interface {
		searchJSON(val string)
		WritJSON(val midel.Blueprints) error
	}
}
