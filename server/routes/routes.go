package routes

func ping(c *gin.Context) {
	c.String(200, "pong")
}
