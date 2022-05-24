package internal

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/json-iterator/go"
	"github.com/pkg/errors"
	"io"
	"math"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"unsafe"

	_ "github.com/matsuwin/errcause"
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

func String(fp *string) string {
	a := Bytes(fp)
	if len(a) > 32 {
		return *(*string)(unsafe.Pointer(&a))
	}
	return string(a)
}

func Bytes(fp *string) []byte {
	data, err := os.ReadFile(*fp)
	if err != nil {
		Stderr(err.Error())
	}
	return data
}

func Wcl(fp *string) int {
	fis, err := os.Open(*fp)
	if err != nil {
		return 0
	}
	defer fis.Close()

	buf := make([]byte, 1024*32)
	sep := []byte{'\n'}
	wcl := 0
	n := 0
	for {
		n, err = fis.Read(buf)
		wcl += bytes.Count(buf[:n], sep)
		switch {
		case err == io.EOF:
			return wcl
		case err != nil:
			return wcl
		}
	}
}

func MD5sumChunked(fp *string) (os.FileInfo, string, error) {
	fis, err := os.Open(*fp)
	if err != nil {
		return nil, "", errors.New(err.Error())
	}
	defer fis.Close()

	info, _ := fis.Stat()
	if info.IsDir() {
		return info, "", errors.New(fmt.Sprintf("%s is a directory", *fp))
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
		if _, err = io.WriteString(hash, string(buf)); err != nil {
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

func CommandArgs(dir string, args []string) (_ string) {
	if len(args) != 0 {
		stdout, stderr := &bytes.Buffer{}, &bytes.Buffer{}
		cmd := exec.Command(args[0])
		cmd.Stdout, cmd.Stderr = stdout, stderr
		cmd.Args = args
		if dir != "" {
			cmd.Dir = filepath.Dir(dir)
		}
		_ = cmd.Run()
		return strings.TrimSpace(commandMerge(stdout, stderr))
	}
	return
}

func Json(a interface{}) []byte {
	data, err := JsonIter().MarshalIndent(a, "", "  ")
	if err != nil {
		Stderr(err.Error())
	}
	return data
}

func JsonIter() jsoniter.API {
	return jsoniter.ConfigFastest
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

const _KB = 1024
const _MB = 1024 * 1024
const _GB = 1024 * 1024 * 1024
