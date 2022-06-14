package actions

import (
	"fmt"
	"net/http"
	"projectfileb/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func NameFile(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("files/index.plush.html"))
}

func ReceiveFiles(c buffalo.Context) error {
	_, fileInfo, _ := c.Request().FormFile("document")

	nameFile := fileInfo.Filename

	file := models.File{
		FileName: nameFile,
	}

	tx := c.Value("tx").(*pop.Connection)

	err := tx.Create(&file)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/name")
}

func Name(c buffalo.Context) error {
	files := []models.File{}

	tx := c.Value("tx").(*pop.Connection)

	err := tx.All(&files)
	if err != nil {
		return err
	}

	c.Set("files", files)

	fmt.Println(files)

	return c.Render(http.StatusOK, r.HTML("files/name.plush.html"))
}
