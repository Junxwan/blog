# 組合語言(汇编语言)

- [閱讀前提](#閱讀前提])
- [進位](#進位)
- [寄存器](#寄存器)
- [指令](#指令)
- [參考](#參考)

## 閱讀前提

具備以下條件才能閱讀

1. 從事golang開發

   會這樣認為是因為golang一些重點如內存如何管理、GMP等通常都在runtime源碼內，可能單純看runtime源碼就可以懂，但你可能會一支半解，因為runtime一些method或是變數是用組合語言寫的，你可能問或許理解這個可能不重要，如果你問我非要這樣做的原因是什麼，我想應該是這個[答案1](https://www.zhihu.com/question/19712941/answer/12878634) [答案2](https://www.zhihu.com/question/23088538/answer/23717201)，當然如果是php、java開發者想往比較底層方向走 or 想了解源碼也是可以，個人是php -> golang．

2. 不講太細節

   個人是為了能夠看懂golang runtime源碼而寫的紀錄並非工作上真的要寫組合語言，單純只是為了看懂、理解，所以你真的想了解關於組合語言跟CPU之間的細節，可能不適合你，本篇單純入門而已．

   

## 寄存器

### cs



### ip



### ss



### sp



## 指令

組合語言的動作表示，以下都是以AT&T格式做說明，下方推薦的汇编语言(第3版)-王爽書則是Intel格式，[差異](https://blog.csdn.net/kennyrose/article/details/7575952)

### mod



### add



### sub



### push



### pop



## 參考

1. 必看

   [汇编语言(第3版)-王爽](#https://book.douban.com/subject/25726019/)

2. gdb 操作

   https://wizardforcel.gitbooks.io/100-gdb-tips/print-registers.html

   https://sourceware.org/gdb/current/onlinedocs/gdb/