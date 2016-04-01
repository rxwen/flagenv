package flagenv

import (
	"os"
)

type flagRecord struct {
	Name         string
	DefaultValue string
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
			*record.Value.(*string) = env
		} else {
			*record.Value.(*string) = record.DefaultValue
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
