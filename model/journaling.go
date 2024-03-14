package model

import (
	"time"

	"github.com/google/uuid"
)

// REQUEST MODEL
type JournalingEntryReq struct {
	UserID  uuid.UUID                        `json:"user_id"`
	Answers map[string][]JournalingAnswerReq `json:"answers"`
	Mood    MoodReq                          `json:"mood"`
}

type JournalingAnswerReq struct {
	Answer string `json:"answer"`
}

type MoodReq struct {
	MoodID int `json:"mood_id"`
}

// RESPONSE MODEL
type JournalingEntryResponse struct {
	ID        uuid.UUID                  `json:"id"`
	UserID    uuid.UUID                  `json:"user_id"`
	CreatedAt time.Time                  `json:"created_at"`
	Answers   []JournalingAnswerResponse `json:"answers"`
}

type JournalingAnswerResponse struct {
	Question JournalingQuestionResponse `json:"question"`
}

type JournalingQuestionResponse struct {
	QuestionID   int      `json:"question_id"`
	QuestionText string   `json:"question_text"`
	Answer       []string `json:"question_answers"`
}
