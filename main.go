package main

import (
  "github.com/gin-gonic/gin"

  "sequis/controllers"

  "sequis/models"

)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.POST("/add-friend", controllers.AddFriendHandler)
	r.POST("/accept-reject", controllers.AcceptRejectHandler)
	r.POST("/list-friend-request", controllers.ListFriendsHandler)
	r.POST("/friend-list", controllers.FriendListHandler)
	r.POST("/common-friend-list", controllers.CommonFriendListHandler)
	r.POST("/block-user", controllers.BlockUserHandler)

	r.Run()
}