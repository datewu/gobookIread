System log file analysis is one of the most important tasks when analyzing the
system. In fact, looking at the system log files should be the first thing to do
when maintaining or troubleshooting a system.
Since the move to systemd, kernel messages and messages of system services 
registered with systemd are logged in systemd journal. Other log files (mainly
those of system applications) are written in plain text and can be easily read
using an editor or pager. It is also possible to parse them using scripts. This
allow you to filter their content.
