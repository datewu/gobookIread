A carefully planned installation ensures that the system is set up
exactly as you need it for the given purpose. It also saves considerable
time when fine tuning the system.

1. Partitioning
Depending on the server's range of applications and the hardware layout,
the partitioning scheme can influence the machine's performance(although
to a lesser extent only). It is beyond the scope of this manual to 
suggest different partitioning schemes for particular workloads. However
the following rules will positively affect performance. They do not apply
when using an external storage system.
  - Make sure there always is some free space available on the disk, 
      since a full disk delivers inferior performance.
  - Disperse simultaneous read and write access onto different disks by,
      for example:
      - using separate disks for the operating system,data,and log files
      - placing a mail server's spool directory on a separate disk
      - distributing the user directories of a home server between
          different disks

2. Installation Scope
The installation scope has no direct influence on the machine's 
performance, but a carefully chosen scope of packages has advantages. It
is recommended to install the minimum of packages needed to run the 
server. A system with a minimum set of packages is easier to maintain and
has fewer potential security issues. Furthermore, a tailor made 
installation scope also ensures that no unnecessary services are started

