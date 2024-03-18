package rest

import (
	"fmt"
	"net/http"

	"github.com/Syahreza-Ferdian/heal-in/entity"
	"github.com/Syahreza-Ferdian/heal-in/model"
	"github.com/Syahreza-Ferdian/heal-in/pkg/web/response"
	"github.com/gin-gonic/gin"
)

func (r *Rest) NewPodcast(ctx *gin.Context) {
	var podcastReq model.NewPodcastRequest

	err := ctx.Bind(&podcastReq)
	if err != nil {
		response.OnFailed(ctx, http.StatusBadRequest, "failed to bind podcast request", err)
		return
	}

	podcastFile, err := ctx.FormFile("podcast")
	if err != nil {
		response.OnFailed(ctx, http.StatusBadRequest, "failed to get podcast file", err)
		return
	}

	podcastReq.Podcast = podcastFile

	thumbnailFile, err := ctx.FormFile("thumbnail")
	if err != nil {
		response.OnFailed(ctx, http.StatusBadRequest, "failed to get thumbnail file", err)
		return
	}

	podcastReq.Thumbnail = thumbnailFile

	podcast, err := r.service.PodcastService.NewPodcast(podcastReq)
	if err != nil {
		response.OnFailed(ctx, http.StatusInternalServerError, "failed to create a new podcast", err)
		return
	}

	response.OnSuccess(ctx, http.StatusCreated, "podcast created", podcast)
}

func (r *Rest) GetAllPodcastsBasedOnUserStatus(ctx *gin.Context) {
	currUser, ada := ctx.Get("user")
	if !ada {
		response.OnFailed(ctx, http.StatusUnauthorized, "user not found", fmt.Errorf("you need to logged in to access this feature"))
		return
	}

	userID := currUser.(*entity.User).ID

	podcast, err := r.service.PodcastService.GetAllPodcastsBasedOnUserStatus(userID.String())
	if err != nil {
		response.OnFailed(ctx, http.StatusInternalServerError, "failed to get podcasts", err)
		return
	}

	response.OnSuccess(ctx, http.StatusOK, "podcasts found", podcast)
}

func (r *Rest) GetPodcastByID(ctx *gin.Context) {
	podcastID := ctx.Param("id")

	podcast, err := r.service.PodcastService.GetPodcastByID(podcastID)
	if err != nil {
		response.OnFailed(ctx, http.StatusInternalServerError, "failed to get podcast", err)
		return
	}

	response.OnSuccess(ctx, http.StatusOK, "podcast found", podcast)
}
