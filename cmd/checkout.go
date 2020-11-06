package cmd

import (
	"os"

	"github.com/ipfs/interface-go-ipfs-core/path"
	"github.com/multiverse-vcs/go-multiverse/config"
	"github.com/multiverse-vcs/go-multiverse/core"
	"github.com/spf13/cobra"
)

var checkoutCmd = &cobra.Command{
	Use:          "checkout [ref]",
	Short:        "Copy changes from a commit to the local repo.",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE:         executeCheckout,
}

func init() {
	rootCmd.AddCommand(checkoutCmd)
}

func executeCheckout(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	cfg, err := config.Open(cwd)
	if err != nil {
		return err
	}

	c, err := core.NewCore(ctx)
	if err != nil {
		return err
	}

	commit, err := c.Reference(ctx, path.New(args[0]))
	if err != nil {
		return err
	}

	if err := c.Checkout(ctx, commit, cfg.Path); err != nil {
		return err
	}

	cfg.Base = commit.Cid()
	return cfg.Write()
}
