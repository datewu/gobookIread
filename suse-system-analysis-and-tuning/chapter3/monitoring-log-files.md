Monitoring Log Files with logwatch
logwatch is a customizable, pluggable log-monitoring script. It parse system logs
extracts the important information and presents them in a human readable manner.
To use logwatch, install logwatch package.

logwatch can either be used at the command line to generate on-the-fly reports,
or via cron to regularly create custom reports. Reports can either be printed on 
the screen, saved to a file, or be mailed to a specified address. The latter is
especially useful when automatically generating reports via cron.

On the command line, you can tell logwatch for which service and time span to 
generate a report and how much detail should be included:
# Detailed report on all kernel messages form yesterday
logwatch --service kernel --detail High --range Yesterday --print

The --range option has got a complex syntax -- see logwatch --range help for
details. A list of all services that can be queried is available with the
following command:
ls /usr/share/logwatch/default.conf/services/ | sed 's/\.conf//g'

logwatch can be customized to great detail. However, the default configuration
should usually be sufficient. The default configuration files are located under
/usr/share/logwatch/default.conf/. Never change them because they would get
overwritten again with the next update. Rather place custom configuration in
/etc/logwatch/conf/ (you may use the default configuration fileas a template,
though). A detailed HOWTO on customizing logwatch is available at 
/usr/share/doc/packages/logwatch/HOWTO-Customize-LogWatch. The following 
configuration files exist:
  - logwatch.conf : The main configuration file. The default version 
      is extensively commented. Each configuration option can be overwritten on
      the command line.
  - ignore.conf : Filter for all lines that should globally be ignored by logwatch
  - services/*.conf : The service directory holds configuration files for each 
      service you can generate a report for.
  - logfiles/*.conf : Specifications on which log files should be parsed for 
      each service.
