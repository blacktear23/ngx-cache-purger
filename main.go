package main

import (
	"flag"
	"fmt"
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

	ctx := &WalkCtx{
		Path:   path,
		Prefix: prefix,
	}

	walkPath(ctx)
}
