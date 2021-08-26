package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username        string `json:"username" gorm:"primary_key"`
	Nickname        string `json:"nickname"`
	Email           string `json:"email" gorm:"primary_key"`
	DisplayEmail    string `json:"displayEmail"`
	PasswdHash      string `json:"passwdHash"`
	Avatar          string `json:"avatar"`
	AvatarUrl       string `json:"avatar_url"`
	NoteNum         uint   `json:"note_num"`
	CurrentNoteNum  uint   `json:"current_note_num"`
	GiveUpNoteNum   uint   `json:"give_up_note_num"`
	CompleteNoteNum uint   `json:"complete_note_num"`
}
