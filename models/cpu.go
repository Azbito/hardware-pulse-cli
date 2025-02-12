package models

import (
	"syscall"
	"time"
	"unsafe"
)

type CPU struct {
	idle   uint64
	kernel uint64
	user   uint64
}

var (
	modKernel32        = syscall.NewLazyDLL("kernel32.dll")
	procGetSystemTimes = modKernel32.NewProc("GetSystemTimes")
)

func (cpu *CPU) GetSystemTimes() error {

	var idleTime, kernelTime, userTime syscall.Filetime

	ret, _, err := procGetSystemTimes.Call(
		uintptr(unsafe.Pointer(&idleTime)),
		uintptr(unsafe.Pointer(&kernelTime)),
		uintptr(unsafe.Pointer(&userTime)),
	)

	if ret == 0 {
		return err
	}

	cpu.idle = (uint64(idleTime.HighDateTime) << 32) | uint64(idleTime.LowDateTime)
	cpu.kernel = (uint64(kernelTime.HighDateTime) << 32) | uint64(kernelTime.LowDateTime)
	cpu.user = (uint64(userTime.HighDateTime) << 32) | uint64(userTime.LowDateTime)

	return nil
}

func (cpu *CPU) CalculateCPUUsage() (float64, error) {
	err := cpu.GetSystemTimes()

	if err != nil {
		return 0.0, err
	}

	idle1, kernel1, user1 := cpu.idle, cpu.kernel, cpu.user

	time.Sleep(1 * time.Second)

	err = cpu.GetSystemTimes()
	if err != nil {
		return 0.0, err
	}

	idle2, kernel2, user2 := cpu.idle, cpu.kernel, cpu.user

	idleDelta := idle2 - idle1
	kernelDelta := kernel2 - kernel1
	userDelta := user2 - user1

	total := kernelDelta + userDelta

	if total == 0 {
		return 0.0, nil
	}

	return (float64(total-idleDelta) / float64(total)) * 100, nil
}
