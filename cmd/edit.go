package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:     "edit [OPTIONS] FILE VALUE",
	Aliases: []string{"ed"},
	Short:   "Edit the content of a csv file",
	Long: `Edit the content of a csv file, replacing the content of the selected cells with the given VALUE.

Example:

# Consider the following csv:
> csvkit list shopping_list.csv
Product,Quantity,Price
Apple,0.75KG,2.23€
Milk,6,4.31€
Rice,2,3.20€

> csvkit edit --selection A2:B3 shopping_list.csv "1 u."
Product,Quantity,Price
Apple,0.75KG,2.23€
Milk,1 u.,4.31€
Rice,1 u.,3.20€

> csvkit edit --selection A2:B3 shopping_list.csv "1 u." | csvit list --pretty
| Product | Quantity | Price |
| Apple   | 0.75KG   | 2.23€ |
| Milk    | 1 u.     | 4.31€ |
| Rice    | 1 u.     | 3.20€ |`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO - edit called")
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
