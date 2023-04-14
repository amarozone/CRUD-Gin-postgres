package main

import (
	"golang-crud-gin/config"
	"golang-crud-gin/controller"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"
	"golang-crud-gin/repository"
	"golang-crud-gin/router"
	"golang-crud-gin/service"
	"net/http"

	// "github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Started Server!")

	//Database
	db :=  config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})


	//Repository
	tagsRepository := repository.NewTagsRepositoryImpl(db)


	//Service 
	tagsService := service.NewTagsServiceImpl(tagsRepository,validate)


	//controller
	TagsController:=controller.NewTagsController(tagsService)


	//Router
	routes := router.NewRouter(TagsController)

	server := &http.Server{
		Addr: ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}