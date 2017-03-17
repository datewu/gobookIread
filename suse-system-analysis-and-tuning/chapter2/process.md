1. Interprocess Communication: ipcs
The command ipcs produces a list of the IPC resources currently in use.

2. Process List: ps
The command ps produces a list of processes. Most parameters must be written without
a minus sign. Refer to ps --help for a brief help or to the man page for extensive
help.
  - To list all processes with user and command line information, use ps axu
  - To check how many sshd process are running, use the option -p together with the 
command pidof/pgrep, which list the process IDs of the given processes.
The process list can be formatted according to your needs. The option L returns
a list of all keywords. 

USEFUL ps CALLS
  ps aux --sort column
    Sort the output by column. Replace column with
      pmen: for physical memory ratio
      pcpu: for CPU ratio
      rss: for resident set size (non-swapped physical memory)
  ps axo pid,%cpu,rss,vsz,args,wchan
    Shows every process, their PID,CPU usage ratio, memory size (resident and
    virtual), name, and their syscall.
  ps axfo pid,args
    Show a process tree.

3. Process Tree: pstree
The command pstree produces a list of processes in the form of a tree.
The parameter -p adds the process ID to a give name. To have the command lines
displayed as well, use the -a parameter.

4. Table of Processes: top
The command top (an abbreviation of table of processes) displays a list of 
processes that is refreshed every two seconds. To terminate the program, press
q. The parameter -n 1 terminates the program after a single display of the 
process list.
By default the output is sorted by CPU usage (column %CPU, shortcut shift+p).
Use the following key combinations to change the sort field:
  - Shift + m : Resident Memory (RES)
  - Shift + n : Process ID(PID)
  - Shift + t : Time(TIME+)
To use any other field for sorting, press f and select a field form the list.
To toggle the sort order, Use Shift + r.
The parameter -U UID monitors only the processes associated with a particular
user. Replace UID with the user ID of the user.
Use top -U $(id -u) to show processes of the current user.

5. z Systems Hypervisor Monitor: hyptop

6. A top-like I/O Monitor: iotop
The iotop utility displays a table of I/O usage by processes or thread.
iotop displays columns for the I/O bandwidth read and written by each process
during the sampling period. It also displays the percentage of time the process
spent while swapping in and while waiting on I/O. For each proces, its I/O
priority (class/level) is shown. In addition, the total I/O bandwidth read 
and written during the sampling period is displayed at the sampling period is
displayed at the top of the interface.
  - r: reverses the sort order.
  - o: toggles between showing all processes and threads (default view) and 
       showing only those donig I/O. (This function is similar to adding
       --only on command line.)
  - p: toggles between showing threads (default view) and processes. (This
       function is similar to --only.)
  - a: toggles between showing the current I/O bandwidth (default view) and
       accumulated I/O operatins since iotop was started. (This function is
       similar to --accumulated.)
  - i: let you change the priority of a thread or a process's threads.
  - q: quits iotop.

7. Modify a process's niceness: nice and renice
The kernel determines which processes require more CPU time than others by the
process's nice level, also called niceness. The higher the "nice" level of a 
process is, the less CPU time it will take from other processes. Nice levels
range form -20 (the least "nice" level) to 19. Negative values can only
be set by root.
Adjusting the niceness level is useful when running a non time-critical process
that lasts long and uses large amounts of CPU time. For example, compiling a
kernel on a system that also performs other tasks. Making such a process
"nicer", ensures that the other tasks, for example a Web server, will have
a higher priority.
Calling nice without any parameters prints the current niceness.
Running nice pid increates increments the current nice level for the given
command by 10.
Using nice -n level pid lets you specify a new relative to the current one.
To renice all processes owned by a specific user, use the -u user option.
Process groups are reniced by the option -g process group id.
