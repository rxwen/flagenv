package flagenv_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/rxwen/flagenv"
)

var testCases = []struct {
	envs         map[string]string
	flagName     string
	defaultValue interface{}
	result       interface{}
}{
	{
		map[string]string{
			"hello": "42",
		},
		"hello",
		3,
		42,
	},
	{
		map[string]string{},
		"hello",
		3,
		3,
	},
	{
		map[string]string{
			"hello": "zzz",
		},
		"hello",
		"mmm",
		"zzz",
	},
	{
		map[string]string{},
		"hello",
		"mmm",
		"mmm",
	},
}

func TestFlagenv(t *testing.T) {
	for _, testCase := range testCases {
		for key, value := range testCase.envs {
			_ = os.Setenv(key, value)
		}
		switch testCase.result.(type) {
		case int:
			var variable int
			flagenv.IntVar(&variable, testCase.flagName, testCase.defaultValue.(int), "")
			flagenv.Parse()
			fmt.Println("variable is %d", variable)
			if testCase.result.(int) != variable {
				t.Error(fmt.Sprintf("expect %d, got %d", testCase.result, variable))
			}
		case string:
			var variable string
			flagenv.StringVar(&variable, testCase.flagName, testCase.defaultValue.(string), "")
			flagenv.Parse()
			if strings.Compare(testCase.result.(string), variable) != 0 {
				t.Error(fmt.Sprintf("expect %s, got %s", testCase.result, variable))
			}
		}
		for key, _ := range testCase.envs {
			_ = os.Unsetenv(key)
		}
	}
}
