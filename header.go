package main

import (
	"encoding/binary"
	"fmt"
	"time"
)

type NgxCacheHeader struct {
	Version uint64
	Date    int64
	Key     string
}

func (h *NgxCacheHeader) ParseHeader(hdr []byte) error {
	ver := binary.LittleEndian.Uint64(hdr[0:8])
	date := binary.LittleEndian.Uint64(hdr[40:48])
	h.Version = ver
	h.Date = int64(date)
	return nil
}

func (h *NgxCacheHeader) String() string {
	date := time.Unix(h.Date, 0)
	return fmt.Sprintf("Date: %v, %s", date, h.Key)
}
