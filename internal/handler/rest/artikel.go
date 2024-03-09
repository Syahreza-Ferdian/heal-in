package rest

import (
	"net/http"

	"github.com/Syahreza-Ferdian/heal-in/model"
	"github.com/Syahreza-Ferdian/heal-in/pkg/web/response"
	"github.com/gin-gonic/gin"
)

func (r *Rest) NewArtikel(ctx *gin.Context) {
	// err := ctx.Request.ParseMultipartForm(32 << 20)
	// if err != nil {
	// 	response.OnFailed(ctx, http.StatusInternalServerError, "failed to parse multipart form", err)
	// 	return
	// }

	var artikelReq model.NewArtikelRequest

	err := ctx.Bind(&artikelReq)

	if err != nil {
		response.OnFailed(ctx, http.StatusBadRequest, "failed to bind artikel request", err)
		return
	}

	artikel, err := r.service.ArtikelService.NewArtikel(artikelReq)

	if err != nil {
		response.OnFailed(ctx, http.StatusInternalServerError, "failed to create artikel", err)
		return
	}

	form, _ := ctx.MultipartForm()
	imgs := form.File["image"]

	for _, file := range imgs {
		err := r.service.ArtikelService.UploadArtikelImage(artikel.ID.String(), model.ArtikelUploadImageParam{
			Image: file,
		})

		if err != nil {
			response.OnFailed(ctx, http.StatusInternalServerError, "failed to upload image", err)
			return
		}
	}

	response.OnSuccess(ctx, http.StatusCreated, "artikel created", artikel)
}

func (r *Rest) GetArtikel(ctx *gin.Context) {
	id := ctx.Param("id")

	artikel, err := r.service.ArtikelService.GetArtikel(id)

	if err != nil {
		response.OnFailed(ctx, http.StatusInternalServerError, "failed to get artikel", err)
		return
	}

	response.OnSuccess(ctx, http.StatusOK, "artikel found", artikel)
}

func (r *Rest) GetAllArtikel(ctx *gin.Context) {
	artikels, err := r.service.ArtikelService.GetAllArtikel(ctx)

	if err != nil {
		response.OnFailed(ctx, http.StatusInternalServerError, "failed to get artikel", err)
		return
	}

	response.OnSuccess(ctx, http.StatusOK, "artikel found", artikels)
}

func (r *Rest) GetFewSampleArtikel(ctx *gin.Context) {
	artikels, err := r.service.ArtikelService.GetFewSampleArtikel()

	if err != nil {
		response.OnFailed(ctx, http.StatusInternalServerError, "failed to get artikel", err)
		return
	}

	response.OnSuccess(ctx, http.StatusOK, "artikel found", artikels)
}
