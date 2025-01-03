package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	host     string
	port     int
	user     string
	password string
)

var rootCmd = &cobra.Command{
	Use:   "myvisor",
	Short: "MySQL Performance Advisor",
	Long:  "myvisor is a tool to help you analyze the performance of your MySQL database in real time.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("myvisor is a tool to help you analyze the performance of your MySQL database in real time using the MySQL Performance Schema.")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&host, "host", "H", "localhost", "MySQL host")
	rootCmd.PersistentFlags().IntVarP(&port, "port", "P", 3306, "MySQL port")
	rootCmd.PersistentFlags().StringVarP(&user, "user", "u", "root", "MySQL user")
	rootCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "MySQL password")
}
