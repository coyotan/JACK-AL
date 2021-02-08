# Google Calendar Module

This module provides JACK-AL bots with a method of interacting with the Google Calendar system. With this module, JACK-AL can performed timed operations which can be programed through the Google Calendar and Google Tasks API, through https://calendar.google.com.

##T ODO:

REMEMBER THAT THIS DESIGN MUST SUPPORT MULTIPLE CALENDARS AND MULTIPLE GUILDS. USE DATABUCKET TO CORRELATE INFORMATION.

Design this module. The design layout should be as follows:
1) Decode Google Calendar data as structures.
    1) Including overarching Calendar Title, Description, etc.
   ```go
    //Structure of Main Calendar
    type Calendar struct {
       Title           string  `json:"Title"`
       Description     string  `json:"Description"`
       Date            string  `json:"Date"`
    }
    ```
    2) Decode Calendar events including Title, Description, Time.
    ```go
    //Structure of Calendar Event
    type Event struct {
       Title           string  `json:"Title"`
       Description     string  `json:"Description"`
       Date            string  `json:"Date"`
    }
    ```
2) Decode Google Tasks Data as structures.
    1) Including coorelation with Google Calendar dates.
    2) Decode Calendar Tasks including Title, Description, Time. We will store this data in respective fields.
    ```go
    type Task struct {
       Title           string  `json:"Title"`
       Description     string  `json:"Description"`
       Date            string  `json:"Date"`
    }
    ```
    <br>
3) Create a method to allow for scheduled tasks to be completed when an event is fired by the Google API for 
Tasks and Calendar Events. This can be done using the following structure in the Description field of the Task or Event.
    ```go
    /*
    Content can be decoded differently based on the type of information 
    described by the "Event" field. We can use this to select which module/function
    is triggered when this Task/Event is recognized by the Google Module similar to 
    the way that the command module for the JACK-AL Bot Framework functions.
    */
    type Description struct {
        Event   string      `json:"Event"`
        Content interface{} `json:"Content"`
    }
    ```