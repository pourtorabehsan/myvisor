package advisor

import "database/sql"

type Collector interface {
	Collect(db *sql.DB) Diagnostic
}
