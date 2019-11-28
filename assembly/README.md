# 組合語言(汇编语言)

- [閱讀前提](#閱讀前提])

- [環境](#環境)

- [進位](#進位)

- [寄存器](#寄存器)

  - [通用寄存器](#通用寄存器)
  - [段寄存器](#段寄存器)
  - [標記寄存器](#標記寄存器)

- [區塊](#AT&T)

  - [段](#段)
  - [入口](#入口)
  
- [型別](#型別)

- [指令](#指令)

- [參考](#參考)

  

## 閱讀前提

1. 從事golang開發

   會這樣認為是因為golang一些重點如內存如何管理、GMP等通常都在runtime源碼內，可能單純看runtime源碼就可以懂，但你可能會一支半解，因為runtime一些method或是變數是用組合語言寫的，你可能問或許理解這個可能不重要，如果你問我非要這樣做的原因是什麼，我想應該是這個[答案1](https://www.zhihu.com/question/19712941/answer/12878634) [答案2](https://www.zhihu.com/question/23088538/answer/23717201)，當然如果是php、java開發者想往比較底層方向走 or 想了解源碼也是可以，個人是php -> golang．

2. 不講太細節

   個人是為了能夠看懂golang runtime源碼而寫的紀錄並非工作上真的要寫組合語言，單純只是為了看懂、理解，所以你真的想了解關於組合語言跟CPU之間的細節，可能不適合你，本篇單純入門而已．

3. 由於不同CPU或作業系統對於組合語言寫法或指令都各有不同，本章節介紹的只是一普遍概念，主要目地只是入門，往後看到類似不懂或是有疑惑的地方再自行深入學習，如果已有大量組合語言實戰者可能不適合觀看．

4. 組合語言有分成兩派如AT&T、intel，本章節主要以AT&T風格來編寫

5. 本章節有任何有誤之處請儘管提出來，如果我有空則會回覆



## 環境



## 進位



## 寄存器



### 通用寄存器

#### ax 

#### bx

#### cx

#### dx

#### di

#### si

#### sp

#### bp



### 段寄存器

#### cs

#### ds

#### ss

#### es

### fs

#### gs



#### ip

指令指針寄存器，存放的是CPU需要執行的下一條指令地址，每當執行完一條指令之後，這個寄存器會自動移動到下一條指令的地址，不需要人為介入由CPU自行實作

```assembly
.section .text

.globl _start

_start:
    movq $1, %rax
    movq $2, %rax
    movq $3, %rax
```

上述範例可以知道入口點是`_start`，以gdb設置中斷點並執行

```bash
(gdb) b *_start
Breakpoint 1 at 0x401000: file test.s, line 6.
(gdb) r
Breakpoint 1, _start () at test.s:6
```

入口點是0x401000，看一下rip是指向0x401000，從0x401000確認後面的指令地址

```bash
(gdb) i r $rip
rip            0x401000            0x401000 <_start>
(gdb) x /3i 0x401000
=> 0x401000 <_start>:       mov    $0x1,%rax
   0x401007 <_start+7>:     mov    $0x2,%rax
   0x40100e <_start+14>:    mov    $0x3,%rax

```

這時候單部執行看看rip是指向0x401007，也就是`movq $2, %rax`

```bash
(gdb) n
7           movq $2, %rax
rip            0x401007            0x401007 <_start+7>
```

另外gdb內有一個`$pc`跟rip是一樣指向下一次要執行的指令地址，但$pc在組合語言是不存在，只是gdb內有這樣一個變數可以用

```bash
(gdb) x /3i $pc
   0x401000 <_start>:       mov    $0x1,%rax
=> 0x401007 <_start+7>:     mov    $0x2,%rax
   0x40100e <_start+14>:    mov    $0x3,%rax
(gdb) p $pc
$2 = (void (*)()) 0x401007 <_start+7>
```





### 標記寄存器

#### 



#### 



| 寄存器     | 16位 | 32位 | 64位 |
| :--------- | :--- | :--- | :--- |
| 累加寄存器 | AX   | EAX  | RAX  |
| 基址寄存器 | BX   | EBX  | RBX  |
| 计数寄存器 | CX   | ECX  | RCX  |
| 数据寄存器 | DX   | EDX  | RDX  |
| 堆栈基指针 | BP   | EBP  | RBP  |
| 变址寄存器 | SI   | ESI  | RSI  |
| 堆栈顶指针 | SP   | ESP  | RSP  |
| 指令寄存器 | IP   | EIP  | RIP  |



## AT&T

### 段

#### .section

定義一個段落



#### .data

該段落放置已有初始化的變數

| 命令  | 介紹 |
| ----- | ---- |
| .fill |      |



#### .bss

該段落放置未有初始化的變數

| 命令   | 介紹 |
| ------ | ---- |
| .comm  |      |
| .lcomm |      |



#### .text

該段落放置指令代碼



### 入口

#### .globl

#### _start



#### .locmm



### .comm



## 型別

| 命令    | 型別       |
| ------- | ---------- |
| .ascii  | string     |
| .byte   | byte       |
| .double | double     |
| .float  | float      |
| .int    | 64位元整數 |
| .long   | 64位元整數 |
| .short  | 16位整数   |
| .equ    | 常數       |

### 

## 指令

組合語言的動作表示，指令後綴詞代表不同位元操作

| 後綴詞 | 位元 |
| ------ | ---- |
| b      | 8    |
| w      | 16   |
| l      | 32   |
| q      | 64   |

Ex: movb（8位）、movw（16位）、movl（32位）、movq（64位）



### mod

資料傳送指令，如a = b + c，將b+c結果賦予a 

格式為: `mod <source>,<destination`> 等於  `destination` = `source` + `destination`

規則如下

1. 直接傳送數據給寄存器，數據用`$`表示，這種數據稱為立即數

   | 範例             | rax值(十進位) |
   | ---------------- | ------------- |
   | movq $0, %rax    | 0             |
   | movq $0x01, %rax | 1             |

   ```assembly
   .section .text
   
   .globl _start
   
   _start:
       movq $0, %rax
       movq $0x01, %rax
   ```
   

   
2. 直接傳送數據給內存(變數)，這邊的內存其實就是values變數的內存地址，執行時會自動分配一塊內存給values，所以只用values就可以表示該變數的內存地址

   | 範例             | values位置內的值(十進位) |
   | ---------------- | ------------------------ |
   | movq $10, values | 10                       |

   ```assembly
   .section .data
   values:
      .int 0
   
   .section .text
   
   .globl _start
   
   _start:
       movq $10, values
   ```
   

實際上`movq $10, values`該段會被解析成` movq 0x402000, %rax `

3. 寄存器傳給寄存器

    | 範例            | rax值 | rbx值   |
    | :-------------- | ----- | ------- |
    | movq %rax, %rbx | 不變  | rax的值 |

    ```assembly
    .section .text
    
    .globl _start
    
    _start:
        movq $1, %rax
        movq $2, %rbx
        movq %rax, %rbx
    ```
    

    
4. 寄存器傳給內存(變數地址)

   | 範例              | values位置內的值 |
   | :---------------- | ---------------- |
   | movq %rax, values | rax的值          |

   ```assembly
   .section .data
   values:
      .int 0
   
   .section .text
   
   .globl _start
   
   _start:
       movq $1, %rax
       movq %rax, values
   ```
   

   
5. 內存(變數地址)傳給寄存器
   | 範例              | rax值      |
   | :---------------- | ---------- |
   | movq values, %rax | values的值 |

   ```assembly
   .section .data
   values:
      .int 1
   
   .section .text
   
   .globl _start
   
   _start:
       movq $0, %rax
       movq values, %rax
   ```
   

   
6. 變數內存地址傳給寄存器，這種稱為内存變址尋址(Indexed Addressing Mode)

   規則表示為`base_address(offset_address, index, size)`

   計算方式為`base_address + offset_address + index * size`

   | 欄位           | 說明                             | 值       |
   | -------------- | -------------------------------- | -------- |
   | base_address   | 變數一開始的地址                 | 0x000000 |
   | offset_address | 想要位移的地址，只能用寄存器表示 | 0x000004 |
   | index          | 想要偏移次數，只能用寄存器表示   | 1        |
   | size           | 每次偏移多少                     | 4        |

   假設value變數內存地址從0x000000開始，value有三個int值(陣列)分別為10,20,30，%rax是0x000004、%rbx是0x000001、size是4，則value內存結構為下

   | 地址     | 值   |
   | -------- | ---- |
   | 0x000000 | 10   |
   | 0x000001 | 0    |
   | 0x000002 | 0    |
   | 0x000003 | 0    |
   | 0x000004 | 20   |
   | 0x000005 | 0    |
   | 0x000006 | 0    |
   | 0x000007 | 0    |
   | 0x000008 | 30   |
   | 0x000009 | 0    |
   | 0x00000a | 0    |
   | 0x00000b | 0    |

   使用一個範例說明

   | 範例                             | rax值 | rbx值 | rcx值 |
   | :------------------------------- | ----- | ----- | ----- |
   | movq values(%rax, %rbx, 4), %rcx | 不變  | 不變  | 30    |

   ```assembly
   .section .data
   values:
      .int 10,20,30
   
   .section .text
   
   .globl _start
   
   _start:
       movq $4, %rax
       movq $1, %rbx
       movq values(%rax, %rbx, 4), %rcx
   ```
   

根據規則分別步驟如下

1. `base_address`是value變數地址是0x000000
   2. `offset_address`是%rax為0x000004
   3. `base_address` + `offset_address` = 0x000004
   4. `index` * `size` = 4，以十六進位表示為0x000004
   5. `base_address + offset_address + index * size` = 0x000004 + 0x000004 = 0x000008
   6. 對照內存結構表0x000008上的值是30
   7. %rcx是64位元所以值會是0x000008 ～ 0x00000b，答案就是30
   
7. 內存地址傳給寄存器

   首先比較下方兩個範例差異，values是一個int變數值為10

   | 範例               | rax值    |
   | :----------------- | -------- |
   | movq values, %rax  | 10       |
   | movq $values, %rax | 0x402000 |

   很明顯可以看出來是`values`與`$values`之間有不同的意義

   `movq values, %rax `代表將`values的值`傳給`rax`

   `movq $values, %rax `代表將`values的地址`傳給`rax`，所以`rax`內真正儲存的是地址，代表`rax`已變成一個指標(pointer or 指針)

   ```assembly
   .section .data
   values:
      .int 10
   
   .section .text
   
   .globl _start
   
   _start:
       movq values, %rax
       movq $values, %rax
   ```
   

   
8. 寄存器間接尋址

   把一個寄存器當作指標使用，也就是寄存器不存儲一般數據值而是儲存內存地址，在利用另一個寄存器傳送數據給指標寄存器

   | 範例              | rbx的值 |
   | ----------------- | ------- |
   | movq %rax, (%rbx) | 不變    |
   
   如下範例會將values資料改成5，不直接對values改值而是透過寄存器去更改
   
   ```assembly
   .section .data
   values:
      .int 10
   
   .section .text
   
   .globl _start
   
   _start:
       movq $5, %rax
       movq $values, %rbx
       movq %rax, (%rbx)
   ```
   
   設定偏移量來更改，如下範例values資料會變成`10,5`
   
   ```assembly
   .section .data
   values:
      .int 10,20
   
   .section .text
   
   .globl _start
   
   _start:
       movq $5, %rax
       movq $values, %rbx
       movq %rax, 4(%rbx)
   ```
   
   為什麼會改動values中的第二個位置的值(20)而非第一個位置的值(10)，原因在尋址的時候多偏移了4個位元`4(%rbx)`，以下先給出values內存結構
   | 地址     | 值   |
   | -------- | ---- |
   | 0x000000 | 10   |
   | 0x000001 | 0    |
   | 0x000002 | 0    |
   | 0x000003 | 0    |
   | 0x000004 | 20   |
   | 0x000005 | 0    |
   | 0x000006 | 0    |
   | 0x000007 | 0    |
   
   (%rbx)值表示values內存地址，初始地址則是0x000000，這時前方多了個數字表示在從該地址往後偏移幾個位元，範例是4所以`0x000000 + 0x000004` = `0x000004`，所以是從20開始．也就將%rax值覆蓋到0x000004上替換掉原先的值(20)，改成5
   
   
   

### add



### sub



### div



### push

將資料放到stack中

格式為: `push <source>`

可接受的source有以下

1. 寄存器
2. 內存
3. 立即數

實際範例如下

```assembly
.section .data

data:
   .int 3

.section .text

.globl _start

_start:
    movq $2, %rax
    pushq $1
    pushq %rax
    pushq data
```

看看push完後stack內存會是如何，每個數字是64位元所以站8個byte，範例中總共push 三個64位元數字順序是1 -> 2 -> 3，從`0x7fffffffed08`到`0x7fffffffed18`佔用24byte，因為3 * 8

```bash
(gdb) x /24db $rsp
0x7fffffffed08: 3       0       0       0       0       0       0       0
0x7fffffffed10: 2       0       0       0       0       0       0       0
0x7fffffffed18: 1       0       0       0       0       0       0       0
```

stack特性是先進後出的概念，看一下rsp寄存器目前指向0x7fffffffed08，代表是數字3，而3是最後才push進stack的，所以rsp會一直指向最後push進入的資料內存地址

```bash
(gdb) i f
Stack level 0, frame at 0x7fffffffed10:
 rip = 0x40100c; saved rip = 0x3
 called by frame at 0x7fffffffed18
 Arglist at 0x7fffffffed00, args: 
 Locals at 0x7fffffffed00, Previous frame's sp is 0x7fffffffed10
 Saved registers:
 rip at 0x7fffffffed08
```

### pop

從stack拿出資料

格式為: `pop <destination`>

可接受的destination有以下

1. 寄存器
2. 內存

實際範例如下

```assembly
.section .data

data:
   .int 3

.section .text

.globl _start

_start:
    movq $2, %rax
    pushq $1
    pushq %rax
    pushq data
    popq %rbx
    popq %rbx
    popq %rbx
```

先看看執行到第一個`popq %rbx`後，rsp寄存器指向0x7fffffffed10，原先是指向0x7fffffffed08，代表pop一次則減0x08，而rbx存放著3，這邊以ebx表示32位元更能清楚顯示

```bash
(gdb) x /24b $rsp
0x7fffffffed10: 2       0       0       0       0       0       0       0
0x7fffffffed18: 1       0       0       0       0       0       0       0
0x7fffffffed20: 1       0       0       0       0       0       0       0
(gdb) i r $rbx
rbx            0x100000003         4294967299
(gdb) i r $ebx
ebx            0x3                 3
```

等到pop都執行完後看看rsp寄存器指向0x7fffffffed20，之前存進去的1,2,3數字都已從stack取出了

```bash
(gdb) i f
Stack level 0, frame at 0x7fffffffed28:
 rip = 0x401014; saved rip = 0x1
 called by frame at 0x7fffffffed30
 Arglist at 0x7fffffffed18, args: 
 Locals at 0x7fffffffed18, Previous frame's sp is 0x7fffffffed28
 Saved registers:
 rip at 0x7fffffffed20
```



### jmp

跳轉指令

格式為: jmp <location>`



### inc



### dec



### loop



### and



### or



### xchg

交換兩個數據，交換時會lock所以效能會有所影響

1. 寄存器之間交換

   | 範例            | rax的值 | rbx的值 |
   | --------------- | ------- | ------- |
   | xchg %rax, %rbx | rbx的值 | rax的值 |
   
   ```assembly
   .section .text
   
   .globl _start
   
   _start:
      movq $100, %rax
      movq $200, %rbx
      xchg %rax, %rbx
   ```
   
   
   
2. 寄存器與內存交換，不能兩個內存做交換
   | 範例              | values的值 | rax的值    |
   | ----------------- | ---------- | ---------- |
   | xchg values, %rax | rax的值    | values的值 |
   
   ```assembly
   .section .data
   values:
      .int 100
   
   .section .text
   
   .globl _start
   
   _start:
      movq $200, %rax
      xchg values, %rax
   ```






## 參考

1. 必看

   [汇编语言(第3版)-王爽](#https://book.douban.com/subject/25726019/) 基礎入門，使用intel寫法較簡單，我個人覺得當學習概念就好不用全懂，主要是有一個帶入感讓你覺得CPU是怎樣了解程式語言

   [汇编语言程序设计](https://book.douban.com/subject/1446250/) AT&T寫法入門，跟intel相比是另一派風格也較難一點，linux主要都是AT&T 寫法，基本上這本必看，本書例子是32位元但是開發機大部分都是64位元所以給出來的範例會有少許錯誤，請自行google解決 ex: push在64與32位元上的[差異](https://stackoverflow.com/questions/5485468/x86-assembly-pushl-popl-dont-work-with-error-suffix-or-operands-invalid)

   

   閱讀順序是`汇编语言(第3版)` -> `汇编语言程序设计`

2. gdb 操作

   https://wizardforcel.gitbooks.io/100-gdb-tips/print-registers.html

   https://sourceware.org/gdb/current/onlinedocs/gdb/

3. 操作

   http://ouonline.net/tag/att-asm