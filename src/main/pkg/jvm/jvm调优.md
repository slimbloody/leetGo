### JVM Tuning ###

有两个堆吧, 一个java堆一个native堆, 他的意思是两个加起来小于70%么?

甚至出现内存泄漏（每次垃圾收集使用的时间越来越长，垃圾收集频率越来越高，每次垃圾收集清理掉的垃圾数据越来越少）

②堆栈错误信息：当系统出现异常后，可以根据堆栈信息初步定位问题所在，比如根据可以判断是堆内存溢出；根据“java.lang.StackOverflowError”可以判断是栈溢出；根据“java.lang.OutOfMemoryError: PermGen space”可以判断是方法区溢出等。

可以在合适的场景（如实现缓存）采用软引用、弱引用，比如用软引用来为ObjectA分配实例：SoftReference objectA=new SoftReference(); 在发生内存溢出前，会将objectA列入回收范围进行二次回收，如果这次回收还没有足够内存，才会抛出内存溢出的异常。
避免产生死循环，产生死循环后，循环体内可能重复产生大量实例，导致内存空间被迅速占满。
尽量避免长时间等待外部资源（数据库、网络、设备资源等）的情况，缩小对象的生命周期，避免进入老年代，如果不能及时返回结果可以适当采用异步处理的方式等。

#### 预准备 ####
1. <<深入理解java虚拟机>> 里面讲过, 大部分young区的代码都是朝生夕灭的, 能GC进入到老年代的大部分都是因为代码层面的问题导致的
```
之前activity-support full-gc 特别频繁, 半小时一次, 
```
2. 在JAVA中, JVM内存指的是堆内存. 机器内存中, 不属于堆内存的部分即为堆外内存.
3. 

#### 为什么需要调优 ####
1. jvm在垃圾回收时有STW现象, 所有工作线程都停止了,  `单次`回收时长 和 `单次`回收效率 之间要平衡好
2. Full GC 次数频繁
3. 应用出现 OutOfMemory 等内存异常
4. 应用中有使用本地缓存且占用大量内存空间


#### 调优工具参数 ####
XX:+PrintFlagsInitial 查看jvm参数初始的默认值
XX:+PrintFlagsFinal 主要是查看jvm参数修改更新后的值

#### 调优点 ####
内存占用: 程序正常运行需要的内存大小
延迟: 由于垃圾收集而引起的程序停顿时间
吞吐量: 用户程序运行时间占用户程序和垃圾收集占用总时间的比值

吞吐量, 延迟, 内存占用三者类似CAP, 能选择其中两个进行调优

其实调优的本质目的是保证程序的稳定性, 附带的好处是可以在硬件环境不变的情况下,  获取更大的吞吐量或者响应速度. 例如一些大数据计算引擎相关的项目, 吞吐量极高, 会产生大量的短生命周期对象和大量的长生命周期对象, 或者极高并发量导致内存溢出等情况, 才会适用JVM调优

1. 吞吐量调优
尽可能避免或者很少发生FullGC或者Stop-The-World压缩式垃圾收集（CMS），因为这两种方式都会造成应用程序吞吐降低。尽量在MinorGC 阶段回收更多的对象，避免对象提升过快到老年代。

#### 日志参数 ####
-XX:+PrintGCDetails 和 -Xloggc:/data/jvm/gc.log 可以在程序运行时把gc的详细过程记录下来，或者直接配置"-verbose:gc"参数把gc日志打印到控制台，通过记录的gc日志可以分析每块内存区域gc的频率、时间等

#### 常见调优参数 ####
-Xms(初始内存)和-Xmx(最大内存)的值设置成相等, 以避免每次垃圾回收完成后JVM重新分配内存, 堆大小默认为-Xms指定的大小, 默认空闲堆内存小于40%时, JVM会扩大堆到-Xmx指定的大小; 空闲堆内存大于70%, JVM会减小堆到-Xms指定的大小. 如果在Full GC后满足不了内存需求会动态调整，这个阶段比较耗费资源.
-Xmn: 年轻代大小: eden + 2 survivor space, Sun官方推荐配置为整个堆的3/8 

