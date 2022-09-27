package main

import (
	"errors"
	"os"
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

type CacheFile struct {
	FileName string
	Header   *NgxCacheHeader
}

func ReadCacheFile(fname string) (*CacheFile, error) {
	fp, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	buf := make([]byte, 4096)
	count, err := fp.Read(buf)
	if err != nil {
		return nil, err
	}
	if count < HEADER_SIZE {
		return nil, errInvalidHeader
	}
	if buf[HEADER_POS] != '\n' {
		return nil, errInvalidHeader
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
		return nil, errInvalidKey
	}
	keyBuf := buf[HEADER_POS+1 : keyEnd]
	key := string(keyBuf)
	hdr := &NgxCacheHeader{
		Key: key,
	}
	hdr.ParseHeader(buf[0:HEADER_POS])
	return &CacheFile{
		FileName: fname,
		Header:   hdr,
	}, nil
}

func (f *CacheFile) checkKeyPrefix(prefix string) bool {
	key := strings.TrimPrefix(f.Header.Key, "KEY: ")
	return strings.HasPrefix(key, prefix)
}

func (f *CacheFile) checkDate(date int64) bool {
	// Any of them got 0 should delete
	if date == 0 || f.Header.Date == 0 {
		return true
	}
	// Delete file date less or equals than check date
	return date >= f.Header.Date
}

func (f *CacheFile) CheckPurge(keyPrefix string, date int64) bool {
	keyMatch := f.checkKeyPrefix(keyPrefix)
	if keyMatch {
		return f.checkDate(date)
	}
	return false
}
