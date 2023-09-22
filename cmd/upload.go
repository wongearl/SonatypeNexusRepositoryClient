/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wongearl/SonatypeNexusRepositoryClient/pkg/upload"
)

func createUploadCmd(rootOpt *rootOption) (c *cobra.Command) {
	opt := &uploadOption{rootOption: rootOpt}
	c = &cobra.Command{
		Use:   "upload",
		Short: "upload file",
		RunE:  opt.runE,
	}
	flags := c.Flags()
	opt.AddFlags(flags)
	flags.StringVarP(&opt.nexusURL, "nexusURL", "n", "", "The mode nexusURL")
	flags.StringVarP(&opt.username, "username", "u", "", "The username")
	flags.StringVarP(&opt.password, "password", "p", "", "The password")
	flags.StringVarP(&opt.repositoryName, "repository", "r", "", "The repository")
	flags.StringVarP(&opt.filePath, "file-path", "f", "", "The file path")
	return
}

func (u *uploadOption) runE(c *cobra.Command, args []string) (err error) {
	err = upload.Upload(u.username, u.password, u.filePath, u.repositoryName, u.nexusURL)
	return
}

type uploadOption struct {
	*rootOption
	nexusURL       string
	username       string
	password       string
	repositoryName string
	filePath       string
}
