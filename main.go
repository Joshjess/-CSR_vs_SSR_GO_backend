package main

import "github.com/gin-gonic/gin"
import "net/http"
import "github.com/gin-contrib/cors"


func homePage(c *gin.Context){
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis sit amet ex eros. Cras ac lorem eros. Donec sollicitudin quam vel massa accumsan laoreet. Proin malesuada magna quis augue faucibus tincidunt. Integer eleifend porta magna, eu feugiat erat dictum a. In hac habitasse platea dictumst. Donec justo nunc, semper sit amet nulla ut, varius finibus purus. Ut hendrerit est in sem malesuada, ultricies tincidunt sem auctor. Morbi aliquet, sapien vel cursus porttitor, diam purus aliquet justo, eu malesuada sem nisl eu dolor. Aliquam id tempus elit. Aliquam non risus sem. Nam ut sapien leo. Nullam vitae rhoncus ligula. Mauris facilisis lectus dui, sit amet laoreet diam volutpat ac. Cras elementum arcu nisi, quis tincidunt turpis gravida non. In consectetur ex tellus, ut pellentesque ligula pellentesque vel. Duis velit ante, interdum id euismod id, tincidunt laoreet lorem. Nunc eget odio quam. Maecenas nunc risus, volutpat at maximus vitae, ornare eu dui. Donec sed nisl in nisl tincidunt blandit. Donec nunc elit, convallis ultrices ex a, sagittis tristique mauris.Ut elementum arcu lobortis augue scelerisque, vel venenatis dui imperdiet. Mauris iaculis ac ligula id mollis. Nulla nec luctus enim, ac scelerisque risus. Sed condimentum tincidunt eros, eget vulputate nisl tincidunt quis. Morbi scelerisque ante ex. Phasellus vel commodo arcu, ac imperdiet elit. Suspendisse potenti. Nulla sit amet leo facilisis, dignissim enim id, eleifend elit. Aliquam eget auctor libero. Morbi nisi neque, fermentum quis nulla eget, eleifend bibendum est. Nam semper dictum dolor id tincidunt. Proin sollicitudin varius vestibulum. Nunc justo turpis, volutpat pharetra libero vitae, facilisis tincidunt metus. Pellentesque semper euismod nunc, sagittis ultrices ligula pretium sit amet. Proin in sodales est, id rutrum ipsum. Fusce sed mi auctor, fringilla ipsum et, malesuada ex. Curabitur odio dolor, euismod vitae ligula ac, feugiat facilisis augue. Curabitur ac arcu in neque elementum commodo. Maecenas commodo orci ut varius hendrerit. Duis sit amet est rutrum, rhoncus lacus id, tempus lacus. Cras euismod metus vitae metus pretium aliquam. Aenean sollicitudin turpis ex, non consequat ipsum faucibus id. Sed eget pretium erat. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Suspendisse tincidunt gravida sapien, at vehicula erat efficitur eget. Nullam ultricies quam eget massa aliquet maximus. Curabitur tempus, nunc ut suscipit molestie, ex sem egestas nisi, vitae scelerisque velit dolor sit amet enim. Maecenas et nibh nisi. Mauris quis lectus ut enim placerat hendrerit at vitae eros. Etiam mi dui, dapibus vitae ultricies vitae, euismod non diam. Aenean blandit ligula vel interdum pretium. Suspendisse efficitur non libero eget tempus. Etiam magna diam, cursus eget felis sit amet, maximus condimentum neque. Sed gravida maximus massa id volutpat."
	c.JSON(http.StatusOK, gin.H{
		"message" : text,
	})

}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}



func main() {
	router := gin.Default()

	router.Use(CORS())
	router.Use(cors.Default())
	router.GET("/gethomepage", homePage)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
}