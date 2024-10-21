package terminal

import (
	"fmt"
	"log"
	"os/exec"
	"reflect"
	"runtime"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func Confirm(msg string) bool {
	ok := true
	prompt := &survey.Confirm{
		Message: msg,
		Default: true,
	}
	if err := survey.AskOne(prompt, &ok); err != nil {
		panic(err)
	}
	return ok
}

// return index of array opts selected
func SelectIndexAnswer(msg string, opts []string, pageSize int) int {
	if len(opts) == 0 {
		log.Panicf("nothing to select")
	}
	ans := -1 // non select, -1 index
	prompt := &survey.Select{
		Message: msg,
		Options: opts,
		VimMode: true,
	}
	askOpts := []survey.AskOpt{survey.WithPageSize(pageSize)}
	if err := survey.AskOne(prompt, &ans, askOpts...); err != nil {
		panic(err)
	}
	return ans
}

func StringKeyVals(msg string, obj any) string {
	s := reflect.ValueOf(obj).Elem()
	typeOfT := s.Type()

	r := msg + ": \n"
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		r += fmt.Sprintf("  %s: %v\n", typeOfT.Field(i).Name, f.Interface())
	}
	return strings.TrimSpace(r)
}

func OpenURL(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
		args = []string{url}
	default: // "linux", "freebsd", "openbsd", "netbsd"
		// Check if running under WSL
		if isWSL() {
			// Use 'cmd.exe /c start' to open the URL in the default Windows browser
			cmd = "cmd.exe"
			args = []string{"/c", "start", url}
		} else {
			// Use xdg-open on native Linux environments
			cmd = "xdg-open"
			args = []string{url}
		}
	}
	if len(args) > 1 {
		// args[0] is used for 'start' command argument, to prevent issues with URLs starting with a quote
		args = append(args[:1], append([]string{""}, args[1:]...)...)
	}
	return exec.Command(cmd, args...).Start()
}

// isWSL checks if the Go program is running inside Windows Subsystem for Linux
func isWSL() bool {
	releaseData, err := exec.Command("uname", "-r").Output()
	if err != nil {
		return false
	}
	return strings.Contains(strings.ToLower(string(releaseData)), "microsoft")
}

func ShortenString(s string, maxLength int) string {
    s = strings.ReplaceAll(s, "\n", " ")
	if len(s) > maxLength {
		return s[:maxLength] + "..."
	}
	return s
}
