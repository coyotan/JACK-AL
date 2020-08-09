Package - Coyotech: LogUtil
- 
The Logutil package from Coyotech is designed to be easily implemented into Golang projects. It is used as a dependency for some Coyotech Packages.

Dependencies 
- 
Coyotech - Config

Methods
-
<pre>
<code>InitLoggers(logFile string) (Console *log.Logger, Info *log.Logger, Warn *log.Logger, Error *log.Logger)</code>
</pre>

The InitLoggers method is intended to configure formatted logging support for the specified file located at the path provided. If the file does not exist, it will attempt to create one. <br>
<br>
InitLoggers will return 4 of *log.Logger, each with specialized formatted tags to match their descriptions. These can be customized from their defaults if desired.

<pre>
<code>VerifyFile(fName string) (fExists bool)</code>
</pre>
VerifyFile attempts to validate the existence of the file at the path provided. It will return a boolean value based on the status of the file. 
* True: File exists.
* False: File does not exist.
<pre>
<code>CreateFile(fName string) (fHandle *os.File, err error)</code>
</pre>
CreateFile will attempt to create a file at the path provided. It may return two variables. Upon successful creation of the provided file, it will return the file handle along with nil, in place of the error.