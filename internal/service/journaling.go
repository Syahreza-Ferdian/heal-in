package service

import (
	"fmt"
	"strconv"

	"github.com/Syahreza-Ferdian/heal-in/entity"
	"github.com/Syahreza-Ferdian/heal-in/internal/repository"
	"github.com/Syahreza-Ferdian/heal-in/model"
	"github.com/google/uuid"
)

type InterfaceJournalingService interface {
	NewJournalingEntry(journalingReq model.JournalingEntryReq) (*entity.JournalingEntry, error)
	GetJournalingEntryByID(id string) (model.JournalingEntryResponse, error)
	GetJournalingEntriesByUserID(userID string) ([]model.JournalingEntryResponse, error)
}

type JournalingService struct {
	ja repository.InterfaceJournalingAnsRepository
	je repository.InterfaceJournalingEntryRepository
	jq repository.InterfaceJournalingQuestionRepository
	ur repository.InterfaceUserRepository
}

func NewJournalingService(ja repository.InterfaceJournalingAnsRepository, je repository.InterfaceJournalingEntryRepository, jq repository.InterfaceJournalingQuestionRepository, ur repository.InterfaceUserRepository) InterfaceJournalingService {
	return &JournalingService{
		ja: ja,
		je: je,
		jq: jq,
		ur: ur,
	}
}

func (js *JournalingService) NewJournalingEntry(journalingReq model.JournalingEntryReq) (*entity.JournalingEntry, error) {
	var subsStatus int
	var journalingEntryCount int

	userID := journalingReq.UserID

	if userID == uuid.Nil {
		return nil, fmt.Errorf("cannot find current user")
	}

	subsStatus, err := js.ur.GetUserSubscriptionStatus(userID.String())
	if err != nil {
		return nil, err
	}

	journalingEntryCount, err = js.ur.GetUserJournalingCount(userID.String())
	if err != nil {
		return nil, err
	}

	if subsStatus == 0 && journalingEntryCount >= 3 {
		return nil, fmt.Errorf("user has reached maximum journaling limit. Subscribe to our feature to get unlimited journaling")
	}

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

	user := &entity.User{
		JournalingEntryCount: journalingEntryCount + 1,
	}

	err = js.ur.UpdateUserColoumn("journaling_entry_count", user, userID.String())
	if err != nil {
		return nil, err
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
		MoodID:    entry.Mood,
	}

	return response, nil
}

func (js *JournalingService) GetJournalingEntriesByUserID(userID string) ([]model.JournalingEntryResponse, error) {
	entries, err := js.je.GetJournalingEntriesByUserID(userID)

	if err != nil {
		return nil, err
	}

	var responses []model.JournalingEntryResponse

	for _, entry := range entries {
		perQuestionAnswersMap := make(map[int][]string)

		for _, ans := range entry.Answers {
			perQuestionAnswersMap[ans.QuestionID] = append(perQuestionAnswersMap[ans.QuestionID], ans.Answer)
		}

		answers := make([]model.JournalingAnswerResponse, 0)

		for questionID, ans := range perQuestionAnswersMap {
			questionText, err := js.jq.GetJournalingQuestionText(questionID)
			if err != nil {
				return nil, err
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
			MoodID:    entry.Mood,
		}

		responses = append(responses, response)
	}

	return responses, nil
}
