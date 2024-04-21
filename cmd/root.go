/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
)

const (
	addTaskQuery = `insert into (id, task_name) values (:id, :task_name)`
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tasker",
	Short: "Tasker to create and manage tasks",
}

var addTaskCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new expense",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		_, err := dbConn.ExecContext(ctx, addTaskQuery, 1, "sample_task_name")
		if err != nil {
			log.Fatalln("Error in inserting task to database: ", err)
		}

		fmt.Println("Task created")
	},
}

var finishTaskCmd = &cobra.Command{
	Use:   "finish",
	Short: "Move the task to completed state",
}

var deleteTaskCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task from the task list",
}

var dbConn *sqlx.DB

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(databaseConn *sqlx.DB) {
	dbConn = databaseConn
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.expense_tracker.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(addTaskCmd)
	rootCmd.AddCommand(finishTaskCmd)
	rootCmd.AddCommand(deleteTaskCmd)
}
