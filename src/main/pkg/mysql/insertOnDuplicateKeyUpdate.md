REPLACE
works exactly like INSERT , except that if an old row in the table has the same value as a new row for a PRIMARY KEY or a UNIQUE index, the old row is deleted before the new row is inserted. 

如果有主键或者唯一键冲突, 那么就删除旧数据,并新插入一条数据(2 rows affected).


---------------------------------------
性能优化: 关闭死锁检测


通过SQL语句查询锁表相关信息:
1. 查询表打开情况
SHOW OPEN TABLES WHERE IN_USE > 0

2. 查询锁情况列表
SELECT * FROM INFORMATION_SCHEMA.INNODB_LOCKS

3. 查询锁等待信息, 其中blocking_lock_id是当前事务在等待的事务
SELECT * FROM INFORMATION_SCHEMA.INNODB_LOCK_WAITS

4. 查询死锁日志
SHOW ENGINE INNODB STATUS
这条语句只能显示最新的一条死锁, 无法完全捕获到系统发生的所有死锁信息

如果想要记录所有的死锁日志，需要打开innodb_print_all_deadlocks参数，将所有的死锁日志记录到errorlog中

5. 查询锁等待时间
SHOW STATUS LIKE '%lock%'


---------------------------------------
死锁避免
---------------------------------------
当有死锁发生时，通常是由于项目的程序中出现了冗长的事务，或是由于隔离级别设置的不合适等。

我们需要在事务使用中注意以下几点：

1. 尽量保持事务的短小精悍，做出一系列关联的更新操作后立即提交事务，以降低死锁的可能性。特别是不要让有关联的MySQL会话长时间挂起未提交的事务。
2. 建议使用更低的隔离级别，如READ COMMITTED。
3. 在同一事务内修改多张表，或一张表内的不同行时，每次以相同的顺序执行操作。以便让事务形成清晰的锁操作队列而规避死锁

-----------------------------
死锁解决:
-----------------------------

MySQL数据库通过死锁检测（innodb_deadlock_detect）和死锁超时时间（innodb_lock_wait_timeout）这两个参数来进行死锁解决。

死锁检测（innodb_deadlock_detect)：在MySQL 8.0中，增加了一个新的动态变量innodb_deadlock_detect，用来控制InnoDB是否执行死锁检测。

该参数的默认值为ON，即打开死锁检测。开启后InnoDB在加锁的时候会检测加锁后是否会造成死锁，如果会加锁，就回滚代价最小的那一个事务。

死锁超时时间（innodb_lock_wait_timeout）：这个参数可以用来处理检测不出来的死锁，或是避免长时间等待较长的事务的情况。

对于高并发的系统，当大量线程等待同一个锁时，死锁检测可能会导致性能的下降。

此时，如果禁用死锁检测，而改为依靠参数innodb_lock_wait_timeout来释放长时间占用锁资源的事务可能会更加高效。

也就是说，在确认死锁检测功能影响了系统的性能并且禁用死锁检测不会带来负面影响时，可以尝试关闭innodb_deadlock_detect选项。

另外，如果禁用了InnoDB死锁检测，需要及时调整参数innodb_lock_wait_timeout的值，以满足实际的需求。

---------------------------------------




