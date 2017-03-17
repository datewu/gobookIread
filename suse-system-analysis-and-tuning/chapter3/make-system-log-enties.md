Using logger to Make System Log Entries
logger is a tool for making entries in the system log. It provides a shell 
command interface to the rsyslogd system log module. For example, the following
line outputs its message in /var/log/messages or directly in the journal (if
no logging facility is running):
logger -t Test "This messages come form &USER"
Depending on the current user and host name, the log contains a line similar to
this: Mar 17 16:18:19 nginx Test: This message comes from root
