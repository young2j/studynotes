/*
 * File: main2.go
 * Created Date: 2022-05-07 04:10:34
 * Author: ysj
 * Description:
 */

package main

import (
	"fmt"
	"rs/types"
	"syscall"
	"unsafe"
)

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
	read()
}
