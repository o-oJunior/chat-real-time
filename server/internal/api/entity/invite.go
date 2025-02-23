package entity

import "fmt"

type Invite struct {
	ID                    string `json:"id,omitempty" bson:"_id,omitempty"`
	InviteStatus          string `json:"inviteStatus,omitempty" bson:"inviteStatus,omitempty"`
	InvitedAtMilliseconds int64  `json:"omitempty" bson:"invitedAt,omitempty"`
	InvitedAt             string `json:"invitedAt,omitempty" bson:"omitempty"`
	UserIdInvited         string `json:"userIdInvited,omitempty" bson:"userIdInvited,omitempty"`
	UserIdInviter         string `json:"userIdInviter,omitempty" bson:"userIdInviter,omitempty"`
}

func (invite *Invite) ValidateRegisterInvite() error {
	if invite.UserIdInvited == "" || invite.InvitedAt == "" || invite.InviteStatus == "" {
		return fmt.Errorf("o corpo da requisição está vazio ou mal formatado")
	}
	return nil
}
