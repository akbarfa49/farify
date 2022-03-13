/*
Copyright © 2021 Akbar Fauzi <fauzia84@ymail.com>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/akbarfa49/farify/doublequote"
	"github.com/spf13/cobra"
)

var (
	filename string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: `farify`,

	Short: `a tools for purify your code from contamination`,
	Long:  `a tools for purify your code from contamination`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// switch {
		// case wish != ``:
		// 	switch wish {
		// 	case `url`:
		// 		GetWishUrl()
		// 	}

		// }

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	err := rootCmd.Execute()
	return err
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.genshin-impact.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	// rootCmd.AddCommand(&cobra.Command{Short: `esteh`, Args: rootCmd.Args})

	// rootCmd.SetUsageTemplate(`[command] [fileName]`)

	dq := &cobra.Command{
		Use:   `doublequote`,
		Short: `remove Double Quote (") and replace it with Grave accent(` + "`" + `) for your string`,
		Long:  `remove Double Quote (") and replace it with Grave accent(` + "`" + `) for your string`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			f, err := os.OpenFile(filename, os.O_RDWR, os.ModeType)
			if err != nil {
				fmt.Println(err)
				return
			}
			doublequote.PurifyDq(f)
		}}
	dq.Flags().StringVarP(&filename, `file`, `f`, ``, `filename`)
	rootCmd.AddCommand(dq)
}
