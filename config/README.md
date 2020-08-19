Requirements
-
Coyotech Config library requires the use of the following
* Coyotech Logging<br>

Without these libraries, this file library will return critical errors.

Methods
-
<pre>
    <code>config.Init(core interface{}) (console *log.Logger, info *log.Logger, warn *log.Logger, err *log.Logger)</code>
</pre>

The Init function requires an interface containing the method "LogFile". This method MUST return a string that points to the file path of the log file. If one does not exist where the path indicates, then one will be created.

<pre>
    <code>config.LoadCfg(filename string, config interface{}) (err error)</code>
</pre> 
The LoadCfg function requires a filename and an interface to be provided. The function will parse the json file located at the specified position and attempt to load it into the interface provided. It is recommended that the interface is referenced as a pointer using the & keyrune.

Error Codes
-
- 10: Failure opening config file.
- 11: Failure reading from config file.
- 12: Could not find User Config Directory
- 13: Failure writing data to config file.