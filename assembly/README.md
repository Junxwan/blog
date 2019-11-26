# 組合語言(汇编语言)

- [閱讀前提](#閱讀前提])
- [進位](#進位)
- [寄存器](#寄存器)
  - [通用寄存器](#通用寄存器)
  - [段寄存器](#段寄存器)
  - [標記寄存器](#標記寄存器)
- [指令](#指令)
- [AT&T](#AT&T)
- [參考](#參考)

## 閱讀前提

具備以下條件才能閱讀

1. 從事golang開發

   會這樣認為是因為golang一些重點如內存如何管理、GMP等通常都在runtime源碼內，可能單純看runtime源碼就可以懂，但你可能會一支半解，因為runtime一些method或是變數是用組合語言寫的，你可能問或許理解這個可能不重要，如果你問我非要這樣做的原因是什麼，我想應該是這個[答案1](https://www.zhihu.com/question/19712941/answer/12878634) [答案2](https://www.zhihu.com/question/23088538/answer/23717201)，當然如果是php、java開發者想往比較底層方向走 or 想了解源碼也是可以，個人是php -> golang．

2. 不講太細節

   個人是為了能夠看懂golang runtime源碼而寫的紀錄並非工作上真的要寫組合語言，單純只是為了看懂、理解，所以你真的想了解關於組合語言跟CPU之間的細節，可能不適合你，本篇單純入門而已．

   

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

#### 

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



## 指令

組合語言的動作表示，以下都是以AT&T格式做說明，下方推薦的汇编语言(第3版)-王爽書則是Intel格式，[差異](https://blog.csdn.net/kennyrose/article/details/7575952)

指令後綴詞代表不同位元操作

| 後綴詞 | 位元 |
| ------ | ---- |
| b      | 8    |
| w      | 16   |
| l      | 32   |
| q      | 64   |

Ex: movb（8位）、movw（16位）、movl（32位）、movq（64位）



### mod



### add



### sub



### div



### push



### pop



### jmp



### inc



### dec



### loop



### and



### or



## AT&T

### .section



### .data

放置已有初始化的變數



### .bss

放置未有初始化的變數



### .text



### .globl



### .ascii 

變數型別為字串



### .short 

### .int .long

變數型別為int or long



### .byte

變數型別為byte



### .float

變數型別為float



### .double 

變數型別為double



### .fill

變數型別為const



### .comm



### _start

### 



## 參考

1. 必看

   [汇编语言(第3版)-王爽](#https://book.douban.com/subject/25726019/) 基礎入門，使用intel寫法，我個人覺得學習概念就好

   [汇编语言程序设计](https://book.douban.com/subject/1446250/) AT&T寫法入門，跟intel相比是另一派風格，linux主要都是AT&T 寫法

2. gdb 操作

   https://wizardforcel.gitbooks.io/100-gdb-tips/print-registers.html

   https://sourceware.org/gdb/current/onlinedocs/gdb/

3. 操作

   http://ouonline.net/tag/att-asm