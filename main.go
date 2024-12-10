package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Robot struct {
	Code        string   `json:"code"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Model       string   `json:"model"`
	Tech        []string `json:"tech"`
	Status      string   `json:"status"`
}

var robots []Robot
var models = []string{"car", "transformation", "humanoid"}
var techs = []string{"AI", "car", "robot", "cyborg"}
var statuses = []string{"progress", "active", "inactive"}

var statusTranslations = map[string]map[string]string{
	"progress": {
		"en": "progress",
		"id": "proses",
	},
	"active": {
		"en": "active",
		"id": "aktif",
	},
	"inactive": {
		"en": "inactive",
		"id": "tidak aktif",
	},
}

var techTranslations = map[string]map[string]string{
	"AI": {
		"en": "AI",
		"id": "Kecerdasan Buatan",
	},
	"car": {
		"en": "car",
		"id": "mobil",
	},
	"robot": {
		"en": "robot",
		"id": "robot",
	},
	"cyborg": {
		"en": "cyborg",
		"id": "cyborg",
	},
}

var modelTranslations = map[string]map[string]string{
	"car": {
		"en": "car",
		"id": "mobil",
	},
	"transformation": {
		"en": "transformation",
		"id": "transformasi",
	},
	"humanoid": {
		"en": "humanoid",
		"id": "humanoid",
	},
}

func init() {
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, &robots); err != nil {
		panic(err)
	}
}

func main() {
	r := gin.Default()

	r.GET("/robots", getRobots)
	r.GET("/robots/:code", getRobotByCode)
	r.POST("/robots", createRobot)
	r.PUT("/robots/:code", updateRobot)
	r.DELETE("/robots/:code", deleteRobot)
	r.GET("/robots/filters", filterRobots)

	r.Run()
}

func getRobots(c *gin.Context) {
	lang := c.DefaultQuery("lang", "en") // Default ke English
	var translatedRobots []map[string]interface{}

	for _, robot := range robots {
		translatedRobot := map[string]interface{}{
			"code":        robot.Code,
			"name":        robot.Name,
			"description": robot.Description,
			"model":       translate(robot.Model, modelTranslations, lang),
			"tech":        translateSlice(robot.Tech, techTranslations, lang),
			"status":      translate(robot.Status, statusTranslations, lang),
		}
		translatedRobots = append(translatedRobots, translatedRobot)
	}

	c.JSON(http.StatusOK, translatedRobots)
}

func getRobotByCode(c *gin.Context) {
	code := c.Param("code")
	lang := c.DefaultQuery("lang", "en")

	for _, robot := range robots {
		if robot.Code == code {
			translatedRobot := map[string]interface{}{
				"code":        robot.Code,
				"name":        robot.Name,
				"description": robot.Description,
				"model":       translate(robot.Model, modelTranslations, lang),
				"tech":        translateSlice(robot.Tech, techTranslations, lang),
				"status":      translate(robot.Status, statusTranslations, lang),
			}
			c.JSON(http.StatusOK, translatedRobot)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Robot not found"})
}

func createRobot(c *gin.Context) {
	var newRobot Robot
	if err := c.BindJSON(&newRobot); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, robot := range robots {
		if robot.Code == newRobot.Code {
			c.JSON(http.StatusConflict, gin.H{"error": "same code"})
			return
		}
	}
	if !contains(statuses, newRobot.Status) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Status"})
		return
	}
	if !contains(models, newRobot.Model) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Model"})
		return
	}
	for _, tech := range newRobot.Tech {
		if !contains(techs, tech) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Tech"})
			return
		}
	}
	robots = append(robots, newRobot)
	c.JSON(http.StatusCreated, newRobot)
}

func updateRobot(c *gin.Context) {
	code := c.Param("code")
	var updatedRobot Robot
	if err := c.BindJSON(&updatedRobot); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, robot := range robots {
		if robot.Code == code {
			robots[i] = updatedRobot
			c.JSON(http.StatusOK, updatedRobot)
			return
		}
	}
	if !contains(statuses, updatedRobot.Status) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Status"})
		return
	}
	if !contains(models, updatedRobot.Model) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Model"})
		return
	}
	for _, tech := range updatedRobot.Tech {
		if !contains(techs, tech) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Tech"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Robot not found"})
}
func deleteRobot(c *gin.Context) {
	code := c.Param("code")
	for i, robot := range robots {
		if robot.Code == code {
			robots = append(robots[:i], robots[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Robot deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Robot not found"})
}

func translate(key string, translations map[string]map[string]string, lang string) string {
	if val, exists := translations[key]; exists {
		if translated, ok := val[lang]; ok {
			return translated
		}
	}
	return key // Kembalikan nilai asli jika tidak ada terjemahan
}

func translateSlice(keys []string, translations map[string]map[string]string, lang string) []string {
	var translated []string
	for _, key := range keys {
		translated = append(translated, translate(key, translations, lang))
	}
	return translated
}
func filterRobots(c *gin.Context) {
	model := c.Query("model")
	techFilters := c.QueryArray("tech")
	var filtered []Robot
	for _, robot := range robots {
		if (model == "" || robot.Model == model) || (len(techFilters) == 0 || hasTech(robot.Tech, techFilters)) {
			filtered = append(filtered, robot)
		}
	}
	c.JSON(http.StatusOK, filtered)
}

func hasTech(robotTech []string, filters []string) bool {
	for _, filter := range filters {
		if !contains(robotTech, filter) {
			return false
		}
	}
	return true
}
func contains(slice []string, item string) bool {
	for _, v := range slice {
		if strings.ToLower(v) == strings.ToLower(item) {
			return true
		}
	}
	return false
}
