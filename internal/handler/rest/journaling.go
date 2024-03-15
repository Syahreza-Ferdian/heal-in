package rest

import (
	"fmt"
	"net/http"

	"github.com/Syahreza-Ferdian/heal-in/entity"
	"github.com/Syahreza-Ferdian/heal-in/model"
	"github.com/Syahreza-Ferdian/heal-in/pkg/web/response"
	"github.com/gin-gonic/gin"
)

func (r *Rest) NewJournalingEntry(ctx *gin.Context) {
	currUser, ada := ctx.Get("user")

	if !ada {
		response.OnFailed(ctx, http.StatusUnauthorized, "Unauthorized", fmt.Errorf("you need to logged in to access this feature"))
		return
	}

	var entryReq model.JournalingEntryReq

	entryReq.UserID = currUser.(*entity.User).ID

	err := ctx.BindJSON(&entryReq)
	if err != nil {
		response.OnFailed(ctx, http.StatusBadRequest, "Failed to bind request", err)
		return
	}

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

func (r *Rest) GetCurrentUserJournalingEntries(ctx *gin.Context) {
	currUser, ada := ctx.Get("user")

	if !ada {
		response.OnFailed(ctx, http.StatusUnauthorized, "Unauthorized", fmt.Errorf("you need to logged in to access this feature"))
		return
	}

	userID := currUser.(*entity.User).ID

	entries, err := r.service.JournalingService.GetJournalingEntriesByUserID(userID.String())

	if err != nil {
		response.OnFailed(ctx, http.StatusInternalServerError, "Failed to get journaling entries", err)
		return
	}

	response.OnSuccess(ctx, http.StatusOK, "Journaling entries retrieved", entries)
}
