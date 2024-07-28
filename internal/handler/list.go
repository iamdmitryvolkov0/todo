package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo/internal/domain"
)

func (h *Handler) createList(c *gin.Context) {
	userID, err := h.getUserId(c)
	if err != nil {
		return
	}

	var input domain.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoList.Create(userID, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type allListsResponse struct {
	Data []domain.TodoList `json:"data"`
}

func (h *Handler) allLists(c *gin.Context) {
	userID, err := h.getUserId(c)
	if err != nil {
		return
	}

	lists, err := h.services.TodoList.All(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, allListsResponse{
		Data: lists,
	})
}

func (h *Handler) listById(c *gin.Context) {
	userID, err := h.getUserId(c)
	if err != nil {
		return
	}

	listID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	list, err := h.services.TodoList.ListByID(userID, listID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) updateList(c *gin.Context) {
	userID, err := h.getUserId(c)
	if err != nil {
		return
	}

	listID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
	}

	var input domain.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.TodoList.Update(userID, listID, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "updated",
	})
}

func (h *Handler) deleteList(c *gin.Context) {
	userID, err := h.getUserId(c)
	if err != nil {
		return
	}

	listID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	err = h.services.TodoList.Delete(userID, listID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "deleted",
	})
}
