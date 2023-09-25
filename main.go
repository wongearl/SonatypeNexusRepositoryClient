/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"

	fakeruntime "github.com/linuxsuren/go-fake-runtime"
	"github.com/wongearl/SonatypeNexusRepositoryClient/cmd"
)

func main() {
	err := cmd.CreateRootCmd(fakeruntime.DefaultExecer{}).Execute()
	if err != nil {
		os.Exit(1)
	}
}
