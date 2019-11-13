# 環境

```bash
> go version
go version go1.13.3 linux/amd64

> cat /etc/os-release 
VERSION_ID=3.10.3
PRETTY_NAME="Alpine Linux v3.10"

> gdb --version
GNU gdb (GDB) 8.3
```



## main 入口點探討

先寫一段範例

```go
package main

import "fmt"

func main() {
	fmt.Println("test")
}
```

編譯 go build -gcflags "-N -l" -o main main.go

```ba
> go build -gcflags "-N -l" -o main main.go
> ls 
main     main.go
```

執行gdb，找出entry point，如下為`0x454dd0`

```ba
> gdb ./main
(gdb) info files
Symbols from "/var/www/runtime/main".
Local exec file:
        `/var/www/runtime/main', file type elf64-x86-64.
        Entry point: 0x454dd0
        0x0000000000401000 - 0x000000000048cfc3 is .text
        0x000000000048d000 - 0x00000000004dc5d0 is .rodata
        0x00000000004dc7a0 - 0x00000000004dd40c is .typelink
        0x00000000004dd410 - 0x00000000004dd460 is .itablink
        0x00000000004dd460 - 0x00000000004dd460 is .gosymtab
        0x00000000004dd460 - 0x0000000000548978 is .gopclntab
        0x0000000000549000 - 0x0000000000549020 is .go.buildinfo
        0x0000000000549020 - 0x00000000005560f8 is .noptrdata
        0x0000000000556100 - 0x000000000055d150 is .data
        0x000000000055d160 - 0x00000000005789d0 is .bss
        0x00000000005789e0 - 0x000000000057b148 is .noptrbss
        0x0000000000400f9c - 0x0000000000401000 is .note.go.buildid
```

> TODO 之後介紹[gdb]()

entry point做斷點，可找到程式入口點

```ba
> (gdb) b *0x454dd0
Breakpoint 1 at 0x454dd0: file /usr/local/go/src/runtime/rt0_linux_amd64.s, line 8.
```

打開`rt0_linux_amd64.s`查看組合語言(大陸用詞為汇编)，TEST代表一個function，可以看到名稱為`main`

就是golang main入口點

```assembly
TEXT main(SB),NOSPLIT|NOFRAME,$0
	MOVD	$runtime·rt0_go(SB), R2
	BL	(R2)
exit:
	MOVD $0, R0
	MOVD	$94, R8	// sys_exit
	SVC
	B	exit
```

> TODO 之後介紹[組合語言]()跟[plan9]()

利用gdb查看`runtime·rt0_go`在哪裡，看到是`asm_amd64.s`

```bash
(gdb) b runtime.rt0_go
Breakpoint 3 at 0x451410: file /usr/local/go/src/runtime/asm_amd64.s, line 89.
```

打開`asm_amd64.s`

```assembly

```

