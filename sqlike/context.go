package sqlike

import (
	"context"

	"github.com/RevenueMonster/sqlike/sqlike/primitive"
)

const contextResolutionKey = "_sqlike_context_query"

func (tb *Database) InjectResolution(ctx context.Context, queries ...primitive.Group) context.Context {
	query := extractResolution(ctx)
	query = append(query, queries...)
	return context.WithValue(ctx, contextResolutionKey, query)
}

func (tb *Table) InjectResolution(ctx context.Context, queries ...primitive.Group) context.Context {
	query := extractResolution(ctx)
	query = append(query, queries...)
	return context.WithValue(ctx, contextResolutionKey, query)
}

func extractResolution(ctx context.Context) []primitive.Group {
	iQuery := ctx.Value(contextResolutionKey)
	query, ok := iQuery.([]primitive.Group)
	if ok {
		return query
	}
	return []primitive.Group{}
}
