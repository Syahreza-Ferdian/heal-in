package service

import (
	"strconv"

	"github.com/Syahreza-Ferdian/heal-in/entity"
	"github.com/Syahreza-Ferdian/heal-in/internal/repository"
	"github.com/Syahreza-Ferdian/heal-in/model"
	"github.com/google/uuid"
)

type InterfaceJournalingService interface {
	NewJournalingEntry(journalingReq model.JournalingEntryReq) (*entity.JournalingEntry, error)
	// GetJournalingEntryByID(id string) (*entity.JournalingEntry, error)
	GetJournalingEntryByID(id string) (model.JournalingEntryResponse, error)
}

type JournalingService struct {
	ja repository.InterfaceJournalingAnsRepository
	je repository.InterfaceJournalingEntryRepository
	jq repository.InterfaceJournalingQuestionRepository
}

func NewJournalingService(ja repository.InterfaceJournalingAnsRepository, je repository.InterfaceJournalingEntryRepository, jq repository.InterfaceJournalingQuestionRepository) InterfaceJournalingService {
	return &JournalingService{
		ja: ja,
		je: je,
		jq: jq,
	}
}

func (js *JournalingService) NewJournalingEntry(journalingReq model.JournalingEntryReq) (*entity.JournalingEntry, error) {
	entityJournaling := &entity.JournalingEntry{
		ID:     uuid.New(),
		UserID: journalingReq.UserID,
		Mood:   journalingReq.Mood.MoodID,
	}

	output, err := js.je.NewJournalingEntry(entityJournaling)
	if err != nil {
		return nil, err
	}

	for questionKey, answers := range journalingReq.Answers {
		for _, answer := range answers {
			questionID, _ := strconv.Atoi(questionKey)
			entityAnswer := &entity.JournalingAnswer{
				EntryID:    entityJournaling.ID,
				QuestionID: questionID,
				Answer:     answer.Answer,
			}
			_, err := js.ja.NewJournalingAns(entityAnswer)

			if err != nil {
				return nil, err
			}
		}
	}

	return output, nil
}

func (js *JournalingService) GetJournalingEntryByID(id string) (model.JournalingEntryResponse, error) {
	entry, err := js.je.GetJournalingEntryByID(id)

	if err != nil {
		return model.JournalingEntryResponse{}, err
	}

	perQuestionAnswersMap := make(map[int][]string)

	for _, ans := range entry.Answers {
		perQuestionAnswersMap[ans.QuestionID] = append(perQuestionAnswersMap[ans.QuestionID], ans.Answer)
	}

	answers := make([]model.JournalingAnswerResponse, 0)

	for questionID, ans := range perQuestionAnswersMap {
		questionText, err := js.jq.GetJournalingQuestionText(questionID)
		if err != nil {
			return model.JournalingEntryResponse{}, err
		}

		question := model.JournalingQuestionResponse{
			QuestionID:   questionID,
			QuestionText: questionText,
			Answer:       ans,
		}

		answer := model.JournalingAnswerResponse{
			Question: question,
		}

		answers = append(answers, answer)
	}

	response := model.JournalingEntryResponse{
		ID:        entry.ID,
		UserID:    entry.UserID,
		CreatedAt: entry.CreatedAt,
		Answers:   answers,
	}

	return response, nil
}
