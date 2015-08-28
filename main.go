package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
	"time"
)

// Create Models

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
}

type Goal struct {
	ID        int
	Content   string
	CreatedAt time.Time
	Complete  bool
	UserID    int
}

// Declare Global Variables

var db = initDb()

// Database Initializer
func initDb() gorm.DB {
	db, err := gorm.Open("sqlite3", "conquistador.db")
	checkErr(err, "There was an error loading the database")

	db.AutoMigrate(&User{}, &Goal{})
	return db
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Printf(msg, err)
	}
}

// Route Methods

func index(c *gin.Context) {
	content := gin.H{"Hello": "World"}
	c.JSON(200, content)
}

func goalsGet(c *gin.Context) {
    goals := getAllGoals()
    content := gin.H{"goals": goals,}
    c.JSON(200, content)
}

func goalsDetail(c *gin.Context) {
    goalId := c.Params.ByName("id")
    gId, _ := strconv.Atoi(goalId)
    goal := readGoal(gId)
    content := gin.H{"goal": goal}
    c.JSON(200, content)
}

func goalsPost(c *gin.Context) {
	var json Goal

    c.Bind(&json)
    goal := createGoal(json.Content, json.UserID)
    if goal.Content == json.Content {
        content := gin.H{
            "result": "success",
            "goal": goal,
        }
        c.JSON(201, content)
    } else {
        c.JSON(500, gin.H{"result": "An error occurred"})
    }
}

// User CRUD

func createUser(firstName, lastName, email string) User {
	user := User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		CreatedAt: time.Now(),
	}
	db.Create(&user)

	return user
}

func readUser(id int) User {
	var user User
	db.First(&user, id)

	return user
}

func updateUser(id int, newUser User) User {
	var user User
	db.First(&user, id)
	user = newUser
	db.Save(&user)

	return user
}

func deleteUser(id int) User {
	var user User
	db.Delete(db.First(&user, id))

	return user
}

// Goal CRUD

func createGoal(content string, userId int) Goal {
	goal := Goal{
		Content:   content,
		CreatedAt: time.Now(),
		Complete:  false,
		UserID:    userId,
	}
	db.Create(&goal)

	return goal
}

func readGoal(id int) Goal {
	var goal Goal
	db.First(&goal, id)

	return goal
}

func updateGoal(id int, newGoal Goal) Goal {
	var goal Goal
	db.First(&goal, id)
	goal = newGoal
	db.Save(&goal)

	return goal
}

func deleteGoal(id int) Goal {
	var goal Goal
	db.Delete(db.First(&goal, id))

	return goal
}

func getAllGoals() []Goal {
    var goals []Goal
    db.Find(&goals)
    
    return goals
}

func main() {
	app := gin.Default()

	app.GET("/", index)

	// Define Routes
	app.GET("/goals", goalsGet)
	app.POST("/goals", goalsPost)
	app.GET("/goals/:id", goalsDetail)

	// Define Port
	app.Run(":8080")
}
