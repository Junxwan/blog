以下是準備寫成文章的主題，先做個紀錄與reference

寫bolg原因主要是紀錄加學習，有些主題知道原理但別人問不一定能講出細解，下列reference是已經看過的文章覺得不錯可以重覆閱讀，個人習慣重覆看兩次，第一次看了解大致方向，第二次動手研究寫成一篇自己看的懂的文章以利重覆學習，書籍也是如此．

文章完成後會先公布在此再發佈至https://junxwan.github.io/

長期研究主題主要是後端

1. golang源碼(特別是runtime)

2. Kafka源碼(broker部分)

3. redis,nginx,etcd等開源專案源碼

4. 資料結構

5. C++相關知識

6. 股票理財相關知識(基本面,技術面,籌碼面)

   

## TODO主題

1. 进程 线程 协程 Process v.s. Thread v.s. Coroutine

   https://juejin.im/post/5b0014b7518825426e023666

   https://www.imooc.com/article/31751

   https://zhuanfou.com/ask/78120839_783

   https://www.jianshu.com/p/f11724034d50

   https://zhuanlan.zhihu.com/p/81390586

   https://zhuanlan.zhihu.com/p/37754274

2. 用户线程(user thread)与核心线程(kernel thread)

3. Golang CSP

4. Green Threads

5. work stealing

6. 并发与并行

7. HTTP VS HTTPS

8. TCP v.s. UDP

9. 死鎖（Deadlock)

10. golang GMP

   https://colobu.com/2017/05/04/go-scheduler/

   https://juejin.im/post/5b7678f451882533110e8948

   https://juejin.im/post/5d9a9c12e51d45781420fb7e

   https://studygolang.com/articles/11627

   https://tonybai.com/2017/06/23/an-intro-about-goroutine-scheduler/

   http://skoo.me/go/2013/11/29/golang-schedule?utm_campaign=studygolang.com&utm_medium=studygolang.com&utm_source=studygolang.com

   https://www.zhihu.com/question/20862617/answer/131341519

   https://www.zhihu.com/question/20862617/answer/27964865

   https://blog.csdn.net/heiyeshuwu/article/details/51178268

11. golang defer

    https://juejin.im/post/5c964437e51d456d6d4fae46

12. goalng Channel

13. 調度策略（Scheduler）

    https://wudaijun.com/2018/11/scheduler-blabla/

14. RadixTree

15. Segregated Free List

16. tcmalloc

    https://zhuanlan.zhihu.com/p/59437135

    https://wallenwang.com/2018/11/tcmalloc/

    https://zhuanlan.zhihu.com/p/29216091

    https://dirtysalt.github.io/html/tcmalloc.html

    http://gao-xiao-long.github.io/2017/11/25/tcmalloc/

17. jemalloc

18. Mark And Sweep

    https://liujiacai.net/blog/2018/07/08/mark-sweep/

    https://juejin.im/post/5c8525666fb9a049ea39c3e6

    https://www.memorymanagement.org/index.html

    https://www.jianshu.com/p/bfc3c65c05d1

    https://github.com/golang/proposal/blob/master/design/17503-eliminate-rescan.md

    http://lessisbetter.site/2019/10/20/go-gc-1-history-and-priciple/

