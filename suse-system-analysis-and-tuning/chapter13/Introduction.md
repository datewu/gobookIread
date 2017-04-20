Tuning the Task Scheduler
Modern operating systems normally run many different tasks at the same 
time. To provide each task with its required system resources, the Linux
kernel needs a tool to distribute available system resources to individual
tasks. And this is exactly what the task scheduler does.
The following sections explain the most important terms related to process
scheduling. They also introduce information about the task scheduler 
policy, scheduling algorithm, description of the task scheduler used by 
SUSE Linux Enterprise Server, and references to other sources of relevant
information.


1. Introduction
The Linux kernel controls the way that tasks(or processes) are managed on
the system. The task scheduler, sometimes called process scheduler, is 
the part of the kernel that decides which task to run next. It is 
responsible for best using system resources to guarantee that multiple 
tasks are being executed simultaneously. This makes it a core component 
of any multitasking operating system.

1.1 Preemption
The theory behind task scheduling is very simple. If there are runnable 
processes in a system, at least one process must always be running. If
there are more runnable processes that processors in a system, not all
the processes can be running all the time.
Therefore, some processes need to be stopped temporarily, or suspended,
so that others can be running again. The scheduler decides what process 
in the queue will run next.
As already mentioned, Linux, like all other Unix variants, is a 
multitasking operating system. That means that several tasks can be 
running at the same time. Linux provides a so called preemptive
multitasking, where the scheduler decides when a process is suspended.
This forced suspension is called preemption. All Unix flavors have been
providing preemptive multitasking since the beginning.

1.2 Timeslice
The time period for which a process will be the running before it is 
preemptive is defined in advance. It is called a timeslice of a process
and represents the amount of processor time that is provided to each 
process. By assigning timeslices, the scheduler make global decisions for
the running system, and prevents individual processes from dominating 
over the processor resources.

1.3 Process Priority
The scheduler evaluates processed based on their priority. To calculate 
the current priority of a process, the task scheduler uses complex 
algorithms. As a result, each process is given a value according to which
it is allowed to run on a processor.
