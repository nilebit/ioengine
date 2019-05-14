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

// +build linux
// +build i386 arm mips

package ioengine

import (
	"unsafe"
)

type iocb struct {
	data   unsafe.Pointer
	pad1   uint32
	key    uint32
	pad2   uint32
	opcode int16
	prio   int16
	fd     uint32
	buf    unsafe.Pointer
	pad3   uint32
	nbytes uint64
	offset int64
	pad4   int64
	flags  uint32
	resfd  uint32
}

type event struct {
	data unsafe.Pointer
	pad1 uint32
	obj  *iocb
	pad2 uint32
	res  int64
	pad3 uint32
	res2 int64
	pad4 uint32
}
