package test

import "github.com/stellaraf/go-utils/environment"

type Environment struct {
	URL      string `env:"NFA_URL"`
	Username string `env:"NFA_USERNAME"`
	Password string `env:"NFA_PASSWORD"`
}

var Env Environment

func init() {
	err := environment.Load(&Env)
	if err != nil {
		panic(err)
	}
}
