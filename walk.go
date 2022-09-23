package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	HEADER_SIZE = 0x151 // 337
	HEADER_POS  = 0x150 // 336
)

var (
	errInvalidHeader = errors.New("Invalid Header")
	errInvalidKey    = errors.New("Invalid Key")
)

type WalkCtx struct {
	Path   string
	Prefix string
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
	rkey, err := c.getFileKey(fpath)
	if err != nil {
		log.Println("Get file key got error:", err)
		return nil
	}
	if c.checkKey(rkey) {
		log.Println("Should Delete File:", fpath)
		c.deleteFile(fpath)
	}
	return nil
}

func (c *WalkCtx) checkKey(rkey string) bool {
	key := strings.TrimPrefix(rkey, "KEY: ")
	return strings.HasPrefix(key, c.Prefix)
}

func (c *WalkCtx) getFileKey(path string) (string, error) {
	fp, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer fp.Close()
	buf := make([]byte, 4096)
	count, err := fp.Read(buf)
	if err != nil {
		return "", err
	}
	if count < HEADER_SIZE {
		return "", errInvalidHeader
	}
	if buf[HEADER_POS] != '\n' {
		return "", errInvalidHeader
	}
	var keyEnd int = 0
	for i := HEADER_POS + 1; i < count-1; i++ {
		char := buf[i]
		if char == '\n' {
			keyEnd = i - 1
			break
		}
	}
	if keyEnd < HEADER_POS+1 {
		return "", errInvalidKey
	}
	keyBuf := buf[HEADER_POS+1 : keyEnd]
	return string(keyBuf), nil
}

func (c *WalkCtx) deleteFile(path string) {
	os.Remove(path)
}

func (c *WalkCtx) getFilePath(path string) string {
	root := strings.TrimSuffix(c.Path, "/")
	return fmt.Sprintf("%s%s", root, path)
}

func walkPath(ctx *WalkCtx) {
	err := filepath.Walk(ctx.Path, ctx.walkFunc)
	if err != nil {
		log.Println("Walk directory got error:", err)
	}
}