19. Memory Leaking

    https://go101.org/article/memory-leaking.html

    https://colobu.com/2019/08/28/go-memory-leak-i-dont-think-so/

    [https://www.do1618.com/archives/1328/go-%E5%86%85%E5%AD%98%E9%80%83%E9%80%B8%E8%AF%A6%E7%BB%86%E5%88%86%E6%9E%90/](https://www.do1618.com/archives/1328/go-内存逃逸详细分析/)

20. Prometheus 和 Grafana

    [https://site-optimize-note.tk/%E4%BC%BA%E6%9C%8D%E5%99%A8%E6%95%88%E8%83%BD%E7%9B%A3%E6%8E%A7prometheus-%E8%88%87-grafana-%E5%AE%89%E8%A3%9D%E6%95%99%E5%AD%B8/](https://site-optimize-note.tk/伺服器效能監控prometheus-與-grafana-安裝教學/)

    https://prometheus.io/docs/prometheus/latest/installation/

21. 組合語言(汇编语言)

    https://davidwong.fr/goasm/

    https://github.com/teh-cmc/go-internals/blob/master/chapter1_assembly_primer/README.md

    https://golang.org/doc/asm

    http://blog.rootk.com/post/golang-asm.html

    https://syslog.ravelin.com/anatomy-of-a-function-call-in-go-f6fc81b80ecc

22. 指標 (pointer) 和 參考 (reference)與傳遞

    https://www.zhihu.com/question/31203609/answer/576030121

    https://www.zhihu.com/question/20628016/answer/576031541

    https://www.zhihu.com/question/20628016/answer/29032988

    https://dotblogs.com.tw/brian/2012/10/18/77588

    http://wp.mlab.tw/?p=176

    https://sanyuesha.com/2017/08/10/go-no-reference-type/

    https://dave.cheney.net/2017/04/29/there-is-no-pass-by-reference-in-go

23. GC

    https://www.zhihu.com/question/51244545/answer/126055789

24. Memory Heap VS Stack

    https://www.zhihu.com/question/281940376

    https://www.zhihu.com/question/36103513

    https://www.zhihu.com/question/37413173

    https://www.zhihu.com/question/34341582

    https://www.zhihu.com/question/318874857

    https://www.zhihu.com/question/29005517

    https://www.zhihu.com/question/29833675/answer/82661572

    https://www.zhihu.com/question/34499262/answer/59415153

    https://www.zhihu.com/question/41691246/answer/98069640

    https://zhuanlan.zhihu.com/p/78478567

    http://wp.mlab.tw/?p=312

    https://nwpie.blogspot.com/2017/05/5-stack-heap.html

    http://guang.logdown.com/posts/236293-c-programs-memory-usage

    https://www.geeksforgeeks.org/memory-layout-of-c-program/

    https://zhuanlan.zhihu.com/p/28409657

25. Memory Allocator

    https://zhuanlan.zhihu.com/p/51056407

    https://zhuanlan.zhihu.com/p/51855842

    https://juejin.im/post/5c888a79e51d456ed11955a8#heading-5

    http://lessisbetter.site/2019/07/06/go-memory-allocation/

    http://dmitrysoshnikov.com/compilers/writing-a-memory-allocator/

26. 內存

    https://yq.aliyun.com/articles/652551

    https://studygolang.com/articles/11978

    https://zhuanlan.zhihu.com/p/76802887

    https://www.infoq.cn/article/IEhRLwmmIM7-11RYaLHR

    https://juejin.im/post/5c888a79e51d456ed11955a8

    https://www.jianshu.com/p/47691d870756

    https://zhuanlan.zhihu.com/p/59125443

    https://zhuanlan.zhihu.com/p/27807169

    https://medium.com/@kai.chihkaiyu/golang-memory-management-based-on-1-12-5-51fcc97f3c92

    https://blog.yiz96.com/golang-mm-gc/

    https://mp.weixin.qq.com/s/3gGbJaeuvx4klqcv34hmmw

    https://www.jianshu.com/p/1ffde2de153f

    https://www.twblogs.net/a/5cb97871bd9eee0eff45da7e

    [https://xenojoshua.com/2019/03/golang-memory/#1-%E5%89%8D%E8%A8%80](https://xenojoshua.com/2019/03/golang-memory/#1-前言)

    https://yq.aliyun.com/blog/652551

    https://mp.weixin.qq.com/s/3gGbJaeuvx4klqcv34hmmw

27. go gctrace

    https://www.jishuwen.com/d/2KE4

    https://medium.com/square-corner-blog/always-be-closing-3d5fda0e00da

    [http://cbsheng.github.io/posts/godebug%E4%B9%8Bgctrace%E8%A7%A3%E6%9E%90/](http://cbsheng.github.io/posts/godebug之gctrace解析/)

    https://golang.org/pkg/runtime/#hdr-Environment_Variables

    https://www.jishuwen.com/d/2KE4

    https://github.com/golang/go/issues/32284

    https://github.com/golang/go/issues/14521

    https://stackoverflow.com/questions/37382600/cannot-free-memory-once-occupied-by-bytes-buffer/37383604#37383604

28. cpu密集型 io密集型

29. 虚拟内存

30. go源碼閱讀參考

    https://juejin.im/post/5d661014f265da03f04cdddc

    https://github.com/tiancaiamao/go-internals

    https://github.com/zboya/golang_runtime_reading

    https://jingwei.link/2018/07/01/runtime-of-goroutine-creation.html

    https://www.cnblogs.com/zkweb/p/7777525.html

31. Plan9

32. Linux huge page

    https://zhuanlan.zhihu.com/p/34659353

33. Transparent Huge pages

    http://blog.itpub.net/26736162/viewspace-2214374/

    http://www.zendei.com/article/37419.html

