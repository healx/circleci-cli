package cmd

import (
	"context"

	"github.com/CircleCI-Public/circleci-cli/api"
	"github.com/spf13/cobra"
)

// Path to the config.yml file to operate on.
var configPath string

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Operate on build config files",
}

var validateCommand = &cobra.Command{
	Use:     "validate",
	Aliases: []string{"check"},
	Short:   "Check that the config file is well formed.",
	RunE:    validateConfig,
}

var expandCommand = &cobra.Command{
	Use:   "expand",
	Short: "Expand the config.",
	RunE:  expandConfig,
}

func init() {
	configCmd.PersistentFlags().StringVarP(&configPath, "path", "p", ".circleci/config.yml", "path to build config")
	configCmd.AddCommand(validateCommand)
	configCmd.AddCommand(expandCommand)
}

func validateConfig(cmd *cobra.Command, args []string) error {

	ctx := context.Background()
	response, err := api.ConfigQuery(ctx, Logger, configPath)

	if err != nil {
		return err
	}

	if !response.Valid {
		return response.ToError()
	}

	Logger.Infof("Config file at %s is valid", configPath)
	return nil
}

func expandConfig(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	response, err := api.ConfigQuery(ctx, Logger, configPath)

	if err != nil {
		return err
	}

	if !response.Valid {
		return response.ToError()
	}

	Logger.Info(response.OutputYaml)
	return nil
}
