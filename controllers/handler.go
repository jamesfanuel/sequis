package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
)

type RequestorToParam struct {
	Requestor string `json:"requestor" binding: "required" validate:"required"`
	To string `json:"to" binding: "required" validate:"required"`
}

type RequestorStatusResponse struct {
    Requestor string `json:"requestor"`
    Status  string `json:"status"`
}

type EmailParam struct {
    Email string `json:"email"`
}

type RequestorBlockParam struct {
    Email string `json:"email"`
}

func AddFriendHandler(c *gin.Context) {
	var requestorToParam RequestorToParam

	if err := c.ShouldBindJSON(&requestorToParam); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
    }

	if err := validator.Validate(requestorToParam); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
    }

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func AcceptRejectHandler(c *gin.Context) {
	var requestorToParam RequestorToParam

	if err := c.ShouldBindJSON(&requestorToParam); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
    }

	if err := validator.Validate(requestorToParam); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
    }

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func ListFriendsHandler(c *gin.Context) {
	var emailParam EmailParam

	if err := c.ShouldBindJSON(&emailParam); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
    }

	if err := validator.Validate(emailParam); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
	}
	
	var data = []RequestorStatusResponse{
		RequestorStatusResponse{"andy@example.com", "accepted"},
		RequestorStatusResponse{"joe@example.com", "rejected"},
		RequestorStatusResponse{"grace@example.com", "pending"},
	}

	c.JSON(http.StatusOK, gin.H{"requests": data})
}

func FriendListHandler(c *gin.Context) {
	var emailParam EmailParam

	if err := c.ShouldBindJSON(&emailParam); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
    }

	if err := validator.Validate(emailParam); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
	}
	
	var data = []string{"john@example.com", "joe@example.com"}


	c.JSON(http.StatusOK, gin.H{"friends": data})
}

func CommonFriendListHandler(c *gin.Context) {
	var emailParam EmailParam

	if err := c.ShouldBindJSON(&emailParam); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
    }

	if err := validator.Validate(emailParam); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
	}
	
	var data = []string{"frank@example.com"}


	c.JSON(http.StatusOK, gin.H{"success": true, "friends": data, "count": len(data)})
}

func BlockUserHandler(c *gin.Context) {
	var requestorBlockParam RequestorBlockParam

	if err := c.ShouldBindJSON(&requestorBlockParam); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
    }

	if err := validator.Validate(requestorBlockParam); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}