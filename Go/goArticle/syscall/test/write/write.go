package main

import (
	"fmt"
	"rs/types"
	"syscall"
	"unsafe"
)

func write() {
	shmid, _, err := syscall.Syscall(syscall.SYS_SHMGET, types.KEY, types.SIZE, types.IPC_CREATE|0644)
	if err != 0 {
		panic(err)
	}
	fmt.Printf("shmid: %v\n", shmid)

	shmaddr, _, err := syscall.Syscall(syscall.SYS_SHMAT, shmid, 0, 0)
	if err != 0 {
		panic(err)
	}
	fmt.Printf("shmaddr: %v \n", shmaddr)

	data := types.Data{
		F1: "field1",
		F2: 22,
		F3: []string{"色情", "涉政"},
		F4: map[int]int{0: 0, 1: 1},
	}
	*(*types.Data)(unsafe.Pointer(uintptr(shmaddr))) = data

	shmdt, _, err := syscall.Syscall(syscall.SYS_SHMDT, shmaddr, 0, 0)
	if err != 0 {
		panic(err)
	}
	fmt.Printf("shmdt: %v\n", shmdt)
}

func read() {
	shmid, _, err := syscall.Syscall(syscall.SYS_SHMGET, types.KEY, types.SIZE, types.IPC_CREATE|0644)
	if err != 0 {
		panic(err)
	}
	fmt.Printf("shmid: %v\n", shmid)

	shmaddr, _, err := syscall.Syscall(syscall.SYS_SHMAT, shmid, 0, 0)
	if err != 0 {
		panic(err)
	}
	fmt.Printf("shmaddr: %v\n", shmaddr)

	data := *(*types.Data)(unsafe.Pointer(uintptr(shmaddr)))
	fmt.Printf("data: %#v\n", data)

	shmdt, _, err := syscall.Syscall(syscall.SYS_SHMDT, shmaddr, 0, 0)
	if err != 0 {
		panic(err)
	}
	fmt.Printf("shmdt: %v\n", shmdt)
}

func main() {
	write()
	read()
}
