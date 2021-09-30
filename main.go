package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Username string `json:"username"`
	Follower int    `json:"followers"`
}

func main() {

	app := fiber.New()

	app.Get("/follower/:username", func(ctx *fiber.Ctx) error {
		response, err := http.Get("https://jsonkeeper.com/b/DMXK")
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		user := map[string]User{}
		json.Unmarshal(data, &user)
		var result string
		usernameSearch := ctx.Params("username")

		for _, v := range user {
			if v.Username == usernameSearch {
				result = "followers: " + strconv.Itoa(v.Follower)
				break
			}
		}

		return ctx.SendString(result)
	})

	app.Get("/:userid/detail", func(ctx *fiber.Ctx) error {

		response, err := http.Get("https://jsonkeeper.com/b/DMXK")
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		user := map[string]User{}
		json.Unmarshal(data, &user)
		userid := ctx.Params("userid")
		result := user[userid]

		return ctx.JSON(result)
	})

	app.Listen(":9000")

}
