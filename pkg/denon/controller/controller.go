package controller

import (
	"github.com/Ladicle/denon-avr/pkg/denon"
)

const cmdPower = "PW"

// GetPowerInfo returns power status information.
func GetPowerInfo(c *denon.Client) (string, error) {
	return runPowerCmd(c, "?")
}

// PowerOn turns on the AVR power.
func PowerOn(c *denon.Client) (string, error) {
	return runPowerCmd(c, "ON")
}

// PowerStandby turns the AVR power to standby.
func PowerStandby(c *denon.Client) (string, error) {
	return runPowerCmd(c, "STANDBY")
}

func runPowerCmd(c *denon.Client, param string) (string, error) {
	c.Write(cmdPower, param)
	return c.Read()
}
