package flagenv

import (
	"os"
	"strconv"
)

type flagRecord struct {
	Name         string
	DefaultValue interface{}
	Usage        string
	Value        interface{}
}

var flagRecords = make(map[string]flagRecord)

// Parse parses the flags from os.Environ.  Must be called
// after all flags are defined and before flags are accessed by the program.
func Parse() {
	for _, record := range flagRecords {
		env, exist := os.LookupEnv(record.Name)
		if exist {
			switch record.Value.(type) {
			case *string:
				*record.Value.(*string) = env
			case *int:
				val, _ := strconv.Atoi(env)
				*record.Value.(*int) = val
			case *bool:
				val, _ := strconv.ParseBool(env)
				*record.Value.(*bool) = val
			}
		} else {
			switch record.Value.(type) {
			case *string:
				*record.Value.(*string) = record.DefaultValue.(string)
			case *int:
				*record.Value.(*int) = record.DefaultValue.(int)
			case *bool:
				*record.Value.(*bool) = record.DefaultValue.(bool)
			}
		}
	}
}

// StringVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func StringVar(p *string, name string, value string, usage string) {
	flagRecords[name] = flagRecord{
		Name:         name,
		DefaultValue: value,
		Usage:        usage,
		Value:        p,
	}
}

// IntVar defines a int flag with specified name, default value, and usage string.
// The argument p points to a int variable in which to store the value of the flag.
func IntVar(p *int, name string, value int, usage string) {
	flagRecords[name] = flagRecord{
		Name:         name,
		DefaultValue: value,
		Usage:        usage,
		Value:        p,
	}
}

// BoolVar defines a int flag with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
func BoolVar(p *bool, name string, value bool, usage string) {
	flagRecords[name] = flagRecord{
		Name:         name,
		DefaultValue: value,
		Usage:        usage,
		Value:        p,
	}
}
