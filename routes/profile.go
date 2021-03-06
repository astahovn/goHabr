package routes

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "github.com/astahovn/ninja/lib/session"
  "github.com/astahovn/ninja/lib/db"
)

// Profile index page
func ProfileIndex(c *gin.Context) {
  c.HTML(http.StatusOK, "profile/index.tmpl", gin.H{
    "title": "Profile",
    "auth": session.GetActiveUser(c),
  })
}

// Profile edit page
func ProfileEdit(c *gin.Context) {
  authUser := session.GetActiveUser(c)

  c.HTML(http.StatusOK, "profile/edit.tmpl", gin.H{
    "title": "Profile edit",
    "nick": authUser.Nick,
    "auth": authUser,
  })
}

// Profile edit page form saving
func ProfileEditSave(c *gin.Context) {
  nick := c.PostForm("nick")
  authUser := session.GetActiveUser(c)

  if nick == "" {
    c.HTML(http.StatusOK, "profile/edit.tmpl", gin.H{
      "title": "Profile edit",
      "auth": authUser,
      "errors": "Nick is empty",
    })
    return
  }

  authUser.Nick = nick
  db.GetInstance().Save(&authUser)

  c.Redirect(http.StatusFound, "/profile")
}
