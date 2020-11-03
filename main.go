package main

import (
	"fmt"
	cli "github.com/urfave/cli/v2"
	"github.com/google/gopacket/pcap"
	"os"
)

func main() {

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "print-version",
		Aliases: []string{"V"},
		Usage:   "print only the version",
	}

	app := &cli.App{
		Name:    "partay",
		Version: "v19.99.0",
		Commands: []*cli.Command{
		  {
		  	Name: "pcapv",
		  	Aliases: []string{"pcv"},
		  	Usage: "show libpcap version used",
		  	Action:  func(c *cli.Context) error {
		  		version := pcapv()
		  		fmt.Println("pcap version: ", version)
		  		return nil
		  	},
		  },
		},
	}
	app.Run(os.Args)

}

func pcapv() (version string) {
	version = pcap.Version()
	return
}
