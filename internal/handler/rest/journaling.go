package rest

import (
	"net/http"

	"github.com/Syahreza-Ferdian/heal-in/entity"
	"github.com/Syahreza-Ferdian/heal-in/model"
	"github.com/Syahreza-Ferdian/heal-in/pkg/web/response"
	"github.com/gin-gonic/gin"
)

func (r *Rest) NewJournalingEntry(ctx *gin.Context) {
	currUser, ada := ctx.Get("user")

	if !ada {
		response.OnFailed(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	var entryReq model.JournalingEntryReq

	entryReq.UserID = currUser.(*entity.User).ID

	err := ctx.BindJSON(&entryReq)
	if err != nil {
		response.OnFailed(ctx, http.StatusBadRequest, "Failed to bind request", err)
		return
	}

	// for question_key, answers := range entryReq.Answers {
	// 	log.Printf("Question: %s\n", question_key)
	// 	for _, answer := range answers {
	// 		log.Printf("Answer: %s\n", answer.Answer)
	// 	}
	// }

	newEntry, err := r.service.JournalingService.NewJournalingEntry(entryReq)

	if err != nil {
		response.OnFailed(ctx, http.StatusInternalServerError, "Failed to create new journaling entry", err)
		return
	}

	response.OnSuccess(ctx, http.StatusCreated, "Journaling entry created", newEntry)
}

func (r *Rest) GetJournalingEntryByID(ctx *gin.Context) {
	entryID := ctx.Param("id")

	entry, err := r.service.JournalingService.GetJournalingEntryByID(entryID)

	if err != nil {
		response.OnFailed(ctx, http.StatusInternalServerError, "Failed to get journaling entry", err)
		return
	}

	response.OnSuccess(ctx, http.StatusOK, "Journaling entry retrieved", entry)
}
