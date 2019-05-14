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

// +build darwin

package ioengine

import (
	"os"
	"syscall"
)

const (
	// AlignSize OSX doesn't need any alignment
	AlignSize = 0
	// BlockSize direct IO minimum number of bytes to write
	BlockSize = 4096
)

// OpenFileWithDIO open files with no cache on darwin.
func OpenFileWithDIO(name string, flag int, perm os.FileMode) (*os.File, error) {
	fd, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	// set no cache
	_, _, er := syscall.Syscall(syscall.SYS_FCNTL, uintptr(fd.Fd()), syscall.F_NOCACHE, 1)
	if er != 0 {
		fd.Close()
		return nil, os.NewSyscallError("Fcntl:NoCache", er)
	}

	return fd, nil
}

// WriteAtv simulate writeatv by calling writev serially and dose not change the file offset.
func (dio *DirectIO) WriteAtv(bs [][]byte, off int64) (int, error) {
	return genericWriteAtv(dio, bs, off)
}

// Append write data to the end of file.
func (dio *DirectIO) Append(bs [][]byte) (int, error) {
	return genericAppend(dio, bs)
}
