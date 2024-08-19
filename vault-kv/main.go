// Package main provides the Vault kv (Key-Value) secrets engine module
//
// This module uses the Vault commands (CLI) to communicate with the Vault instance.
// The functions can be called from the dagger CLI or from one of the SDKs.
//
// Check the original documentation: https://developer.hashicorp.com/vault/docs/commands/
package main

import (
	"context"
	"vault-kv/internal/dagger"
)

// Contains the Vault K/V command options.
type VaultKv struct {
	// Vault address
	Address string
	// Vault access token used for login
	Token   string
	// `kv get` command option: mount
	Mount   string
	// path / key referencing the K/V secret
	Path    string
	// `kv get` output option: field
	Field   string
}

// Configure Vault address to connect to
func (m *VaultKv) NewForAddress(
	// Vault address
	address string,
) (*VaultKv, error) {
	m.Address = address
	return m, nil
}

// Configure access token to be used for Vault login
func (m *VaultKv) Login(
	// Vault access token used for login
	token string,
) (*VaultKv, error) {
	m.Token = token
	return m, nil
}

// The `kv get` command retrieves the value from K/V secrets from Vault.
func (m *VaultKv) GetKV(
	ctx context.Context,
	// The path where the KV backend is mounted
	mount string,
	// path / key referencing the K/V secret
	path string,
	// Print only the field with the given name
	field string,
) (string, error) {
	return m.createContainer().
		WithExec([]string{"vault", "kv", "get", "-mount", mount, "-field", field, path}).
		Stdout(ctx)
}

// The `kv put` command creates a secret in Vault.
func (m *VaultKv) PutKV(
	ctx context.Context,
    // The path where the KV backend is mounted
    mount string,
	// path / key referencing the K/V secret
	path string,
	// Field to be created
	field string,
	// Value to be assigned
	value string,
) (string, error) {
	return m.createContainer().
		WithExec([]string{"vault", "kv", "put", "-mount", mount, path, field + "=" + value}).
		Stdout(ctx)
}

func (m *VaultKv) createContainer() *dagger.Container {
    return dag.Container().From("hashicorp/vault:1.17.3").
        WithEnvVariable("VAULT_ADDR", m.Address).
        WithEnvVariable("SKIP_SETCAP", "1").
        WithExec([]string{"vault", "login", "-non-interactive", m.Token})
}
