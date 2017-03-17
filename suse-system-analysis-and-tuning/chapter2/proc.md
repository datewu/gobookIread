0. The /proc File System 
The /proc file system is a pseudo file system in which the kernel reserves 
important information in the form of virtual files. For example:
Display the CPU type with this command: cat /proc/cpuinfo
Query the allocation and use of interruputs with: cat /proc/interrupts
The address assignment of executables and libraries is contained in the maps file:
cat /proc/self/maps
A lot more information can be obtained from the /proc file system. Some of the
important files and their contents are:
  - /proc/devices: Available devices
  - /proc/modules: Kernel modules loaded
  - /proc/cmdline: Kernel command line
  - /proc/meminfo: Detailed information about memory usage
  - /proc/config.gz: gzip-compressed configuration file of the kernel currently running
  - /proc/PID/ : information about processes currently runnig in the /proc/NNN 
      directories, where NNN is the process ID (PID) of the relevant process. Every
      process can find its own characteristics in /proc/self/.
Furter information is available in the text file /usr/src/linux/Documentation/filesystems/proc.txt (this file is available when the package kernel-source is installed),

1. procinfo
Important information from the /proc file system is summarized by the command procinfo.
To see all the information, use the parameter -a. The parameter -nN produces
updates of the information every N seconds. In this case, terminate the program
by pressing q.
By default, the cumulative values are displayed. The parameter -d produces the
differential values. procinfo -dnt dispalys the values that have changed in 
the last five seconds.

2. System Control Parameters: /proc/sys/
System control parameters are used to modify the Linux kernel parameters at runtime
They reside in /pro/sys/ and can be viewed and modified with the sysctl command. To
list all parameters, run sysctl -a. A single parameter can be listed with 
sysctl parameter name.
Parameters are grouped into categories and can be listed with sysctl category or by
listing the contents of the respective directories. The most important categories
are listed below. The links to further readings require the installation of
the package kernel-source.
  - sysctl dev (/proc/sys/dev/) : Device-specific information.
  - sysctl fs (/proc/sys/fs/) : Used file handles, quotas, and other file
      system-oriented parameters.
  - sysctl kernel (/proc/sys/kernel/) : Information about the task scheduler,
      system shared memory, and other kernel-related parameters.
  - sysctl net (/proc/sys/net/) : Information about network bridges, and general
      parameters (mainly the ipv4/ subdirectory).
  - sysctl vm (/proc/sys/vm/) : Entries in this path relate to information about
      the virtual memory, swapping, and caching.
To set or change a parameter for the current session, use the command
sysctl -w parameter = value. 
To permanently change a setting, add a line parameter = value to /etc/sysctl.conf
the run sysctl -p /etc/sysctl.conf.
