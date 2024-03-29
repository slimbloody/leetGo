IOPS过高时的处理方案
针对undo/redo log 和 binlog的写入磁盘机制， mysql其实提供了参数可以进行配置：

innodb_flush_log_at_trx_commit配置
此项配置用来针对undo/redo log的磁盘写入配置。他有3个类型:
设置为0: 会每隔1秒把缓存中的undo/redo log写入到磁盘。
设置为1: 每次提交事务（一般的insert和update都有事务）写入到磁盘，该方案最安全，也是最慢的。
设置为2: 会写入系统的缓存，但会每隔一秒才调用文件系统的“flush”将缓存刷新到磁盘上去。这样mysql即使崩了，系统缓存还在，比0的方案优。

所以通过配置注释看到，如果我们可以在数据库服务器宕机的时候，允许有1秒的数据丢失，其实用设置为2是最优的方案，可以提高性能。
查看当前配置的代码：

show variables like 'innodb_flush_log_at_trx_commit';

--------------------


查看当前配置的代码：
    show variables like 'sync_binlog';
设置的代码(立即生效，无需重启)：
    set global sync_binlog=100;

关系型数据库最注重的是ACID特性，按照优化的配置也就破坏了这个特性，所以有利的时候同时也存在着弊端，只适合数据允许接收丢失的情况。如果实在不能接收丢失，只要依靠分布式的方式，也就是增加处理能力，来支持业务所需的IOPS。

qjl的设置都是1
活动规格
规格：16核128 GB
磁盘：102400GB
32000






Threads_cached	57	The number of threads in the thread cache
Threads_connected	1268	The number of currently open connections.
Threads_created	31715	The number of threads created to handle connections.
Threads_running	1	The number of threads that are not sleeping.

thread_cache_size的意义：每创建一个链接，都须要一个线程来与之匹配，此参数用来缓存空闲的线程，以致不被销毁，若是线程缓存中有空闲线程，这时候若是创建新链接，MYSQL就会很快的响应链接请求。
mysql创建链接很是消耗资源，因此就有了thread_cache，当已有链接再也不使用以后，mysql server不是直接断开链接，而是将已有链接转入到thread_cache中，以便下次在有create thread的需求时，能够在cache中复用，提升性能，下降资源消耗。
threads_cached :表明当前此时此刻线程缓存中有多少空闲线程。
Threads_connected :表明当前已创建链接的数量，由于一个链接就须要一个线程，因此也能够当作当前被使用的线程数。
Threads_created :表明从最近一次服务启动，已建立线程的数量。
Threads_running :表明当前激活的（非睡眠状态）线程数。并非表明正在使用的线程数，有时候链接已创建，可是链接处于sleep状态，这里相对应的线程也是sleep状态。
四者之间的关系：
running和其余三个状态关系不大，但确定不会超过thread_connected
(new_con-old_con)=create+(old_cache-new_cache)
从上面公式能够看出，若是create等于0，那么thread_connected减小的和thread_cached增长的相等，thread_connected增长的和thread_cached减小的相等。（其实这也就是thread_cached存在的意义，资源能够复用）














----------------------------------


压力测试过程中，如果因为资源使用瓶颈等问题引发最直接性能问题是业务交易响应时间偏大，TPS逐渐降低等。而问题定位分析通常情况下，
1. 最优先排查的是监控服务器资源利用率，例如先用TOP 或者nmon等查看CPU、内存使用情况，
2. 然后在排查IO问题，例如网络IO、磁盘IO的问题。 如果是磁盘IO问题，一般问题是SQL语法问题、MYSQL参数配置问题、服务器自身硬件瓶颈导致IOPS吞吐率问题。

1. 打开日志跟踪引起的磁盘IO问题
例如：MySQL的日志包括错误日志（ErrorLog），更新日志（UpdateLog），二进制日志（Binlog），查询日志（QueryLog），慢查询日志（SlowQueryLog）等，正常情况下，在生产系统或者压力测试环境中很少有系统会时时打开查询日志。因为查询日志打开之后会将MySQL中执行的每一条Query都记录到日志中，会该系统带来比较大的IO负担，而带来的实际效益却并不是非常大。

2. SQL写法问题引起磁盘IO高
例如：曾经在做某一个项目时，在看到数据库磁盘IO使用率偏高，前端查询业务交易loadrunner显示事物响应时间偏长，通过监控工具抓取对应SQL，通过计划分析，发现该SQL 中使用distinct 又多表关联且是大表、然后使用order by，最终显示10笔数据，而在产生中间过程数据进行筛选时，使用的是临时表，并把数据放入临时表中，内存刚好设置不大，于是放到磁盘中导致IO偏高。

