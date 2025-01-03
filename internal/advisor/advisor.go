package advisor

import (
	"database/sql"
	"fmt"
	"slices"
	"sync"
)

func Run(collectors []Collector, db *sql.DB) {
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	diagnostics := []Diagnostic{}

	for _, collector := range collectors {
		wg.Add(1)
		go func(collector Collector) {
			defer wg.Done()
			diagnostic := collector.Collect(db)
			mu.Lock()
			diagnostics = append(diagnostics, diagnostic)
			mu.Unlock()
		}(collector)
	}
	wg.Wait()

	slices.SortFunc(diagnostics, func(a, b Diagnostic) int {
		return a.ID - b.ID
	})

	for _, diagnostic := range diagnostics {
		fmt.Printf("%s\n", diagnostic)
	}
}
