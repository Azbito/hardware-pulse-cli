package models

import (
	"fmt"
	"syscall"
	"unsafe"
)

type Memory struct {
	Length        uint32
	MemoryLoad    uint32
	TotalPhys     uint64
	AvailPhys     uint64
	TotalPageFile uint64
	AvailPageFile uint64
	TotalVirtual  uint64
	AvailVirtual  uint64
	SecurityInfo  uint64
}

var procGlobalMemoryStatusEx = modKernel32.NewProc("GlobalMemoryStatusEx")

func (mem *Memory) GetMemoryStatus() error {
	mem.Length = uint32(unsafe.Sizeof(*mem))

	ret, _, err := syscall.SyscallN(
		procGlobalMemoryStatusEx.Addr(),
		uintptr(unsafe.Pointer(mem)),
	)

	if ret == 0 {
		return fmt.Errorf("erro ao obter status de memória: %v", err)
	}

	return nil
}
