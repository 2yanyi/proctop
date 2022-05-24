package cat

import (
	"github.com/json-iterator/go"
	"github.com/matsuwin/cat/internal"
	"os"
)

func SizeFormat(bytes float64) (_ string) { return internal.SizeFormat(bytes) }

func Stderr(err string) { internal.Stderr(err) }

func String(fp string) string { return internal.String(&fp) }

func Bytes(fp string) []byte { return internal.Bytes(&fp) }

func Wcl(fp string) int { return internal.Wcl(&fp) }

func MD5sumChunked(fp string) (os.FileInfo, string, error) { return internal.MD5sumChunked(&fp) }

func FileExist(fp string) bool { return internal.FileExist(fp) }

func CommandArgs(dir string, args []string) (_ string) { return internal.CommandArgs(dir, args) }

func Json(a interface{}) []byte { return internal.Json(a) }

func JsonIter() jsoniter.API { return internal.JsonIter() }

func LanAddress() []string { return internal.LanAddress() }

func SystemInfo() *internal.Environment { return internal.SystemInfo() }
