package main

import (
	"context"
)

type VaultKv struct{}

type VaultContainer struct {
	*Container
}

func (m *VaultKv) New(address string) *VaultContainer {
	return &VaultContainer{dag.Container().From("vault:latest").WithEnvVariable("VAULT_ADDR", address)}
}

func (c *VaultContainer) Login(ctx context.Context, token string) *VaultContainer {
	return &VaultContainer{c.WithExec([]string{"vault", "login", "-non-interactive", token})}
}

func (c *VaultContainer) GetKV(ctx context.Context, mount string, path string, field string) (string, error) {
	return c.WithExec([]string{"vault", "kv", "-mount", mount, "-field", field, path}).Stdout(ctx)
}
