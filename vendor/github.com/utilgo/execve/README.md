# execve

```go
func main() {
    text := execve.Args("", []string{"ls", "-l"})
    fmt.Println(text)
}
```
