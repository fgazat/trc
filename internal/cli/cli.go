package cli

import (
	"fmt"
	"reflect"
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