1. -XX:MaxTenuringThreshold=xx 该参数只有在串行GC时才有效.
对象经过垃圾回收进入老年代的最大年龄, 如果设置为0的话，则年轻代对象不经过Survivor区，直接进入年老代.

该参数在实战中基本不用. jvm本身有动态年龄晋升机制, jvm会对每个年龄段的对象大小分别累加, 每次进行youngGC时会把每个年龄段的对象所占的空间从小到大累加起来, 累加途中, 若发现累加的值超过了young区的一半, 就会把这个年龄和MaxTenuringThreshold的值做对比, 取两者中更小的那个作为youngGC的年龄参数.

2. 
-XX:+UseBiasedLocking	锁机制的性能改善
如果说轻量级锁是在无竞争的情况下使用CAS操作消除同步使用的互斥量, 那偏向锁就是在无竞争的情况下将整个同步都消除掉, 连CAS都不操作了. 偏向的意思就是这个对象的锁会偏向于第一个获取到它的线程, 如果再接下来的过程中, 该锁没有被其他线程获取, 则持有偏向锁的线程将永远不需要同步.

-Xss: 每个线程的堆栈大小,根据应用的线程所需内存大小进行调整.在相同物理内存下,减小这个值能生成更多的线程.但是操作系统对一个进程内的线程数还是有限制的,不能无限生成,经验值在3000~5000左右
一般小的应用, 如果栈不是很深, 应该是128k够用的. 大的应用建议使用256k. 这个选项对性能影响比较大，需要严格的测试.
#### 通用措施 ####
1. 控制FullGC发生的时间大于24h, 在低峰期主动触发FullGC(定时任务每天凌晨4点system.gc()手动触发full-gc, 比较相对来说会比较吃内存, young区和old区相对来说都会比较大)
2. 让单次FullGC的时间相对较短


#### JVM监控参数 ####

##### JVM Misc #####

##### JVM Memory Pools(Non-Heap) #####
Metaspace:
1. java8移除Permgen（永久代），使用Metaspace
2. 主要存储类的元数据，比如类的各种描述信息，类名、方法、属性、访问限制等，都按照一定结构存在metaspace里

Compressed Class Space:

Code Cache:
java代码在执行时一旦被编译器编译为机器码，下一次执行的时候就会直接执行编译后的代码，也就是编译后的代码被缓存了起来。Code Cache就是缓存编译后的代码的内存区域。另外，除了JIT编译的代码外，java使用的本地代码（JNI）也会存在Code Cache中

##### Classloading #####

##### Buffer Pools #####




#### JVM监控指标 ####
概览
1. 进程启动时长
2. 进程启动时间
3. 堆内存使用率
4. 非内存使用率

服务黄金指标
1. QPS(1分钟平均)
2. error数(1分钟平均)
3. 请求耗时(1分钟平均)

jvm非堆内存详细
1. metaspace
2. compressed class space
3. code cache

垃圾回收(gc)
1. gc次数
2. 暂停时间
3. 内存分配/晋升

类加载
1. 已加载的类数量
2. 加载类数量的变化

Buffer Pools
1. direct buffers
A direct buffer is a chunk of memory typically used to interface Java to the OS I/O subsystems, for example as a place where the OS writes data as it receives it from a socket or disk, and from which Java can read directly.
直接管理操作系统中的内存, 不需要从操作系统再拷贝一次到jvm堆中
2. mapped buffers


G1:
G1 Evacuation Pause
G1 Humongous Allocation
GCLocker Initiated GC


end of major GC (Ergonomics)
end of minor GC (Allocation Failure)

#### 小技巧 ####
一个中间件如果不知道观察什么指标重要, 可以观察
1. 对应的工业级中间件商业化后的产品提供的监控大盘包括了哪些指标(eg: PolarDB)
2. 开源监控大盘取了哪些指标(eg: grafana dashboards)
了解了监控指标的意义, 离分析出当时的问题也不远了.








AdaptiveSizePolicy(自适应大小策略):
JDK 1.8 默认使用 UseParallelGC 垃圾回收器, 该垃圾回收器默认启动了 AdaptiveSizePolicy, 会根据GC的情况自动计算计算 Eden, From 和 To 区的大小

-XX:+UseAdaptiveSizePolicy 开启
-XX:-UseAdaptiveSizePolicy 关闭

