package cmd

import (
	"fmt"
	"github.com/deepaksinghvi/perfrun/pkg/service"
	"github.com/spf13/cobra"
)

var targeturl *string
// perfrunsCmd represents the perfruns command
var perfrunsCmd = &cobra.Command{
	Use:   "perfruns",
	Short: "A brief description of your command",
	Long: `A longer description.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("perfruns called")
		perfRuns := service.GetPerfRuns(*targeturl)
		fmt.Println(perfRuns)
	},
}

func init() {
	rootCmd.AddCommand(perfrunsCmd)

	// Here you will define your flags and configuration settings.
	targeturl = perfrunsCmd.Flags().String("targeturl","", "target url for perfrun apis")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// perfrunsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// perfrunsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
