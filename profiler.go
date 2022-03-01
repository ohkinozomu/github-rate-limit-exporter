package main

import (
	"context"
	"errors"
	"os"

	ncp "github.com/ohkinozomu/neutral-cp"
)

func StartProfiler() error {
	var registry ncp.Registry
	if os.Getenv("PROFILER_REGISTRY") == "pyroscope" {
		registry = ncp.PYROSCOPE
	} else if os.Getenv("PROFILER_REGISTRY") == "cloud_profiler" {
		registry = ncp.CLOUD_PROFILER
	} else {
		return errors.New("undefined profiler registry")
	}

	c := ncp.Config{
		Registry:        registry,
		ApplicationName: "github-rate-limit-exporter",
		ServerAddress:   os.Getenv("SERVER_ADDRESS"),
		Version:         "0.0.8",
	}
	ncp := ncp.NeutralCP{Config: c}

	ctx := context.Background()
	err := ncp.Start(ctx)
	if err != nil {
		return err
	}
	return nil
}
