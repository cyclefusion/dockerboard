package cmd

import (
	"os"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"

	"github.com/dockerboard/dockerboard/app"
)

var CmdServer = cli.Command{
	Name:   "server",
	Usage:  "Start DockerBoard web server",
	Action: runServer,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "static, s",
			Value: "dist",
			Usage: "Static directory",
		},
		cli.StringFlag{
			Name:  "port, p",
			Value: "8001",
			Usage: "Port to run the server on",
		},
	},
}

func runServer(c *cli.Context) {
	// Set bluewhale dir from ENV BLUEWHALE_DIST or /bluewhale/dist.
	bluewhale := defaultTo(c.String("static"), os.Getenv("BLUEWHALE_DIST"))
	bluewhale = defaultTo(bluewhale, "/bluewhale/dist/")
	bluewhale, _ = filepath.Abs(bluewhale)
	port := c.String("port")
	log.Info("Open http://0.0.0.0:" + port)
	app.Run(bluewhale, port)
}

func defaultTo(v, dv string) string {
	if v == "" {
		return dv
	}
	return v
}