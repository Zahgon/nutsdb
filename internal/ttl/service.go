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
	"context"
	"time"

	"github.com/nutsdb/nutsdb/internal/core"
)

type BatchExpiredCallback func(events []*ExpirationEvent)

// Service provides TTL checking and batched expiration deletion.
// Expired keys are detected lazily during reads and collected into a queue,
// then batch-deleted by a background goroutine.
type Service struct {
	lifecycle       core.ComponentLifecycle
	checker         *Checker
	clock           Clock
	expiredCallback BatchExpiredCallback
	queue           *expirationQueue
	batchSize       int
	batchTimeout    time.Duration
	wheelManager    *TimingWheelManager // nil if timing wheel is disabled
}

func NewService(clk Clock, config Config, callback BatchExpiredCallback) *Service {
	_ = "STUB: not implemented"
	return nil
}

// NowMillis returns the current time in milliseconds via the internal clock.
// This facade method avoids exposing the Clock hierarchy to callers.
func (s *Service) NowMillis() int64 { _ = "STUB: not implemented"; return 0 }

// NowSeconds returns the current time in seconds via the internal clock.
// This facade method avoids exposing the Clock hierarchy to callers.
func (s *Service) NowSeconds() int64 { _ = "STUB: not implemented"; return 0 }

// IsExpired checks whether a record is expired via the internal checker.
// This facade method avoids exposing the Checker hierarchy to callers.
func (s *Service) IsExpired(ttl uint32, timestamp uint64) bool {
	_ = "STUB: not implemented"
	return false
}

// GetChecker returns the internal checker.
// Internal use only for internal data structures.
// Deprecated: root package code should use Service.IsExpired instead.
func (s *Service) GetChecker() *Checker {
	_ = "STUB: not implemented"

	// SetClock updates the service clock and propagates it to the checker.
	// Intended for tests that need deterministic time control.
	return nil
}

func (s *Service) SetClock(clk Clock) { _ = "STUB: not implemented"; return }

func (s *Service) onExpired(bucketId uint64, key []byte, ds uint16, timestamp uint64) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) processExpirationEvents(ctx context.Context) { _ = "STUB: not implemented"; return }

func (s *Service) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *Service) Stop(timeout time.Duration) error { _ = "STUB: not implemented"; return nil }

func (s *Service) Name() string { _ = "STUB: not implemented"; return "" }

// RegisterKeyForActiveExpiration registers a key in the timing wheel for active expiration.
// Should be called when a key with TTL is written to the database.
// If the timing wheel is disabled, this is a no-op.
func (s *Service) RegisterKeyForActiveExpiration(
	bucketId uint64,
	key []byte,
	ds uint16,
	ttl uint32,
	timestamp uint64,
) {
	_ = "STUB: not implemented"
	return
}

// DeregisterKeyFromActiveExpiration removes a key from the timing wheel.
// Should be called when a key is deleted before expiration or its TTL is updated.
// If the timing wheel is disabled, this is a no-op.
func (s *Service) DeregisterKeyFromActiveExpiration(bucketId uint64, key []byte) {
	_ = "STUB: not implemented"
	return
}
