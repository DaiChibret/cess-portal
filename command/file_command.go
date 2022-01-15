package command

import (
	"dapp_cess_client/client"
	"fmt"
	"github.com/spf13/cobra"
)

func NewFileCommand() *cobra.Command {
	fc := &cobra.Command{
		Use:   "file <subcommand>",
		Short: "File related commands",
	}

	fc.AddCommand(NewFileUploadCommand())
	fc.AddCommand(NewFileDownloadCommand())

	return fc
}

func NewFileUploadCommand() *cobra.Command {
	cc := &cobra.Command{
		Use:   "upload <filepath>",
		Short: "upload refers to the upload file",
		Long:  `Price command send local source files to scheduling nodes.`,

		Run: FileUploadCommandFunc,
	}

	return cc
}

func FileUploadCommandFunc(cmd *cobra.Command, args []string) {
	fmt.Println("there is File Upload command!")

	client.FileUpload()
}

func NewFileDownloadCommand() *cobra.Command {
	cc := &cobra.Command{
		Use:   "download <fileid>",
		Short: "download refers to the download file",
		Long:  `Download command download file based on fileId.`,

		Run: FileDownloadCommandFunc,
	}

	return cc
}

func FileDownloadCommandFunc(cmd *cobra.Command, args []string) {
	fmt.Println("there is File Download command!")
}
