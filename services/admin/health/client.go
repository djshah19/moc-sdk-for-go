// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license

package health

import (
	"context"

	"github.com/microsoft/moc-sdk-for-go/services/admin/health/internal"
	"github.com/microsoft/moc/pkg/auth"
	"github.com/microsoft/moc/rpc/common"
)

// Service interfacetype Service interface {
type Service interface {
	CheckHealth(ctx context.Context, timeoutSeconds uint32) error
	GetAgentInfo(context.Context) (*common.NodeInfo, error)
}

// Client structure
type HealthClient struct {
	internal Service
}

// NewClient method returns new client
func NewHealthClient(cloudFQDN string, authorizer auth.Authorizer) (*HealthClient, error) {
	c, err := internal.NewHealthClient(cloudFQDN, authorizer)
	return &HealthClient{c}, err
}

// CheckHealth
func (c *HealthClient) CheckHealth(ctx context.Context, timeoutSeconds uint32) error {
	return c.internal.CheckHealth(ctx, timeoutSeconds)
}

// GetAgentInfo
func (c *HealthClient) GetAgentInfo(ctx context.Context) (*common.NodeInfo, error) {
	return c.internal.GetAgentInfo(ctx)
}
