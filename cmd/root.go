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
	addTaskQuery    = `insert into tasks (id, task_name) values (?, ?)`
	finishTaskQuery = `delete from tasks`
	showTasksQuery  = `select * from tasks`
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

var showTasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "Show all tasks in To:Do",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		res, err := dbConn.QueryContext(ctx, showTasksQuery)
		if err != nil {
			log.Fatalln("Error in fetching tasks from database: ", err)
		}

		if res != nil {
			defer res.Close()
		}

		for res.Next() {
			fmt.Println(res)
		}

		fmt.Println()
		fmt.Println("Tasks printed successfully")
	},
}

var finishTaskCmd = &cobra.Command{
	Use:   "finish",
	Short: "Move the task to completed state",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		_, err := dbConn.ExecContext(ctx, finishTaskQuery)
		if err != nil {
			log.Fatalln("Error in finishing task from database: ", err)
		}
		fmt.Println("Task finished")
	},
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(addTaskCmd)
	rootCmd.AddCommand(finishTaskCmd)
	rootCmd.AddCommand(showTasksCmd)
}
