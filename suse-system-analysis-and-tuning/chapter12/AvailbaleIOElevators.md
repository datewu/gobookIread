In the following elevators available on SUSE Linux Enterprise Server are
listed. Each elevator has a set of tunable parameters, which can be set 
with the following command:
`echo VALUE > /sys/block/DEVICE/queue/iosched/TUNABLE`
where VALUE is the desired value for the TUNABLE and DEVICE the block 
device.
To find out which elevator is the current default, run the following 
command. The currently selected scheduler is listed in brackets:

```
[root@nginx ~]# ls  /sys/block/sdb/queue/iosched/
back_seek_max      fifo_expire_sync  low_latency  slice_async_rq
back_seek_penalty  group_idle        quantum      slice_idle
fifo_expire_async  group_isolation   slice_async  slice_sync
[root@nginx ~]# cat /sys/block/sdb/queue/scheduler
noop anticipatory deadline [cfq]
```
This file can also contain string none meaning that I/O scheduling does 
not happen for this device. This is usually because the device uses
multi-queue queueing mechanism.

1. CFQ(Completely Fair Queuing)
CFQ is a fairness-oriented scheduler and is used by default on SUSE Linux
Server. The algorithm assigns each thread a time slice in which it is 
allowed to submit I/O to disk. This way each thread gets a fair share of
I/O throughput. It also allows assigning tasks I/O priorities which are
taken into account during scheduling decisions. The CFQ scheduler has
the following tunable parameters:
  slice_idle_us : When a task has no more I/O to submit in its time slice
    the I/O scheduler waites for a while before scheduling the next
    thread. The slice_idle_us is the time in microseconds the I/O 
    scheduler waits. File slice_idle controls the same tunable but in
    millisecond units. The waiting for more I/O from a thread can improve
    locality of I/O. Additionally, it avoids starving processes doing 
    dependent I/O. A process does dependent I/O if it needs a result of
    one I/O in order to submit another I/O. For example, if you first 
    need to read an index block in order to find out a data block to 
    read, thest two reads form a dependent I/O. For media where locality
    does not play a big role (SSDs, SANs with lots of disks) setting
    slice_idle_us to 0 can improve the throughput considerably.
  quantum :  This option limits the maximum number of requests that are
    being processed at once by the device. The default value is 4. For a
    storage with several disks, this setting can unnecessarily limit 
    parallel processing of requests. Therefore, increasing the value can
    improve performance. However, it can alse cause latency of certain 
    I/O operations to increase because more requests are buffered inside
    the storage. When changing this value, you can also consied tuning
    slice_async_rq(the default value is 2). This limits the maximum 
    number of asynchronous requests, usually write requests, that are
    submitted in one time slice.
  low_latency : When enabled(which is the default on SUSE Linux 
    Enterprise Server) the scheduler may dynamically adjust the length of
    the time slice by aiming to meet a tuning parameter called the 
    target_latency. Time slices are recomputed to meet this target_latency
    and ensure that processes get fair access within a bounded lengeth of
    time.
  target_latency : Contains an estimated latency time for the CFQ. CFQ
    will use it to calculate the time slice used for every task.
  group_idle_us : To avoid starving of blkio cgroups doing dependent I/O,
    CFQ waits a bit after completion of I/O for one blkio cgroup before
    scheduling I/O for a differnet blkio cgroup. When slice_idle_us is
    set, this parameter does not have a big impact. Howevevr, for fast 
    media, the overhead slice_idle_us is generally indesirable. Disabling
    slice_idle_us and setting group_idle_us is a method to avoid 
    starvation of blkio cgroups doing dependent I/O with lower overhead.
    Note that the file group_idle controls the same tunable however with
    millisecond granularity.

2. NOOP
A trivial scheduler that only passed down the I/O that comets to it. 
Useful for checking whether complex I/O scheduling decisions of other
schedulers are causing I/O performance regressions. This scheduler is
recommended for setups with devices that do I/O scheduling themselves,
such as intelligent storage or in multipathing environments. If you 
If you choose a more complicated scheduler on the host, the scheduler of 
the host and the scheduler of the storage device compete with each other.
This can decrease performance. The storage device can usually determine
best how to schedule I/O.
For similar reasons, this schedule is also recommended for use within
virtual machines.
The NOOP scheduler can be useful for devices that do not depend on 
mechanical movement, like SSDs. Usually, the DEADLINE I/O scheduler is 
better choice for these devices. However, NOOP creates less overhead and 
thus can on certain workloads increase performance.

3. DEADLINE
DEADLINE is a latency-oriented I/O scheduler. Each I/O request is assigned
a deadline. Usually, requests are stored in queues(read and write) sorted
by sector numbers. The DEADLINE algorithm maintains two additional queues
(read and write) in which requests are sorted by deadline. As long as no
request has timed out, the "sector" queue is used. When timeouts occur,
requests from the "deadline" queue are served until there are no more 
expired requests. Generally, the algorithm prefers reads over writes.
This scheduler can provide a superior throughput over the CFQ I/O 
scheduler in cases where several threads read and write and fairness is
not an issue. For example, for several parallel readers from a SAN and
for databases (especially when using "TCQ" disks). The DEADLINE scheduler
has the following tunable parameters:
  writes_starved :Controls how many reads can be sent to disk before it is
    possible to send writes. A value of 3 means, that three read operation
    are carried out for one write operation.
  read_expire : Sets the deadline (current time plus the read_expire value
    for read operations in milliseconds. The default is 500.
  wirte_expire : Sets the deadline (current time plus write_expire value)
    for write operations in milliseconds. The default is 5000.
