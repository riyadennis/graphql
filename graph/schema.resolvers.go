package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/riyadennis/graphql/graph/generated"
	"github.com/riyadennis/graphql/graph/model"
)

func (r *queryResolver) Countries(ctx context.Context) ([]*model.Country, error) {
	return []*model.Country{
		{
			"United Kingdom",
		},
		{
			"United States",
		},
		{
			"India",
		},
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
