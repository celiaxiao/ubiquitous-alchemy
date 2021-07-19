package application

import (
	"ubiquitous-biz-server/app/domain/entity"
	"ubiquitous-biz-server/app/domain/repository"
)

type innApp struct {
	ir repository.InnRepository
}

var _ InnAppInterface = &innApp{}

type InnAppInterface interface {
	SaveTag(*entity.Tag) (*entity.Tag, error)
	GetAllTag() ([]entity.Tag, error)
	UpdateTag(*entity.Tag) (*entity.Tag, error)
	DeleteTag(uint64) error

	SaveArticle(*entity.Article) (*entity.Article, error)
	GetAllArticle(*entity.PaginationM10) ([]entity.Article, error)
	UpdateArticle(*entity.Article) (*entity.Article, error)
	DeleteArticle(uint64) error
}

func (ia *innApp) SaveTag(tag *entity.Tag) (*entity.Tag, error) {
	return ia.ir.SaveTag(tag)
}

func (ia *innApp) GetAllTag() ([]entity.Tag, error) {
	return ia.ir.GetAllTag()
}

func (ia *innApp) UpdateTag(tag *entity.Tag) (*entity.Tag, error) {
	return ia.ir.UpdateTag(tag)
}

func (ia *innApp) DeleteTag(id uint64) error {
	return ia.ir.DeleteTag(id)
}

func (ia *innApp) SaveArticle(article *entity.Article) (*entity.Article, error) {
	return ia.ir.SaveArticle(article)
}

func (ia *innApp) GetAllArticle(pagination *entity.PaginationM10) ([]entity.Article, error) {
	return ia.ir.GetAllArticle(pagination)
}

func (ia *innApp) UpdateArticle(article *entity.Article) (*entity.Article, error) {
	return ia.ir.UpdateArticle(article)
}

func (ia *innApp) DeleteArticle(id uint64) error {
	return ia.ir.DeleteArticle(id)
}
