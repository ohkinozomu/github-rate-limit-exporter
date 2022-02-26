package main

import (
	"context"
	"os"

	ncp "github.com/ohkinozomu/neutral-cp"
)

func StartProfiler() error {
	c := ncp.Config{
		Registry:        ncp.PYROSCOPE,
		ApplicationName: "github-rate-limit-exporter",
		ServerAddress:   os.Getenv("SERVER_ADDRESS"),
	}
	ncp := ncp.NeutralCP{Config: c}

	ctx := context.Background()
	err := ncp.Start(ctx)
	if err != nil {
		return err
	}
	return nil
}
