# 淺談main 入口點探討

非必要前置知識，有時間撰寫會介紹
1. 組合語言(中國用詞為汇编) [參考](https://github.com/Junxwan/blog/tree/master/assembly)
2. golang plan9 
3. gdb用法
4. golang delve
5. golang GMP or Tcmalloc
6. stack與heap概念
7. linux page size

觀看方式請打開源碼根據以下步驟實際動手做，只用眼睛看你有很高的機率看不懂

如使用mac or window，所分析的method可能會跟以下步驟不一樣，
因為golang會根據作業系統來決定執行哪一支.go檔，建議是在ubuntu or alpine or centOS環境下，
不要在mac or window 下做操作

如linux => os_linux.go，window => os_window.go

# 目錄

- [環境](#環境)
- [entry point](#entry-point)
  - [gdb](#gdb)
  - [delve](#delve)
- [rt0_linux_amd64.s](#rt0_linux_amd64.s)
- [asm_amd64.s](#asm_amd64.s)
  - [runtime·args](#runtime·args)
    - [sysargs](#sysargs)
  - [runtime·osinit](#runtime·osinit)
    - [getproccount](#getproccount)
    - [getHugePageSize](#getHugePageSize)
  - [runtime·schedinit](#runtime·schedinit)
    - [tracebackinit](#tracebackinit)
    - [moduledataverify](#moduledataverify)
    - [stackinit](#stackinit)

# 環境

```bash
$ go version
go version go1.13.3 linux/amd64

$ cat /etc/os-release 
VERSION_ID=3.10.3
PRETTY_NAME="Alpine Linux v3.10"

$ gdb --version
GNU gdb (GDB) 8.3

$ dlv version
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
$ go build -gcflags "-N -l" -o main main.go
$ ls 
main     main.go
```

以gdb做debug，執行gdb找出entry point，先看看目前停留main，但主要是找到真正的入口點，執行`info files`如下為`0x454dd0`這就是真正的入口點

```bash
$ gdb ./main
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
        Entry point: 0x454dd0 # 入口點
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
(gdb) b *0x454dd0
Breakpoint 1 at 0x454dd0: file /usr/local/go/src/runtime/rt0_linux_amd64.s, line 8.
```

### delve

不需要編譯main.go，直接做debug，可以看到delve會直接停在程式的入口點是`rt0_linux_amd64.s`，跟gdb不一樣

```bash
$ dlv debug main.go
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
# 省略有機會再研究
...
執行main()前的初始化
CALL	runtime·args(SB)
CALL	runtime·osinit(SB)
CALL	runtime·schedinit(SB)

// create a new goroutine to start program
# 建立一個main goroutine
MOVQ	$runtime·mainPC(SB), AX		// entry
PUSHQ	AX
PUSHQ	$0			// arg size
CALL	runtime·newproc(SB)
POPQ	AX
POPQ	AX

# 執行golang GMP中的 M
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

#### sysargs

分析`runtime.sysargs`，初略觀察是分析`command line `參數，之後待研究(TODO)

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

#### getproccount

`runtime.getproccount`，是一個計算CPU Core數量，之後待研究(TODO)

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
(dlv) c # 如果先前已有多個斷點則需要一值透過c跳到下一個斷點
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
..... # 由於是單步執行所以省略中間
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

#### getHugePageSize

`runtime.getHugePageSize`，用於取得Huge Page Size，之後待研究(TODO)

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

為甚麼是0呢？因為系統禁用Transparent Huge pages

```bash
$ grep -i HugePages_Total /proc/meminfo  
HugePages_Total:       0  # 0代表禁用
```

### runtime·schedinit

查看`CALL runtime·schedinit(SB)`

```bash
(dlv) c
> runtime.schedinit() /usr/local/go/src/runtime/proc.go:529 (hits total:1) (PC: 0x4309d3)
```

打開`proc.go`找到`schedinit()`

```go
func schedinit() {
  // raceinit must be the first call to race detector.
  // In particular, it must be done before mallocinit below calls racemapshadow.
  // 拿出G
  _g_ := getg()
  if raceenabled {
    _g_.racectx, raceprocctx0 = raceinit()
  }

  // 設定Ｍ最大數量
  sched.maxmcount = 10000

  // 好像沒作用，主要是跟print trace stack log有關 
  // 至從該版本後好像就沒用了 https://github.com/golang/go/issues/19348
  // 但不知道為什麼還留著，設計跟有關 https://github.com/golang/proposal/blob/master/design/19348-midstack-inlining.md
  tracebackinit()

  // 待研究
  moduledataverify()

  // 初始化stack
  stackinit()
  mallocinit()
  mcommoninit(_g_.m)
  cpuinit()       // must run before alginit
  alginit()       // maps must not be used before this call
  modulesinit()   // provides activeModules
  typelinksinit() // uses maps, activeModules
  itabsinit()     // uses activeModules

  msigsave(_g_.m)
  initSigmask = _g_.m.sigmask

  goargs()
  goenvs()
  parsedebugvars()
  gcinit()

  sched.lastpoll = uint64(nanotime())
  procs := ncpu
  if n, ok := atoi32(gogetenv("GOMAXPROCS")); ok && n > 0 {
    procs = n
  }
  if procresize(procs) != nil {
    throw("unknown runnable goroutine during bootstrap")
  }

  // For cgocheck > 1, we turn on the write barrier at all times
  // and check all pointer writes. We can't do this until after
  // procresize because the write barrier needs a P.
  if debug.cgocheck > 1 {
    writeBarrier.cgo = true
    writeBarrier.enabled = true
    for _, p := range allp {
      p.wbBuf.reset()
    }
  }

  if buildVersion == "" {
    // Condition should never trigger. This code just serves
    // to ensure runtime·buildVersion is kept in the resulting binary.
    buildVersion = "unknown"
  }
  if len(modinfo) == 1 {
    // Condition should never trigger. This code just serves
    // to ensure runtime·modinfo is kept in the resulting binary.
    modinfo = ""
  }
}
```

簡略流程

1. 透過`getg()`拿出當前要運行的goroutine(G)
2. 

#### tracebackinit

TODO

#### moduledataverify

TODO

#### stackinit

`runtime.stackinit`主要是初始化stack

```go
func stackinit() {
  // _StackCacheSize => 32768 bit => 32K
  // _PageMask => 8191 bit => 1k - 1 bit
  // TODO
  if _StackCacheSize&_PageMask != 0 {
    throw("cache size must be a multiple of page size")
  }

  // stackpool意思是global stack pool，主要是span List
  // init只是將span設為nil
  for i := range stackpool {
    stackpool[i].init()
  }

  // TODO
  for i := range stackLarge.free {
    stackLarge.free[i].init()
  }
}
```

#### mallocinit

`runtime.mallocinit`，

```go
func mallocinit() {
  // 在TCMalloc內存分配管理有所謂的小(0 ~ 256KB)、中(256KB ~ 1MB)、大對象(1MB以上)
  // golang中的小對象有名叫Tiny，意思是16KB以下的對象
  // class_to_size是一個[]uint16
  // _TinySizeClass值是in8(2)
  // _TinySize值是16代表Tiny對象最大為16KB
  // 這裡主要檢查Tiny對象最大容量有沒有等於16KB
  if class_to_size[_TinySizeClass] != _TinySize {
    throw("bad TinySizeClass")
  }

  // 待研究
  testdefersizes()

  if heapArenaBitmapBytes&(heapArenaBitmapBytes-1) != 0 {
    // heapBits expects modular arithmetic on bitmap
    // addresses to work.
    throw("heapArenaBitmapBytes not a power of 2")
  }

  // Copy class sizes out for statistics table.
  // 將class size複製一份至memstats，主要是做統計用的
  // memstats是runtime.mstats，如果有做過golang metrics因該會覺得跟
  // runtime.MemStats很像，因為MemStats資料就是來自mstats
  for i := range class_to_size {
    memstats.by_size[i].size = uint32(class_to_size[i])
  }

  // Check physPageSize.
  // physPageSize指的是系統物理page size
  // 等於4096(4KB)，每個平台不見得都是4KB
  if physPageSize == 0 {
    // The OS init code failed to fetch the physical page size.
    throw("failed to get system page size")
  }

  // minPhysPageSize指的是系統物理page size下限
  // physPageSize不能小於minPhysPageSize
  if physPageSize < minPhysPageSize {
    print("system page size (", physPageSize, ") is smaller than minimum page size (", minPhysPageSize, ")\n")
    throw("bad system page size")
  }

  // physPageSize必須是2的次方
  if physPageSize&(physPageSize-1) != 0 {
    print("system page size (", physPageSize, ") must be a power of 2\n")
    throw("bad system page size")
  }

  // physHugePageSize必須是2的次方，0也算 2^0
  if physHugePageSize&(physHugePageSize-1) != 0 {
    print("system huge page size (", physHugePageSize, ") must be a power of 2\n")
    throw("bad system huge page size")
  }

  // 如果physHugePageSize不為零則要算出physHugePageSize是2的幾次方
  if physHugePageSize != 0 {
    // Since physHugePageSize is a power of 2, it suffices to increase
    // physHugePageShift until 1<<physHugePageShift == physHugePageSize.
    for 1<<physHugePageShift != physHugePageSize {
      physHugePageShift++
    }
  }

  // Initialize the heap.
  // 初始化heap
  mheap_.init()
  
  // 取得G (golang GMP)
  _g_ := getg()
  _g_.m.mcache = allocmcache()

  // Create initial arena growth hints.
  if sys.PtrSize == 8 {
    // On a 64-bit machine, we pick the following hints
    // because:
    //
    // 1. Starting from the middle of the address space
    // makes it easier to grow out a contiguous range
    // without running in to some other mapping.
    //
    // 2. This makes Go heap addresses more easily
    // recognizable when debugging.
    //
    // 3. Stack scanning in gccgo is still conservative,
    // so it's important that addresses be distinguishable
    // from other data.
    //
    // Starting at 0x00c0 means that the valid memory addresses
    // will begin 0x00c0, 0x00c1, ...
    // In little-endian, that's c0 00, c1 00, ... None of those are valid
    // UTF-8 sequences, and they are otherwise as far away from
    // ff (likely a common byte) as possible. If that fails, we try other 0xXXc0
    // addresses. An earlier attempt to use 0x11f8 caused out of memory errors
    // on OS X during thread allocations.  0x00c0 causes conflicts with
    // AddressSanitizer which reserves all memory up to 0x0100.
    // These choices reduce the odds of a conservative garbage collector
    // not collecting memory because some non-pointer block of memory
    // had a bit pattern that matched a memory address.
    //
    // However, on arm64, we ignore all this advice above and slam the
    // allocation at 0x40 << 32 because when using 4k pages with 3-level
    // translation buffers, the user address space is limited to 39 bits
    // On darwin/arm64, the address space is even smaller.
    // On AIX, mmaps starts at 0x0A00000000000000 for 64-bit.
    // processes.
    for i := 0x7f; i >= 0; i-- {
      var p uintptr
      switch {
        case GOARCH == "arm64" && GOOS == "darwin":
        p = uintptr(i)<<40 | uintptrMask&(0x0013<<28)
        case GOARCH == "arm64":
        p = uintptr(i)<<40 | uintptrMask&(0x0040<<32)
        case GOOS == "aix":
        if i == 0 {
          // We don't use addresses directly after 0x0A00000000000000
          // to avoid collisions with others mmaps done by non-go programs.
          continue
        }
        p = uintptr(i)<<40 | uintptrMask&(0xa0<<52)
        case raceenabled:
        // The TSAN runtime requires the heap
        // to be in the range [0x00c000000000,
        // 0x00e000000000).
        p = uintptr(i)<<32 | uintptrMask&(0x00c0<<32)
        if p >= uintptrMask&0x00e000000000 {
          continue
        }
        default:
        p = uintptr(i)<<40 | uintptrMask&(0x00c0<<32)
      }
      hint := (*arenaHint)(mheap_.arenaHintAlloc.alloc())
      hint.addr = p
      hint.next, mheap_.arenaHints = mheap_.arenaHints, hint
    }
  } else {
    // On a 32-bit machine, we're much more concerned
    // about keeping the usable heap contiguous.
    // Hence:
    //
    // 1. We reserve space for all heapArenas up front so
    // they don't get interleaved with the heap. They're
    // ~258MB, so this isn't too bad. (We could reserve a
    // smaller amount of space up front if this is a
    // problem.)
    //
    // 2. We hint the heap to start right above the end of
    // the binary so we have the best chance of keeping it
    // contiguous.
    //
    // 3. We try to stake out a reasonably large initial
    // heap reservation.

    const arenaMetaSize = (1 << arenaBits) * unsafe.Sizeof(heapArena{})
    meta := uintptr(sysReserve(nil, arenaMetaSize))
    if meta != 0 {
      mheap_.heapArenaAlloc.init(meta, arenaMetaSize)
    }

    // We want to start the arena low, but if we're linked
    // against C code, it's possible global constructors
    // have called malloc and adjusted the process' brk.
    // Query the brk so we can avoid trying to map the
    // region over it (which will cause the kernel to put
    // the region somewhere else, likely at a high
    // address).
    procBrk := sbrk0()

    // If we ask for the end of the data segment but the
    // operating system requires a little more space
    // before we can start allocating, it will give out a
    // slightly higher pointer. Except QEMU, which is
    // buggy, as usual: it won't adjust the pointer
    // upward. So adjust it upward a little bit ourselves:
    // 1/4 MB to get away from the running binary image.
    p := firstmoduledata.end
    if p < procBrk {
      p = procBrk
    }
    if mheap_.heapArenaAlloc.next <= p && p < mheap_.heapArenaAlloc.end {
      p = mheap_.heapArenaAlloc.end
    }
    p = round(p+(256<<10), heapArenaBytes)
    // Because we're worried about fragmentation on
    // 32-bit, we try to make a large initial reservation.
    arenaSizes := []uintptr{
      512 << 20,
      256 << 20,
      128 << 20,
    }
    for _, arenaSize := range arenaSizes {
      a, size := sysReserveAligned(unsafe.Pointer(p), arenaSize, heapArenaBytes)
      if a != nil {
        mheap_.arena.init(uintptr(a), size)
        p = uintptr(a) + size // For hint below
        break
      }
    }
    hint := (*arenaHint)(mheap_.arenaHintAlloc.alloc())
    hint.addr = p
    hint.next, mheap_.arenaHints = mheap_.arenaHints, hint
  }
}
```

