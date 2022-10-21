package menu

import (
	"database/sql"

	"github.com/Theofilush/warung-makan/model"
)

type MenuRepository interface {
	Create(newMenu model.Menu) error
	RetrieveAll() ([]model.Menu, error)
	FindById(id string) (model.Menu, error)
	Delete(id string) error
	Update(menu model.Menu) error
	FindImage(image string) (model.Menu2, error)
}

type menuDbRepository struct {
	db *sql.DB
}

func (c *menuDbRepository) Create(newMenu model.Menu) error {
	insertStatement := "insert into m_menu(menu_name, price, images) values($1,$2,$3)"
	_, err := c.db.Exec(insertStatement, newMenu.Menu_name, newMenu.Price, newMenu.Image)
	if err != nil {
		return err
	}
	return nil
}

func (c *menuDbRepository) RetrieveAll() ([]model.Menu, error) {
	rows, err := c.db.Query("select id, menu_name, price, images from m_menu")
	if err != nil {
		return nil, err
	}
	var menus []model.Menu
	for rows.Next() {
		var id string
		var menuName string
		var price int
		var image sql.NullString
		err = rows.Scan(&id, &menuName, &price, &image)
		if err != nil {
			return nil, err
		}

		menu := model.NewMenu(id, menuName, image, price)
		menus = append(menus, menu)
	}
	return menus, nil
}

func (c *menuDbRepository) FindById(id string) (model.Menu, error) {
	rows, err := c.db.Query("select id, menu_name, price, images from m_menu where id=$1", id)
	if err != nil {
		return model.Menu{}, err
	}
	var menu model.Menu
	for rows.Next() {
		var id string
		var menu_name string
		var price int
		var image sql.NullString
		err = rows.Scan(&id, &menu_name, &price, &image)
		if err != nil {
			panic(err)
		}

		menu = model.NewMenu(id, menu_name, image, price)
	}
	return menu, nil
}

func (c *menuDbRepository) FindImage(image string) (model.Menu2, error) {
	rows, err := c.db.Query("select images from m_menu where images=$1", image)
	if err != nil {
		return model.Menu2{}, err
	}
	var menu model.Menu2
	for rows.Next() {
		var image sql.NullString
		err = rows.Scan(&image)
		if err != nil {
			panic(err)
		}

		menu = model.NewMenu2(image)
	}
	return menu, nil
}

func (c *menuDbRepository) Delete(id string) error {
	deleteStatement := "delete from m_menu where id = $1"
	_, err := c.db.Exec(deleteStatement, id)

	if err != nil {
		return err
	}
	return nil
}

func (c *menuDbRepository) Update(menu model.Menu) error {
	updateStatement := "update m_menu set menu_name=$2, price=$3, images=$4 where id = $1"
	_, err := c.db.Exec(updateStatement, menu.Id, menu.Menu_name, menu.Price, menu.Image)

	if err != nil {
		return err
	}
	return nil
}

// Func seperti consturcor
func NewMenuDbRepository(db *sql.DB) MenuRepository {
	return &menuDbRepository{
		db: db,
	}
}
