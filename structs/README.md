# Package: JACK-AL Structs

This package containers the core structures for the JACK-AL Bot Framework and associated systems. Here, configuration structures, the core structure, and other centralized program-specific objects, and their child methods, can be found and manipulated.
<br>

# Structures
#### Core
<pre>
<code>//This structure is saved to JSON and loaded by the config lib.
type CoreCfg struct {...}</code>
</pre>
The CoreCfg structure is the heart of the JACK-AL Framework, and contains a large portion of information that the bot (and associated services) cannot run without. This includes information such as the Discord Session, Token, Database references, Command Dispatchers and Responders, among other fields and methods.
<pre>
<code>(core *CoreCfg) LogFile() (fPath string)</code>
</pre> 
A method of CoreCfg, this function returns the string contained in the field "LogFilePath" of the CoreCfg structure. This is REQUIRED to be present by the Coyotech Config library.
<pre>
<code>(core *CoreCfg) LogConsole() (console *log.Logger)</code>
</pre>
A method of CoreCfg, this function returns a pointer referencing an instance of a Logger from the Golang "log" library. It is configured to display debugging information to the Console at Stdout in date, time, filename format.
<pre>
<code>(core *CoreCfg) LogInfo() (info *log.Logger)</code>
</pre> 
A method of CoreCfg, this function returns a pointer referencing an instance of a Logger from the Golang "log" library. It is configured to display INFO level information to the configured logfile identified in CoreCfg in date, time, filename format.
<pre>
<code>(core *CoreCfg) LogWarn() (warn *log.Logger)</code>
</pre>
A method of CoreCfg, this function returns a pointer referencing an instance of a Logger from the Golang "log" library. It is configured to display WARN level information to the configured logfile identified in CoreCfg in date, time, filename format.
<pre>
<code>(core *CoreCfg) LogError() (warn *log.Logger)</code>
</pre>
A method of CoreCfg, this function returns a pointer referencing an instance of a Logger from the Golang "log" library. It is configured to display ERROR level information to the configured logfile identified in CoreCfg in date, time, filename format.

Error Codes
-
