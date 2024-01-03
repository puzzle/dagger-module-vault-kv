package main

import (
	"context"
)

type VaultKv struct {
	Address string
	Token   string
	Mount   string
	Path    string
	Field   string
}

func (c *VaultKv) NewForAddress(address string) (*VaultKv, error) {
	c.Address = address
	return c, nil
}

func (c *VaultKv) Login(token string) (*VaultKv, error) {
	c.Token = token
	return c, nil
}

func (c *VaultKv) GetKV(ctx context.Context, mount string, path string, field string) (string, error) {
	return dag.Container().From("vault:1.13.3").
		WithEnvVariable("VAULT_ADDR", c.Address).
		WithEnvVariable("SKIP_SETCAP", "1").
		WithExec([]string{"vault", "login", "-non-interactive", c.Token}).
		WithExec([]string{"vault", "kv", "get", "-mount", mount, "-field", field, path}).Stdout(ctx)
}
