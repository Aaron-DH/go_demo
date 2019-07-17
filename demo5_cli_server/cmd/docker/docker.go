package docker

import (
	"github.com/urfave/cli"
)

var Command = cli.Command{
	Name : "docker",
	Usage : "Manage docker cmd",
	Subcommands : []cli.Command{
		Images,
		//Networks,
		Pull,
		//Ps,
	},
}
var (
	Images = cli.Command{
		Name : "images",
		Usage : "List Docker Images",
		Action : listimages,
	}

	Pull = cli.Command{
	    Name : "pull",
	    Usage : "Pull Docker Images",
	    Action : pullimages,
	}

	/**
	Networks = cli.Command{
		Name : "networks",
		Usage : "List Docker Networks",
		Action : listnets,
	}

	Ps = cli.Command{
	    Name : "ps",
	    Usage : "List Docker Containers",
	    Action : listcontainers,
	}
	*/
)
