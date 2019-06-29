package daemon

import (
	"github.com/pkg/errors"

	"github.com/spf13/cobra"

	"github.com/havoc-io/mutagen/cmd"
	"github.com/havoc-io/mutagen/pkg/daemon"
)

func unregisterMain(command *cobra.Command, arguments []string) error {
	// Validate arguments.
	if len(arguments) != 0 {
		return errors.New("unexpected arguments provided")
	}

	// Attempt deregistration.
	if err := daemon.Unregister(); err != nil {
		return err
	}

	// Success.
	return nil
}

var unregisterCommand = &cobra.Command{
	Use:   "unregister",
	Short: "Unregisters automatic Mutagen daemon start-up",
	Run:   cmd.Mainify(unregisterMain),
}

var unregisterConfiguration struct {
	// help indicates whether or not help information should be shown for the
	// command.
	help bool
}

func init() {
	// Grab a handle for the command line flags.
	flags := unregisterCommand.Flags()

	// Disable alphabetical sorting of flags in help output.
	flags.SortFlags = false

	// Manually add a help flag to override the default message. Cobra will
	// still implement its logic automatically.
	flags.BoolVarP(&unregisterConfiguration.help, "help", "h", false, "Show help information")
}