# JACK-AL Command Framework Module
## Responsibilities
The JACK-AL Command Framework module is responsible for allowing the simple addition, modification, moderation, and integration of commands.
## Requirements
Coyotech JACK-AL Bot Dispatch Framework requires the use of the following
* https://github.com/bwmarrin/discordgo<br>
* https://github.com/bwmarrin/dgvoice

Without these libraries, this file library will return critical errors.

## Constants 
<pre><code>const (
   	//Jackal Command Constants for Context and Permission bitmasks.
   	Enabled  int = 0x1 
   	Discord  int = 0x2  
   	Guild    int = 0x4
   	Group    int = 0x8
   	DM       int = 0x10
   	Terminal int = 0x20
   	API      int = 0x40
   	RPC      int = 0x80
)</code></pre>
JACK-AL uses a bitmask system to determine if a command can be run in the current environment (eg. Guild, Group DM, etc), if a command can be run in a server, and if a user can run the specified command.

Command overrides are stored in the JACK-AL database under the GuildCommands and UserCommands table. The results from query are combined using an AND operation to check if the command can be run by the user, in the context of the guild.

When Command constants are used as context values, they determine if a command can be in a specific environment. See the below example for more a reference.

When Command Constants are used for permissions values, they determine if a command can be run in a specific guild or channel, or by a specific user. See the below example for a reference. 

####Context and Permissions Example:
<pre><code>var (
    MyGroup = CommandFramework.CommandGroup{Name: "MyTestGroup"}
)
           
func init() {
  MyGroup.NewCommand("ping", newPingResponder, CommandFramework.Enabled+CommandFramework.Discord+CommandFramework.Guild)
  
  err := MyGroup.RegisterAllCommands()
  
  if err != nil {
    return err
  }
}
   
func newPingResponder(_ *discordgo.Session, created *discordgo.MessageCreate) (err error) {
   
  message := created
 
  jackal.Logger.Console.Println("Discord message received: ", message.Content)
 
  //JACK-AL has a fall through for command
  if strings.ToLower(message.Content[len(jackal.Discord.CommandPrefix):]) == "ping" {
      _, err = jackal.Discord.Session.ChannelMessageSend(message.ChannelID, "Pong!")
      jackal.Logger.Info.Println("Received ping, Pong!")
  }
   
  return
}
</code></pre> 
## Methods

<pre>
    <code>//Placeholder</code>
</pre>


## Error Codes

##### Pre-Init:


##### Post-Init:
