package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"jencli/pkg/printer"
)

var getConfigCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"g", "list", "l"},
	Short:   "Get config",
	RunE:    getConfig,
}

func getConfig(cmd *cobra.Command, args []string) error {
	keys := viper.AllKeys()

	var data [][]string
	data = append(data, []string{"Project Key", "Value"})

	for _, key := range keys {
		row := []string{key, viper.GetString(key)}
		data = append(data, row)
	}

	return printer.PrintYaml(viper.AllSettings())
}
