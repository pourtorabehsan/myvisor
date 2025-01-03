package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/pourtorabehsan/myvisor/internal/advisor"
	"github.com/pourtorabehsan/myvisor/internal/mysql"
	"github.com/pourtorabehsan/myvisor/internal/sos"
	"github.com/spf13/cobra"
)

var sosSeconds int

// sosCmd represents the sos command
var sosCmd = &cobra.Command{
	Use:   "sos",
	Short: "Get SOS report for the MySQL server",
	Long:  "Generate a SOS report for the MySQL server to help you analyze the performance of your MySQL database in real time.",
	Run: func(cmd *cobra.Command, args []string) {
		sosReport()
	},
}

func init() {
	rootCmd.AddCommand(sosCmd)
	sosCmd.Flags().IntVar(&sosSeconds, "seconds", 5, "Number of seconds to collect data")
}

func sosReport() {
	db, err := mysql.Open(user, password, host, port)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	sosHeader(db)
	advisor.Run(sos.Collectors(), db)
	sosFooter()
}

func sosHeader(db *sql.DB) {
	fmt.Println("================= MYVISOR SOS ==================")
	fmt.Println("Hostname:", mysql.Hostname(db))
	fmt.Println("Time:", time.Now().Format(time.RFC3339))
	fmt.Printf("================================================\n\n")
}

func sosFooter() {
	fmt.Printf("\n================================================\n")
}
