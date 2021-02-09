package dnd5e

import (
	"errors"
)

//Add custom/predefined Header and Footer.
///TODO Add Validation function.
type BasicNotif struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Picture     string `json:"Picture, omitempty"`
	Color       string `json:"Color, omitempty"`
}

//Add custom/predefined Header and Footer.
///TODO Add Validation function.
type QuestBoard struct {
	Title       string   `json:"Title"`
	Description string   `json:"Description"`
	Picture     string   `json:"Picture, omitempty"`
	Color       string   `json:"Color, omitempty"`
	QuestID     string   `json:"QuestID"`
	Cost        []string `json:"Cost, omitempty"`
	Reward      []string `json:"Reward, omitempty"`
	Difficulty  string   `json:"Difficulty"`
	Location    string   `json:"Location"`
	Origin      string   `json:"Origin"`
	PartySize   string   `json:"PartySize"`
	React       string   `json:"React, omitempty"`
	QuestRole   string   `json:"QuestRole, omitempty"`
	Mentions    []string `json:"Mentions, omitempty"`
}

//Add custom/predefined Header and Footer.
///TODO Add Validation function.
type Daily struct {
	Role []string `json:"Role, omitempty"`
	BasicNotif
}

type Taxes struct {
	Announced   bool   `json:"Announced"`
	Title       string `json:"Title, omitempty"`
	Description string `json:"Description, omitempty"`
	Percentage  int    `json:"Percentage, omitempty"`
	Amount      int    `json:"Amount, omitempty"`
}

//Validate ensures that the fields are properly populated, to prevent errors from occurring when we attempt to generate an embedded message from this.
func (t *Taxes) Validate() (valid bool, reason error) {
	if t.Announced {
		if len(t.Title) == 0 {
			return false, errors.New("announced is true. title cannot be empty")
		}
	} else if t.Percentage == 0 && t.Amount == 0 {
		return false, errors.New("percentage and amount are both null. this is either a mistake or this event is useless")
	}
	return true, nil
}

type IncomePayment struct {
	Announced   bool   `json:"Announced"`
	Title       string `json:"Title, omitempty"`
	Description string `json:"Description, omitempty"`
	IncomeTax   int    `json:"IncomeTax, omitempty"`
}

//Validate ensures that the fields are properly populated, to prevent errors from occurring when we attempt to generate an embedded message from this.
func (i *IncomePayment) Validate() (valid bool, reason error) {
	if i.Announced {
		if len(i.Title) == 0 {
			return false, errors.New("announced is true. title cannot be empty")
		}
	}
	return true, nil
}
