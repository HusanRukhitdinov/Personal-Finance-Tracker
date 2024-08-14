package handler

import (
	"auth/api/token"
	pb "auth/genproto/user"
	"auth/service"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticaionHandler interface {
	RegisterUser(c *gin.Context)
	LoginUser(c *gin.Context)
	Logout(c *gin.Context)
	// UserProfile(c *gin.Context)
	// UpdateUserProfile(c *gin.Context)
}

type AuthenticaionHandlerImpl struct {
	UserManage service.UserManagementService
	Logger     *slog.Logger
}

func NewAuthenticaionHandler(userManage service.UserManagementService, logger *slog.Logger) AuthenticaionHandler {
	return &AuthenticaionHandlerImpl{UserManage: userManage, Logger: logger}
}

// @Summary      Register a new user
// @Description  Register a new user account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body     user.RegisterRequest  true  "User registration request"
// @Success      200   {object}  user.RegisterResponse
// @Failure      400   {object}  map[string]string    "Bad request"
// @Failure      500   {object}  map[string]string    "Internal server error"
// @Router       /register [post]
func (h *AuthenticaionHandlerImpl) RegisterUser(c *gin.Context) {
	req := pb.RegisterRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := h.UserManage.RegisterUser(c, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}


// @Summary      Login a new user
// @Description  Login a new user account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body     user.LoginRequest  true  "User login request"
// @Success      200   {object}  user.LoginResponse
// @Failure      400   {object}  map[string]string    "Bad request"
// @Failure      500   {object}  map[string]string    "Internal server error"
// @Router       /login [post]
func (h *AuthenticaionHandlerImpl) LoginUser(c *gin.Context) {
	req := pb.LoginRequest{}
	
	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := h.UserManage.LoginUser(c, &req)
	if err != nil {
		fmt.Println("^%$22222222",err)
		fmt.Println(resp)
		h.Logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	tokenString := token.GenerateJWT(resp)

	c.JSON(http.StatusOK, tokenString)
}

// @Summary      Logout a new user
// @Description  Logout a new user account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body     user.LogoutRequest  true  "User logout request"
// @Success      200   {object}  user.Message
// @Failure      400   {object}  map[string]string    "Bad request"
// @Failure      500   {object}  map[string]string    "Internal server error"
// @Router      /logout [post]
func (h *AuthenticaionHandlerImpl) Logout(c *gin.Context) {
	req := pb.LogoutRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}
	
	resp, err := h.UserManage.Logout(c, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	
	c.JSON(http.StatusOK, resp)
}

// // @Summary      UserProfile a user
// // @Description  Retrieve a user profile by user ID
// // @Tags         auth
// // @Accept       json
// // @Produce      json
// // @Param        id  path     string  true  "User ID"
// // @Success      200   {object}  user.ProfileResponse
// // @Failure      400   {object}  map[string]string    "Bad request"
// // @Failure      500   {object}  map[string]string    "Internal server error"
// // @Router       /userprofile/{id} [GET]
// func (h *AuthenticaionHandlerImpl) UserProfile(c *gin.Context) {
//     id := c.Param("id") 
// 	fmt.Println("))))",id)
//     if id == "" {
//         errMsg := "User ID is required"
//         h.Logger.Error(errMsg)
//         c.JSON(http.StatusBadRequest, map[string]string{"error": errMsg})
//         return
//     }
	
//     req := &pb.ProfileRequest{Id: id}
	
// 	fmt.Println("+++",req)
//     resp, err := h.UserManage.GetUserProfile(c, req)
//     if err != nil {
//         h.Logger.Error(err.Error())
//         c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
//         return
//     }

//     c.JSON(http.StatusOK, resp)
// }


// // @Summary      Update user Profile
// // @Description  Update user Profile
// // @Tags         auth
// // @Accept       json
// // @Produce      json
// // @Param        user  body     user.ProfileUpdateRequest  true  "User update profile request"
// // @Success      200   {object}  user.ProfileUpdateResponse
// // @Failure      400   {object}  map[string]string    "Bad request"
// // @Failure      500   {object}  map[string]string    "Internal server error"
// // @Router       /updateprofile [PUT]
// func (h *AuthenticaionHandlerImpl) UpdateUserProfile(c *gin.Context) {
// 	req := pb.ProfileUpdateRequest{}
// 	fmt.Println("-----",&req)

// 	err := json.NewDecoder(c.Request.Body).Decode(&req)
// 	if err != nil {
// 		h.Logger.Error(err.Error())
// 		c.JSON(http.StatusBadRequest, err)
// 		return
// 	}

// 	resp, err := h.UserManage.UpdateUserProfile(c, &req)
// 	if err != nil {
// 		h.Logger.Error(err.Error())
// 		c.JSON(http.StatusInternalServerError, err)
// 		return
// 	}

// 	c.JSON(http.StatusOK, resp)
// }
