package lead

import (
	// import database here
	"github.com/jinzhu/gorm"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm/dialects/sqlite"
)

type Lead struct{
	gorm.Model
	Name 		string		`json:"name"`
	Company 	string		`json:"company"`
	Email 		string		`json:"email"`
	Phone 		int			`json:"phone"`
}

func GetLeads(c *fiber.Ctx){
	db := Databse.DBConn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)

}

func GetLead(){
	id := c.Params["id"]
	db := Databse.DBConn
	var lead lead
	db.Find(&lead, id)
	c.JSON(lead)
}

func NeLead(){
	db := databse.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil{
		c.Status(503).Send(err)
		return
	}
}

func DeleteLead(){
	id := c.Params["id"]
	db := database.DBConn

	var lead Lead
	db.First(&lead, id)
	if lead.Name == ""{
		c.Status(500).Send("No lead found with id")
		return
	}
	db.Delete(&lead)
	c.Send("Lead successfully deleted")
}