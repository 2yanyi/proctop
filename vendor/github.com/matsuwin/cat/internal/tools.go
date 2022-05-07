package internal

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/json-iterator/go"
	"github.com/matsuwin/stringx"
	"github.com/pkg/errors"
	"io"
	"math"
	"net"
	"os"
	"strings"
)

func SizeFormat(bytes float64) (_ string) {
	if bytes >= _GB {
		return fmt.Sprintf("%.1fG", bytes/_GB)
	} else if bytes >= _MB {
		return fmt.Sprintf("%.1fM", bytes/_MB)
	} else if bytes >= _KB {
		return fmt.Sprintf("%.1fK", bytes/_KB)
	}
	return
}

func Stderr(err string) {
	if err != "" {
		_, _ = fmt.Fprintf(os.Stderr, "error: %s", err)
	}
}

func String(fp string) string {
	return stringx.BytesToString(Bytes(fp))
}

func Bytes(fp string) []byte {
	data, err := os.ReadFile(fp)
	if err != nil {
		Stderr(err.Error())
	}
	return data
}

func MD5sumChunked(fp string) (os.FileInfo, string, error) {
	fis, err := os.Open(fp)
	if err != nil {
		return nil, "", errors.New(err.Error())
	}
	defer fis.Close()

	info, _ := fis.Stat()
	if info.IsDir() {
		return info, "", errors.New(fmt.Sprintf("%s is a directory", fp))
	}

	// Chunked calculations
	size := info.Size()
	blocks := uint64(math.Ceil(float64(size) / float64(_MB)))
	hash := md5.New()
	for i := uint64(0); i < blocks; i++ {
		blockSize := int(math.Min(_MB, float64(size-int64(i*_MB))))
		buf := make([]byte, blockSize)
		if _, err = fis.Read(buf); err != nil {
			return info, "", errors.New(err.Error())
		}
		if _, err = io.WriteString(hash, stringx.BytesToString(buf)); err != nil {
			return info, "", errors.New(err.Error())
		}
	}
	sum := hex.EncodeToString(hash.Sum(nil))

	return info, sum, nil
}

func FileExist(fp string) bool {
	_, err := os.Lstat(fp)
	return !os.IsNotExist(err)
}

func Json(a interface{}) []byte {
	data, err := JsonIter().MarshalIndent(a, "", "  ")
	if err != nil {
		Stderr(err.Error())
	}
	return data
}

func LanAddress() []string {
	address := make([]string, 0, 10)
	nia, _ := net.InterfaceAddrs()
	for i := range nia {
		if addr, has := nia[i].(*net.IPNet); has {
			ipv4 := addr.IP.String()
			if ipv4 == "127.0.0.1" || strings.Contains(ipv4, ":") {
				continue
			}
			address = append(address, ipv4)
		}
	}
	return address
}

func JsonIter() jsoniter.API {
	return jsoniter.ConfigFastest
}

const _KB = 1024
const _MB = 1024 * 1024
const _GB = 1024 * 1024 * 1024
