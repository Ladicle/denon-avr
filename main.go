package main

import (
	"flag"
	"time"

	"github.com/golang/glog"
	"github.com/spf13/pflag"
)

const defaultTimeout = 3 * time.Second

var (
	// overwritten in built time
	version = "0.1"
	gitRepo = "https://github.com/Ladicle/denon-avr"

	flags = pflag.NewFlagSet("", pflag.ExitOnError)
)

func init() {
	// defaults to log to standard error instead of files
	flag.Set("logtostderr", "true")
}

func main() {
	flags.AddGoFlagSet(flag.CommandLine)
	flag.Parse()
	if flag.NArg() != 1 {
		glog.Fatalf("Invalid argument of numbers.")
	}
	host := flag.Arg(0)

	glog.Infof("Using build: %v - %v", gitRepo, version)

	c, err := Dial(host, defaultTimeout)
	if err != nil {
		glog.Fatalf("Could not dial %v: %v", host, err)
	}
	defer c.Close()
	glog.Infof("Success to dial %v", host)

	c.Write("PW", "?")
	data, err := c.Read()
	if err != nil {
		glog.Fatalf("Could not read response: %v", err)
	}
	glog.Infof("AVR: %v", string(data))
}
