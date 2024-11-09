/*
Copyright © 2024 Onigns <keita.onigns@outlook.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"onigns.io/keita/todotasks/models"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		taskMgr := models.GetTaskManager()

		printTasksInTable(*taskMgr)
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

func printTasksInTable(tasks models.TaskManager) {
	// Create a new table writer
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Title", "Description", "Created At", "Completed At", "Updated At"})

	// Add each task to the table
	for _, task := range tasks {
		var completedAt string
		if task.CompletedAt != nil {
			completedAt = task.CompletedAt.Format("2006-01-02 15:04:05")
		} else {
			completedAt = "Not completed"
		}

		var updatedAt string
		if task.UpdatedAt != nil {
			updatedAt = task.UpdatedAt.Format("2006-01-02 15:04:05")
		} else {
			updatedAt = "Not updated"
		}

		// Add task row to table
		table.Append([]string{
			fmt.Sprintf("%d", task.Id),
			task.Title,
			task.Description,
			task.CreatedAt.Format("2006-01-02 15:04:05"),
			completedAt,
			updatedAt,
		})
	}

	// Render the table to stdout
	table.Render()
}
