package main

import (
	"fmt"
)

// OrganizingTeam represents the team organizing the event
type OrganizingTeam struct {
	TeamMembers    []string
	PrimaryContact string
}

// NSPEvent represents an event managed by the NSP (New Service Provider)
type NSPEvent struct {
	EventDate       string
	EventName       string
	PrimaryContact  string
	OrganizingTeam  OrganizingTeam // Embedded OrganizingTeam
}

// AddTeamMember adds a new team member to the organizing team
func (ot *OrganizingTeam) AddTeamMember(member string) {
	ot.TeamMembers = append(ot.TeamMembers, member)
}

// SetPrimaryContact sets the primary contact for the organizing team
func (ot *OrganizingTeam) SetPrimaryContact(contact string) {
	ot.PrimaryContact = contact
}

// NewNSPEvent creates a new NSPEvent instance
func NewNSPEvent(date, name, contact string, teamMembers []string) *NSPEvent {
	event := &NSPEvent{
		EventDate:       date,
		EventName:       name,
		PrimaryContact:  contact,
	}

	// Initialize the organizing team
	event.OrganizingTeam = OrganizingTeam{
		TeamMembers:    teamMembers,
		PrimaryContact: contact, // Assuming primary contact is initially the same as the event contact
	}

	return event
}

func main() {
	// Create a new NSP event
	event := NewNSPEvent("2024-02-01", "NSP Conference", "John Doe", []string{"Alice", "Bob", "Charlie"})

	// Add a new team member
	event.OrganizingTeam.AddTeamMember("David")

	// Set a new primary contact for the organizing team
	event.OrganizingTeam.SetPrimaryContact("Alice")

	// Display event details
	fmt.Println("Event Details:")
	fmt.Println("Date:", event.EventDate)
	fmt.Println("Name:", event.EventName)
	fmt.Println("Primary Contact:", event.PrimaryContact)
	fmt.Println("Organizing Team Members:", event.OrganizingTeam.TeamMembers)
	fmt.Println("Organizing Team Primary Contact:", event.OrganizingTeam.PrimaryContact)
}
