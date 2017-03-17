Graph Your Data: RRDtool
There are a lot of data in the world around you, which can be easily
measured in time. For example, changes in the temperature, or the number
of data sent or received by your computer's network interface. RRDtool
can help you store and visualize such data in detailed and customizable
graphs.
Tip: Bindings. There are Perl, Python, Ruby, and PHP bindings available
for RRDtool, so that you can write your own monitoring scripts in your
preferred scripting language.

1. How RRDtool Works
RRDtool is an abbreviation of ( Round Robin Database tool). Round Robin
is a method for manipulating with a constant amount of data. It uses the
principle of a circular buffer, where there is no end nor beginning to
the data row which is being read. RRDtool uses Round Robin Database to
store and read its data.
As mentioned above, RRDtool is designed to work with data that change in
time. The ideal case is a sensor which repeatedly reads measured data (
like temperature, speed etc) in constant periods of time, and then 
exports them in a given format. Such data are perfectly ready for RRDtool
and it easy to process them and create the desired output.
Sometimes it is not possible to obtain the data automatically and 
regularly. Their format needs to be pre-processed before it is supplied 
to RRDtool, and often you need to manipulate RRDtool even manually.

2. A Practical Example
  tldt;
