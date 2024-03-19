package rest

import (
	"net/http"

	"github.com/Syahreza-Ferdian/heal-in/model"
	"github.com/Syahreza-Ferdian/heal-in/pkg/web/response"
	"github.com/gin-gonic/gin"
)

func (r *Rest) NewEvent(ctx *gin.Context) {
	var eventReq model.NewEventRequest

	err := ctx.Bind(&eventReq)
	if err != nil {
		response.OnFailed(ctx, http.StatusBadRequest, "failed to bind event request", err)
		return
	}

	event, err := r.service.EventService.NewEvent(eventReq)
	if err != nil {
		response.OnFailed(ctx, http.StatusInternalServerError, "failed to create event", err)
		return
	}

	form, _ := ctx.MultipartForm()
	imgs := form.File["image"]

	for _, fileImage := range imgs {
		err := r.service.EventService.EventImageUploader(event.ID.String(), model.EventImageUploadParam{
			Image: fileImage,
		})

		if err != nil {
			response.OnFailed(ctx, http.StatusInternalServerError, "failed to upload image", err)
			return
		}
	}

	response.OnSuccess(ctx, http.StatusCreated, "event created", event)
}

func (r *Rest) GetEventByID(ctx *gin.Context) {
	id := ctx.Param("id")

	event, err := r.service.EventService.GetEvent(id)
	if err != nil {
		response.OnFailed(ctx, http.StatusInternalServerError, "failed to get event", err)
		return
	}

	response.OnSuccess(ctx, http.StatusOK, "event found", event)
}

func (r *Rest) GetAllEvents(ctx *gin.Context) {
	events, err := r.service.EventService.GetAllEvents()
	if err != nil {
		response.OnFailed(ctx, http.StatusInternalServerError, "failed to get events", err)
		return
	}

	response.OnSuccess(ctx, http.StatusOK, "events found", events)
}
