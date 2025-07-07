package rest

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sachatarba/course-db/internal/delivery/v2/dto"
	"github.com/sachatarba/course-db/internal/entity"
	"github.com/sachatarba/course-db/internal/service"
)

type ClientHandler struct {
	clientService service.IClientService
}

func NewClientHandler(clientService service.IClientService) *ClientHandler {
	return &ClientHandler{
		clientService: clientService,
	}
}

func (h *ClientHandler) GetClients(ctx *gin.Context) {
	log.Print("GetClients:", ctx.Request)

	clients, err := h.clientService.ListClients(ctx.Request.Context())
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"clients": clients})
}

func (h *ClientHandler) GetClientByLogin(ctx *gin.Context) {
	log.Print("GetClientByLogin:", ctx.Request)

	login := ctx.Param("login")

	client, err := h.clientService.GetClientByLogin(ctx.Request.Context(), login)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"client": client})
}

func (h *ClientHandler) GetClientByID(ctx *gin.Context) {
	id := ctx.Param("clientId")

	uuID, err := uuid.Parse(id)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})

		return
	}

	client, err := h.clientService.GetClientByID(ctx.Request.Context(), uuID)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"client": client})
}

func (h *ClientHandler) PutClient(ctx *gin.Context) {
	log.Print("PutClient: ", ctx.Request)

	var req dto.PutClient
	err := ctx.BindJSON(&req)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})

		return
	}

	id := ctx.Param("clientId")

	uuID, err := uuid.Parse(id)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}

	err = h.clientService.ChangeClient(ctx.Request.Context(), entity.Client{
		ID:        uuID,
		Login:     req.Login,
		Password:  req.Password,
		Fullname:  req.Fullname,
		Email:     req.Email,
		Phone:     req.Phone,
		Birthdate: req.Birthdate,
	})
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})

		return
	}

	ctx.Status(http.StatusOK)
}

func (h *ClientHandler) DeleteClient(ctx *gin.Context) {
	log.Print("DeleteClient: ", ctx.Request)

	id := ctx.Param("clientId")

	uuID, err := uuid.Parse(id)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})

		return
	}

	err = h.clientService.DeleteClient(ctx.Request.Context(), uuID)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})

		return
	}

	ctx.Status(http.StatusOK)
}
