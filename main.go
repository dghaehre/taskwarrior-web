package main

import (
	"embed"

	"github.com/gin-gonic/gin"
)

//go:embed templates
var templates embed.FS

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	t := NewTaskwarrior()

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	r.GET("/htmx/today.html", func(c *gin.Context) {
		tasks, err := t.GetTodayTasks()
		if err != nil {
			c.HTML(200, "error.html", gin.H{"error": err})
			return
		}
		c.HTML(200, "htmx-today.html", tasks)
	})

	r.Run(":8080")
}
