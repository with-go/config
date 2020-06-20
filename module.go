// Copyright © 2020 The With-Go Authors. All rights reserved.
// Licensed under the BSD 3-Clause License.
// You may not use this file except in compliance with the license
// that can be found in the LICENSE.md file.

package config

import (
	"os"
	"strings"
)

func NewModule(name string) *Module {
	return &Module{ Name: name, Config: map[string]string{} }
}

// Type Module defines a named modular configuration.
//
// Property Name defines the Module's Name as string.
//
// Property Config defines the key-value pair of configuration list as map string
// of string.
type Module struct {
	Name	string
	Config	map[string]string
}

// Delete module configuration with specified key.
func (m *Module) Delete(key string) {
	delete(m.Config, key)
}

// Get module configuration with specified key.
// Returns the string value of the module configuration. When no configuration is
// saved with that specified key, it will returns an empty string.
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

// The same as Get(key string), but it will returns defaultValue instead of empty
// string when no configuration is saved with that specified key.
func (m *Module) GetWithDefault(key string, defaultValue string) string {
	value := m.Get(key)
	if value != "" {
		return value
	}
	return defaultValue
}

// Set module configuration with specified key to a specified value.
func (m *Module) Set(key string, value string) {
	m.Config[key] = value
}
