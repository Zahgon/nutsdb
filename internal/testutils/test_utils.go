// Copyright 2023 The PromiseDB Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testutils

import (
	"testing"

	"github.com/nutsdb/nutsdb/internal/core"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GetTestBytes(i int) []byte {
	_ = "STUB: not implemented"
	// Optimized version without fmt.Sprintf to reduce benchmark overhead
	// Format: "nutsdb-000000000" (7 prefix + 9 digits = 16 bytes)
	return nil
}

// Convert i to 9-digit string with leading zeros

func GetRandomBytes(length int) []byte { _ = "STUB: not implemented"; return nil }

func AssertErr(t *testing.T, err error, expectErr error) { _ = "STUB: not implemented"; return }

func GenerateRecords(count int) []*core.Record { _ = "STUB: not implemented"; return nil }
