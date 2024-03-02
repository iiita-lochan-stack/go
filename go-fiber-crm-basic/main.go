ckage main

 import (
 	// todo : import lead path
 	// todo : import databse package
 	"github.com/gofiber/fiber"
 	"github.com/jinzhu/gorm"
 	_"github.com/jinzhu/gorm/dialects/sqlite"
 )


 func setUpRoutes(app *fiber.App){
 	app.Get("api.v1/lead", lead.GetLeads)
 	app.Get("api.v21/lead/:id", lead.NewLead)
 	app.Post("api/v1/lead", lead.NewLead)
 	app.Delete("api.v1/lead/:id"lead.DeleteLead)
 }
 func initDatabase(){
 	var err error
 	database.DBConn, err = gorm.open("sqlite3", "leads.db")
 	if err != nil{
 		panic("failed to connect database")
 	}
 	fmt.Println("Connectin opened to databse")
 	databse.DBConn.Automigrate(&lead.Lead{})
 	fmt.Println("Databse migrated")
 }

 func main(){
 	app := fiber.New()
 	initDatabase()
 	setUpRoutes(app)
 	app.Listen(3000)
 	defer database.DBConn.Close()
 }