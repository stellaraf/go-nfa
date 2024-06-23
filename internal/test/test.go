package test

import "github.com/stellaraf/go-utils/environment"

type Test struct {
	Prefixes []string `env:"PREFIXES"`
	Exclude  []string `env:"EXCLUDE"`
}

type Environment struct {
	URL      string `env:"NFA_URL"`
	Username string `env:"NFA_USERNAME"`
	Password string `env:"NFA_PASSWORD"`
	Test     Test   `envPrefix:"NFA_TEST_"`
}

var Env Environment

func init() {
	err := environment.Load(&Env)
	if err != nil {
		panic(err)
	}
}
