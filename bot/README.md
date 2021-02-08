# JACK-AL Bot Component
## Responsibilities: 
## Requirements

Coyotech JACK-AL Bot Framework requires the use of the following
* Coyotech Logging<br>
* Coyotech Config<br>
* BWMarrin DiscordGo<br>

Without these libraries, this file library will return critical errors.

## Methods

<pre>
    <code>bot.Init(core *structs.CoreCfg)</code>
</pre>
The Init function provides the JACK-AL Bot Framework with the core configuration required to begin interacting with Discord. This function should only be called once.

<pre>
    <code>bot.GetInput() string</code>
</pre>

The GetInput function reads data in from the terminal and returns a trimmed version of the inputed string.

## Error Codes
- 100: Closed peacefully by JACK-AL Bot Framework
- 101: Failure to open an active connection with Discord.