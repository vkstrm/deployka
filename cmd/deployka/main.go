package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/vkstrm/deployka/internal"
)

func main() {

	var blockCmd = &cobra.Command{
		Use:   "block",
		Short: "Block one or several pipes",
		Long:  "Block one or several pipes by writing their names",
		Run: func(cmd *cobra.Command, args []string) {
			internal.CheckConfig()
			err := internal.BlockPipes(args)
			if err != nil {
				fmt.Printf("error: %v\n", err)
				os.Exit(1)
			}
		},
	}
	addCommonFlags(blockCmd)

	var unblockCmd = &cobra.Command{
		Use:   "unblock",
		Short: "Unblock one or several pipes",
		Long:  "Unblock one or several pipes by writing their names",
		Run: func(cmd *cobra.Command, args []string) {
			internal.CheckConfig()
			err := internal.UnblockPipes(args)
			if err != nil {
				fmt.Printf("error: %v\n", err)
				os.Exit(1)
			}
		},
	}
	addCommonFlags(unblockCmd)

	var initCmd = &cobra.Command{
		Use:   "config",
		Short: "Configure this client",
		Long:  "Configure your API key and the endpoint URL",
		Run: func(cmd *cobra.Command, args []string) {
			err := internal.Config(os.Stdin)
			if err != nil {
				fmt.Printf("error: %v\n", err)
				os.Exit(1)
			}
		},
	}
	addCommonFlags(initCmd)

	var rootCmd = &cobra.Command{
		Use:   "deployka",
		Short: "Deployka will show you if you can deploy",
		Long: "Looking to deploy? Deployka will tell you if any teammate is currently blocking a pipeline.\n" +
			"Or block it yourself to keep their inferior code away from the production environment.\n" +
			"Of course, anyone is free to ignore blockage but Deployka doesn't support that kind of behaviour.\n",
		Run: func(cmd *cobra.Command, args []string) {
			internal.CheckConfig()
			err := internal.FetchPipes()
			if err != nil {
				fmt.Printf("error: %v\n", err)
				os.Exit(1)
			}
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			verbose, err := cmd.Flags().GetBool("verbose")
			if err != nil {
				panic(err)
			}

			if !verbose {
				log.SetOutput(ioutil.Discard)
			}
		},
	}
	addCommonFlags(rootCmd)

	rootCmd.AddCommand(blockCmd)
	rootCmd.AddCommand(unblockCmd)
	rootCmd.AddCommand(initCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func addCommonFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("verbose", "v", false, "Whether to output detailed information")
}
