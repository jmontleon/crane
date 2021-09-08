package main

import (
	"os"

	"github.com/konveyor/crane/cmd/apply"
	export "github.com/konveyor/crane/cmd/export"
	transfer_pvc "github.com/konveyor/crane/cmd/transfer-pvc"
	"github.com/konveyor/crane/cmd/transform"
	tunnel_api "github.com/konveyor/crane/cmd/tunnel-api"
	"github.com/konveyor/crane/internal/flags"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func main() {
	f := &flags.GlobalFlags{}
	root := cobra.Command{
		Use: "crane",
	}
	f.ApplyFlags(&root)
	root.AddCommand(export.NewExportCommand(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}, f))
	root.AddCommand(transfer_pvc.NewTransferOptions(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}))
	root.AddCommand(tunnel_api.NewTunnelAPIOptions(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}))
	root.AddCommand(transform.NewTransformCommand(f))
	root.AddCommand(apply.NewApplyCommand(f))
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
