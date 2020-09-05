Requirements
-
Coyotech JACK-AL Bot Dispatch Framework requires the use of the following
* BWMarrin DiscordGo<br>

Without these libraries, this file library will return critical errors.

Methods
-
<pre>
    <code>addCreateListener(name string, responder func(message *discordgo.Message) (err error))</code>
</pre>
addCreateListener is a non-expported function which can be used by responders compiled in this package to add themselves into the 10-8 pool, or the Dispatch Pool. Once added to the pool, these listeners will receive updates each time an event they're subscribed to is fired.

Error Codes
-
- 100: Closed peacefully by JACK-AL Bot Framework