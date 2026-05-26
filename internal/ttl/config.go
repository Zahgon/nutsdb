// Copyright 2025 The nutsdb Author. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ttl

import "time"

// Config contains TTL-related configuration for batch deletion.
type Config struct {
	BatchSize    int           // Batch size for deletion (default: 100, optimal balance)
	BatchTimeout time.Duration // Batch timeout (default: 1s, reasonable latency tolerance)
	QueueSize    int           // Event queue size (default: 1000, handles burst traffic)

	// Timing wheel configuration
	EnableTimingWheel bool          // Enable active deletion via timing wheel (default: true, <1% overhead)
	WheelSlotDuration time.Duration // Duration of each wheel slot (default: 1s, optimal performance-precision balance)
	WheelSize         int           // Number of slots in the wheel (default: 3600, covers 1 hour with minimal overhead)
}

// DefaultConfig returns the default TTL configuration.
func DefaultConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

// Validate validates the configuration and fills in defaults for zero values.
func (c *Config) Validate() { _ = "STUB: not implemented"; return }
