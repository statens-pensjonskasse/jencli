package deploy

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"jencli/pkg/api/jenkins/types"
	"jencli/pkg/common"
	"jencli/pkg/printer"
	"os"
	"path/filepath"
)

var Cmd = &cobra.Command{
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag(common.ImageFlag, cmd.Flags().Lookup(common.ImageFlag))
		viper.BindPFlag(common.BranchFlag, cmd.Flags().Lookup(common.BranchFlag))
		viper.BindPFlag(common.TagFlag, cmd.Flags().Lookup(common.TagFlag))
		viper.BindPFlag(common.UseBranchTagFlag, cmd.Flags().Lookup(common.UseBranchTagFlag))
		viper.BindPFlag(common.UseBranchPostfixFlag, cmd.Flags().Lookup(common.UseBranchPostfixFlag))
		viper.BindPFlag(common.SwarmFlag, cmd.Flags().Lookup(common.SwarmFlag))
		viper.BindPFlag(common.EnvFlag, cmd.Flags().Lookup(common.EnvFlag))
		viper.BindPFlag(common.StackFlag, cmd.Flags().Lookup(common.StackFlag))
		viper.BindPFlag(common.DryRunFlag, cmd.Flags().Lookup(common.DryRunFlag))
		viper.BindPFlag(common.UseHarborFlag, cmd.Flags().Lookup(common.UseHarborFlag))
	},
	Use:   "deploy",
	Short: "Deploy an application using JPL manual deploy job",
	Run:   deploy,
}

func init() {
	Cmd.Flags().String(common.ImageFlag, "", "Image name to deploy, if empty try to guess from .git/config")
	Cmd.Flags().String(common.BranchFlag, "", "Branch to deploy, if empty try to guess from .git/HEAD")
	Cmd.Flags().String(common.TagFlag, "latest", "Tag to deploy")
	Cmd.Flags().Bool(common.UseBranchTagFlag, false, "Use latest_<branch> tag. Overrides 'tag'")
	Cmd.Flags().Bool(common.UseBranchPostfixFlag, false, "Use branch postfix i image name")
	Cmd.Flags().String(common.SwarmFlag, "", "Which swarm to use ['utv', 'team', 'test']")
	Cmd.Flags().String(common.EnvFlag, "", "Which environment to use, e.g. 'utv', 'tmmmed1'")
	Cmd.Flags().String(common.StackFlag, "", "Which stack config to use. Defaults to environment.")
	Cmd.Flags().Bool(common.UseHarborFlag, false, "Use Harbor instead of Old Dockerhub")

	Cmd.Flags().Bool(common.DryRunFlag, false, "Only render output")
}

func deploy(cmd *cobra.Command, args []string) {
	jenkinsUrl := viper.GetString(common.JenkinsUrlFlag)
	deployJob := viper.GetString(common.JenkinsJPLDeployJobFlag)
	username := viper.GetString(common.UsernameFlag)
	token := viper.GetString(common.TokenFlag)
	slack := viper.GetString(common.SlackFlag)
	image := viper.GetString(common.ImageFlag)
	branch := viper.GetString(common.BranchFlag)
	tag := viper.GetString(common.TagFlag)
	useBranchTag := viper.GetBool(common.UseBranchTagFlag)
	useBranchPostfix := viper.GetBool(common.UseBranchPostfixFlag)
	swarm := viper.GetString(common.SwarmFlag)
	env := viper.GetString(common.EnvFlag)
	stack := viper.GetString(common.StackFlag)
	dryRun := viper.GetBool(common.DryRunFlag)
	useHarbor := viper.GetBool(common.UseHarborFlag)

	var err error

	if len(env) < 1 {
		cobra.CheckErr(errors.New("env cannot be empty"))
	}
	if len(swarm) < 1 {
		cobra.CheckErr(errors.New("swarm cannot be empty"))
	}

	if len(image) < 1 {
		cwd, err := os.Getwd()
		cobra.CheckErr(err)
		image = filepath.Base(cwd)
	}

	if len(branch) < 1 {
		cwd, _ := os.Getwd()
		branch, err = common.GetCurrentBranch(cwd)
		cobra.CheckErr(err)
	}
	normalisedBranch := common.NormaliseBranchName(branch)

	if useBranchTag {
		tag = fmt.Sprintf("latest_%s", normalisedBranch)
	}

	if len(stack) < 1 {
		stack = env
	}

	fullImage := ""
	cwd, _ := os.Getwd()
	project, err := common.GetProject(cwd)
	if useHarbor {
		cobra.CheckErr(err)
		fullImage += "cr.spk.no/" + project + "/"
	} else {
		fullImage += "old-dockerhub.spk.no:5000/"
	}
	fullImage += image
	if useBranchPostfix {
		fullImage += fmt.Sprintf("/%s", normalisedBranch)
	}
	fullImage += fmt.Sprintf(":%s", tag)

	params := types.JPLDeploy{
		Swarm:            swarm,
		Image:            image,
		Branch:           branch,
		UseBranchPostfix: useBranchPostfix,
		Tag:              tag,
		UseBranchTag:     useBranchTag,
		FullImageName:    fullImage,
		Environment:      env,
		StackConfig:      stack,
		Slack:            slack,
		UseHarbor:        useHarbor,
		Project:          project,
	}

	cobra.CheckErr(printer.PrintYaml(params))

	if !dryRun {
		url := fmt.Sprintf("%s/%s/buildWithParameters", jenkinsUrl, deployJob)
		fmt.Printf("Sending POST to %s\n", url)
		resp, err := common.PostRequest(url, username, token, params.ToParamMap())
		cobra.CheckErr(err)
		fmt.Printf("Reply: %s\n", resp.Status)
		switch resp.StatusCode {
		case 201:
			fmt.Printf("Check deployment status at %s/%s\n", jenkinsUrl, deployJob)
		case 401:
			fallthrough
		case 403:
			fmt.Println("Is the API token valid?")
		case 404:
			fmt.Println("Is the job URL correct?")
		default:
			fmt.Println("Unexpected reply from Jenkins")
		}
	}
}
