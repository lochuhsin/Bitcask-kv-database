package main

type stringArrFlags []string

func (s *stringArrFlags) String() string {
	return "an array of flags"
}

func (s *stringArrFlags) Set(path string) error {
	*s = append(*s, path)
	return nil
}

var envPaths stringArrFlags
