package main

import "fmt"

func main() {
	t()
	fmt.Println("0x7f")
}

const (
	GOARCH      string = "arm64"
	GOOS        string = "darwin"
	PtrSize            = 4 << (^uintptr(0) >> 63)
	uintptrMask        = 1<<(8*PtrSize) - 1
)

func t() {
	for i := 0x7f; i >= 0; i-- {
		p := uintptr(i)<<40 | uintptrMask&(0x0013<<28)
		fmt.Println(p)
		//var p uintptr
		//switch {
		//case GOARCH == "arm64" && GOOS == "darwin":
		//	p = uintptr(i)<<40 | uintptrMask&(0x0013<<28)
		//case GOARCH == "arm64":
		//	p = uintptr(i)<<40 | uintptrMask&(0x0040<<32)
		//case GOOS == "aix":
		//	if i == 0 {
		//		// We don't use addresses directly after 0x0A00000000000000
		//		// to avoid collisions with others mmaps done by non-go programs.
		//		continue
		//	}
		//	p = uintptr(i)<<40 | uintptrMask&(0xa0<<52)
		//default:
		//	p = uintptr(i)<<40 | uintptrMask&(0x00c0<<32)
		//}
		//hint := (*arenaHint)(mheap_.arenaHintAlloc.alloc())
		//hint.addr = p
		//hint.next, mheap_.arenaHints = mheap_.arenaHints, hint
	}
}
