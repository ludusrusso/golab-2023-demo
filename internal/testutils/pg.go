package testutils

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
)

func InitPostgres() (string, PurgeFunc, error) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		return "", nil, fmt.Errorf("Could not connect to docker: %w", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := createResource(pool)
	if err != nil {
		return "", nil, fmt.Errorf("Could not start resource: %w", err)
	}
	resource.Expire(60)

	port := resource.GetPort("5432/tcp")
	connString := fmt.Sprintf("postgresql://test:test@localhost:%v/test", port)

	if err = waitForStart(pool, connString); err != nil {
		return "", nil, fmt.Errorf("Could not connect to docker: %w", err)
	}

	var purgeFunc PurgeFunc = func() error {
		err := pool.Purge(resource)
		if err != nil {
			return fmt.Errorf("Could not purge resource: %w", err)
		}
		return nil
	}

	return connString, purgeFunc, nil
}

func createResource(pool *dockertest.Pool) (*dockertest.Resource, error) {
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "13-alpine",
		Env: []string{
			"POSTGRES_USER=test",
			"POSTGRES_PASSWORD=test",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	return resource, err
}

func waitForStart(pool *dockertest.Pool, url string) error {
	return pool.Retry(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		db, err := pgxpool.New(ctx, url)
		if err != nil {
			return err
		}
		return db.Ping(ctx)
	})
}
