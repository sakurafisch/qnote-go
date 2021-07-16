package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username        string `json:"username"`
	Nickname        string `json:"nickname"`
	Email           string `json:"email"`
	DisplayEmail    string `json:"displayEmail"`
	Password        string `json:"password"`
	Avatar          string `json:"avatar"`
	AvatarUrl       string `json:"avatar_url"`
	NoteNum         uint   `json:"note_num"`
	CurrentNoteNum  uint   `json:"current_note_num"`
	GiveUpNoteNum   uint   `json:"give_up_note_num"`
	CompleteNoteNum uint   `json:"complete_note_num"`
}

func (user *User) Init(username string, password string) {
	user.Username = username
	user.Password = password
}
