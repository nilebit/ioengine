// Copyright 2019 shimingyah.
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
// ee the License for the specific language governing permissions and
// limitations under the License.

package ioengine

import (
	"testing"
)

func TestNewFileLock(t *testing.T) {
	// test abs path
	fl00, err := NewFileLock("/tmp/flock/abspath", true)
	if err != nil {
		t.Fatalf("new file lock: %v", err)
	}
	defer fl00.Release()

	// test relative path
	fl01, err := NewFileLock("relativepath", true)
	if err != nil {
		t.Fatalf("new file lock: %v", err)
	}
	defer fl01.Release()
}

func TestRelease(t *testing.T) {
	fl, err := NewFileLock("/tmp/flock/release", true)
	if err != nil {
		t.Fatalf("new file lock: %v", err)
	}
	if err := fl.Release(); err != nil {
		t.Fatalf("release file lock: %v", err)
	}
}

func TestFLock(t *testing.T) {
	fl, err := NewFileLock("/tmp/flock/abspath", true)
	if err != nil {
		t.Fatalf("new file lock: %v", err)
	}
	defer fl.Release()

	if err := fl.FLock(); err != nil {
		t.Fatalf("fllock: %v", err)
	}
}

func TestFUnlock(t *testing.T) {
	fl, err := NewFileLock("/tmp/flock/abspath", true)
	if err != nil {
		t.Fatalf("new file lock: %v", err)
	}
	defer fl.Release()

	if err := fl.FLock(); err != nil {
		t.Fatalf("fllock: %v", err)
	}

	if err := fl.FUnlock(); err != nil {
		t.Fatalf("fllock: %v", err)
	}
}
