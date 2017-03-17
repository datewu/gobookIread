
While most Linux system monitoring tools monitor only a single aspect of
the system, there are a few tools with a broader scope. To get an
overview and find out which part of the system to examine further, use
these tools first.

1. vmstat
vmstat collects information about process, memory, I/O, interrupts and 
CPU. If called without a sampling rate, it displays average values since
the last reboot. When called with a sampling rate, it displays actual 
samples:

[root@nginx ~]# vmstat 2 	### every 2 seconds output a sample
procs -----------memory---------- ---swap-- -----io---- --system-- -----cpu-----
 r  b   swpd   free   buff  cache   si   so    bi    bo   in   cs us sy id wa st
 0  0      0 147928 146180 574800    0    0     0    17    3    3  1  0 99  0  0
 0  0      0 147920 146180 574804    0    0     0    40  275  543  2  0 98  1  0
 0  0      0 147920 146180 574804    0    0     0     8  216  445  0  0 100  0  0
 0  0      0 147796 146180 574804    0    0     0    30  268  435  1  1 98  0  0
 0  0      0 147796 146180 574804    0    0     0     8  243  476  1  0 99  0  0

 Tip: First line of Output (Line 16 in this document)
The first line of the vmstat always display average values since the last reboot.

The columns show the following:
r
  Shows the number of processes in a runnable state. These processes are either
  executing or waiting for a free CPU slot. If the number of process in this
  column is constantly higher than the number of CPUs available, this may be an
  indication of insufficient CPU power.
b
  Shows the number of processes waiting for a resource other than a CPU. A high
  number in this column may indicate an I/O problem (network or disk).
swpd
  The amount of swap space (KB) currently used.
free
  The amount of unused memory (KB).
inact
  Recently unused memory that can be reclaimed. This column is only visible when
  calling vmstat with the parameter -a (recommended).
active
  Recently used memory that normally does not get reclaimed. This column is only
  visible when calling vmstat with the parameter -a (recommended).
buff
  File buffer cache (KB) in RAM that contains file system metadata. This column 
  is not visible when calling vmstat with the parameter -a.
cache
  Page cache (KB) in RAM with the actual contents of files. This column is not
  visible when calling vmstat with the parameter -a.
si/so
  Amount of data (KB) that is moved from swap to RAM (si) or from RAM to swap
  (so) per second. High so values over a long period of time may indicate that
  an application is leaking memory and the leaked memory is being swapped out.
  High si values over a long period of time could mean that an application was
  inactive for a very long time is active again. Combined high si and so values
  for prolonged periods of time are evidence of swap thrashing and may indicate 
  that more RAM needs to be installed in the system because there is not enough
  memory to hold the working set size.
bi
  Number of blocks per second received from a block device (for example, a disk
  read). Note that swapping also impacts the values shown here. The block size
  may vary between filesystem but can be determined using the stat utility. If
  If throughput is required the iostat may be used.
bo
  Number of blocks per second sent to a block device (for example, a disk write)
  Note that swapping also impacts the values shown here.
in
  Interrupts per second. A high value may indicate a high I/O level (network and
  /or disk), but could also be triggered for other reasons such as
  inter-processor interrupts triggered by another activity. Make sure to alse
  check /proc/interrupts to idendity the source of interrupts.
cs
  Number of context switches per second. This is the number of times that the 
  kernel replaces executable code of one program in memory with that of 
  another program.
us
  Percentage of CPU usage executing application code.
sy
  Percentage of CPU usage executing kernel code.
id
  Percentage of CPU time spent idling. If this value is zero over a longer
  period of time, your CPU(s) are working full capacity. This is not necessarily
  a bad sign -- rather refer to the values in columns r and b to determine if 
  your machine is equipped with sufficient CPU power.
wa
  If "wa" time is non-zero, it indicates throughput lost because of waiting for
  I/O. This may be inevitable, for example, if a file is being read for the 
  first time, background writeback cannot keep up, and so on. It can also be
  an indicator for a hardware bottlenect (network or hard disk). Lastly, it can
  indicate a potential for tuning the virtual memory manager.
st
  Percentage of CPU time stolen from a virtual machine.

2. dstat
dstat is a replacemeant for tools such as vmstat, iostat, netstat, or ifstat.
It is written in Python and can be enhanced with plug-ins.

3. System Activity Information: sar
sar can generate extensive reports on almost all important system activities,
among them CPU, memory, IRQ usage, IO,or networking. It can also generate 
reports on the fly. sar gathers all their data from the /pro file system.

NOTE: sar is a part of the sysstat package.

To generate reports on the fly, call sar with an interval (seconds) and a
count. To generate reports from files specify a file name with the option
-f instead of interval and count. If file name, interval and count are
not specified, sar attempts to generate a report from /var/log/sa/saDD,
where DD stands for the current day. This is the default location to
where sadc (the system activity data collector) writes its data. Query 
multiple files with multiple -f options.

When called with no options, sar shows a basic report about CPU usage. On
multi-processor machines, results for all CPUs are summarized. Use the
option -P ALL to also see statistics for individual CPUs.

Memory Usage Reports: sar -r
Paging Statistics Reports: sar -B
Block Device Statistics Reports: sar -d
Network Statistics Reports: sar -n KEYWORD
    - DEV: for all network devices
    - EDEV: error statistics for all network devices
    - NFS: for an NFS client
    - NFSD: for an NFS server
    - SOCK: on sockets
    - ALL: Generates all network statistic reports
NOTE: kSar, a Java application visulizing your sar data,creates easy-to-read
graphs. It can even generate PDF reports.
