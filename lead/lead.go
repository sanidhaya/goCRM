package lead

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/sanidhaya/goCRM/database"
)

type Lead struct {
	gorm.Model
	Name    string
	Company string
	Email   string
	Phone   int
}

func GetLeads(c *fiber.Ctx) {
	db := database.DBConn
	var leads []Lead

}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find((&lead), id)
	c.JSON(lead)
}

func NewLeads(c *fiber.Ctx) {

}

func DeleteLeads(c *fiber.Ctx) {

}
