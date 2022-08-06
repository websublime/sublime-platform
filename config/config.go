/*
Copyright © 2022 Websublime.dev organization@websublime.dev

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package config

import (
	"os"

	"github.com/spf13/viper"
)

type Environment struct {
	WsEnvironment string `env:"WS_ENV" mapstructure:"env"`
	WsApiUrl      string `env:"WS_API_URL" mapstructure:"api_url"`
	WsApiKey      string `env:"WS_API_KEY" mapstructure:"api_key"`
	WsApiSecret   string `env:"WS_API_SECRET" mapstructure:"api_secret"`
	WsHost        string `env:"WS_HOST" mapstructure:"host"`
	WsPort        string `env:"WS_PORT" mapstructure:"port"`
	IsProduction  bool   `env:"-" mapstructure:"is_production"`
}

func Config() Environment {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	env := Environment{}

	viper.AddConfigPath(dir)
	viper.SetEnvPrefix("ws")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	viper.Unmarshal(&env)

	if env.WsEnvironment == "production" {
		env.IsProduction = true
	} else {
		env.IsProduction = false
	}

	return env
}
