# plan9

- [閱讀前提](#閱讀前提)

- [寄存器](#寄存器)

- [變數](#變數)
  - [作用範圍](#作用範圍)
  - [標誌](#標誌)
  - [int](#int)
  - [string](#string)
  - [array](#array)
  - [bool](#bool)
  - [float](#float)
  - [slice](#slice)
  - [map](#map)
  - [channel](#channel)
- [函數](#函數)

- [參考](#參考)



## 閱讀前提

1. 對於組合語言(汇编语言)有基礎概念者才能閱讀，可[參考](https://github.com/Junxwan/blog/tree/master/assembly)
2. 本章節內容有任何錯誤請儘管提出來，如果我有空則會回覆
3. 本篇單純入門而已



## 寄存器



## 變數

定義變數規則:`DATA symbol+offset(SB)/width, value`

```assembly
DATA ·Name+0(SB)/8,$0x4D2
```

宣告一個名叫`Name`變數，是一個64位元變數，初始值為`0x4D2`

| 欄位   | 定義                           | 值    |
| ------ | ------------------------------ | ----- |
| DATA   | 宣告做變數初始化               |       |
| symbol | 變數名稱                       | ·Name |
| offset | 從SB內存地址開始偏移多少個byte | 0     |
| width  | 該變數使用多少byte             | 8     |
| value  | 變數值                         | 0x4D2 |

以下是該變數示意內存結構，從內存結構可以得知`Name`變數從SB起始地址開始往後佔用8個byte

| 內存地址       | 值   |            |
| -------------- | ---- | ---------- |
| 0xffffffff0000 | 0xD2 | SB起始地址 |
| 0xffffffff0008 | 0x04 |            |
| 0xffffffff0010 | 0x00 |            |
| 0xffffffff0018 | 0x00 |            |
| 0xffffffff0020 | 0x00 |            |
| 0xffffffff0028 | 0x00 |            |
| 0xffffffff0030 | 0x00 |            |
| 0xffffffff0038 | 0x00 |            |

除了一次性設定初始值外也可以分批設定，針對每個byte做設定

```assembly
DATA ·Name+0(SB)/1,$0xD2
DATA ·Name+1(SB)/1,$0x04
DATA ·Name+2(SB)/1,$0x00
DATA ·Name+3(SB)/1,$0x00
DATA ·Name+4(SB)/1,$0x00
DATA ·Name+5(SB)/1,$0x00
DATA ·Name+6(SB)/1,$0x00
DATA ·Name+7(SB)/1,$0x00

```

其中value可以接受以下類型

| 類型         | 值         |
| ------------ | ---------- |
| 十進位       | $1234      |
| 十六進位     | $0x04D2    |
| 浮點數       | $1234.56   |
| 字元         | $'1'       |
| 字串         | $"1234"    |
| 變數內存地址 | $·Name(SB) |

當設定完變數後如果想讓golang中也可以使用到則須定義成外部檔案也可以存取

定義變數規則: `GLOBL symbol(SB), width`

```assembly
DATA ·Name+0(SB)/8,$0x4D2
GLOBL ·Name(SB),$8
```

| 欄位   | 定義               | 值    |
| ------ | ------------------ | ----- |
| GLOBL  | 外部可以存取       |       |
| symbol | 要開放外部的變數   | ·Name |
| width` | 該變數佔用多少byte | $8    |

以下兩段golang範例分別宣告一個`Name`全域int變數以及print `Name`變數而初始值是0

```go
// main.go
package main

import "test/test"

func main() {
	println(test.Name)
}
```

```go
// test/test.go
package test

var Name int
```

執行`main.go`可以看到 `Name`變數是0沒錯

```bash
/var/www # go run main.go 
0
```

這時候增加一個名叫`test_amd64.s`檔案，後綴詞`_amd64`是代表系統環境是`linux 64位元`，這樣執行時才會找到該檔案，另外注意最後一行要空一格不然會出現`asm: assembly of test/test_amd64.s failed`

```assembly
# test_amd64.s

DATA ·Name+0(SB)/8,$0x4D2
GLOBL ·Name(SB),$8

```

執行`main.go`可以看到 1234也就是`0x04d2`，初始值直接由plan9決定而非golang

```bash
/var/www # go run main.go 
1234
```



symbol加上 `·` = `·symbol`的意思換成`·Name`其實代表當前package內的`Name`變數

```go
println(test.Name) // ·Name = test.Name
```

將`test_amd64.s`改寫成

```assembly
# test/test_amd64.s

DATA Name+0(SB)/8,$0x4D2
GLOBL Name(SB),$8
```

可以看到`Name`初始值為0而不是1234，因為沒有`Name`變數前綴詞沒有加上`·`

```bash
/var/www # go run main.go 
0
```



利用plan9將某個變數值放到某變數上做初始化

```go
// main.go
package main

import (
	"test/test"
)

func main() {
	println(*test.Name)
}
```

```go
// test/test.go
package test

var Name *int
var Value = 4321
```

```assembly
# test/test_amd64.s

DATA ·Name+0(SB)/8,$·Value(SB)
GLOBL ·Name(SB),$8

```

上述先將`Name`變數改成指標，因為使用DATA標籤賦予變數初始值是放入該變數的內存地址，所以`Name`變數值其實是存放`Value`內存地址，執行後`Name`變數值就是4321

```bash
/var/www # go run main.go 
4321
```



### 作用範圍

在之前有提到GLOBL可以定義一個變數讓golang可以引用，但GLOBL其實還包含了plan9其他文件中使用也可以使用此變數，看下方範例，golang範例跟之前並無差異

```go
// main.go
package main

import (
	"test/test"
)

func main() {
	println(test.Name)
}
```

```go
// test/test.go
package test

var Name int
```

這裡`Name`變數值來自於`Value`變數內存地址

```assembly
# test/test_amd64.s
DATA ·Name+0(SB)/8,$Value(SB)
GLOBL ·Name(SB),$8

```

初始化`Value`變數，這邊可以看到`NOPTR`與`#include "textflag.h"`，這稍後章節在討論

```assembly
# test/value_amd64.s

#include "textflag.h"
DATA Value+0(SB)/8,$0x4d2
GLOBL Value(SB),NOPTR,$8

```

執行看看`Name`變數值為`17568096`，從這得知`test_amd64.s`可以看到`value_amd64.s`內的變數而`test.go`可以看到`test_amd64.s`的變數

```bash
/var/www # go run main.go 
17568096 # Value變數內存十六進位地址轉成十進位結果
```



Value`變數不在定義`GLOBL`

```assembly
# test/value_amd64.s

DATA Value+0(SB)/8,$0x4d2

```

執行時`Name`變數告知找不到`Value`變數來參考，從這得出GLOBL不只是影響golang變數是否能參考到plan9的變數，是否plan9之間不同文件的變數也會互相影響到能見度呢？

```bash
/var/www # go run main.go 
# command-line-arguments
test/test.Name: relocation target Value not defined

```



這裡在做個實驗，將`value_amd64.s`變數移至`test/test_amd64.s`

```assembly
# test/value_amd64.s
# 清空
```

```assembly
# test/test_amd64.s
DATA Value+0(SB)/8,$0x4d2

DATA ·Name+0(SB)/8,$Value(SB)
GLOBL ·Name(SB),$8

```

執行後的結果也也一樣，代表GLOBL在plan9內變數能見度不分文件只看代碼

```bash
/var/www # go run main.go 
# command-line-arguments
test/test.Name: relocation target Value not defined
```



上述範例是`Value`變數並非是`·Value`，哪如果改成`·Value`會怎麼樣？

```assembly
# test/test_amd64.s
DATA ·Value+0(SB)/8,$0x4d2

DATA ·Name+0(SB)/8,$·Value(SB)
GLOBL ·Name(SB),$8

```

結果也是一樣，只是這次`Name`變數想要參考的是`test/test.Value`而非是plan9內定義的`Value`變數

```bash
/var/www # go run main.go 
# command-line-arguments
test/test.Name: relocation target test/test.Value not defined

```



如果golang增加一個`Value`變數

```go
// test/test.go
package test

var Name *int // 由於plan9內賦予的初始值是Value的內存地址，所以把Name變成指標
var Value int
```

```go
// main.go
package main

import (
	"test/test"
)

func main() {
	println(*test.Name) // Name變數是指標，加上*取變數值
}
```

執行後的結果不是`1234(0x4d2)`而是0，這代表`DATA ·Name+0(SB)/8,$·Value(SB)`，內的`Value`是參考golang中的`Value`變數而非plan9的`Value`變數

```bash
/var/www # go run main.go 
0
```

如果將`·Value`多定義一個GLOBL則golang `Value`變數會先參考到plan9的`Value`變數，之後`Name`變數被賦予golang `Value`變數內存地址時就會得到`1234(0x4d2)`的結果

```assembly
# test/test_amd64.s
DATA ·Value+0(SB)/8,$0x4d2

DATA ·Name+0(SB)/8,$·Value(SB)
GLOBL ·Name(SB),$8
```

```bash
/var/www # go run main.go 
1234
```

因此得出下列結論

1. main.go設定全域變數是否能參考到plan9定義的變數呢？
2. DATA ·Name+0(SB),[Value]，中DATA對於`Value`變數的能見度呢？

| 文件/代碼        | GLOBL Name | GLOBL Value | GLOBL ·Name | GLOBL ·Value | Name   | Value  | ·Name  | ·Value |
| ---------------- | ---------- | ----------- | ----------- | ------------ | ------ | ------ | ------ | ------ |
| main.go          | 不可見     | 不可見      | 可見        | 可見         | 不可見 | 不可見 | 不可見 | 不可見 |
| DATA ·Name+0(SB) |            | 可見        |             | 可見         |        | 不可見 |        | 不可見 |

總結

1. golang想要引用到plan9中定義的變數，plan9內變數需以`·symbol`方式定義且加上GLOBL
2. plan9內變數之間要互相使用需定義成GLOBL





## 參考

[plan9 assembly 完全解析]([https://github.com/cch123/golang-notes/blob/master/assembly.md#plan9-assembly-%E5%AE%8C%E5%85%A8%E8%A7%A3%E6%9E%90](https://github.com/cch123/golang-notes/blob/master/assembly.md#plan9-assembly-完全解析))

[Go Assembly 示例](https://colobu.com/goasm/)

[Go语言高级编程](https://chai2010.cn/advanced-go-programming-book/ch3-asm/readme.html)

[A Quick Guide to Go's Assembler](https://golang.org/doc/asm#architectures)

[golang plan9 指令清單](https://golang.org/src/cmd/internal/obj/x86/anames.go)

[Go functions in assembly language](https://github.com/golang/go/files/447163/GoFunctionsInAssembly.pdf)