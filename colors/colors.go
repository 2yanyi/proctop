package colors

import "strings"

const _H = "\u001B["
const _T = "\u001B[0m"

var j = strings.Join

func Black(s, m string) string   { return j([]string{_H, m, ";30;48m", s, _T}, "") } // 黑色
func Red(s, m string) string     { return j([]string{_H, m, ";31;48m", s, _T}, "") } // 红色
func Green(s, m string) string   { return j([]string{_H, m, ";32;48m", s, _T}, "") } // 绿色
func Yellow(s, m string) string  { return j([]string{_H, m, ";33;48m", s, _T}, "") } // 黄色
func Blue(s, m string) string    { return j([]string{_H, m, ";34;48m", s, _T}, "") } // 蓝色
func Fuchsia(s, m string) string { return j([]string{_H, m, ";35;48m", s, _T}, "") } // 紫红色
func Cyan(s, m string) string    { return j([]string{_H, m, ";36;48m", s, _T}, "") } // 青蓝色
func White(s, m string) string   { return j([]string{_H, m, ";37;48m", s, _T}, "") } // 白色

const (
	Zero          = "0"
	Bold          = "1" // 粗体
	Dark          = "2" // 暗色
	Italic        = "3" // 斜体
	Underscore    = "4" // 下划线
	Flicker       = "5" // 闪烁
	Flicker2      = "6" // 闪烁
	Invert        = "7" // 颜色取反
	Hide          = "8" // 隐藏
	Strikethrough = "9" // 删除线
)
