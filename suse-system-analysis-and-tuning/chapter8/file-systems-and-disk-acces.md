Hard disks are the slowest components in a computer system and therefore
often the cause for a bottleneck. Using the file system that best suits
your workload helps to improve performance.
Using special amount options or prioritizing a process's I/O priority are
further means to speed up the system.

1. File Systems
SUSE Linux Enterprise Server ships with several file systems, including
BtrFS, Ext4, Ext3, Ext2, ReiserFs, and XFS. Each file system has its own
advantages and disadvantages.

1.1 NFS
NFS (Version 3) tuning is covered in detail in the NFS Howto at http://blabla.bla/. 
The first things to experiment with when mounting NFS shares is
increasing the read write blocksize to 32768 by using the mount options
wsize and rsize.

2 Time Stamp Update Policy
Each file and directory in a file system has three time stamps associated
with it. Atime when the file was last read called access time, a time
when the file data was last modified called modification time, and a time
when the file metadata was last modified called change time. Keeping 
access time always uptodate has significant performance overhead since 
since every read-only access will incur a write operation. Thus by 
default every file system updates access time only if current file access
time is older than a day or if it is older than file modification or 
change time. This feature is called relative access time and the 
corresponding mount option is relatime. Update of access time can be 
completely disabled using noatime mount option however you need to verity
your applications do not use it. This can be true for file and Web 
servers or for network storage. If the default relative access time 
update policy is not suitable for your applications, use 
strictatime mount option.
Some file systems (for example ext4) also support lazy time stamp updates
When this feature is enabled using lazytime mount option, updates of all
time stamps happen in memory but they are not written to disk. That 
happens only in response to fsync or sync system calls, when the file 
information is written duo to outher reason such as file size update, 
when time stamps are older than 24 hours, or when cached file information
needs to be evicted from memory.
To update mount options used for a file system, either edit /etc/fstab,
or umount mount commands

3 Prioritizing Disk Access with inoice
The ionice command lets you prioritize disk access for single processes.
This enables you to give less I/O priority to background process with
heavy disk access that are not time-critical, such as backup jobs.
ionice also lets you raise the I/O priority for a specific process to
make sure this process always has immediate access to the disk.
The caveat of this feature is that standard writes are cached in the page
cache and are written back to persistent storage only later by an 
independent kernel process. Thus the I/O priority setting generally does
not apply for these writes. Also be aware that I/O class and priority 
setting is obeyed only by CFQ I/O scheduler.
You can set the following three scheduling classes:
  - Idle : A process from the idle scheduling class is only granted disk
      access when no other process has asked for disk I/O.
  - Best effort : The default scheduling class used for any process that
      has not asked for a specific I/O priority. Priority within this 
      class can be adjust to a lecel from 0 to 7 (with 0 being the 
      highest priority). Program running at the same best-effort priority
      are served in a round-robin fashion. Some kernel versions treat
      priority within the best-effort calss differently , for details,
      refer to the ionice(1) man page.
  - Real-time : Processes in this class are always granted disk access 
      first. Fine-tune the priority level from 0 to 7 (with 0 being the 
      highest priority). Use with care, since it can starve other process

For more details and the exact command syntax refer to the ionice(1) man
page. If need more reliable control over bandwidth available to each
application, please use Kernel Control Groups.
