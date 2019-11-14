

# 淺談main 入口點探討

```bash
非必要前置知識，有時間撰寫會介紹
1. 組合語言(中國用詞為汇编)
2. golang plan9 
3. gdb用法
4. golang delve

觀看方式請打開源碼根據以下步驟實際動手做，只用眼睛看你有很大機率看不懂

如使用mac or window，所分析的method可能會跟以下步驟不一樣，因為golang會根據作業系統來決定執行哪一支.go檔，建議是在ubuntu or alpine or centOS環境下，不要在mac or window 下做操作

如linux => os_linux.go，window => os_window.go
```

#目錄

- [環境](#環境)
- [entry point](#entry-point)
  - [gdb](#gdb)
  - [delve](#delve)
- [rt0_linux_amd64.s](#rt0_linux_amd64.s)
- [asm_amd64.s](#asm_amd64.s)
  - [runtime·args](#runtime·args)
  - [runtime·osinit](#runtime·osinit)
  - [runtime·schedinit](#runtime·schedinit)

# 環境

```bash
> go version
go version go1.13.3 linux/amd64

> cat /etc/os-release 
VERSION_ID=3.10.3
PRETTY_NAME="Alpine Linux v3.10"

> gdb --version
GNU gdb (GDB) 8.3

> dlv version
Delve Debugger
Version: 1.3.2
```

## entry point

寫一段範例

```go
package main

import "fmt"

func main() {
	fmt.Println("test")
}
```

#### gdb

以gdb做示範，先編譯 go build -gcflags "-N -l" -o main main.go

```bash
> go build -gcflags "-N -l" -o main main.go
> ls 
main     main.go
```

以gdb做debug，執行gdb找出entry point，先看看目前停留main，但主要是找到真正的入口點，執行`info files`如下為`0x454dd0`這就是真正的入口點

```bash
> gdb ./main
(gdb) l
1       package main
2       
3       import "fmt"
4       
5       func main() {
6               fmt.Println("test")
7       }
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

針對entry point做斷點，找到程式入口點是`rt0_linux_amd64.s`

```bash
> (gdb) b *0x454dd0
Breakpoint 1 at 0x454dd0: file /usr/local/go/src/runtime/rt0_linux_amd64.s, line 8.
```

### delve

不需要編譯main.go，直接做debug，可以看到delve會直接停在程式的入口點是`rt0_linux_amd64.s`，跟gdb不一樣

```bash
> dlv debug main.go
Type 'help' for list of commands.
(dlv) l
> _rt0_amd64_linux() /usr/local/go/src/runtime/rt0_linux_amd64.s:8 (PC: 0x45c420)
Warning: debugging optimized function
     3: // license that can be found in the LICENSE file.
     4: 
     5: #include "textflag.h"
     6: 
     7: TEXT _rt0_amd64_linux(SB),NOSPLIT,$-8
=>   8:         JMP     _rt0_amd64(SB)
     9: 
    10: TEXT _rt0_amd64_linux_lib(SB),NOSPLIT,$0
    11:         JMP     _rt0_amd64_lib(SB)
(dlv)

```

使用gdb or delve都可以，我個人在開發golang主要使用delve，因為delve就是專門針對golang做debug用，支援會更好，接下來debug都會使用delve 

## rt0_linux_amd64.s

打開`rt0_linux_amd64.s`查看組合語言(中國用詞為汇编)，TEST代表一個function，可以看到名稱為`main`

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

利用delve查看`runtime·rt0_go`所在位置是`asm_amd64.s`

```bash
(dlv) b runtime.rt0_go
Breakpoint 1 set at 0x458a50 for runtime.rt0_go() /usr/local/go/src/runtime/asm_amd64.s:89
```

## asm_amd64.s

打開`asm_amd64.s`，先看以下重點

```assembly
TEXT runtime·rt0_go(SB),NOSPLIT,$0
...
省略有機會在研究
...
執行main()前的初始化
CALL	runtime·args(SB)
CALL	runtime·osinit(SB)
CALL	runtime·schedinit(SB)

// create a new goroutine to start program
// 建立一個main goroutine
MOVQ	$runtime·mainPC(SB), AX		// entry
PUSHQ	AX
PUSHQ	$0			// arg size
CALL	runtime·newproc(SB)
POPQ	AX
POPQ	AX

// 執行golang GMP中的 M
// start this M
CALL	runtime·mstart(SB)

CALL	runtime·abort(SB)	// mstart should never return
RET
```

### runtime·args

查看`CALL runtime·args(SB)`

```bash
(dlv) b runtime.args
Breakpoint 2 set at 0x43ca5f for runtime.args() /usr/local/go/src/runtime/runtime1.go:60
```

打開`runtime1.go`，找到`runtime.args`

```go
func args(c int32, v **byte) {
	argc = c
	argv = v
	sysargs(c, v)
}
```

分析`sysargs`，初略觀察是分析`command line `參數，之後待研究(TODO)

```go
func sysargs(argc int32, argv **byte) {
	// skip over argv, envv and the first string will be the path
	n := argc + 1
	for argv_index(argv, n) != nil {
		n++
	}
	executablePath = gostringnocopy(argv_index(argv, n+1))

	// strip "executable_path=" prefix if available, it's added after OS X 10.11.
	const prefix = "executable_path="
	if len(executablePath) > len(prefix) && executablePath[:len(prefix)] == prefix {
		executablePath = executablePath[len(prefix):]
	}
}
```

### runtime·osinit

查看`CALL runtime·osinit(SB)`

```bash
(dlv) b runtime.osinit
Breakpoint 3 set at 0x42bc8f for runtime.osinit() /usr/local/go/src/runtime/os_linux.go:289
```

打開`os_linux.go`找到`osinit`

```go
func osinit() {
	ncpu = getproccount()
	physHugePageSize = getHugePageSize()
}
```

`getproccount`，是一個計算CPU Core數量，之後待研究(TODO)

```go
func getproccount() int32 {
	// This buffer is huge (8 kB) but we are on the system stack
	// and there should be plenty of space (64 kB).
	// Also this is a leaf, so we're not holding up the memory for long.
	// See golang.org/issue/11823.
	// The suggested behavior here is to keep trying with ever-larger
	// buffers, but we don't have a dynamic memory allocator at the
	// moment, so that's a bit tricky and seems like overkill.
	const maxCPUs = 64 * 1024
	var buf [maxCPUs / 8]byte
	
	// sched_getaffinity使用組合語言撰寫
	r := sched_getaffinity(0, unsafe.Sizeof(buf), &buf[0])
	if r < 0 {
		return 1
	}
	n := int32(0)
	for _, v := range buf[:r] {
		for v != 0 {
			n += int32(v & 1)
			v >>= 1
		}
	}
	if n == 0 {
		n = 1
	}
	return n
}
```

查看`getproccount()`執行後的值是多少，答案是2，代表我現在的cpu在系統內顯示的是2 core

```bash
(dlv) c #如果先前已有多個斷點則需要一值透過c跳到下一個斷點
> runtime.osinit() /usr/local/go/src/runtime/os_linux.go:289 (hits total:1) (PC: 0x42bc8f)
Warning: debugging optimized function
   284:                 return 0
   285:         }
   286:         return uintptr(v)
   287: }
   288: 
=> 289: func osinit() { # 由於是斷點在osinit()所以停留在這
   290:         ncpu = getproccount()
   291:         physHugePageSize = getHugePageSize()
   292: }
   293: 
   294: var urandom_dev = []byte("/dev/urandom\x00")
(dlv) n
..... # 由於是單部執行所以省略中間
(dlv) n
> runtime.osinit() /usr/local/go/src/runtime/os_linux.go:291 (PC: 0x42bcab)
Warning: debugging optimized function
   286:         return uintptr(v)
   287: }
   288: 
   289: func osinit() {
   290:         ncpu = getproccount()
=> 291:         physHugePageSize = getHugePageSize()
   292: }
   293: 
   294: var urandom_dev = []byte("/dev/urandom\x00")
   295: 
   296: func getRandomData(r []byte) {
(dlv) p ncpu
2 # getproccount()結果
(dlv) 
```

> 我實驗的電腦是4 core，由於是在docker環境下沒做過任何設定所以限制成原本的一半也就是2 core

`getHugePageSize`，用於取得Huge Page Size，之後待研究(TODO)

```go
func getHugePageSize() uintptr {
	var numbuf [20]byte
	fd := open(&sysTHPSizePath[0], 0 /* O_RDONLY */, 0)
	if fd < 0 {
		return 0
	}
	n := read(fd, noescape(unsafe.Pointer(&numbuf[0])), int32(len(numbuf)))
	closefd(fd)
	if n <= 0 {
		return 0
	}
	l := n - 1 // remove trailing newline
	v, ok := atoi(slicebytetostringtmp(numbuf[:l]))
	if !ok || v < 0 {
		v = 0
	}
	if v&(v-1) != 0 {
		// v is not a power of 2
		return 0
	}
	return uintptr(v)
}
```

查看`getHugePageSize`執行後的值是0

```bash
(dlv) n
> runtime.osinit() /usr/local/go/src/runtime/os_linux.go:292 (PC: 0x42bcbb)
Warning: debugging optimized function
   287: }
   288: 
   289: func osinit() {
   290:         ncpu = getproccount()
   291:         physHugePageSize = getHugePageSize()
=> 292: }
   293: 
   294: var urandom_dev = []byte("/dev/urandom\x00")
   295: 
   296: func getRandomData(r []byte) {
   297:         if startupRandomData != nil {
(dlv) p physHugePageSize
0
```

為甚麼是0呢？，因為系統禁用Transparent Huge pages

```bash
> grep -i HugePages_Total /proc/meminfo  
HugePages_Total:       0  # 0代表禁用
```

### runtime·schedinit

查看`CALL runtime·schedinit(SB)`

```bash
(dlv) c
> runtime.schedinit() /usr/local/go/src/runtime/proc.go:529 (hits total:1) (PC: 0x4309d3)
```

打開`proc.go:529`找到`schedinit()`

```go

```

