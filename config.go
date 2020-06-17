package config

func NewConfig() *Config {
	return &Config{ Modules: []*Module{} }
}

type Config struct {
	Modules []*Module
}

func (c *Config) AppendModule(module *Module) *Config {
	c.Modules = append(c.Modules, module)
	return c
}

func (c *Config) CreateModule(name string) *Module {
	m := NewModule(name)
	c.AppendModule(m)
	return m
}

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

func (c *Config) IsModuleExists(name string) bool {
	for _, moduleConfig := range c.Modules {
		if moduleConfig.Name == name {
			return true
		}
	}
	return false
}

func (c *Config) OnModule(name string) *Module {
	for _, moduleConfig := range c.Modules {
		if moduleConfig.Name == name {
			return moduleConfig
		}
	}
	return c.CreateModule(name)
}
