package gcore

import (
	"encoding/json"
	"strings"
	"time"
)

var Cal = E5Core.dCore.GetGCore()

//GetDailyEventSummary is called by a 24 hour recurring function. It MUST run before 0600, but specific timing is otherwise unimportant.
func GetDailyEventSummary(calenderID string) {

	est, _ := time.LoadLocation("America/New_York")
	t0 := time.Now().In(est)
	y, m, d := t0.Date()
	t0d := time.Date(y, m, d, 0, 0, 0, 0, est)
	t1 := t0d.Add(time.Hour * 24).Format(time.RFC3339)

	//Verify the guild is legit
	gCal, err := verifyGuildCalender(calenderID)

	if err != nil {
		jackal.Logger.Error.Printf("There was a critical error verifying the Guild Calendar\n %v", err)
	}

	evts, err := Cal.Events.List(calenderID).ShowDeleted(false).TimeMax(t1).TimeMin(t0d.Format(time.RFC3339)).SingleEvents(true).OrderBy("startTime").Do()

	if err != nil {
		jackal.Logger.Error.Println("There was a error reading events from the calendar: ", err)
	}

	dailyEvts := &DailyCalendarEvent{
		GuildCal: gCal,
	}

	for _, v := range evts.Items {
		if len(v.Description) < 10 {

			//Condition for "This is probably not a valid json". This should only fire if the user has not properly formatted their Google Calendar event. We should skip it to avoid errors.
			jackal.Logger.Error.Printf("Event %s ID(%v) does not meet the requirements for a DND5e Module event. Skipping. Please review the Readme docuemtnation.", v.Summary, v.Id)
			jackal.Logger.Console.Printf("Skipping event %s. Check error log.", v.Summary)

		} else if strings.ToLower(string(v.Summary[0:10])) == "daily event" {
			/*
				This condition assumes that the data you received in the description of the calendar event is for a daily event. Daily events contain information about upcoming events and daily conditions which players should be aware of.
			*/
			err = json.Unmarshal([]byte(v.Description), dailyEvts)

			if err != nil {
				jackal.Logger.Error.Printf("There was a criticl error when attempting to unmarshal today's event for %s.\n%v", calenderID, err)
			}

		} else {
			/*
				This portion of code assumes that the data received in the description of the event is for a Special Event. Special events are scheduled incidents such as the Bee and Barley Job Fair.
			*/
			newSpecEvt := SpecialEvent{}

			err = json.Unmarshal([]byte(v.Description), &newSpecEvt)

			if err != nil {
				jackal.Logger.Error.Printf("There was an error attempting to unmarshal an event in the %s Calendar.\n %v", calenderID, err)
			}

			newSpecEvt.GuildCal = gCal
			E5Core.GuildCalendars[calenderID].SpecEvts = append(E5Core.GuildCalendars[calenderID].SpecEvts, newSpecEvt)

			//We will need to create goroutines for each of these later, so we can perform the appropriate actions
		}
	}
}

func verifyGuildCalender(calenderID string) (newCal *GuildCalender, err error) {
	guildCal, err := Cal.Calendars.Get(calenderID).Do()

	if err != nil {
		return
	}

	newCal = &GuildCalender{}

	err = json.Unmarshal([]byte(guildCal.Description), &newCal)

	if err != nil {
		return
	} else {
		E5Core.GuildCalendars[calenderID] = newCal
	}
	return
}
