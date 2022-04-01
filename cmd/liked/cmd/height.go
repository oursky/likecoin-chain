package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func ShowHeightCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-height",
		Short: "Show the latest block height in local database for export",

		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			homeDir, err := cmd.Flags().GetString(flags.FlagHome)
			if err != nil {
				return err
			}
			dataDir := filepath.Join(homeDir, "data")
			db, err := sdk.NewLevelDB("application", dataDir)
			if err != nil {
				return err
			}
			cms := store.NewCommitMultiStore(db)
			_, err = fmt.Println(cms.LastCommitID().Version)
			if err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}
