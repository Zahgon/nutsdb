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

import (
	"sync"
	"time"
)

// Clock provides time operations for TTL calculations.
// All timestamps are in milliseconds for internal consistency.
type Clock interface {
	// NowMillis returns the current time in milliseconds since Unix epoch.
	NowMillis() int64

	// NowSeconds returns the current time in seconds since Unix epoch.
	NowSeconds() int64

	// SetOnAdvance sets a callback to be called when the clock time is advanced.
	// This is primarily used by MockClock for testing.
	SetOnAdvance(onAdvance func(newTimeMillis int64))

	// AdvanceTime moves the clock forward by the specified duration.
	AdvanceTime(duration time.Duration)

	// SetTime sets the clock to a specific time in milliseconds since Unix epoch.
	SetTime(timeMillis int64)
}

// RealClock implements Clock using system time.
type RealClock struct{}

// NewRealClock creates a new RealClock instance.
func NewRealClock() Clock {
	_ = "STUB: not implemented"
	return *

	// NowMillis returns the current system time in milliseconds since Unix epoch.
	new(Clock)
}

func (rc *RealClock) NowMillis() int64 { _ = "STUB: not implemented"; return 0 }

// NowSeconds returns the current system time in seconds since Unix epoch.
func (rc *RealClock) NowSeconds() int64 { _ = "STUB: not implemented"; return 0 }

func (rc *RealClock) SetOnAdvance(onAdvance func(newTimeMillis int64)) {
	_ = "STUB: not implemented"
	// No-op for RealClock
	return
}

func (rc *RealClock) AdvanceTime(duration time.Duration) {
	_ = "STUB: not implemented"
	// No-op for RealClock
	return
}

func (rc *RealClock) SetTime(timeMillis int64) {
	_ = "STUB: not implemented"
	// No-op for RealClock
	return
}

// MockClock implements Clock with controllable time for testing.
type MockClock struct {
	mu        sync.RWMutex
	time      int64 // milliseconds since Unix epoch
	onAdvance func(newTimeMillis int64)
}

// NewMockClock creates a new MockClock instance with the specified initial time.
func NewMockClock(initialTime int64) Clock { _ = "STUB: not implemented"; return *new(Clock) }

// SetOnAdvance sets a callback to be called when the clock time is advanced.
func (mc *MockClock) SetOnAdvance(onAdvance func(newTimeMillis int64)) {
	_ = "STUB: not implemented"
	return
}

// NowMillis returns the current mock time in milliseconds since Unix epoch.
func (mc *MockClock) NowMillis() int64 { _ = "STUB: not implemented"; return 0 }

// NowSeconds returns the current mock time in seconds since Unix epoch.
func (mc *MockClock) NowSeconds() int64 { _ = "STUB: not implemented"; return 0 }

// AdvanceTime moves the clock forward by the specified duration.
func (mc *MockClock) AdvanceTime(duration time.Duration) { _ = "STUB: not implemented"; return }

// SetTime sets the clock to a specific time in milliseconds since Unix epoch.
func (mc *MockClock) SetTime(timeMillis int64) { _ = "STUB: not implemented"; return }
