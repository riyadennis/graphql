//go:generate go run github.com/99designs/gqlgen

package graph

import "github.com/riyadennis/graphql/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	countries []*model.Country
}


