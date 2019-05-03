package Zendesk

import (
	"mega/atendimento/DashBoardSync/Zendesk/Group"
	"mega/atendimento/DashBoardSync/Zendesk/Organization"
	"mega/atendimento/DashBoardSync/Zendesk/Ticket"
	"mega/atendimento/DashBoardSync/Zendesk/User"
)

// ---------------------- Ticket ----------------------------

func GetIncExportTicket(tempoInicio string, once bool) {
	Ticket.GetIncExportTickets(tempoInicio, once)
}

func GetLastDataTimeTicket() string {
	return Ticket.GetLastDataTime()
}

func SetTicketLinkIssue() {
	Ticket.SetLinkIssueTicketOracle()
}

// ---------------------- Organization ----------------------------

func GetIncExportOrganization(tempoInicio string, once bool) {
	Organization.GetIncExportOrganization(tempoInicio, once)
}

func GetLastDataTimeOrganization() string {
	return Organization.GetLastDataTime()
}

// ---------------------- User ----------------------------

func GetIncExportUser(tempoInicio string, once bool) {
	User.GetIncExportUser(tempoInicio, once)
}

func GetLastDataTimeUser() string {
	return User.GetLastDataTime()
}

// ---------------------- Group ----------------------------

func GetIncExportGroup(tempoInicio string, once bool) {
	Group.GetIncExportGroup(tempoInicio, once)
}
