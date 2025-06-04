// Package main provides tests for the Vault kv (Key-Value) secrets engine module
//
// The functions can be called from the dagger CLI or from one of the SDKs.
package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"main/internal/dagger"
	"time"
)

type Tests struct {
	// Predefined token to be used.
	Token string
}

func New() *Tests {
	return &Tests{
		Token: "sesame",
	}
}

// The `test` command creates and starts a vault server instance, creates a new secret and reads it afterwards.
func (m *Tests) Test(ctx context.Context) error {
	secretPath := "/secret/test"
	secretKey := fmt.Sprintf("%x", (md5.Sum([]byte(time.Now().String()))))[0:8]
	secretValue := "expected"
	service := m.vaultServer()
	service.Start(ctx)
	endpoint, err := service.Endpoint(ctx)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("http://%s", endpoint)
	dag.VaultKv().NewForAddress(url).Login(m.Token).PutKv(ctx, "", secretPath, secretKey, secretValue)
	actualSecretValue, err := dag.VaultKv().NewForAddress(url).Login(m.Token).GetKv(ctx, "", secretPath, secretKey)
	defer service.Stop(ctx)
	if err != nil {
		return err
	}
	if actualSecretValue != secretValue {
		return fmt.Errorf("Unexpected secret value '%s' for key '%s'", actualSecretValue, secretKey)
	}
	return nil
}

func (m *Tests) vaultServer() *dagger.Service {
	return dag.Container().
		From("hashicorp/vault:1.19.5").
		WithEnvVariable("VAULT_DEV_ROOT_TOKEN_ID", m.Token).
		WithEnvVariable("VAULT_DEV_LISTEN_ADDRESS", "0.0.0.0:8200").
		WithDefaultArgs([]string{"vault", "server", "-dev"}).
		WithExposedPort(8200).
		AsService()
}
