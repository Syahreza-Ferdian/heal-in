package rest

import (
	"fmt"
	"net/http"

	"github.com/Syahreza-Ferdian/heal-in/entity"
	"github.com/Syahreza-Ferdian/heal-in/model"
	"github.com/Syahreza-Ferdian/heal-in/pkg/web/response"
	"github.com/gin-gonic/gin"
)

func (r *Rest) NewVideo(ctx *gin.Context) {
	var videoReq model.NewVideoRequest

	err := ctx.Bind(&videoReq)

	if err != nil {
		response.OnFailed(ctx, http.StatusBadRequest, "failed to bind video request", err)
		return
	}

	videoFile, err := ctx.FormFile("video")

	if err != nil {
		response.OnFailed(ctx, http.StatusBadRequest, "failed to get video file", err)
		return
	}

	videoReq.Video = videoFile

	video, err := r.service.VideoService.NewVideo(videoReq)

	if err != nil {
		response.OnFailed(ctx, http.StatusInternalServerError, "failed to create video", err)
		return
	}

	response.OnSuccess(ctx, http.StatusCreated, "video created", video)
}

func (r *Rest) NewVideoWithLink(ctx *gin.Context) {
	var videoReq model.NewVideoRequestAlt

	err := ctx.Bind(&videoReq)
	if err != nil {
		response.OnFailed(ctx, http.StatusBadRequest, "failed to bind video request", err)
		return
	}

	video, err := r.service.VideoService.NewVideoWithLink(videoReq)

	if err != nil {
		response.OnFailed(ctx, http.StatusInternalServerError, "failed to create video", err)
		return
	}

	response.OnSuccess(ctx, http.StatusCreated, "video created", video)
}

func (r *Rest) GetAllVideos(ctx *gin.Context) {
	currUser, ada := ctx.Get("user")

	if !ada {
		response.OnFailed(ctx, http.StatusUnauthorized, "user not found", fmt.Errorf("you need to logged in to access this feature"))
		return
	}

	videos, err := r.service.VideoService.GetAllVideos(currUser.(*entity.User).ID.String())

	if err != nil {
		response.OnFailed(ctx, http.StatusInternalServerError, "failed to get videos", err)
		return
	}

	response.OnSuccess(ctx, http.StatusOK, "videos found", videos)
}

func (r *Rest) GetSpecificVideo(ctx *gin.Context) {
	videoID := ctx.Param("id")

	video, err := r.service.VideoService.GetVideo(videoID)

	if err != nil {
		response.OnFailed(ctx, http.StatusInternalServerError, "failed to get video", err)
		return
	}

	response.OnSuccess(ctx, http.StatusOK, "video found", video)
}
