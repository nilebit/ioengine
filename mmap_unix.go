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

// +build linux darwin

package ioengine

import (
	"os"

	"golang.org/x/sys/unix"
)

// Mmap use the mmap system call to memory mapped file or device.
func Mmap(fd *os.File, offset int64, length int, writable bool) ([]byte, error) {
	prot := unix.PROT_READ
	if writable {
		prot |= unix.PROT_WRITE
	}
	return unix.Mmap(int(fd.Fd()), offset, length, prot, unix.MAP_SHARED)
}

// Madvise advises the kernel about how to handle the mapped slice.
func Madvise(b []byte) error {
	return unix.Madvise(b, unix.MADV_NORMAL)
}

// Lock locks the maped slice, preventing it from being swapped out.
func Lock(b []byte) error {
	return unix.Mlock(b)
}

// Unlock unlocks the mapped slice, allowing it to swap out again.
func Unlock(b []byte) error {
	return unix.Munlock(b)
}

// Sync flushes mmap slice's all changes back to the device.
func Sync(b []byte) error {
	return unix.Msync(b, unix.MS_SYNC)
}

// Munmap unmaps mapped slice, this will also flush any remaining changes.
func Munmap(b []byte) error {
	return unix.Munmap(b)
}

// WriteAtv like linux pwritev, write to the specifies offset and dose not change the file offset.
func (mmap *MemoryMap) WriteAtv(bs [][]byte, off int64) (int, error) {
	return genericWriteAtv(mmap, bs, off)
}

// Append write data to the end of file.
func (mmap *MemoryMap) Append(bs [][]byte) (int, error) {
	return genericAppend(mmap, bs)
}
