/**
 * Copyright © 2020 The With-Go Authors. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License.
 * You may not use this file except in compliance with the license
 * that can be found in the LICENSE.md file.
 */
package config

import (
	"github.com/joho/godotenv"
	"os"
)

// Load configuration based on OS environment variables or
// based on .env files using github.com/joho/godotenv package.
//
// The load order of .env files is following the convention for
// managing multiple environments.
// Ref: https://github.com/bkeepers/dotenv#what-other-env-files-can-i-use
func init() {
	env := os.Getenv("GO_ENV")
	if "" == env {
		env = "development"
	}
	_ = godotenv.Load(".env." + env + ".local")
	if "test" != env {
		_ = godotenv.Load(".env.local")
	}
	_ = godotenv.Load(".env." + env)
	_ = godotenv.Load()
}