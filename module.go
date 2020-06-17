package config

import (
	"os"
	"strings"
)

func NewModule(name string) *Module {
	return &Module{ Name: name, Config: map[string]string{} }
}

type Module struct {
	Name	string
	Config	map[string]string
}

func (m *Module) Delete(key string) {
	delete(m.Config, key)
}

func (m *Module) Get(key string) string {
	value := os.Getenv(strings.ToUpper(m.Name) + "_" + strings.ToUpper(key))
	if value != "" {
		return value
	}
	for k, v := range m.Config {
		if k == key {
			return v
		}
	}
	return value
}

func (m *Module) GetWithDefault(key string, defaultValue string) string {
	value := m.Get(key)
	if value != "" {
		return value
	}
	return defaultValue
}

func (m *Module) Set(key string, value string) {
	m.Config[key] = value
}
