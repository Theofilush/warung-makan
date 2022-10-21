package menu

import (
	"github.com/Theofilush/warung-makan/model"
	"github.com/Theofilush/warung-makan/repository/menu"
	"github.com/Theofilush/warung-makan/utils/authenticator"
)

type MenuUsecase interface {
	UserAuth(user model.UserCredential) (token string, err error)
	RegisterMenu(menu model.Menu) error
	FindMenuById(id string) (model.Menu, error)
	GetAllMenu() ([]model.Menu, error)
	UpdateMenu(menu model.Menu) error
	DeleteMenu(id string) error
	FindMenuImage(image string) (model.Menu2, error)
}
type menuUsecase struct {
	repo         menu.MenuRepository
	tokenService authenticator.AccessToken
}

func (c *menuUsecase) UserAuth(user model.UserCredential) (token string, err error) {
	if user.Username == "enigma" && user.Password == "123" {
		token, err := c.tokenService.CreateAccessToken(&user)
		if err != nil {
			return "", err
		}
		return token, nil
	} else {
		return "", nil
	}
}

func (c *menuUsecase) RegisterMenu(menu model.Menu) error {
	return c.repo.Create(menu)
}

func (c *menuUsecase) FindMenuById(id string) (model.Menu, error) {
	return c.repo.FindById(id)
}

func (c *menuUsecase) FindMenuImage(image string) (model.Menu2, error) {
	return c.repo.FindImage(image)
}

func (c *menuUsecase) GetAllMenu() ([]model.Menu, error) {
	return c.repo.RetrieveAll()
}

func (c *menuUsecase) UpdateMenu(menu model.Menu) error {
	return c.repo.Update(menu)
}

func (c *menuUsecase) DeleteMenu(id string) error {
	return c.repo.Delete(id)
}

func NewMenuUseCase(repo menu.MenuRepository, service authenticator.AccessToken) MenuUsecase {
	menuUsecase := new(menuUsecase)
	menuUsecase.tokenService = service
	menuUsecase.repo = repo

	return menuUsecase
}
