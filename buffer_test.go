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
	"io"
	"testing"
	"unsafe"
)

type blockAlign struct {
	blockSize uint
	alignSize uint
}

var bas = []blockAlign{
	blockAlign{8, 4},
	// blockAlign{7, 4},
	// blockAlign{7, 0},
}

func TestMemAlign(t *testing.T) {
	for _, ba := range bas {
		b, err := MemAlignWithBase(ba.blockSize, ba.alignSize)
		if err != nil {
			t.Fatal(err)
		}
		if uint(len(b)) != ba.blockSize {
			t.Fatal("unmatched block size")
		}
		if uint(uintptr(unsafe.Pointer(&b[0]))&uintptr(ba.alignSize-1)) != 0 {
			t.Fatal("start address is not multiple of align size")
		}
	}
}

func TestBufferReadWrite(t *testing.T) {
	bs := NewBuffers()
	bs.Write([]byte("hello")).Write([]byte("world"))
	tmp := make([]byte, 10, 10)
	n, err := bs.Read(tmp)
	if err != nil && err != io.EOF {
		t.Fatal(err)
	}
	if n != 10 {
		t.Fatal("unmathed length")
	}
	if string(tmp) != "helloworld" {
		t.Fatal("unmatched content")
	}
}
