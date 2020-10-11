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
			//Condition for "This is probably not a valid json"
		} else if strings.ToLower(string(v.Summary[0:11])) == "daily event" {
			err = json.Unmarshal([]byte(v.Description), dailyEvts)

			if err != nil {
				jackal.Logger.Error.Printf("There was a criticl error when attempting to unmarshal today's event for %s.\n%v", calenderID, err)
			}

		} else {
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
		jackal.Logger.Error.Println("There was a critical error verifying the DND Calender JSON for ", calenderID)
	}

	newCal = &GuildCalender{}

	err = json.Unmarshal([]byte(guildCal.Description), &newCal)

	if err != nil {
		jackal.Logger.Error.Println("There was a critical error when trying to unmarshal the guildCalender Description", err)
		return
	} else {
		E5Core.GuildCalendars[calenderID] = newCal
	}
	return
}
