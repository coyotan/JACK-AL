package gcore

import (
	"encoding/json"
	"time"
)

var Cal = e5Core.dCore.GetGCore()

func GetDailyEventSummary(calenderID string) {

	est, _ := time.LoadLocation("America/New_York")
	t0 := time.Now().In(est)
	t1 := t0.Add(time.Hour * 24).Format(time.RFC3339)

	y, m, d := t0.Date()
	t0d := time.Date(y, m, d, 0, 0, 0, 0, est)

	//Verify the guild is legit

	evts, err := Cal.Events.List(calenderID).ShowDeleted(false).TimeMax(t1).TimeMin(t0d.Format(time.RFC3339)).SingleEvents(true).OrderBy("startTime").Do()

	if err != nil {
		jackal.Logger.Error.Println("There was a error reading events from the calendar: ", err)
	}

	for k, v := range evts.Items {
		if len(v.Description) < 10 {
			//Condition for "This is probably not a valid json"
		} else {

		}
	}

}

func verifyGuildCalender(calenderID string) (err error) {
	guildCal, err := Cal.Calendars.Get(calenderID).Do()

	if err != nil {
		jackal.Logger.Error.Println("There was a critical error verifying the DND Calender JSON for ", calenderID)
	}

	newCal := &GuildCalender{}

	err = json.Unmarshal([]byte(guildCal.Description), &newCal)

	if err != nil {
		jackal.Logger.Error.Println("There was a critical error unmarshalling the guildCalender Description", err)
		return
	} else {
		e5Core.GuildCalendars[calenderID] = newCal
	}
	return
}
