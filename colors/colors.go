package colors

import "strings"

var j = strings.Join

func Black(s string) string   { return j([]string{"\u001B[0;30;48m", s, "\u001B[0m"}, "") } // 黑色(30)
func Red(s string) string     { return j([]string{"\u001B[0;31;48m", s, "\u001B[0m"}, "") } // 红色(31)
func Green(s string) string   { return j([]string{"\u001B[0;32;48m", s, "\u001B[0m"}, "") } // 绿色(32)
func Yellow(s string) string  { return j([]string{"\u001B[0;33;48m", s, "\u001B[0m"}, "") } // 黄色(33)
func Blue(s string) string    { return j([]string{"\u001B[0;34;48m", s, "\u001B[0m"}, "") } // 蓝色(34)
func Fuchsia(s string) string { return j([]string{"\u001B[0;35;48m", s, "\u001B[0m"}, "") } // 紫红色(35)
func Cyan(s string) string    { return j([]string{"\u001B[0;36;48m", s, "\u001B[0m"}, "") } // 青蓝色(36)
func White(s string) string   { return j([]string{"\u001B[0;37;48m", s, "\u001B[0m"}, "") } // 白色(37)
