/**
 * Copyright © 2020 The With-Go Authors. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License.
 * You may not use this file except in compliance with the license
 * that can be found in the LICENSE.md file.
 */
package config

// Creates a new Config pointer.
func NewConfig() *Config {
	return &Config{ Modules: []*Module{} }
}

// Type Config defines a manager of modular configuration. It is a struct with Modules
// property as a slice of Module pointer.
type Config struct {
	Modules []*Module
}

// Appends a module configuration to the Config pointer and returns the Config pointer
// for chain-able functions call purposes.
func (c *Config) AppendModule(module *Module) *Config {
	c.Modules = append(c.Modules, module)
	return c
}

// Creates a named module configuration, appends it to the Config, and returns the Module
// pointer.
func (c *Config) CreateModule(name string) *Module {
	m := NewModule(name)
	c.AppendModule(m)
	return m
}

// Deletes a named module configuration from the Config pointer and returns the Config
// pointer for chain-able functions call purposes.
func (c *Config) DeleteModule(name string) *Config {
	for index, moduleConfig := range c.Modules {
		if moduleConfig.Name == name {
			c.Modules[index] = c.Modules[len(c.Modules)-1]
			c.Modules = c.Modules[:len(c.Modules)-1]
			return c
		}
	}
	return c
}

// Checks whether a named module configuration is exists in the Modules slice in the
// Config pointer. Returns true if the named module is exists, and false otherwise.
func (c *Config) IsModuleExists(name string) bool {
	for _, moduleConfig := range c.Modules {
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
func (c *Config) OnModule(name string) *Module {
	for _, moduleConfig := range c.Modules {
		if moduleConfig.Name == name {
			return moduleConfig
		}
	}
	return c.CreateModule(name)
}