备注: MySQL在执行SQL查询时可能会用到临时表，临时表存储，MySQL会先创建内存临时表，但内存临时表超过配置指定的值后，MySQL会将内存临时表导出到磁盘临时表；

-----------------------------------

MySQL 实例在日常使用中会出现实例 IOPS 使用率高的情况 原因:
1. 实例内存满足不了缓存数据或排序等需要，导致产生大量的物理 IO。
2. 查询执行效率低，扫描过多数据行。



thread_cache_size的意义：每创建一个链接，都须要一个线程来与之匹配，此参数用来缓存空闲的线程，以致不被销毁，若是线程缓存中有空闲线程，这时候若是创建新链接，MYSQL就会很快的响应链接请求。
mysql创建链接很是消耗资源，因此就有了thread_cache，当已有链接再也不使用以后，mysql server不是直接断开链接，而是将已有链接转入到thread_cache中，以便下次在有create thread的需求时，能够在cache中复用，提升性能，下降资源消耗。

threads_cached :表明当前此时此刻 `线程缓存` 中有多少空闲线程。

Threads_connected :表明当前已创建链接的数量，由于一个链接就须要一个线程，因此也能够当作当前被使用的线程数。
Threads_created :表明从最近一次服务启动，已建立线程的数量。
Threads_running :表明当前激活的（非睡眠状态）线程数。并非表明正在使用的线程数，有时候链接已创建，可是链接处于sleep状态，这里相对应的线程也是sleep状态。

四者之间的关系:
running和其余三个状态关系不大，但确定不会超过thread_connected
(new_con-old_con)=create+(old_cache-new_cache)
从上面公式能够看出，若是create等于0，那么thread_connected减小的和thread_cached增长的相等，thread_connected增长的和thread_cached减小的相等。（其实这也就是thread_cached存在的意义，资源能够复用）





排查思路
线上遇到连接数溢出的问题，问题的排查步骤？

1. 查看实例配置（几核几G、支持的最大连接数）
2. 查看当前连接数（show processlist）
3. 排查是什么动作占用了这些连接
4. 分析连接被占用的原因
    1. 慢SQL（缺索引、join太多、查询数据没有分页等）
    2. 长事务
    3. 死锁


==================================================

--------------------------------------


QPS(Query Per Second,既每秒请求、查询次数)
show global status like "Questions";
show global status like "Uptime";

获取这个指标值也很容易在MySQL中执行status命令就可以看到了.不过这个值是在MySQL生命周期内全局指标,可我们的系统不是每时每刻都在忙碌,那么在系统峰值时QPS又是多少,我们只能自己动手算了.当我们执行status的时候有个Questions,尽管它也是全局指标.不过我们可以每隔一秒查询下这个值,并将相邻的两值相减,得到的就是精确的每一秒的实际请求数了.如果MySQL处于繁忙的状态,那么我们获取的值就可以视为MySQL QPS的峰值响应能力了

--------------------------------------

TPS(Transcantion Per Second,既每秒事务数) 至于TPS嘛...
同样是衡量数据库的重要指标.不过MySQL不是每个存储引擎都支持事务.所以就拿InnoDB来说好了
TPS主要涉及提交和回滚
TPS = (Com_commit + Com_rollback) / Seconds

show global status like "Com_commit";
show global status like "Com_rollback";

--------------------------------------


总结:
如果IOPS过高，分析原因：
1. 内存不足，查询数据较多（一般为慢查询，但有时候并不是，单纯的查询数据较多），特别关注一下数据量大还需要排序的分页的，不能通过内存一次完成查询，产生大量的io操作
2. 前面几种情景都没有问题，那可能是你的写操作过多了，从代码、业务或者架构考虑优化
3. 最后的办法..提升mysql、硬件服务器的iops配置，说白了就是换硬件比如机械磁盘换固态

在业务量级没有明显变化的时候主要排查1,2,可以通过优化sql或者对数据量较大的表进行分表处理，3就不说了就是花钱换速度

如果QPS过高，分析原因：
1. 这个一般没什么解决办法，很直观的指标，你的数据库访问次数过多了，可以通过缓存减少查询次数、消息队列削峰等

如果TPS过高，分析原因：
1. 一般也是直观的写操作过度了
2. 也可能是大量的写操作发生回滚


https://cloud.tencent.com/developer/article/1505255

==================================================

