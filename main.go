package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var (
		path   string
		prefix string
	)
	flag.StringVar(&path, "p", "", "Cache path")
	flag.StringVar(&prefix, "k", "", "Key Prefix")

	flag.Parse()

	if path == "" || prefix == "" {
		fmt.Println("Require path and prefix parameter")
		return
	}

	now := time.Now()
	nowUnix := now.Unix()

	ctx := &WalkCtx{
		Path:      path,
		Prefix:    prefix,
		StartDate: nowUnix,
	}

	fmt.Printf("Nginx Cache Path: %s\n", path)
	fmt.Printf("Purge Key Prefix: %s\n", prefix)
	fmt.Printf("Purge Time: %v\n", now)
	fmt.Printf("Start Purge\n")

	ctx.Walk()
}
