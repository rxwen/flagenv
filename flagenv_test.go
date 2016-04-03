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
	defaultValue string
	result       string
}{
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
		var variable string
		for key, value := range testCase.envs {
			_ = os.Setenv(key, value)
		}
		flagenv.StringVar(&variable, testCase.flagName, testCase.defaultValue, "")
		flagenv.Parse()
		if strings.Compare(testCase.result, variable) != 0 {
			t.Error(fmt.Sprintf("expect %s, got %s", testCase.result, variable))
		}
		for key, _ := range testCase.envs {
			_ = os.Unsetenv(key)
		}
	}
}
