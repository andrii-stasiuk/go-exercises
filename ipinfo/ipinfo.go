package main

import "inc"

func main() {
	tmpIP, isGeo := core.GetFlags()
	realIP := core.SetIP(tmpIP)
	ipData := core.GetIPInfo(realIP)
	core.Printer(isGeo, ipData)
}
