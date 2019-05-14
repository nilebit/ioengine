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
	"errors"
	"os"
	"sync"
)

// DirectIO dio mode
type DirectIO struct {
	path string
	opt  Options
	once sync.Once
	*os.File
	*FileLock
}

func newDirectIO(name string, opt Options) (*DirectIO, error) {
	fd, err := OpenFileWithDIO(name, opt.Flag, opt.Perm)
	if err != nil {
		return nil, err
	}

	dio := &DirectIO{path: name, opt: opt, File: fd}

	switch opt.FileLock {
	case None:
	case ReadWrite:
		dio.FileLock, err = NewFileLock(name, true)
		if err != nil {
			fd.Close()
			return nil, err
		}
	case ReadOnly:
		dio.FileLock, err = NewFileLock(name, false)
		if err != nil {
			fd.Close()
			return nil, err
		}
	}

	return dio, nil
}

// FLock a file lock is a recommended lock.
// if file lock not init, we will init once.
func (dio *DirectIO) FLock() (err error) {
	if dio.FileLock == nil {
		dio.once.Do(func() {
			if dio.FileLock == nil {
				dio.FileLock, err = NewFileLock(dio.path, true)
			}
		})
	}
	if err != nil {
		return err
	}
	if dio.FileLock == nil {
		return errors.New("Uninitialized file lock")
	}
	return dio.FileLock.FLock()
}

// FUnlock file unlock
func (dio *DirectIO) FUnlock() error {
	if dio.FileLock == nil {
		return nil
	}
	return dio.FileLock.FUnlock()
}

// Close impl standard File Close method
func (dio *DirectIO) Close() error {
	if dio.FileLock == nil {
		return dio.File.Close()
	}

	dio.FileLock.FUnlock()
	// do we need to release file lock?
	// what will happen if release file lock while other process is using file lock?
	// dio.FileLock.Release()
	return dio.File.Close()
}

// Option return File options
func (dio *DirectIO) Option() Options {
	return dio.opt
}
