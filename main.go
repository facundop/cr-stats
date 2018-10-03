package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"facundoprecentado.com/cr-stats/domain"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var apiToken = ""

func main() {

	// ***** Cargo Token desde config file para no publicar el token
	type Configuration struct {
		Token string `json:"token"`
	}

	file, _ := os.Open("config/config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error loading config json:", err)
	}
	apiToken = configuration.Token
	// *************************************************************

	// Set the router as the default one shipped with Gin
	router := gin.Default()
	router.LoadHTMLFiles("views/index.html", "views/chests.html")
	http.Handle("/static/", http.StripPrefix("/static/api/", http.FileServer(http.Dir("views"))))

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Setup route group for the API
	api := router.Group("/")

	api.GET("/home", homeHandler)
	api.GET("/cards", cardHandler)
	api.GET("/clans/:clanID", clanHandler)
	api.GET("/players/:playerID", playerHandler)
	api.GET("/players/:playerID/chests", playerUpcomingChestsHandler)
	api.GET("/players/:playerID/battles", playerBattleLogHandler)

	// Start and run the server
	router.Run(":8080")

}

func homeHandler(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Clash Royale Companion",
		"chest": "Chests",
	})

}

func cardHandler(c *gin.Context) {
	req, err := http.NewRequest("GET", "https://api.clashroyale.com/v1/cards/", nil)

	req.Header.Set("authorization", "Bearer "+apiToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		decoder := json.NewDecoder(resp.Body)

		data := &domain.Card{}
		err := decoder.Decode(data)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		}

		c.JSON(http.StatusOK, data)
	}

	c.Header("Content-Type", "application/json")
}

func clanHandler(c *gin.Context) {
	//clanTag := "#PP0222JV"
	clanTag := c.Param("clanID")

	clanTag = url.QueryEscape("#" + clanTag)

	req, err := http.NewRequest("GET", "https://api.clashroyale.com/v1/clans/"+clanTag, nil)

	req.Header.Set("authorization", "Bearer "+apiToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		decoder := json.NewDecoder(resp.Body)

		data := &domain.Player{}
		err := decoder.Decode(data)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		}

		c.JSON(http.StatusOK, data)
	}

	c.Header("Content-Type", "application/json")
}

func playerHandler(c *gin.Context) {
	//userTag := "#2LRRQPPJC"
	userTag := c.Param("playerID")

	userTag = url.QueryEscape("#" + userTag)

	req, err := http.NewRequest("GET", "https://api.clashroyale.com/v1/players/"+userTag, nil)

	req.Header.Set("authorization", "Bearer "+apiToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		decoder := json.NewDecoder(resp.Body)

		data := &domain.Player{}
		err := decoder.Decode(data)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		}

		c.JSON(http.StatusOK, data)
	}

	c.Header("Content-Type", "application/json")
}

func playerUpcomingChestsHandler(c *gin.Context) {
	userTag := c.Param("playerID")

	userTag = url.QueryEscape("#" + userTag)

	println("user", userTag)

	req, err := http.NewRequest("GET", "https://api.clashroyale.com/v1/players/"+userTag+"/upcomingchests", nil)

	req.Header.Set("authorization", "Bearer "+apiToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println("Error 1")
	} else {
		decoder := json.NewDecoder(resp.Body)

		data := &domain.UpcomingChests{}
		err := decoder.Decode(data)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			fmt.Println("Error 2")
		}

		fmt.Println("Values:", data.Items)

		c.HTML(http.StatusOK, "chests.html", gin.H{
			"title":  "Player Upcoming Chests",
			"chest":  "Chests",
			"chests": data.Items,
		})
	}

}

func playerBattleLogHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	userTag := c.Param("playerID")

	userTag = url.QueryEscape("#" + userTag)

	println("user", userTag)

	req, err := http.NewRequest("GET", "https://api.clashroyale.com/v1/players/"+userTag+"/battlelog", nil)

	req.Header.Set("authorization", "Bearer "+apiToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {

		decoder := json.NewDecoder(resp.Body)

		data := &domain.BattleLog{}
		err := decoder.Decode(data)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, data)
		}

	}

}
