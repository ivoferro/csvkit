package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list [OPTIONS] FILE",
	Aliases: []string{"ls"},
	Short:   "List the content of a csv file",
	Long: `List the content of a csv file, with a customizable format.

Examples:

> csvkit list shopping_list.csv
Product,Quantity,Price
Apple,0.75 kg,2.23€
Milk,6,4.31€
Rice,2,3.20€

> csvkit list --pretty shopping_list.csv
| Product | Quantity | Price |
| Apple   | 0.75 kg  | 2.23€ |
| Milk    | 6        | 4.31€ |
| Rice    | 2        | 3.20€ |

> csvkit list --pretty --selection A2:B3 shopping_list.csv
| Apple   | 0.75 kg  |
| Milk    | 6        |`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO - list called")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
