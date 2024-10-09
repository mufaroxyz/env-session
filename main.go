package main

import (
	"env-session/lib"
	"os"
)

func main() {
	var config = lib.GetConfig()
	_ = lib.RunPowershellCommand(config)
	os.Exit(0)

	//var err = powershell.Wait()
	//if err != nil {
	//	panic(err)
	//}
}
