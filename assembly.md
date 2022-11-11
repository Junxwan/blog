# 組合語言(汇编语言)

- [閱讀前提](#閱讀前提)

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

  - [mov](#mov)
  - [add](#add)
  - [sub](#sub)
  - [mul](#mul)
  - [div](#div)
  - [and](#and)
  - [or](#or)
  - [not](#not)
  - [inc](#inc)
  - [dec](#dec)
  - [cmp](#cmp)
  - [loop](#loop)
  - [lea](#lea)
  - [push](#push)
  - [pop](#pop)
  - [jmp](#jmp)
  - [call](#call)
  - [ret](#ret)
  - [xchg](#xchg)

- [標誌](#標誌)

- [stack](#stack)

- [參考](#參考)

  

## 閱讀前提

1. 從事golang開發

   會這樣認為是因為golang一些重點如內存如何管理、GMP等通常都在runtime源碼內，可能單純看runtime源碼就可以懂，但你可能會一支半解，因為runtime一些method或是變數是用組合語言寫的，你可能問或許理解這個可能不重要，如果你問我非要這樣做的原因是什麼，我想應該是這個[答案1](https://www.zhihu.com/question/19712941/answer/12878634) [答案2](https://www.zhihu.com/question/23088538/answer/23717201)，當然如果是php、java開發者想往比較底層方向走或想了解源碼也是可以，個人是php -> golang．

2. 不講太細節

   個人是為了能夠看懂golang runtime源碼而寫的紀錄並非工作上真的要寫組合語言，單純只是為了看懂、理解，所以你真的想了解關於組合語言跟CPU之間的細節，可能不適合你，本篇單純入門而已．

3. 由於不同CPU或作業系統對於組合語言寫法或指令都各有不同，本章節介紹的只是一個基礎概念，主要目地只是入門，看到不懂或有疑惑的地方再請自行深入學習，如果已有大量組合語言實戰者可能不適合觀看．

4. 組合語言有分成兩派如AT&T、intel，本章節主要以AT&T風格來編寫

5. 本章節內容有任何錯誤請儘管提出來，如果我有空則會回覆



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

入口點是`0x401000`，看一下`rip`是指向`0x401000`，從`0x401000`確認後面的指令地址，此時還未執行`0x401000`上的指令

```bash
(gdb) i r $rip
rip            0x401000            0x401000 <_start>
(gdb) x /3i 0x401000
=> 0x401000 <_start>:       mov    $0x1,%rax
   0x401007 <_start+7>:     mov    $0x2,%rax
   0x40100e <_start+14>:    mov    $0x3,%rax

```

| 內存地址 | 指令          |      |
| -------- | ------------- | ---- |
| 0x401000 | mov $0x1,%rax | rip  |
| 0x401007 | mov $0x2,%rax |      |
| 0x40100e | mov $0x3,%rax |      |

這時候單部執行`0x401000`上的`mov $0x1,%rax` ，`rip`是指向`0x401007`，也就是`movq $2, %rax`

```bash
(gdb) n
7           movq $2, %rax
(gdb) i r $rip
rip            0x401007            0x401007 <_start+7>
```

| 內存地址 | 指令          |      |
| -------- | ------------- | ---- |
| 0x401000 | mov $0x1,%rax |      |
| 0x401007 | mov $0x2,%rax | rip  |
| 0x40100e | mov $0x3,%rax |      |

另外gdb內有一個`$pc`跟`rip`是一樣指向下一次要執行的指令地址，但`$pc`在組合語言是不存在，只是gdb內有這樣一個變數可以用

```bash
(gdb) x /3i $pc
   0x401000 <_start>:       mov    $0x1,%rax
=> 0x401007 <_start+7>:     mov    $0x2,%rax
   0x40100e <_start+14>:    mov    $0x3,%rax
(gdb) p $pc
$2 = (void (*)()) 0x401007 <_start+7>
```



### 標記寄存器



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



## 區塊

### 段

#### .section

定義一個段落

```assembly
.section .text
```



#### .data

該段落放置已有初始化的變數

```assembly
.section .data
```

| 命令  | 介紹 |
| ----- | ---- |
| .fill |      |



#### .bss

該段落放置未有初始化的變數

```assembly
.section .bss
```

| 命令   | 介紹 |
| ------ | ---- |
| .comm  |      |
| .lcomm |      |



#### .text

該段落放置指令代碼

```assembly
.section .text
```



### 入口

#### .globl

#### _start

#### .locmm

#### .comm



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

| 後綴詞 | 位元 | Data Type |      |
| ------ | ---- | --------- | ---- |
| b      | 8    | [1]byte   | movb |
| w      | 16   | [2]byte   | movw |
| l      | 32   | [4]byte   | movl |
| q      | 64   | [8]byte   | movq |



### mov

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

   假設value變數內存地址從`0x000000`開始，value有三個int值(陣列)分別為10,20,30，`rax`是`0x000004`、`rbx`是`0x000001`、size是4，則value內存結構為下

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

1. `base_address`是value變數地址是`0x000000`
  
   2. `offset_address`是%rax為`0x000004`
   3. `base_address` + `offset_address` = `0x000004`
   4. `index` * `size` = 4，以十六進位表示為`0x000004`
   5. `base_address + offset_address + index * size` = `0x000004 + 0x000004 = 0x000008`
   6. 對照內存結構表`0x000008`上的值是30
   7. `rcx`是64位元所以值會是`0x000008` ～ `0x00000b`，答案就是30
   
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
   
   `(%rbx)`值表示values內存地址，初始地址則是`0x000000`，這時前方多了個數字表示在從該地址往後偏移幾個位元，範例是4所以`0x000000 + 0x000004` = `0x000004`，所以是從20開始．也就將`rax`值覆蓋到`0x000004`上替換掉原先的值(20)，改成5
   
   
   

### add

### sub

### mul

### div

### and

### or

### not

### inc

### dec

### cmp



### loop

迴圈指令

格式為: `loop <location>`

實際範例如下，透過`rcx`設置迴圈次數，每跑一次loop，`rcx`就會遞減一次直到零才不會繼續跑

```assembly
.section .text

.globl _start

_start:
   movq $10, %rcx
   movq $0, %rax
loop1:
   addq $1, %rax
   loop loop1
```

`rax`值為10，因為跑了10次加1

```bash
(gdb) i r $rax
rax            0xa                 10
```



### lea

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

stack特性是先進後出的概念，看一下`rsp`寄存器目前指向`0x7fffffffed08`，代表是數字3，而3是最後才push進stack的，所以`rsp`會一直指向最後push進入的資料內存地址

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

先看看執行到第一個`popq %rbx`後，`rsp`寄存器指向`0x7fffffffed10`，原先是指向`0x7fffffffed08`，代表pop一次則減`0x08`，而`rbx`存放著3，這邊以`ebx`表示32位元更能清楚顯示

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

等到pop都執行完後看看`rsp`寄存器指向`0x7fffffffed20`，之前存進去的1,2,3數字都已從stack取出了

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

格式為: `jmp <location>`

實際範例如下，`int $0x80 `代表exit

```assembly
.section .text

.globl _start

_start:
   movq $1, %rax
   jmp test
   movq $3, %rax
   int $0x80 

test:
   movq $2, %rax
```

最後`rax`值是2，因為`jmp test`會跳到`test`標籤上執行`movq $2, %rax`完後就結束了，jmp是如何知道`test`標籤在哪裡呢？，如下可以看到`jmp 0x401010 <test>`，就是這個`0x401010`代表`test`，直接跳過不執行`mov $0x3,%rax`

```bash
(gdb) x /4i $pc
=> 0x401000 <_start>:   mov    $0x1,%rax
   0x401007 <_start+7>: jmp    0x401010 <test>
   0x401009 <_start+9>: mov    $0x3,%rax
   0x401010 <test>:     mov    $0x2,%rax
```



### call

調用指令

格式為: `call <address>`

實際範例如下

```assembly
.section .text

.globl _start

_start:
   movq $1, %rax
   call test
   addq $3, %rax
   int $0x80 

test:
   addq $2, %rax
   ret
```

最後`rax`值為6，因為調用`test`對`rax`多增加2，首先利用gdb來debug，上述範例檔名為test.s

```bash
as -gstabs -o test.o test.s && ld -o test test.o && gdb ./test
```

斷點入口，查看目前組合語言內存位置與`rax`、 `rip`值

```bash
(gdb) b *_start
Breakpoint 1 at 0x401000: file test.s, line 6.
(gdb) r
Starting program: /var/www/assembly/test 

Breakpoint 1, _start () at test.s:6
6          movq $1, %rax
(gdb) x /6i $pc
=> 0x401000 <_start>:   mov    $0x1,%rax
   0x401007 <_start+7>: callq  0x401012 <test>
   0x40100c <_start+12>:        add    $0x3,%rax
   0x401010 <_start+16>:        int    $0x80
   0x401012 <test>:     add    $0x2,%rax
   0x401016 <test+4>:   retq 
(gdb) i r $rax
rax            0x0                 0
(gdb) i r $rip
rip            0x401000            0x401000 <_start>
(gdb) i r $rsp
rsp            0x7fffffffed10      0x7fffffffed10
(gdb) x /32db $rsp
0x7fffffffed10: 1       0       0       0       0       0       0       0
0x7fffffffed18: -39     -18     -1      -1      -1      127     0       0
0x7fffffffed20: 0       0       0       0       0       0       0       0
0x7fffffffed28: -16     -18     -1      -1      -1      127     0       0
```

| 寄存器 | 值             |
| ------ | -------------- |
| rax    | 0x0            |
| rip    | 0x401000       |
| rsp    | 0x7fffffffed10 |

執行`movq $1, %rax`，針對`rax`+1

```bash
(gdb) n
7          call test
(gdb) x /6i $pc
=> 0x401007 <_start+7>: callq  0x401012 <test>
   0x40100c <_start+12>:        add    $0x3,%rax
   0x401010 <_start+16>:        int    $0x80
   0x401012 <test>:     add    $0x2,%rax
   0x401016 <test+4>:   retq   
   0x401017:    add    %al,(%rcx)
(gdb) i r $rax
rax            0x1                 1
(gdb) i r $rip
rip            0x401007            0x401007 <_start+7>
(gdb) i r $rsp
rsp            0x7fffffffed10      0x7fffffffed10
(gdb) x /32db $rsp
0x7fffffffed10: 1       0       0       0       0       0       0       0
0x7fffffffed18: -39     -18     -1      -1      -1      127     0       0
0x7fffffffed20: 0       0       0       0       0       0       0       0
0x7fffffffed28: -16     -18     -1      -1      -1      127     0       0
```

| 寄存器 | 值             |
| ------ | -------------- |
| rax    | 0x1            |
| rip    | 0x401007       |
| rsp    | 0x7fffffffed10 |

執行`call test`，進到`test`

```bash
(gdb) s
test () at test.s:12
12         addq $2, %rax
(gdb) x /6i $pc
=> 0x401012 <test>:     add    $0x2,%rax
   0x401016 <test+4>:   retq   
   0x401017:    add    %al,(%rcx)
   0x401019:    add    %al,(%rax)
   0x40101b:    add    %al,(%rax)
   0x40101d:    add    %al,(%rdi)
(gdb) i r $rax
rax            0x1                 1
(gdb) i r $rip
rip            0x401012            0x401012 <test>
(gdb) i r $rsp
rsp            0x7fffffffed08      0x7fffffffed08
(gdb) x /32db $rsp
0x7fffffffed08: 12      16      64      0       0       0       0       0
0x7fffffffed10: 1       0       0       0       0       0       0       0
0x7fffffffed18: -39     -18     -1      -1      -1      127     0       0
0x7fffffffed20: 0       0       0       0       0       0       0       0
```

| 寄存器 | 值             |
| ------ | -------------- |
| rax    | 0x1            |
| rip    | 0x401012       |
| rsp    | 0x7fffffffed08 |

這裡可以發現怎直接跳到`0x401012`內存地址上的代碼，而非`0x40100c`，可得知`call`指令具備更改`rip`的操作，更改的值是`call`目標的內存地址

```assembly
movq 0x401012, %rip
```

再看看`rsp`突然從`0x7fffffffed10`變成`0x7fffffffed08`，減少`0x08`，而`0x7fffffffed08`的值代表`0x40100c`，對照一開始查看組合語言代表內存位置上剛好指向`addq $3, %rax`，是執行完`call test`的下一條指令

```bash
(gdb) x /8b 0x7fffffffed08
0x7fffffffed08: 0x0c    0x10    0x40    0x00    0x00    0x00    0x00    0x00
```

意思是在執行`call`時會先將下一條指令push進stack內

```assembly
pushq ％rip
```

總結可以得到`call`操作其實是先將當前`rip`值push進stack後再更改`rip`值為目標內存地址，利用rip跳轉到目標而為什麼要`pushq ％rip`在[ret](#ret)會有詳細解說

```assembly
pushq ％rip
movq 0x401012, %rip
```



### ret

跳轉指令

實際範例如下

```assembly
.section .text

.globl _start

_start:
   movq $1, %rax
   call test
   addq $3, %rax
   int $0x80

test:
   addq $2, %rax
   ret
```

承接[call](#call)說明可以知道執行`call`會有一個動作是將當前`rip`push進入stack，為什麼要這樣做？．其實主因是來自`call`完目標後需要返回原始地方繼續執行接下來的指令，所以需要先保存來源才能返回，表示返回的動作就是`ret`

依據[call](#call)解說進到`test`後

```bash
(gdb) x /5i $pc
=> 0x401012 <test>:     add    $0x2,%rax
   0x401016 <test+4>:   retq   
   0x401017:    add    %al,(%rcx)
   0x401019:    add    %al,(%rax)
   0x40101b:    add    %al,(%rax)
(gdb) x /32db $rsp
0x7fffffffed08: 12      16      64      0       0       0       0       0
0x7fffffffed10: 1       0       0       0       0       0       0       0
0x7fffffffed18: -39     -18     -1      -1      -1      127     0       0
0x7fffffffed20: 0       0       0       0       0       0       0       0
(gdb) i r $rip
rip            0x401012            0x401012 <test>
(gdb) i r $rsp
rsp            0x7fffffffed08      0x7fffffffed08
```

| 寄存器 | 值             |
| ------ | -------------- |
| rip    | 0x401012       |
| rsp    | 0x7fffffffed08 |

等待執行`ret`

```bash
(gdb) n
test () at test.s:13
13         ret
(gdb) x /32db $rsp
0x7fffffffed08: 12      16      64      0       0       0       0       0
0x7fffffffed10: 1       0       0       0       0       0       0       0
0x7fffffffed18: -39     -18     -1      -1      -1      127     0       0
0x7fffffffed20: 0       0       0       0       0       0       0       0
(gdb) i r $rip
rip            0x401016            0x401016 <test+4>
(gdb) i r $rsp
rsp            0x7fffffffed08      0x7fffffffed08
```

| 寄存器 | 值             |
| ------ | -------------- |
| rip    | 0x401016       |
| rsp    | 0x7fffffffed08 |

stack內存結構

| 內存地址       | 值       | 說明      |
| -------------- | -------- | --------- |
| 0x7fffffffed10 | 1        | stack top |
| 0x7fffffffed08 | 0x40100c | 返回地址  |

執行`ret`

```bash
(gdb) n
_start () at test.s:8
8          addq $3, %rax
(gdb) x /5i $pc
=> 0x40100c <_start+12>:        add    $0x3,%rax
   0x401010 <_start+16>:        int    $0x80
   0x401012 <test>:     add    $0x2,%rax
   0x401016 <test+4>:   retq   
   0x401017:    add    %al,(%rcx)
(gdb) x /32db $rsp
0x7fffffffed10: 1       0       0       0       0       0       0       0
0x7fffffffed18: -39     -18     -1      -1      -1      127     0       0
0x7fffffffed20: 0       0       0       0       0       0       0       0
0x7fffffffed28: -16     -18     -1      -1      -1      127     0       0
(gdb) i r $rip
rip            0x40100c            0x40100c <_start+12>
(gdb) i r $rsp
rsp            0x7fffffffed10      0x7fffffffed10
```

| 寄存器 | 值             |
| ------ | -------------- |
| rip    | 0x40100c       |
| rsp    | 0x7fffffffed10 |

stack內存結構

| 內存地址       | 值       | 說明      |
| -------------- | -------- | --------- |
| 0x7fffffffed10 | 1        | stack top |
| 0x7fffffffed08 | 0x40100c | 返回地址  |

可以看到function結束後回到`addq $3, %rax`，而且`rip`指向`0x40100c`就是`call`時事先push進stack內的下一個要執行的指令地址，這時在執行`ret`時被拿出來指向`rip`，同時`rsp`也指回stack top．

從這些跡象來看可以得出`ret`做了以下指令，該指令是拿出當前`rsp`值並傳給`rip`，同時還做`rsp` - `0x08`的動作

```assembly
popq %rip
```

依據上述解說可以得知，`ret`依靠更改`rip`方式返回原來`call`指令完成後的下一個指令，但前提是在執行`ret`時需讓`rsp`指向`返回地址`才有用，也就是一開始執行`call`push 



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



### 



## 標誌

### cf

### of

### pf

### sf

### zf



## stack

本節介紹function如何利用stack保存參數,區域變數,執行結果

以下有一個`add` function，邏輯是傳遞一個int參數並加ㄧ再回傳結果，golang範例如下

```go
package main

import "fmt"

var result = 0

func main() {
	result = add(2)
	result += add(4)

	// output 8
	fmt.Println(result)
}

func add(a int) int {
	b := 1
	return a + b
}
```

在開始討論上述範例中stack變化是如何前，先構想一下會有哪些問題出現

1. stack怎麼儲存參數呢？
2. 當調用 `add(2)`完成後怎麼回到原始的調用位置呢？
3. function怎用從stack拿出參數 ?
4. stack怎麼儲存區域變數呢？
5. function結束後怎樣清空stack？
6. function怎樣回傳結果？

首先從問題(1)點開始解釋，編譯器編譯時會知道`add()`function有一個int參數，所以只要調用`add()`時，stack會優先將int參數先做push的動作 

如下是未調用`add(2)`前的stack內存結構

| 內存地址       | 值   | 說明                  |                   |
| -------------- | ---- | --------------------- | ----------------- |
| 0x7fffffffed20 |      | stack一開始的內存位置 | rsp指向的內存地址 |

當調用`add(2)`後的stack內存結構

```assembly
 pushq $2 # 放到0x7fffffffed18
```

| 內存地址       | 值   | 說明                  |                   |
| -------------- | ---- | --------------------- | ----------------- |
| 0x7fffffffed20 |      | stack一開始的內存位置 |                   |
| 0x7fffffffed18 | 2    | add() int參數         | rsp指向的內存地址 |



問題(2)，當調用`add(2)`時，在跳轉到`add()`function前先向stack push下一個執行點(內存地址)，以剛剛的範例來看，`add(2)`執行完後下一個要執行的是`add(4)`，所以stack push`result += add(4)`這段內存地址，這樣function結束後才知道要回去哪一個內存地址開始繼續執行

```go
result = add(2)  // 地址：0x401002
result += add(4) // 地址：0x401007，此地址被放到0x7fffffffed10
```

| 內存地址       | 值       | 說明                       |                   |
| -------------- | -------- | -------------------------- | ----------------- |
| 0x7fffffffed20 |          | stack一開始的內存位置      |                   |
| 0x7fffffffed18 | 2        | add() int參數              |                   |
| 0x7fffffffed10 | 0x401007 | `result += add(4)`內存地址 | rsp指向的內存地址 |



問題(3)主要是透過`rsp`間接尋址的方式拿到參數，比如先前push的`add()` int參數 是64位元，只要把當前`rsp`位址加上`0x08`就得到該參數的內存地址，就可以拿到該參數的值，如下範例將參數值放到`rax`中

```assembly
movq 8(%rsp), %rax # 0x7fffffffed10 + 0x08 = 0x7fffffffed18 
```

上述使用`rsp`來做間接尋址，但有個問題是在function中操作`push` `pop`這些指令會使`rsp`內存位置又變動，上述案例是加上`0x08`，如`push`就變成加上`0x10`之後又`pop`才會變成加上`0x08`，這樣會讓編譯器無法固定間接尋址偏移量，所以需要把`rsp`內存地址放到另外一個寄存器，也就是`rbp`，但是`rbp`也可能保存著其他重要的值，所以在利用`rbp`儲存`rsp`值之前須把`rbp` push進入stack做保存，這裡假設原先`rbp`值是10

| 寄存器 | 值             |
| ------ | -------------- |
| rsp    | 0x7fffffffed08 |
| rbp    | 10             |

```assembly
pushq %rbp # 放到0x7fffffffed08
```

| 內存地址       | 值       | 說明                       |                   |
| -------------- | -------- | -------------------------- | ----------------- |
| 0x7fffffffed20 |          | stack一開始的內存位置      |                   |
| 0x7fffffffed18 | 2        | add() int參數              |                   |
| 0x7fffffffed10 | 0x401007 | `result += add(4)`內存地址 |                   |
| 0x7fffffffed08 | 10       | 原先rbp值                  | rsp指向的內存地址 |

將rbp原先值push進stack後，rsp值賦予給rbp

```assembly
movq %rsp, %rbp
```

| 寄存器 | 值             |
| ------ | -------------- |
| rsp    | 0x7fffffffed08 |
| rbp    | 0x7fffffffed08 |

所以在function中想從stack拿值就使用rbp做間接尋址，如下範例找到`add()` int參數 

```assembly
movq 16(%rbp), %rax # 0x7fffffffed08 + 0x10 = 0x7fffffffed18 
```



問題(4)在function會有所謂的區域變數，這些變數也都是放在stack中非放到寄存器中，原因是區域變數只作用在function內，無放到寄存器的必要，除非是要做特定運算而該指令只能接受寄存器當參數使用，否則都只放在stack內．依照上述golang範例在`add()`內有一個`b`區域int64變數，將b變數push入stack

```assembly
pushq $1 # 放到0x7fffffffed00
```

| 內存地址       | 值       | 說明                       | 間接尋址 |                   |
| -------------- | -------- | -------------------------- | -------- | ----------------- |
| 0x7fffffffed20 |          | stack一開始的內存位置      | 24(%rbp) |                   |
| 0x7fffffffed18 | 2        | add() int參數              | 16(%rbp) |                   |
| 0x7fffffffed10 | 0x401007 | `result += add(4)`內存地址 | 8(%rbp)  |                   |
| 0x7fffffffed08 | 10       | 原先rbp值                  | (%rbp)   |                   |
| 0x7fffffffed00 | 1        | 區域變數b                  | -8(%rbp) | rsp指向的內存地址 |



問題(5)當function要結束時需要把復原`rsp`，也就是還原到開始真正執行function邏輯前的地址，這地址在`rbp`上

```assembly
movq %rbp, %rsp
```

| 內存地址       | 值       | 說明                       | 間接尋址 |                   |
| -------------- | -------- | -------------------------- | -------- | ----------------- |
| 0x7fffffffed20 |          | stack一開始的內存位置      | 24(%rbp) |                   |
| 0x7fffffffed18 | 2        | add() int參數              | 16(%rbp) |                   |
| 0x7fffffffed10 | 0x401007 | `result += add(4)`內存地址 | 8(%rbp)  |                   |
| 0x7fffffffed08 | 10       | 原先rbp值                  | (%rbp)   | rsp指向的內存地址 |
| 0x7fffffffed00 | 1        | 區域變數b                  | -8(%rbp) |                   |

再來也需要還原`rbp`調用`add()`之前的值

```assembly
popq %rbp # 從0x7fffffffed08拿出值
```

| 內存地址       | 值       | 說明                       | 間接尋址 |                   |
| -------------- | -------- | -------------------------- | -------- | ----------------- |
| 0x7fffffffed20 |          | stack一開始的內存位置      | 24(%rbp) |                   |
| 0x7fffffffed18 | 2        | add() int參數              | 16(%rbp) |                   |
| 0x7fffffffed10 | 0x401007 | `result += add(4)`內存地址 | 8(%rbp)  | rsp指向的內存地址 |
| 0x7fffffffed08 | 10       | 原先rbp值                  | (%rbp)   |                   |
| 0x7fffffffed00 | 1        | 區域變數b                  | -8(%rbp) |                   |

延續問題(2)，function結束了要回到下一個要執行的內存地址，透過`ret`表示function結束了要返回原先調用的地方，目前rsp指向的值就是要返回的地址，會利用pop將返回地址值拿出來

```assembly
ret # 0x401007，可以當作popq %rip
```

| 內存地址       | 值       | 說明                       | 間接尋址 |                   |
| -------------- | -------- | -------------------------- | -------- | ----------------- |
| 0x7fffffffed20 |          | stack一開始的內存位置      | 24(%rbp) |                   |
| 0x7fffffffed18 | 2        | add() int參數              | 16(%rbp) | rsp指向的內存地址 |
| 0x7fffffffed10 | 0x401007 | `result += add(4)`內存地址 | 8(%rbp)  |                   |
| 0x7fffffffed08 | 10       | 原先rbp值                  | (%rbp)   |                   |
| 0x7fffffffed00 | 1        | 區域變數b                  | -8(%rbp) |                   |

由於先前調用`add()`有push一個int參數，所以編譯器在`result = add(2)`時除了編譯成`call add`plan9 其實後面還會多出一個動作來清除push的參數，如下組合語言

```assembly
 call add       
 addq $8, %rsp
```

可以看到`addq $8, %rsp`，這裡意思是要清除最初執行`add()`所push進的int64參數，上述執行完後stack內存結構如下

| 內存地址       | 值       | 說明                       | 間接尋址 |                   |
| -------------- | -------- | -------------------------- | -------- | ----------------- |
| 0x7fffffffed20 |          | stack一開始的內存位置      | 24(%rbp) | rsp指向的內存地址 |
| 0x7fffffffed18 | 2        | add() int參數              | 16(%rbp) |                   |
| 0x7fffffffed10 | 0x401007 | `result += add(4)`內存地址 | 8(%rbp)  |                   |
| 0x7fffffffed08 | 10       | 原先rbp值                  | (%rbp)   |                   |
| 0x7fffffffed00 | 1        | 區域變數b                  | -8(%rbp) |                   |

上述內存結構表可知`rsp`回到的stack top，也就是最原先調用`add()`前的狀態，如果之後執行`result += add(4)`就會在重覆上述流程，這時`0x7fffffffed18` ~ `0x7fffffffed00`上的值就會被覆蓋，所以每次調用function所使用的stack才不會互相不受影響．其實也就是將`rsp`指向最初狀態的地址，這樣就還原stack

問題(5)

TODO



以下使用一個組合語言範例來描述上述golang範例在做什麼，另外該組合語言非利用上述`go build` or `go tool`產生出來的組合語言，只是用組合語言方式呈現剛剛golang的功能

```assembly
.section .data
result:
   .int 0

.section .text

.globl _start

_start:
   pushq $2
   call add
   addq $8, %rsp
   movq %rax, result

   pushq $4
   call add
   addq $8, %rsp
   addq %rax, result

   int $0x80

.type add, @function
add:
   pushq %rbp
   movq %rsp, %rbp
   movq 16(%rbp), %rax
   pushq $1
   addq -8(%rbp), %rax

   movq %rbp, %rsp
   popq %rbp
   ret
```

|function| golang  | 組合語言                                                     |
|-|-----------------|------------------------------------------------------------ | -------- |
|main|result = add(2)  |pushq $2 </br> call add </br> addq $8, %rsp </br> movq %rax, result </br> |
|main|result += add(4) |pushq $4 </br> call add </br> addq $8, %rsp </br> addq %rax, result </br> |
|add| b := 1           |pushq $1                                                     |
|add| a + b            |addq -8(%rbp), %rax                                          |
|add| return           |movq %rbp, %rsp </br> popq %rbp </br> ret </br>              |



利用gdb來debug此範例，範例檔案名稱是`test.s`

```bash
as -gstabs -o test.o test.s && ld -o test test.o && gdb ./test
```

斷點到程式入口點`_start`

```bash
(gdb) b *_start
Breakpoint 1 at 0x401000: file test.s, line 10.
(gdb) r
Starting program: /var/www/assembly/test 

Breakpoint 1, _start () at test.s:10
10         pushq $2
```

查看組合語言在內存上的位置，目前在`pushq $0x2` 也就是`_start`的 `pushq $2`

```bash
(gdb) x /17i $pc
=> 0x401000 <_start>:    pushq  $0x2
   0x401002 <_start+2>:  callq  0x401028 <add>
   0x401007 <_start+7>:  add    $0x8,%rsp
   0x40100b <_start+11>: mov    %rax,0x402000
   0x401013 <_start+19>: pushq  $0x4
   0x401015 <_start+21>: callq  0x401028 <add>
   0x40101a <_start+26>: add    $0x8,%rsp
   0x40101e <_start+30>: add    %rax,0x402000
   0x401026 <_start+38>: int    $0x80
   0x401028 <add>:       push   %rbp
   0x401029 <add+1>:     mov    %rsp,%rbp
   0x40102c <add+4>:     mov    0x10(%rbp),%rax
   0x401030 <add+8>:     pushq  $0x1
   0x401032 <add+10>:    add    -0x8(%rbp),%rax
   0x401036 <add+14>:    mov    %rbp,%rsp
   0x401039 <add+17>:    pop    %rbp
   0x40103a <add+18>:    retq  
```

查看寄存器，可以看到`rsp`在`0x7fffffffed10`，`rbp`則是0

```bash
(gdb) i r
rax            0x0                 0
rbx            0x0                 0
rcx            0x0                 0
rdx            0x0                 0
rsi            0x0                 0
rdi            0x0                 0
rbp            0x0                 0x0
rsp            0x7fffffffed10      0x7fffffffed10
略......
```

| 寄存器 | 值             |
| ------ | -------------- |
| rsp    | 0x7fffffffed10 |
| rbp    | 0x0            |

查看stack狀況，從上述得知`rsp`在`0x7fffffffed10`，代表stack從這個內存地址開始往下，`0x7fffffffed18`後的值可以忽略不看不在討論範圍，因為stack是向下不是向上

```bash
(gdb) x /32db $rsp
0x7fffffffed10: 1       0       0       0       0       0       0       0
0x7fffffffed18: -39     -18     -1      -1      -1      127     0       0
0x7fffffffed20: 0       0       0       0       0       0       0       0
0x7fffffffed28: -16     -18     -1      -1      -1      127     0       0
```

這裡可以看一下`0x7fffffffed10`之前的值，比如從`0x7fffffffed00`開始看，可以看到`0x7fffffffed00`與`0x7fffffffed08`皆為空值，這兩個內存地址等等都會用到，目前都是空的

```bash
(gdb) x /32db 0x7fffffffed00
0x7fffffffed00: 0       0       0       0       0       0       0       0
0x7fffffffed08: 0       0       0       0       0       0       0       0
0x7fffffffed10: 1       0       0       0       0       0       0       0
0x7fffffffed18: -39     -18     -1      -1      -1      127     0       0
```

目前stack內存結構可以看作以下

| 內存地址       | 值   | 說明                  |                   |
| -------------- | ---- | --------------------- | ----------------- |
| 0x7fffffffed10 | 1    | stack一開始的內存位置 | rsp指向的內存地址 |

執行`pushq $2`，查看相關資料

```bash
(gdb) n
_start () at test.s:11
11         call add
(gdb) x /32db $rsp
0x7fffffffed08: 2       0       0       0       0       0       0       0
0x7fffffffed10: 1       0       0       0       0       0       0       0
0x7fffffffed18: -39     -18     -1      -1      -1      127     0       0
0x7fffffffed20: 0       0       0       0       0       0       0       0
(gdb) i r
rax            0x0                 0
rbx            0x0                 0
rcx            0x0                 0
rdx            0x0                 0
rsi            0x0                 0
rdi            0x0                 0
rbp            0x0                 0x0
rsp            0x7fffffffed08      0x7fffffffed08
```

目前stack內存結構可以看作以下

| 內存地址       | 值   | 說明                  |                   |
| -------------- | ---- | --------------------- | ----------------- |
| 0x7fffffffed10 | 1    | stack一開始的內存位置 |                   |
| 0x7fffffffed08 | 2    | add() int參數         | rsp指向的內存地址 |

| 寄存器 | 值             |
| ------ | -------------- |
| rsp    | 0x7fffffffed08 |
| rbp    | 0x0            |

執行`call add`，進到function 

```bash
(gdb) s
add () at test.s:24
24         pushq %rbp
(gdb) x /32db $rsp
0x7fffffffed00: 7       16      64      0       0       0       0       0
0x7fffffffed08: 2       0       0       0       0       0       0       0
0x7fffffffed10: 1       0       0       0       0       0       0       0
0x7fffffffed18: -39     -18     -1      -1      -1      127     0       0
(gdb) i r
rax            0x0                 0
rbx            0x0                 0
rcx            0x0                 0
rdx            0x0                 0
rsi            0x0                 0
rdi            0x0                 0
rbp            0x0                 0x0
rsp            0x7fffffffed00      0x7fffffffed00
```
`0x7fffffffed00`呼印了golang範例所講到的會push該function下一個要執行的指令地址，`0x401007`就是`add $0x8,%rsp`

```bash
(gdb) x /8x 0x7fffffffed00
0x7fffffffed00: 0x07    0x10    0x40    0x00    0x00    0x00    0x00    0x00
```

| 內存地址       | 值       | 說明                        |                   |
| -------------- | -------- | --------------------------- | ----------------- |
| 0x7fffffffed10 | 1        | stack一開始的內存位置       |                   |
| 0x7fffffffed08 | 2        | add() int參數               |                   |
| 0x7fffffffed00 | 0x401007 | fuction執行完返回的內存地址 | rsp指向的內存地址 |

| 寄存器 | 值             |
| ------ | -------------- |
| rsp    | 0x7fffffffed00 |
| rbp    | 0x0            |

執行`pushq %rbp`，因為要暫時使用`rbp`替代`rsp`做間接尋址，所以先保存`rbp`值等function結束後才能還原`rbp`值

```bash
(gdb) n
25         movq %rsp, %rbp
(gdb) x /32db $rsp
0x7fffffffecf8: 0       0       0       0       0       0       0       0
0x7fffffffed00: 7       16      64      0       0       0       0       0
0x7fffffffed08: 2       0       0       0       0       0       0       0
0x7fffffffed10: 1       0       0       0       0       0       0       0
(gdb) i r
rax            0x0                 0
rbx            0x0                 0
rcx            0x0                 0
rdx            0x0                 0
rsi            0x0                 0
rdi            0x0                 0
rbp            0x0                 0x0
rsp            0x7fffffffecf8      0x7fffffffecf8
```

| 內存地址       | 值       | 說明                        |                   |
| -------------- | -------- | --------------------------- | ----------------- |
| 0x7fffffffed10 | 1        | stack一開始的內存位置       |                   |
| 0x7fffffffed08 | 2        | add() int參數               |                   |
| 0x7fffffffed00 | 0x401007 | fuction執行完返回的內存地址 |                   |
| 0x7fffffffecf8 | 0        | rbp原本的值                 | rsp指向的內存地址 |

| 寄存器 | 值             |
| ------ | -------------- |
| rsp    | 0x7fffffffecf8 |
| rbp    | 0x0            |

執行`movq %rsp, %rbp`，`rbp`保存`rsp`值以利做間接尋址

```bash
(gdb) n
26         movq 16(%rbp), %rax
(gdb) i r
rax            0x0                 0
rbx            0x0                 0
rcx            0x0                 0
rdx            0x0                 0
rsi            0x0                 0
rdi            0x0                 0
rbp            0x7fffffffecf8      0x7fffffffecf8
rsp            0x7fffffffecf8      0x7fffffffecf8
```

| 寄存器 | 值             |
| ------ | -------------- |
| rsp    | 0x7fffffffecf8 |
| rbp    | 0x7fffffffecf8 |

執行`movq 16(%rbp), %rax`，將fuction參數放到`rax`

```bash
(gdb) n
27         pushq $1
(gdb) x /32db $rsp
0x7fffffffecf8: 0       0       0       0       0       0       0       0
0x7fffffffed00: 7       16      64      0       0       0       0       0
0x7fffffffed08: 2       0       0       0       0       0       0       0
0x7fffffffed10: 1       0       0       0       0       0       0       0
(gdb) i r
rax            0x2                 2
rbx            0x0                 0
rcx            0x0                 0
rdx            0x0                 0
rsi            0x0                 0
rdi            0x0                 0
rbp            0x7fffffffecf8      0x7fffffffecf8
rsp            0x7fffffffecf8      0x7fffffffecf8
```

| 內存地址       | 值       | 說明                        | 間接尋址 |                   |
| -------------- | -------- | --------------------------- | -------- | ----------------- |
| 0x7fffffffed10 | 1        | stack一開始的內存位置       | 24(%rbp) |                   |
| 0x7fffffffed08 | 2        | add() int參數               | 16(%rbp) |                   |
| 0x7fffffffed00 | 0x401007 | fuction執行完返回的內存地址 | 8(%rbp)  |                   |
| 0x7fffffffecf8 | 0        | rbp原本的值                 | (%rbp)   | rsp指向的內存地址 |

| 寄存器 | 值             |
| ------ | -------------- |
| rsp    | 0x7fffffffecf8 |
| rbp    | 0x7fffffffecf8 |
| rax    | 0x2            |

執行`pushq $1`，將1值放入stack當區域變數

```bash
(gdb) n
28         addq -8(%rbp), %rax
(gdb) x /32db $rsp
0x7fffffffecf0: 1       0       0       0       0       0       0       0
0x7fffffffecf8: 0       0       0       0       0       0       0       0
0x7fffffffed00: 7       16      64      0       0       0       0       0
0x7fffffffed08: 2       0       0       0       0       0       0       0
(gdb) i r
rax            0x2                 2
rbx            0x0                 0
rcx            0x0                 0
rdx            0x0                 0
rsi            0x0                 0
rdi            0x0                 0
rbp            0x7fffffffecf8      0x7fffffffecf8
rsp            0x7fffffffecf0      0x7fffffffecf0
```

| 內存地址       | 值       | 說明                        | 間接尋址 |                   |
| -------------- | -------- | --------------------------- | -------- | ----------------- |
| 0x7fffffffed10 | 1        | stack一開始的內存位置       | 24(%rbp) |                   |
| 0x7fffffffed08 | 2        | add() int參數               | 16(%rbp) |                   |
| 0x7fffffffed00 | 0x401007 | fuction執行完返回的內存地址 | 8(%rbp)  |                   |
| 0x7fffffffecf8 | 0        | rbp原本的值                 | (%rbp)   |                   |
| 0x7fffffffecf0 | １       | 區域變數                    | -8(%rbp) | rsp指向的內存地址 |

這邊可以看到`rsp`與`rbp`已經不一樣了，`rbp`不會因為對stack做什麼動作而異動，所以能使用間接尋址來操作stack

| 寄存器 | 值             |
| ------ | -------------- |
| rsp    | 0x7fffffffecf0 |
| rbp    | 0x7fffffffecf8 |
| rax    | 0x2            |

執行`addq -8(%rbp), %rax`，將區域變數加上`rax`(fuction 參數值) = 1 + 2 = 3

```bash
(gdb) n
30         movq %rbp, %rsp
(gdb) x /32db $rsp
0x7fffffffecf0: 1       0       0       0       0       0       0       0
0x7fffffffecf8: 0       0       0       0       0       0       0       0
0x7fffffffed00: 7       16      64      0       0       0       0       0
0x7fffffffed08: 2       0       0       0       0       0       0       0
(gdb) i r
rax            0x3                 3
rbx            0x0                 0
rcx            0x0                 0
rdx            0x0                 0
rsi            0x0                 0
rdi            0x0                 0
rbp            0x7fffffffecf8      0x7fffffffecf8
rsp            0x7fffffffecf0      0x7fffffffecf0
```

| 寄存器 | 值             |
| ------ | -------------- |
| rsp    | 0x7fffffffecf0 |
| rbp    | 0x7fffffffecf8 |
| rax    | 0x3            |

執行`movq %rbp, %rsp`，fuction要結束了需要把`rsp`還原到最初狀態

```bash
(gdb) n
31         popq %rbp
(gdb) x /32db $rsp
0x7fffffffecf8: 0       0       0       0       0       0       0       0
0x7fffffffed00: 7       16      64      0       0       0       0       0
0x7fffffffed08: 2       0       0       0       0       0       0       0
0x7fffffffed10: 1       0       0       0       0       0       0       0
(gdb) x /40db 0x7fffffffecf0
0x7fffffffecf0: 1       0       0       0       0       0       0       0
0x7fffffffecf8: 0       0       0       0       0       0       0       0
0x7fffffffed00: 7       16      64      0       0       0       0       0
0x7fffffffed08: 2       0       0       0       0       0       0       0
0x7fffffffed10: 1       0       0       0       0       0       0       0
(gdb) i r
rax            0x3                 3
rbx            0x0                 0
rcx            0x0                 0
rdx            0x0                 0
rsi            0x0                 0
rdi            0x0                 0
rbp            0x7fffffffecf8      0x7fffffffecf8
rsp            0x7fffffffecf8      0x7fffffffecf8
```

| 內存地址       | 值       | 說明                        | 間接尋址 |                   |
| -------------- | -------- | --------------------------- | -------- | ----------------- |
| 0x7fffffffed10 | 1        | stack一開始的內存位置       | 24(%rbp) |                   |
| 0x7fffffffed08 | 2        | add() int參數               | 16(%rbp) |                   |
| 0x7fffffffed00 | 0x401007 | fuction執行完返回的內存地址 | 8(%rbp)  |                   |
| 0x7fffffffecf8 | 0        | rbp原本的值                 | (%rbp)   | rsp指向的內存地址 |
| 0x7fffffffecf0 | １       | 區域變數                    | -8(%rbp) |                   |

| 寄存器 | 值             |
| ------ | -------------- |
| rsp    | 0x7fffffffecf8 |
| rbp    | 0x7fffffffecf8 |
| rax    | 0x3            |

執行`popq %rbp`，fuction要結束了需要把`rbp`還原到最初狀態

```bash
(gdb) n
add () at test.s:32
32         ret
(gdb) x /32db $rsp
0x7fffffffed00: 7       16      64      0       0       0       0       0
0x7fffffffed08: 2       0       0       0       0       0       0       0
0x7fffffffed10: 1       0       0       0       0       0       0       0
0x7fffffffed18: -39     -18     -1      -1      -1      127     0       0
(gdb) x /40db 0x7fffffffecf0
0x7fffffffecf0: 1       0       0       0       0       0       0       0
0x7fffffffecf8: 0       0       0       0       0       0       0       0
0x7fffffffed00: 7       16      64      0       0       0       0       0
0x7fffffffed08: 2       0       0       0       0       0       0       0
0x7fffffffed10: 1       0       0       0       0       0       0       0
(gdb) i r
rax            0x3                 3
rbx            0x0                 0
rcx            0x0                 0
rdx            0x0                 0
rsi            0x0                 0
rdi            0x0                 0
rbp            0x0                 0x0
rsp            0x7fffffffed00      0x7fffffffed00
rip            0x40103a            0x40103a <add+18>
```

| 內存地址       | 值       | 說明                        | 間接尋址 |                   |
| -------------- | -------- | --------------------------- | -------- | ----------------- |
| 0x7fffffffed10 | 1        | stack一開始的內存位置       | 24(%rbp) |                   |
| 0x7fffffffed08 | 2        | add() int參數               | 16(%rbp) |                   |
| 0x7fffffffed00 | 0x401007 | fuction執行完返回的內存地址 | 8(%rbp)  | rsp指向的內存地址 |
| 0x7fffffffecf8 | 0        | rbp原本的值                 | (%rbp)   |                   |
| 0x7fffffffecf0 | １       | 區域變數                    | -8(%rbp) |                   |

| 寄存器 | 值             |
| ------ | -------------- |
| rsp    | 0x7fffffffecf8 |
| rbp    | 0x0            |
| rax    | 0x3            |
| rip    | 0x40103a       |

執行`ret`，相當於`popq ％rip`

```bash
(gdb) n
_start () at test.s:12
12         addq $8, %rsp
(gdb) x /32db $rsp
0x7fffffffed08: 2       0       0       0       0       0       0       0
0x7fffffffed10: 1       0       0       0       0       0       0       0
0x7fffffffed18: -39     -18     -1      -1      -1      127     0       0
0x7fffffffed20: 0       0       0       0       0       0       0       0
(gdb) x /40db 0x7fffffffecf0
0x7fffffffecf0: 1       0       0       0       0       0       0       0
0x7fffffffecf8: 0       0       0       0       0       0       0       0
0x7fffffffed00: 7       16      64      0       0       0       0       0
0x7fffffffed08: 2       0       0       0       0       0       0       0
0x7fffffffed10: 1       0       0       0       0       0       0       0
(gdb) i r
rax            0x3                 3
rbx            0x0                 0
rcx            0x0                 0
rdx            0x0                 0
rsi            0x0                 0
rdi            0x0                 0
rbp            0x0                 0x0
rsp            0x7fffffffed08      0x7fffffffed08
rip            0x401007            0x401007 <_start+7>
```

| 內存地址       | 值       | 說明                        | 間接尋址 |                   |
| -------------- | -------- | --------------------------- | -------- | ----------------- |
| 0x7fffffffed10 | 1        | stack一開始的內存位置       | 24(%rbp) |                   |
| 0x7fffffffed08 | 2        | add() int參數               | 16(%rbp) | rsp指向的內存地址 |
| 0x7fffffffed00 | 0x401007 | fuction執行完返回的內存地址 | 8(%rbp)  |                   |
| 0x7fffffffecf8 | 0        | rbp原本的值                 | (%rbp)   |                   |
| 0x7fffffffecf0 | １       | 區域變數                    | -8(%rbp) |                   |

這邊可以看到stack內`0x7fffffffed00`值是`0x401007`，而執行完`ret`後`rip`值也是`0x401007`，所以是利用此方法改變`rip`再跳回原來調用`add()`後的下一條指令

| 寄存器 | 值             |
| ------ | -------------- |
| rsp    | 0x7fffffffed08 |
| rbp    | 0x0            |
| rax    | 0x3            |
| rip    | 0x401007       |

執行`addq $8, %rsp`，可以從上述表中得知`rsp`目前還未指向stack最一開始位置，原因是之前有把參數帶入stack，所以這邊還需把`rsp`往上移還原stack

```bash
(gdb) n
_start () at test.s:13
13         movq %rax, result
(gdb) x /32db $rsp
0x7fffffffed10: 1       0       0       0       0       0       0       0
0x7fffffffed18: -39     -18     -1      -1      -1      127     0       0
0x7fffffffed20: 0       0       0       0       0       0       0       0
0x7fffffffed28: -16     -18     -1      -1      -1      127     0       0
(gdb) x /40db 0x7fffffffecf0
0x7fffffffecf0: 1       0       0       0       0       0       0       0
0x7fffffffecf8: 0       0       0       0       0       0       0       0
0x7fffffffed00: 7       16      64      0       0       0       0       0
0x7fffffffed08: 2       0       0       0       0       0       0       0
0x7fffffffed10: 1       0       0       0       0       0       0       0
(gdb) i r
rax            0x3                 3
rbx            0x0                 0
rcx            0x0                 0
rdx            0x0                 0
rsi            0x0                 0
rdi            0x0                 0
rbp            0x0                 0x0
rsp            0x7fffffffed10      0x7fffffffed10
```

| 內存地址       | 值       | 說明                        | 間接尋址 |                   |
| -------------- | -------- | --------------------------- | -------- | ----------------- |
| 0x7fffffffed10 | 1        | stack一開始的內存位置       | 24(%rbp) | rsp指向的內存地址 |
| 0x7fffffffed08 | 2        | add() int參數               | 16(%rbp) |                   |
| 0x7fffffffed00 | 0x401007 | fuction執行完返回的內存地址 | 8(%rbp)  |                   |
| 0x7fffffffecf8 | 0        | rbp原本的值                 | (%rbp)   |                   |
| 0x7fffffffecf0 | １       | 區域變數                    | -8(%rbp) |                   |

| 寄存器 | 值             |
| ------ | -------------- |
| rsp    | 0x7fffffffed10 |
| rbp    | 0x0            |
| rax    | 0x3            |
| rip    | 0x401007       |

到這裡就知道fuction與stack之間的關係是如何互動，fuction結束後stack內原有的值並未清空只是更改`rsp`，剩餘後面的指令就不多介紹都一樣只是`add()`參數不太一樣，後面指令只要再繼續執行就可以發現原有stack值都被覆蓋掉，這樣舊的stack值就不會影響現在function的操作




## 參考

1. 必看

   [汇编语言(第3版)-王爽](#https://book.douban.com/subject/25726019/) 基礎入門，使用intel寫法較簡單，我個人覺得當學習概念就好不用全懂，主要是有一個帶入感讓你覺得CPU是怎樣了解程式語言

   [汇编语言程序设计](https://book.douban.com/subject/1446250/) AT&T寫法入門，跟intel相比是另一派風格也較難一點，linux主要都是AT&T 寫法，基本上這本必看，本書例子是32位元但是現代大部分都是64位元所以給出來的範例會有少許錯誤，請自行google解決 ex: push在64與32位元上的[差異](https://stackoverflow.com/questions/5485468/x86-assembly-pushl-popl-dont-work-with-error-suffix-or-operands-invalid)，部分進階內容看不懂可以調過，畢竟只是入門學習非實際做開發

   

   閱讀順序是`汇编语言(第3版)` -> `汇编语言程序设计`

2. gdb 操作

   https://wizardforcel.gitbooks.io/100-gdb-tips/print-registers.html

   https://sourceware.org/gdb/current/onlinedocs/gdb/

3. 操作

   http://ouonline.net/tag/att-asm