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

// FileIO Standrad I/O mode
type FileIO struct {
	path string
	opt  Options
	once sync.Once
	*os.File
	*FileLock
}

func newFileIO(name string, opt Options) (*FileIO, error) {
	fd, err := os.OpenFile(name, opt.Flag, opt.Perm)
	if err != nil {
		return nil, err
	}

	fi := &FileIO{path: name, opt: opt, File: fd}

	switch opt.FileLock {
	case None:
	case ReadWrite:
		fi.FileLock, err = NewFileLock(name, true)
		if err != nil {
			fd.Close()
			return nil, err
		}
	case ReadOnly:
		fi.FileLock, err = NewFileLock(name, false)
		if err != nil {
			fd.Close()
			return nil, err
		}
	}

	return fi, nil
}

// FLock a file lock is a recommended lock.
// if file lock not init, we will init once.
func (fi *FileIO) FLock() (err error) {
	if fi.FileLock == nil {
		fi.once.Do(func() {
			if fi.FileLock == nil {
				fi.FileLock, err = NewFileLock(fi.path, true)
			}
		})
	}
	if err != nil {
		return err
	}
	if fi.FileLock == nil {
		return errors.New("Uninitialized file lock")
	}
	return fi.FileLock.FLock()
}

// FUnlock file unlock
func (fi *FileIO) FUnlock() error {
	if fi.FileLock == nil {
		return nil
	}
	return fi.FileLock.FUnlock()
}

// Close impl standard File Close method
func (fi *FileIO) Close() error {
	if fi.FileLock == nil {
		return fi.File.Close()
	}

	fi.FileLock.FUnlock()
	// do we need to release file lock?
	// what will happen if release file lock while other process is using file lock?
	// fi.FileLock.Release()
	return fi.File.Close()
}

// Option return File options
func (fi *FileIO) Option() Options {
	return fi.opt
}
