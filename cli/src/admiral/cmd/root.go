package cmd

import (
	"fmt"

	"admiral/auth"
	"admiral/functions"
	"admiral/help"
	"admiral/paths"

	"github.com/spf13/cobra"
)

var ShowVersion bool
var (
	defaultVersion string = "0.5.0-SNAPSHOT"
	version        string
)

func init() {
	helpToken := "Authorization token."

	AutocompleteCmd.Hidden = true

	RootCmd.AddCommand(AppsRootCmd)
	RootCmd.AddCommand(CertsRootCmd)
	RootCmd.AddCommand(ConfigRootCmd)
	RootCmd.AddCommand(CredentialsRootCmd)
	RootCmd.AddCommand(DeploymentPoliciesRootCmd)
	RootCmd.AddCommand(HostsRootCmd)
	RootCmd.AddCommand(PoliciesRootCmd)
	RootCmd.AddCommand(ResourcePoolsRootCmd)
	RootCmd.AddCommand(TemplatesRootCmd)
	RootCmd.AddCommand(RegistriesRootCmd)
	RootCmd.AddCommand(NetworksRootCmd)
	RootCmd.AddCommand(CustomPropertiesRootCmd)
	RootCmd.AddCommand(AutocompleteCmd)
	RootCmd.AddCommand(GroupsRootCmd)
	RootCmd.Flags().BoolVar(&ShowVersion, "version", false, "Admiral CLI Version.")
	RootCmd.PersistentFlags().BoolVar(&functions.Verbose, "verbose", false, "Showing every request/response json body.")
	RootCmd.PersistentFlags().StringVar(&auth.TokenFromFlagVar, "token", "", helpToken)
	RootCmd.SetUsageTemplate(help.DefaultUsageTemplate)
}

//Root command which add every other commands, but can't be used as standalone.
var RootCmd = &cobra.Command{
	Use:   "admiral",
	Short: "Admiral CLI",
	Long:  "Type \"admiral readme\" for more information about the Admiral CLI.",
	Run: func(cmd *cobra.Command, args []string) {
		if !ShowVersion {
			cmd.Help()
			return
		}
		var versionToPrint string
		if version == "" {
			versionToPrint = defaultVersion
		} else {
			versionToPrint = version
		}
		fmt.Println(admiralLogo)
		fmt.Println("Version: ", versionToPrint)
	},
}

var AppsRootCmd = &cobra.Command{
	Use:   "app",
	Short: "Perform operations with applications.",
}

var CertsRootCmd = &cobra.Command{
	Use:   "cert",
	Short: "Perform operations with certificates.",
}

var ConfigRootCmd = &cobra.Command{
	Use:   "config",
	Short: "Set and get configuration properties.",
}

var CredentialsRootCmd = &cobra.Command{
	Use:   "credentials",
	Short: "Perform operations with credentials.",
}

var DeploymentPoliciesRootCmd = &cobra.Command{
	Use:   "deployment-policy",
	Short: "Perform operations with deployment policies.",
}

var HostsRootCmd = &cobra.Command{
	Use:   "host",
	Short: "Perform operations with hosts.",
}

var PoliciesRootCmd = &cobra.Command{
	Use:   "policy",
	Short: "Perform operations with policies.",
}

var ResourcePoolsRootCmd = &cobra.Command{
	Use:   "resource-pool",
	Short: "Perform operations with resource pools.",
}

var TemplatesRootCmd = &cobra.Command{
	Use:   "template",
	Short: "Perform operations with templates.",
}

var RegistriesRootCmd = &cobra.Command{
	Use:   "registry",
	Short: "Perform operations with registries.",
}

var NetworksRootCmd = &cobra.Command{
	Use:   "network",
	Short: "Perform operations with netoworks.",
}

var CustomPropertiesRootCmd = &cobra.Command{
	Use:   "custom-properties",
	Short: "Perform opertaions with custom properties.",
}

var GroupsRootCmd = &cobra.Command{
	Use:   "group",
	Short: "Perform operations with groups.",
}

var AutocompleteCmd = &cobra.Command{
	Use:   "autocomplete",
	Short: "Generate autocomplete file. It's generated in home/.admiral-cli",
	Run: func(cmd *cobra.Command, args []string) {
		RootCmd.GenBashCompletionFile(paths.CliDir() + "/admiral-cli-autocomplete.sh")
	},
}
