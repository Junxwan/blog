# Process v.s. Thread v.s. Coroutine

- [名詞](#名詞])
- [前置知識](#前置知識)
- [Process](#Process)
- [Thread](#Thread)
- [Coroutine](#Coroutine)
- [用戶級線程模型](#用戶級線程模型)
- [內核級線程模型](#內核級線程模型)
- [兩級線程模型](#兩級線程模型)
- [總結](#總結)
- [參考資料](#參考資料)
- [疑問](#疑問)



## 名詞

網路上或書上常常會看到不同地區的不同翻譯，以下是對照

| 英        | 台灣   | 中國 |
| --------- | ------ | ---- |
| Process   | 程序   | 进程 |
| Thread    | 執行緒 | 线程 |
| Coroutine | 協程   | 协程 |



## 前置知識

閱讀此篇希望具備的知識

1. 並發與並行差異？

2. 知道什麼是Thread，有寫過類似程式？

3. CPU switch context 意義 ？

4. 什麼是lock  ？

5. 什麼是stack ？

6. CPU寄存器 ？

   

## Process

每個應用程式啟動後都可以看成是一種Process，各自都分配到資源且Process之間都是隔離的，作業系統內可以同時啟動多個Process，一般狀況下是不會出現A Process 讀取到 B Process的資料，所以可以各自獨立運行，所以Process是資源分配最基本的單位．

早期同一時間只能有一個Process運作，但這樣太沒效率，所以出現了Process多並發，但是Process的建立、銷毀、切換時後CPU運作的時間太長，尤其是Process並發需要不停地切換Process，看似並發增加效能但一大部分時間都耗費在CPU調度，所以後來衍生出Thread來解決這塊．



## Thread

存在於Process內，主要是解決Process併發時switch context的成本，是CPU調度的最基本單位，一個Process內可以有多個thread且各自共享內存，相反的Process間都是獨立不共享．但由於內存共享也引來需要一些lock機制來確保資料的正確性

thread之間是有依賴關係，如在main thread 啟動sub thread，當main thread掛掉時sub thread也會一起掛掉，相反的sub thread掛掉則main thread不受影響，Process是不會有這個問題，因為都是各自獨立

當啟動一個應用程式建立Process時就一定會有一個thread，因為作業系統把應用程式加載起來時一定需要有某機制去幫助初始化或是做一些其他事情，而負責這件事的就是一個thread，該thread也可以認為是main thread

thread建立與銷毀雖然都比Process消耗的資源來的少很多，但還是有一些缺點需要克服

1. 過多的thread工作時會導致CPU調度的時間增長，列如只有4 core但卻開了100個thread，同時間只有四個thread能夠工作但因為還有其他thread排隊等待CPU執行，所以CPU並需要很頻繁的切換到各個thread，每個thread在CPU上執行都需要把資料複製到CPU寄存器上，當需要切換時就必須在保存回thread後再把要執行的thread資料複製一份至CPU寄存器上等動作，這一連貫的動作也會造成CPU調度的時間增長

2. 每個thread stack都是固定的(看作業系統決定)，這導致可能只使用幾KB但每個thread都先分配幾MB就造成內存浪費或著是分配的內存不夠造成溢出

總合以上觀點，希望能夠減輕CPU調度開銷的時間讓CPU利用率上升且stack可以動態分配解決溢出或是浪費，於是就有了coroutine



## Coroutine

coroutine是一種更輕量的thread，原本thread調度是由CPU控制而coroutine是由用戶控制，用戶自行決定何時要調度不需要交由CPU調度減少了調度時間更能利用CPU，切換非常輕量，coroutine本身stack內存空間由用戶控制，減少內存的浪費提高內存利用率．由於是依附在thread上，如果thread掛掉或是阻塞一樣會無法使用



## 用戶級線程模型

TODO

## 內核級線程模型

TODO

## 兩級線程模型

TODO



## 結論

|           | 內存共享 | 調度開銷 | 併發效率 | 內存消耗 | 調度方式 | stack內存分配 |
| --------- | -------- | -------- | -------- | -------- | -------- | ------------- |
| Process   | NO       | 大       | 低       | 大       | 系統     | 固定          |
| Thread    | YES      | 中       | 中       | 中       | 系統     | 固定          |
| Coroutine | YES      | 小       | 高       | 小       | 用戶     | 動態          |

1. Process與Thread是一對多關係
2. Thread跟Coroutine是一對多關係
3. Process是資源分配最基本單元
4. Thread是CPU調度最基本單元
5. OS無法感知到Coroutine
6. Coroutine能更好利用CPU，減少調度開銷
7. Coroutine能更好利用內存，可以動態分配stack內存



## 參考資料

線程模型

https://taohuawu.club/high-performance-implementation-of-goroutine-pool

https://zhuanlan.zhihu.com/p/81390586



Process v.s. Thread v.s. Coroutine

http://lessisbetter.site/2019/03/10/golang-scheduler-1-history/



Process與Thread switch context時間

https://zhuanlan.zhihu.com/p/79772089

https://zhuanlan.zhihu.com/p/80037638



Thread與Coroutine switch context

https://www.zhihu.com/question/308641794/answer/572499202

https://www.zhihu.com/question/307787570/answer/565079481



## 疑問

1. Process與thread switch context CPU指令步驟差異有些？
2. 建立一個Process與thread作業系統分配的資源有哪些？