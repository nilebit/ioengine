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
// +build amd64 arm64

package ioengine

import "unsafe"

type iocb struct {
	data   unsafe.Pointer
	key    uint64
	opcode int16
	prio   int16
	fd     uint32
	buf    unsafe.Pointer
	nbytes uint64
	offset int64
	pad1   int64
	flags  uint32
	resfd  uint32
}

type event struct {
	data unsafe.Pointer
	obj  *iocb
	res  int64
	res2 int64
}
