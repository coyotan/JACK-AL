DnD5e Google Calendar Module
-
This module provides JACK-AL bots with a method of interacting with the Google Calendar system. Using this module, game masters can configure automatic daily events for their players. This module is primarily developed for JACK-AL Esper, however Archon features will be added at a later date.

TODO:
-
Design this module. The design layout should be as follows:
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