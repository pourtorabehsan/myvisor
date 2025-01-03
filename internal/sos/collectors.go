package sos

import (
	"github.com/pourtorabehsan/myvisor/internal/advisor"
)

func Collectors() []advisor.Collector {
	return []advisor.Collector{
		&TopQueriesCountCollector{},
		&TopQueriesTimeCollector{},
	}
}
