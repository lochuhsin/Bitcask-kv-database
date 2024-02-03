package main

import "flag"

type stringArrFlags []string

func (s *stringArrFlags) String() string {
	return "an array of flags"
}

func (s *stringArrFlags) Set(path string) error {
	*s = append(*s, path)
	return nil
}

type flagOption func(*Flags)

type Flags struct {
	envPaths []string
}

func NewFlags(flagOpts ...flagOption) Flags {
	flags := Flags{}
	for _, fn := range flagOpts {
		fn(&flags)
	}
	return flags
}

func setEnvPaths(envPaths stringArrFlags) flagOption {
	return func(f *Flags) {
		f.envPaths = []string(envPaths)
	}
}

func ParseFlags() Flags {
	var envPaths stringArrFlags
	flag.Var(&envPaths, "envfiles", "Specifies the env files")
	flag.Parse()

	return NewFlags(
		setEnvPaths(envPaths),
	)
}
