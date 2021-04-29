package main

import "os"

func runFeatures() {
	if os.Getenv("CONTINUOUS_PROFILING") == "enabled" {
		err := StartProfiler()
		if err != nil {
			panic(err)
		}
	}
}
