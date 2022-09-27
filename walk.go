package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type WalkCtx struct {
	Path      string
	Prefix    string
	StartDate int64
}

func (c *WalkCtx) walkFunc(fpath string, info fs.FileInfo, err error) error {
	if err != nil {
		log.Printf("Walk path %s got error: %v", fpath, err)
		return nil
	}
	if info.IsDir() {
		// Skip dir
		return nil
	}
	cf, err := ReadCacheFile(fpath)
	if err != nil {
		log.Println("Get file key got error:", err)
		return nil
	}
	if cf.CheckPurge(c.Prefix, c.StartDate) {
		log.Println("Should Delete File:", fpath)
		c.deleteFile(fpath)
	}
	return nil
}

func (c *WalkCtx) deleteFile(path string) {
	os.Remove(path)
}

func (c *WalkCtx) getFilePath(path string) string {
	root := strings.TrimSuffix(c.Path, "/")
	return fmt.Sprintf("%s%s", root, path)
}

func (c *WalkCtx) Walk() {
	err := filepath.Walk(c.Path, c.walkFunc)
	if err != nil {
		log.Println("Walk directory got error:", err)
	}
}
