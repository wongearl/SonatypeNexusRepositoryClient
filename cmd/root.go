/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	flag "github.com/spf13/pflag"

	"github.com/spf13/cobra"

	fakeruntime "github.com/linuxsuren/go-fake-runtime"
)

func CreateRootCmd(execer fakeruntime.Execer) (c *cobra.Command) {
	opt := &rootOption{execer: execer}
	c = &cobra.Command{
		Use:   "SonatypeNexusRepositoryClient",
		Short: "SonatypeNexusRepositoryClient",
	}
	c.AddCommand(createUploadCmd(opt))
	return
}

func (o *rootOption) AddFlags(flags *flag.FlagSet) {

}

type rootOption struct {
	// innder fields
	execer fakeruntime.Execer
}
