package logger

import "fmt"

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"

func Info(lg string) string {
	lg = fmt.Sprintf("%sℹ %s%s", Blue, lg, Reset)
	fmt.Printf("%s\n", lg)
	return lg
}

func Error(lg string) string {
	lg = fmt.Sprintf("%s✖️ %s %s", Red, lg, Reset)
	fmt.Printf("%s\n", lg)
	return lg
}

func Success(lg string) string {
	lg = fmt.Sprintf("%s✔ %s %s", Green, lg, Reset)
	fmt.Printf("%s\n", lg)
	return lg
}
