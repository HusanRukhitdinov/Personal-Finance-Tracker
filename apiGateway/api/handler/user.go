package handler

import (
	pb "api/genproto/user"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary      UserProfile a user
// @Description  Retrieve a user profile by user ID
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        id  path     string  true  "User ID"
// @Success      200   {object}  user.ProfileResponse
// @Failure      400   {object}  map[string]string    "Bad request"
// @Failure      500   {object}  map[string]string    "Internal server error"
// @Router       /userprofile/{id} [GET]
func (h *Handler) UserProfile(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		errMsg := "User ID is required"
		log.Fatalf(errMsg)
		c.JSON(http.StatusBadRequest, map[string]string{"error": errMsg})
		return
	}

	req := &pb.ProfileRequest{Id: id}

	fmt.Println("+++", req)
	resp, err := h.services.UserService().GetUserProfile(c, req)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary      Update user Profile
// @Description  Update user Profile
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body     user.ProfileUpdateRequest  true  "User update profile request"
// @Success      200   {object}  user.ProfileUpdateResponse
// @Failure      400   {object}  map[string]string    "Bad request"
// @Failure      500   {object}  map[string]string    "Internal server error"
// @Router       /updateprofile [PUT]
func (h *Handler) UpdateUserProfile(c *gin.Context) {
	req := pb.ProfileUpdateRequest{}
	fmt.Println("-----", &req)

	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := h.services.UserService().UpdateUserProfile(c, &req)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
