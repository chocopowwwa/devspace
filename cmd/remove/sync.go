package remove

import (
	"github.com/devspace-cloud/devspace/pkg/devspace/config/configutil"
	"github.com/devspace-cloud/devspace/pkg/devspace/configure"
	"github.com/devspace-cloud/devspace/pkg/util/log"
	"github.com/spf13/cobra"
)

type syncCmd struct {
	LabelSelector string
	LocalPath     string
	ContainerPath string
	RemoveAll     bool
}

func newSyncCmd() *cobra.Command {
	cmd := &syncCmd{}

	syncCmd := &cobra.Command{
		Use:   "sync",
		Short: "Remove sync paths from the devspace",
		Long: `
	#######################################################
	############### devspace remove sync ##################
	#######################################################
	Remove sync paths from the devspace

	How to use:
	devspace remove sync --local=app
	devspace remove sync --container=/app
	devspace remove sync --label-selector=release=test
	devspace remove sync --all
	#######################################################
	`,
		Args: cobra.NoArgs,
		Run:  cmd.RunRemoveSync,
	}

	syncCmd.Flags().StringVar(&cmd.LabelSelector, "label-selector", "", "Comma separated key=value selector list (e.g. release=test)")
	syncCmd.Flags().StringVar(&cmd.LocalPath, "local", "", "Relative local path to remove")
	syncCmd.Flags().StringVar(&cmd.ContainerPath, "container", "", "Absolute container path to remove")
	syncCmd.Flags().BoolVar(&cmd.RemoveAll, "all", false, "Remove all configured sync paths")

	return syncCmd
}

// RunRemoveSync executes the remove sync command logic
func (cmd *syncCmd) RunRemoveSync(cobraCmd *cobra.Command, args []string) {
	// Set config root
	configExists, err := configutil.SetDevSpaceRoot()
	if err != nil {
		log.Fatal(err)
	}
	if !configExists {
		log.Fatal("Couldn't find any devspace configuration. Please run `devspace init`")
	}

	err = configure.RemoveSyncPath(cmd.RemoveAll, cmd.LocalPath, cmd.ContainerPath, cmd.LabelSelector)
	if err != nil {
		log.Fatal(err)
	}
}