1. 在 JDK1.8中, 如果使用 CMS, UseAdaptiveSizePolicy只能为false
2. UseAdaptiveSizePolicy不要和SurvivorRatio参数显示设置搭配使用, 一起使用会导致参数失效
3. 由于 AdaptiveSizePolicy 会动态调整 Eden,Survivor 的大小, 有些情况存在Survivor 被自动调为很小(eg: 几MB), 这个时候YGC回收掉 Eden区后, 还存活的对象进入Survivor 装不下, 就会直接晋升到老年代, 导致老年代占用空间逐渐增加, 从而触发FULL GC, 如果一次FULL GC的耗时很长(比如到达几百毫秒), 那么在要求高响应的系统就是不可取的

附: 对于面向外部的大流量, 低延迟系统, 不建议启用此参数, 建议关闭该参数

AdaptiveSizePolicy 有三个目标: 
1. Pause goal: 应用达到预期的 GC 暂停时间
2. Throughput goal: 应用达到预期的吞吐量, 即应用正常运行时间 / (正常运行时间 + GC 耗时)
3. Minimum footprint: 尽可能小的内存占用量

AdaptiveSizePolicy 为了达到三个预期目标，涉及以下操作:
1. 如果 GC 停顿时间超过了预期值, 会减小内存大小. 理论上, 减小内存, 可以减少垃圾标记等操作的耗时, 以此达到预期停顿时间.
2. 如果应用吞吐量小于预期, 会增加内存大小. 理论上, 增大内存, 可以降低 GC 的频率, 以此达到预期吞吐量.
3. 如果应用达到了前两个目标, 则尝试减小内存, 以减少内存消耗.



tenure
英 [ˈtenjə(r)]  美 [ˈtenjər]
n. （土地的）居住权，保有权；（尤指大学教师的）终身职位，长期聘用；（尤指重要政治职务的）任期，任职
v. 授予……终身职位（尤指教师、讲师职位）

tenured
英 [ˈtenjəd]  美 [ˈtenjərd]
adj. （美）享有终身职位的
v. 授予……终身职位（tenure 的过去分词）


=======================================
-XX:+UseGCLogFileRotation
问题1：丢失旧的日志
如果你配置的日志文件个数是5个，一段时间过后就会产生出来5个日志文件，假如最老的是gc.log.0，最近的是gc.log.4，当gc.log.4到达20M以后，日志会重新写入到gc.log.0，gc.log.0之前的内容会被清空掉！

问题2：日志会变混乱
假如现在还是有5个日志文件：gc.log.0到gc.log.4，现在JVM重启了，此时GC的日志会重新从gc.log.0开始写入，但是gc.log.1、gc.log.2、gc.log.3、gc.log.4这里面的日志却还是之前旧的日志！新旧日志就掺杂在了一起！要解决这个问题，在重启服务器之前你需要把老的日志全部迁移到其他地方。

问题3：不方便日志集中管理
这种情况下，当前活动的日志文件的后缀名会被标记为.current,假如当前活动的日志文件是gc.log.3，那么他的文件名会被命名成gc.log.3.current，如果你要把不同机器上的日志文件都放到一个集中的地方去的话，大多数运维人员都会使用rsyslog，但是这种命名方式对rsyslog来说是一个非常大的挑战！

问题4：对日志分析工具不友好
当使用日志分析工具（gceasy、GCViewer）来分析日志的时候需要上传多个日志文件而不是一个！




-XX:+PrintGCDetails
-XX:+PrintGCDateStamps
-Xloggc:/home/GCEASY/gc-%t.log
%t会给文件名添加时间戳后缀，格式是YYYY-MM-DD_HH-MM-SS。这样就非常简单了克服了UseGCLogFileRotation存在的所有的问题！


=======================================

When G1 GC determines that a garbage collection is necessary, it collects the regions with the least live data first (garbage first).



===================================
如何阅读 GC 日志
GC Allocation Failure 是我们经常遇到的一种GC日志
分配失败代表着在JVM的Eden区中没有更多的空间来分配对象了, 这是minor GC的正常日志

2020-03-17T19:03:19.701+0800: 6664.686:
Total time for which application threads were stopped: 0.0313360 seconds , Stopping threads took: 0.0000925 seconds



