package db

import (
	"sync"

	"github.com/Masterminds/squirrel"
)

var queryBuildOnce sync.Once

type QueryBuilder struct {
	squirrel.StatementBuilderType
}

var psql QueryBuilder

func NewQueryBuilder() QueryBuilder {
	queryBuildOnce.Do(func() {
		psql = QueryBuilder{
			StatementBuilderType: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		}
	})
	return psql
}
