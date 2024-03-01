package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"jencli/pkg/cmd/config"
	"jencli/pkg/cmd/deploy"
	"jencli/pkg/cmd/version"
	"jencli/pkg/common"
	"log"
	"path/filepath"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "jencli",
		Short: "jencli — A simple CLI-tool for Jenkins",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, common.ConfigFlag, "", "Base config file (default $HOME/.config/jencli/config.yaml")
	rootCmd.PersistentFlags().String(common.JenkinsUrlFlag, "", "Base URL for Jenkins")
	rootCmd.PersistentFlags().String(common.JenkinsJPLDeployJobFlag, "", "JPL deploy job")
	rootCmd.PersistentFlags().String(common.JenkinsJobFlag, "", "Jenkins job name")
	rootCmd.PersistentFlags().String(common.UsernameFlag, "", "Your username")
	rootCmd.PersistentFlags().String(common.TokenFlag, "", "Your Jenkins API token")
	rootCmd.PersistentFlags().String(common.SlackFlag, "", "Default Slack notification channel")

	viper.BindPFlag(common.JenkinsUrlFlag, rootCmd.PersistentFlags().Lookup(common.JenkinsUrlFlag))
	viper.BindPFlag(common.JenkinsJPLDeployJobFlag, rootCmd.PersistentFlags().Lookup(common.JenkinsJPLDeployJobFlag))
	viper.BindPFlag(common.JenkinsJobFlag, rootCmd.PersistentFlags().Lookup(common.JenkinsJobFlag))
	viper.BindPFlag(common.UsernameFlag, rootCmd.PersistentFlags().Lookup(common.UsernameFlag))
	viper.BindPFlag(common.TokenFlag, rootCmd.PersistentFlags().Lookup(common.TokenFlag))
	viper.BindPFlag(common.SlackFlag, rootCmd.PersistentFlags().Lookup(common.SlackFlag))

	rootCmd.AddCommand(config.Cmd)
	rootCmd.AddCommand(deploy.Cmd)
	rootCmd.AddCommand(version.Cmd)
}

func initConfig() {
	cfgPath, err := common.GetConfigPath()
	cobra.CheckErr(err)

	if cfgFile == "" {
		cfgFile = filepath.Join(cfgPath, "config.yaml")
	}

	cobra.CheckErr(common.CreateDirIfNotExists(cfgFile, 0700))
	cobra.CheckErr(common.CreateFileIfNotExists(cfgFile, 0600))
	cobra.CheckErr(common.CheckFilePermission(cfgFile, 0600))

	viper.AutomaticEnv()

	viper.SetConfigFile(cfgFile)
	viper.SetConfigPermissions(0600)
	cobra.CheckErr(viper.ReadInConfig())
}
