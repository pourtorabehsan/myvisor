package sos

import (
	"database/sql"
	"fmt"

	"github.com/pourtorabehsan/myvisor/internal/advisor"
	"github.com/pourtorabehsan/myvisor/pkg/humanize"
)

type TopQueriesCountCollector struct{}

func (c *TopQueriesCountCollector) Collect(db *sql.DB) advisor.Diagnostic {
	diag := advisor.Diagnostic{
		ID:          1,
		Description: "TOP QUERIES (Top 3 queries by exec count)",
	}

	query := "select query, exec_count from sys.x$statement_analysis order by exec_count desc limit 3"
	rows, err := db.Query(query)
	if err != nil {
		diag.Error = err
		return diag
	}
	defer rows.Close()

	diag.Data = c.format(rows)
	return diag
}

func (c *TopQueriesCountCollector) format(rows *sql.Rows) string {
	data := ""
	i := 1
	for rows.Next() {
		var query string
		var execCount int64
		rows.Scan(&query, &execCount)

		data += fmt.Sprintf("%d. %s - %s\n", i, humanize.SI(float64(execCount)), query)
		i++
	}
	return data
}
