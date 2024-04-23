package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	_ "github.com/trustacks/trustacks/pkg/actions"
)

var (
	rootCmdServer string
)

var rootCmd = &cobra.Command{
	Use:   "tsctl",
	Short: "TruStacks software delivery engine",
}

func commandFailure(err error) {
	fmt.Printf("%s %s\n", lipgloss.NewStyle().Foreground(lipgloss.Color("#ff5555")).Bold(true).Render("[ERR]"), err)
	os.Exit(1)
}

func main() {
	rootCmd.Flags().StringVar(&rootCmdServer, "server", "", "rpc server host")
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("%s executing the command: %s\n", lipgloss.NewStyle().Foreground(lipgloss.Color("#ff5555")).Bold(true).Render("[ERR]"), err)
		os.Exit(1)
	}
}
