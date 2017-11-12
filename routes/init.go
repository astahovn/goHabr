package routes

import (
  "github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine) {
  // init templates
  engine.LoadHTMLGlob("tpl/**/*")

  // init static
  engine.Static("/images", "./assets/images")
  engine.Static("/css", "./assets/css")

  // init dynamic
  engine.GET("/", Index)
  engine.GET("/register", Register)
  engine.POST("/register_user", RegisterPost)
  engine.POST("/login", LoginPost)
  engine.GET("/logout", Logout)

  engine.GET("/profile", ProfileIndex)
  engine.GET("/profile/edit", ProfileEdit)
  engine.POST("/profile/edit_save", ProfileEditSave)
}
