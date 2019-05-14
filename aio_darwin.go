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

// +build !linux

package ioengine

import (
	"errors"
	"os"
)

type AsyncIO struct {
	*os.File
}

func newAsyncIO(name string, opt Options) (*AsyncIO, error) {
	return nil, errors.New("Please use AIO on linux")
}

func (aio *AsyncIO) WriteAtv(bs [][]byte, off int64) (int, error) {
	return 0, nil
}

func (aio *AsyncIO) Append(bs [][]byte) (int, error) {
	return 0, nil
}

func (aio *AsyncIO) FLock() error {
	return nil
}

func (aio *AsyncIO) FUnlock() error {
	return nil
}

func (aio *AsyncIO) Option() Options {
	return Options{}
}
