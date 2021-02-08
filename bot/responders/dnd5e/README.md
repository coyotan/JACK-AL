# Dungeons and Dragons 5th Edition Module - JACK-AL Framework Extension
The DND5e JACK-AL Framework Extension allows for Game Masters who wish to host, or integrate their DND sessions with Discord to employ a valuable kit of tools for running and properly mastering games over Discord.

## Features:
1) Combined player and task scheduling for events and announcements.
    - explain
2) Automatic execution of Tasks.
    - explain
3) Gamemasters can easily define their own events.
    - explain. Demonstrate.
## Event List:
- CallToArms
- ClergyNotif
- Daily
- Decree
- Election
- EndElection
- GuardNotif
- GuildNotif
- IncomeEvent
- NobilityNotif
- QuestBoard
- Restock
- Taxes
- WarDeclaration
- WarCompletion

## Event Structure:
### Basic Event Structures
##### Basic Notification Event Structure:
Below is an example of the basic Notification event structure for the DND5e Framework Extension. The text should be placed into the description section of a Google Calendar Task. (Provide Example).<br><br>
The below applies to the following Events: CallToArms, ClergyNotif, GuardNotif, GuildNotif, and NobilityNotif.<br><br>
Title and Description are required.<br>
Picture and Color are optional.
```json
{
  "Event": "[EventName]",
  "Content": {
    "Title": "Example",
    "Description": "This is an example of a basic event structure. It can be used as a reference to create custom events.",
    "Picture": "https://link.to.picture/dogecandoit",
    "Color": "0x555555"
  }
}
```
### Special Event Structures:
##### SpecialEventNotif Event
The SpecialEventNotif event can be used to RSVP players for specific upcoming events. Players can indicate their reservation by reacting to the announcement. After their reaction, they will be assigned a corresponding role that can be used for tracking. This can be used in combination with the Daily and QuestBoard Event.<br><br>
Title, Description, EventRole, and React are all required.<br>
Picture and Color are optional.<br>
```json
{
  "Event": "SpecialEventNotif",
  "Content": {
    "Title": "Example Special Event Notification!",
    "Description": "Someone should probably put something interesting here!",
    "Picture": "https://link.to.picture",
    "Color": "0x555555",
    "EventRole":  "NameOfEventRole",
    "React":  ":ValidDiscordEmote:"
  }
}
```
If the event role does not exist in the server, JACK-AL will attempt to create one. If one cannot be created, an error will be generated.<br>
##### QuestBoard Event Structure
QuestBoard Events are used to announce special events that players can take part in with a game master. After announcing a job on the quest board, players can click on the reaction to accept the quest. This will add that player to the quest role for tracking.
<br><br>
React, Mentions, Picture, Cost, Reward, and QuestRole are all optional, and can be removed. All remaining fields are required.
```json
{
  "Event":  "QuestBoard",
  "Content": {
    "Title": "JobName",
    "Description": "This is an example job. Players will be able to read this to get an idea of what is so special about the job. Use this section to sell your idea to them, to buy in interest.",
    "Picture": "https://link.to.picture/dogecandoit",
    "Color": "0x555555",
    "QuestID": "SomethingUnique",
    "Cost": ["costItem1","costItem2"],
    "Reward": ["rewardItem1","rewardItem2"],
    "Difficulty": "Easy",
    "Location": "In town",
    "Origin": "The King Himself.",
    "PartySize": "Medium/4",
    "React": ":AnyValidDiscordEmote",
    "QuestRole": "SomeRoleNameFromYourServer",
    "Mentions": ["NamesOf", "Roles", "YouWant","ToPing"]
  }
}
``` 
##### Daily Event
Daily Events are special events which can be used to engage players in RP or dungeon scenarios that are special. They can be random, passive events, or active special events.
<br><br>
Title and Description are required.<br>
Role, Color, and Picture are optional.
```json
{
  "Event": "Daily",
  "Content": {
    "Title": "Example of a Daily Event!",
    "Description": "GMs should change this description to something more interesting than sample text!",
    "Role": ["List","OfRoles","ToPing"],
    "Picture": "https://link.to.picture",
    "Color": "0x555555"
  }
}
```
##### Taxes
This one is pretty obvious. GMs should set this incident up as recurring, so that it only has to be set once. <br>Announced, and Percentage OR Amount are required. <br>If Percentage and Amount are BOTH provided, amount is considered a minimum value.<br>
Title is required if Announced is true.<br>Description is optional.
```json
{
  "Event": "Taxes",
  "Content": {
    "Announced": true,
    "Title": "City of Tevian Tax Service",
    "Description": "Hello, citizens! As a part of the King's plans to continue the development of Tevian, the Tax Service has taken a small portion of all the citizen's income.",
    "Percentage": 15,
    "Amount": 20
  }
}
```
## Todo:
1) Using data buckets, keep track of which guild calendar is associated with which guild. GMs can link a single calendar by using the !dnd calendar set command.
    1) The !dnd calendar set command will take the Google Calendar link as a sole argument. It will then create an entry into the data buckets. <br><br>
2) We will index through the bucket of guild Calendars every midnight. Each guild handler will fetch all events from the appropriate guild calendar for the next 24 hours. These events will be created as objects and stored in an array.
    1) GMs will have the ability to set their timezone. JACK-AL will check every hour to see if any guilds exist in midnight timezones.
    2) Guild Handlers will start countdown timers for the events which were retrieved.
    <br><br>
3) When a guild event has triggered, it will post a notification in the guild's dedicated event channel. If necessary, it will also ping the associated roles.
    1) Guild Events can be created using the !dnd event create command.
    2) GMs can configure which channel event announcements are made in using the !dnd notifications command.
        * Ex: !dnd notifications #event-notifications
    3) When a guild event has been triggered, a timer (attached to the object) will begin. Once this timer has ended, the event will be concluded. The duration of this timer is determined by the "end date" on Google Calendar.