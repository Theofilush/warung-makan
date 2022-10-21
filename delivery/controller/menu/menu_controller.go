package menu

import (
	"database/sql"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Theofilush/warung-makan/delivery/middleware"
	_ "github.com/Theofilush/warung-makan/docs"
	"github.com/Theofilush/warung-makan/model"

	useCaseMen "github.com/Theofilush/warung-makan/usecase/menu"
	"github.com/Theofilush/warung-makan/utils"
	"github.com/gin-gonic/gin"
)

type MenuController struct {
	rgg         *gin.RouterGroup
	menuUsecase useCaseMen.MenuUsecase
}

func (cc *MenuController) userAuth(ctx *gin.Context) {
	var user model.UserCredential
	if err2 := ctx.ShouldBindJSON(&user); err2 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "can't bind struct",
		})
		return
	}

	token, err := cc.menuUsecase.UserAuth(user)

	if err != nil {
		ctx.AbortWithStatus(401)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (cc *MenuController) getAllMenu(ctx *gin.Context) {
	menus, err := cc.menuUsecase.GetAllMenu()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, menus)
}
func (cc *MenuController) getMenuById(ctx *gin.Context) {
	id := ctx.Param("id")
	menus, err := cc.menuUsecase.FindMenuById(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, menus)
}

func (cc *MenuController) getMenuImage(ctx *gin.Context) {
	param := ctx.Param("image")
	menus, err := cc.menuUsecase.FindMenuImage(param)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	ctx.File("files/" + param)

	ctx.JSON(http.StatusOK, menus)
}

func (cc *MenuController) registerMenu(ctx *gin.Context) {
	var menu model.Menu

	menu.Menu_name = ctx.PostForm("menu_name")
	menu.Price, _ = strconv.Atoi(ctx.PostForm("price"))
	file, _ := ctx.FormFile("images")
	fileMenuName := utils.UuidGenerate() + " - " + file.Filename
	menu.Image = sql.NullString{String: fileMenuName, Valid: true}

	dir, err := os.Getwd()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err1": err.Error()})
		return
	}
	fileLocation := filepath.Join(dir, "files", fileMenuName)
	ctx.SaveUploadedFile(file, fileLocation)

	if err := cc.menuUsecase.RegisterMenu(menu); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err3": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, menu)
}

func (cc *MenuController) UpdateMenu(ctx *gin.Context) {
	var menu model.Menu
	menu.Id = ctx.PostForm("id")
	menu.Menu_name = ctx.PostForm("menu_name")
	menu.Price, _ = strconv.Atoi(ctx.PostForm("price"))
	file, _ := ctx.FormFile("images")
	fileMenuName := utils.UuidGenerate() + " - " + file.Filename
	menu.Image = sql.NullString{String: fileMenuName, Valid: true}

	dir, err := os.Getwd()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err1": err.Error()})
		return
	}
	fileLocation := filepath.Join(dir, "files", fileMenuName)
	ctx.SaveUploadedFile(file, fileLocation)

	if err := cc.menuUsecase.UpdateMenu(menu); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, menu)
}

func (cc *MenuController) DeleteMenu(ctx *gin.Context) {
	id := ctx.Param("id")
	err := cc.menuUsecase.DeleteMenu(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, err)
}

func NewMenuController(routerGroup *gin.RouterGroup, usecaseMen useCaseMen.MenuUsecase, tokenMdw middleware.AuthTokenMiddleware) *MenuController {
	controllerr := MenuController{
		rgg:         routerGroup,
		menuUsecase: usecaseMen,
	}

	// controllerr.rgg.POST("/auth", controllerr.userAuth)

	protectedGroup := controllerr.rgg.Group("/private", tokenMdw.RequireToken())

	protectedGroup.GET("/menu", controllerr.getAllMenu)
	protectedGroup.GET("/menu/:id", controllerr.getMenuById)
	protectedGroup.POST("/menu", controllerr.registerMenu)
	protectedGroup.PUT("/menu", controllerr.UpdateMenu)
	protectedGroup.DELETE("menu/:id", controllerr.DeleteMenu)
	controllerr.rgg.GET("/menu/images/:image", controllerr.getMenuImage)
	// protectedGroup.POST("upload", controllerr.UploadMenu)

	return &controllerr
}
