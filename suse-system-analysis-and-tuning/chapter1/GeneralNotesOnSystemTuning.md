
This manual discusses how to find the reasons for performance problems an
d provides means to sole these problems. Before you start tuning your sys
tem, you should make sure you have ruled out common problems and have 
found the cause for the problem. You should also have a detailed plan on
how to tune the system, because applying random tuning tips often will 
not help and could make things worse.

GENERAL APPROACH WHEN TUNING A SYSTEM
1. Specify the problem that needs to be solved.
2. In case the degradation is new, identify any recent changes to the
syetem.
3. Identify why the issue is considered a performance problem.
4. Specify a metric can be used to analyze performance. This metric
could for example be latency, throughput, the maximum number of users
that are simultaneously logged in, or the maximum number of active
users.
5. Measure current performance using the metric from the previous from
the previous step.
6. Identify the subsystem(s) where the application is spending the most
time.
7. a. Monitor the system and/or the application.
   b. Analyze the data, categorize where time is being spent.
8. Tune the subsystem identified in the previous step.
9. Remeasure the current performance without monitoring using the same
metric as before.
10. If performance is still not acceptable, start over with Step 3.



Be Sure What Problem to Slove

Before starting to tuning a system, try to describe the problem as
exactly as possible. A statement like "The system is slow!" is not
a helpful problem description. For example, it could make a 
difference between the system speed needs to be improved in general
or at peak times.
Furthermore, make sure you can apply a measurement to your problem,
otherwise you will not be able to verify if the tuning was a success
or not. You should always be able to compare "before" and "after".
Which metrics to use depends on the scenario or application you are
looking into. Relevant Web server metrics, for example, could be in
terms of
Latency: The time to deliver a page
Throughput: Number of pages served per second or megabytes transferred
per second
Active Users: The maximum number of users can be downloading pages while
still receiving pages within an acceptable latency


Rule Out Common Problems

A performance problem often is caused by network or hardware problems,
bugs, or configuration issues. Make sure to rule out problems such as
the ones listed below before attempting to tune your system:
  - Check the output of the systemd journal for unusual entries.
  - Check (using top or ps_ whether  a certain process misbehaves by 
      eating using unusual amounts of CPU time or memory.
  - Check for network problems by inspecting /proc/net/dev.
  - in case of I/O problems with physical disks. make sure it is not
      caused by hardware problems (check the disk with the smartmontools)
      or by a full disk.
  - Ensure that background jobs are scheduled to be carried out in times
      the server load is low. Those jobs should also run with low
      priority (set via nice).
  - If the machine runs several services using the same resources, 
      considere moving services to another server.
  - Last, make sure your software is up-to-date.


Finding the Bottleneck

Finding the bottleneck very often is the hardest part when tuning a 
system. SUSE Linux Enterprise Server offers many tools to help you with
this task. See Part Two. "System Monitoring" for detailed information on
general system monitoring applications and log file analysis. If the 
problem requires a long-time in-depth analysis, the Linux kernel offers
means to perform such analysis. See Part Three. "Kernel Monitoring" for
coverage.
Once you have collected the data, it needs to be analyzed. First, inspect
if the server's hardware (memory, CPU, bus) and its I/O capacities
(disk, metwork) are sufficient. If these basic conditions are met, the 
system might benefit from tuning.

Step-by-Step Tuning

Make sure to carefully plan the tuning itself. It is of vital importance
to only do one step at a time. Only by doing so you will be able to
measure if the change provided an improvement or even had a negative 
impact. Each runing activity should be measured over a sufficient time
period to ensure you can do an analysis based on significant data.If 
you cannot measure a positive effect, do not make the change permanent.
Chances are, that it might have a negative effect in the future.
