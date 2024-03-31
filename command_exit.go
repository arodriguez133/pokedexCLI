package main

import "os"

func callBackExit(cfg *config, args ...string) error {
	os.Exit(0)
	return nil
}