awk 获取mysql监控数据
https://blog.51cto.com/u_15127579/2726979


==================================================
mysql 指标
临时表数量:
Created_tmp_disk_tables 每秒创建磁盘临时表的次数

执行次数:
Com_select/s: 平均每秒select语句执行次数
Com_insert/s: 平均每秒insert语句执行次数
Com_update/s: 平均每秒update语句执行次数
Com_delete/s: 平均每秒delete语句执行次数
Com_delete_multi 记录多表delete语句执行的次数
eg: delete tb1,tb2 from tb1,tb2 where tb1.id=tb2.id
Com_insert_select: 记录INSERT INTO TABLE table1 SELECT * FROM table2 语句执行的次数
Com_update_multi:
update tb1,tb2 set tb1.id = tb1.id+1, tb2.name='test' where tb1.id=tb2.id;
记录UPDATE tablename1, tablename2... SET... 语句执行的次数;


InnoDB Data 读写吞吐量(KB)
平均每秒读取的数据量： innodb_data_read/s
平均每秒写入的数据量： innodb_data_written/s


InnoDB Buffer Pool 请求次数
平均每秒读的请求数:innodb_buffer_pool_read_requests
平均每秒向缓冲池写的请求数:innodb_buffer_pool_write_requests/s

InnoDB Buffer Pool 命中率
缓冲池的利用率:( 1 - innodb_buffer_pool_pages_free / innodb_buffer_pool_pages_total) * 100
缓冲池中的脏块百分比: innodb_buffer_pool_pages_dirty/innodb_buffer_pool_pages_total
展示所选择节点的缓冲池的脏块率、读命中率、利用率。来源于show global status命令查询结果中的Innodb_buffer_pool_pages_dirty、Innodb_buffer_pool_pages_total、Innodb_buffer_pool_reads、Innodb_buffer_pool_read_requests等指标。


InnoDB redo 写次数
1. mysql.innodb_log_write_requests 平均每秒日志写请求次数
2. mysql.innodb_log_writes 平均每秒向日志文件的物理写次数
3. mysql.innodb_os_log_fsyncs 平均每秒向日志文件完成fsync()写数量


Q: 当集群无业务流量时，为何性能监控中的QPS还显示为大约10次/秒？
A: 因为系统后台存在监控、日志采集以及管控任务，大概每秒会产生10个左右的查询请求，对集群运行基本无影响。

Q: CPU使用率过高时，应该如何解决？
A: 建议您按如下步骤进行排查：
1. 确认是否有大量慢请求，建议优化慢SQL后再进行测试。如何查看和优化慢SQL，请参见慢SQL。
2. 确认CPU曲线是否与QPS或TPS曲线走向一致，若一致，说明该问题是事务高并发导致，可以考虑升级集群配置。
3. 若没有慢请求，CPU曲线与QPS或TPS曲线走向也不一致，请提交工单联系技术支持解决。

Q: 当前连接数远大于活跃连接数时，应该如何解决？
A: 您可以尝试将wait_timeout和interactive_timeout参数值设置的小一些以加速空闲连接的释放。但建议最好在业务端使用完连接后及时关闭回收以减少空闲连接的存在。
interactive_timeout和wait_timeout：在连接空闲阶段(sleep)起作用
即使没有网络问题，也不能允许客户端一直占用连接。对于保持sleep状态超过了wait_timeout（或interactive_timeout，取决于client_interactive标志）的客户端，MySQL会主动断开连接。

https://help.aliyun.com/document_detail/68555.html?utm_content=g_1000230851&spm=5176.20966629.toubu.3.f2991ddcpxxvD1#title-d4b-8ia-hnh


-----------------------------------------
https://zhuanlan.zhihu.com/p/461739068

主键查询慢 1s 2s:
刚好buffer tool满了, 在淘汰页面, 有两次磁盘io所以比较慢  
相当于对buffer处理的速率远远大于磁盘io的速率, 这个时候可以增大buffer pool, 相当于mq把队列扩容了

优化iops方案:
1. 升级硬件, 加快磁盘io, 或者增大内存(提升buffer pool)
2. 存储优化
    1. 数据归档(冷数据)
    2. 分库分表
3. SQL优化
索引下 SQL 语句的优化和索引调整层面的优化。根据具体业务场景及数据调整索引策略，这个方面没什么好说的，尽可能使得扫描的行数降低
4. 配置优化: sync_binlog
此项配置用来针对 binlog 的磁盘写入配置，可以用来配置合并多少条 binlog 一次性写入磁盘