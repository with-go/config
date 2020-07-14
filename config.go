// Copyright © 2020 The With-Go Authors. All rights reserved.
// Licensed under the BSD 3-Clause License.
// You may not use this file except in compliance with the license
// that can be found in the LICENSE.md file.

package config

// Creates a new Config pointer.
func New() *Config {
	return &Config{ Modules: []*Module{} }
}

// Type Config defines a manager of modular configuration. It is a struct with Modules
// property as a slice of Module pointer.
type Config struct {
	Modules []*Module
}

// Appends a module configuration to the Config pointer and returns the Config pointer
// for chain-able functions call purposes.
func (config *Config) AppendModule(module *Module) *Config {
	if config.IsModuleExists(module.Name) {
		config.DeleteModule(module.Name)
	}
	config.Modules = append(config.Modules, module)
	return config
}

// Creates a named module configuration, appends it to the Config, and returns the Module
// pointer.
func (config *Config) CreateModule(name string) *Module {
	module := NewModule(name)
	config.AppendModule(module)
	return module
}

// Deletes a named module configuration from the Config pointer and returns the Config
// pointer for chain-able functions call purposes.
func (config *Config) DeleteModule(name string) *Config {
	for index, moduleConfig := range config.Modules {
		if moduleConfig.Name == name {
			config.Modules[index] = config.Modules[len(config.Modules)-1]
			config.Modules = config.Modules[:len(config.Modules)-1]
			return config
		}
	}
	return config
}

// Checks whether a named module configuration is exists in the Modules slice in the
// Config pointer. Returns true if the named module is exists, and false otherwise.
func (config *Config) IsModuleExists(name string) bool {
	for _, moduleConfig := range config.Modules {
		if moduleConfig.Name == name {
			return true
		}
	}
	return false
}

// Get a named module configuration from the Config pointer. If the module configuration
// with specified name does not exist in the Config's Modules slice, it will then try to
// create a new Module with that name, so that the returns value is predicted will not be
// nil. Returns a Module pointer.
func (config *Config) OnModule(name string) *Module {
	for _, moduleConfig := range config.Modules {
		if moduleConfig.Name == name {
			return moduleConfig
		}
	}
	return config.CreateModule(name)
}
