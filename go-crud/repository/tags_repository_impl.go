package repository

import (
	"errors"
	"golang-crud-gin/data/request"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"

	"gorm.io/gorm"
)

type TagsRepositoryImpl struct {
	Db *gorm.DB
}

func NewTagsRepositoryImpl(db *gorm.DB) TagsRepository {
	return &TagsRepositoryImpl{Db: db}
}

//Delete implements TagsRepository
func (t *TagsRepositoryImpl) Delete(tagsId int) {
	var tags model.Tags
	result :=t.Db.Where("id = ?",tagsId).Delete(&tags)
	helper.ErrorPanic(result.Error)
}

//Find all implements tagsRepository
func(t *TagsRepositoryImpl) FindAll() []model.Tags {
	var tags []model.Tags
	result :=t.Db.Find(&tags)
	helper.ErrorPanic(result.Error)
	return tags
}

//Find by ID implements tagsRepository
func(t *TagsRepositoryImpl) FindById(tagsId int ) (tags model.Tags, err error) {
	var tag model.Tags
	result :=t.Db.Find(&tag, tagsId)
	if result != nil{
		return tag,nil
	}else{
		return tag, errors.New("tag is not found")
	}
	
}

//save implements TagsRepository
func(t *TagsRepositoryImpl) Save (tags model.Tags) {
	result := t.Db.Create(&tags)
	helper.ErrorPanic(result.Error)
}

//update implements TagsRepository
func(t *TagsRepositoryImpl) Update(tags model.Tags) {
	var updateTag = request.UpdateTagsRequest{
		Id: tags.Id,
		Name: tags.Name,
	}
	result := t.Db.Model(&tags).Updates(updateTag)
	helper.ErrorPanic(result.Error)
}