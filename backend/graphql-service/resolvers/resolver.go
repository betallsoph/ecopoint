package resolvers

import (
	"ecopoint-backend/shared/config"
	"ecopoint-backend/graphql-service/clients"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Config      *config.Config
	GRPCClients *clients.GRPCClients
}
