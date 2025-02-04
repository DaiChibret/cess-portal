package command

import (
	"cess-portal/client"
	"cess-portal/conf"
	"cess-portal/tools"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func NewFileCommand() *cobra.Command {
	fc := &cobra.Command{
		Use:   "file <subcommand>",
		Short: "File related commands",
	}

	fc.AddCommand(NewFileUploadCommand())
	fc.AddCommand(NewFileDownloadCommand())
	fc.AddCommand(NewFileDeleteCommand())

	return fc
}

func NewFileUploadCommand() *cobra.Command {
	cc := &cobra.Command{
		Use:   "upload <filepath> <backups> <private key>",
		Short: "upload refers to the upload file",
		Long:  `Upload command send local source files to scheduling nodes.`,

		Run: FileUploadCommandFunc,
	}

	return cc
}

func FileUploadCommandFunc(cmd *cobra.Command, args []string) {
	InitComponents(cmd)
	if len(args) < 2 {
		fmt.Printf("Please enter correct parameters 'upload <filepath> <backups> <private key>'\n")
		os.Exit(conf.Exit_CmdLineParaErr)
	}
	PrivateKey := ""
	if len(args) < 3 {
		fmt.Printf("%s[Warming] Do you want to upload your file without private key (it's means your file status is public)?%s\n", tools.Red, tools.Reset)
		fmt.Printf("%sYou can type the 'private key' or enter with nothing to jump it:%s", tools.Red, tools.Reset)
		fmt.Scanln(&PrivateKey)
	} else {
		PrivateKey = args[2]
	}

	client.FileUpload(args[0], args[1], PrivateKey)
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
	InitComponents(cmd)
	if len(args) == 0 {
		fmt.Printf("Please enter the fileid of the downloaded file 'download <fileid>'\n")
		os.Exit(conf.Exit_CmdLineParaErr)
	}

	client.FileDownload(args[0])
}

func NewFileDeleteCommand() *cobra.Command {
	cc := &cobra.Command{
		Use:   "delete <fileid>",
		Short: "delete refers to the delete the file",
		Long:  `Deleting a file means removing the file from CESS,But there may be a delay.`,

		Run: FileDeleteCommandFunc,
	}

	return cc
}

func FileDeleteCommandFunc(cmd *cobra.Command, args []string) {
	InitComponents(cmd)
	if len(args) == 0 {
		fmt.Printf("Please enter the fileid of the deleted file'delete <fileid>'\n")
		os.Exit(conf.Exit_CmdLineParaErr)
	}
	client.FileDelete(args[0])

}
