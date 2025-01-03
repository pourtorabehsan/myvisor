package sos

import (
	"database/sql"
	"fmt"

	"github.com/pourtorabehsan/myvisor/internal/advisor"
)

type TopQueriesTimeCollector struct{}

func (c *TopQueriesTimeCollector) Collect(db *sql.DB) advisor.Diagnostic {
	diag := advisor.Diagnostic{
		ID:          2,
		Description: "TOP QUERIES (Top 3 queries by total latency)",
	}

	query := "select query, FORMAT_PICO_TIME(total_latency) as total_latency_formatted, FORMAT_PICO_TIME(avg_latency) as avg_latency_formatted from sys.x$statement_analysis order by total_latency desc limit 3"
	rows, err := db.Query(query)
	if err != nil {
		diag.Error = err
		return diag
	}
	defer rows.Close()

	diag.Data = c.format(rows)
	return diag
}

func (c *TopQueriesTimeCollector) format(rows *sql.Rows) string {
	data := ""
	i := 1
	for rows.Next() {
		var query string
		var totalLatencyFormatted string
		var avgLatencyFormatted string
		rows.Scan(&query, &totalLatencyFormatted, &avgLatencyFormatted)

		data += fmt.Sprintf("%d. %s (avg %s) - %s\n", i, totalLatencyFormatted, avgLatencyFormatted, query)
		i++
	}
	return data
}
