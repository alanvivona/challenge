package gdrive

import (
	"api/app/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Auth ...
func Auth(c *gin.Context) {
	// Verify if is a redirect from Gdrive with the authorized token
	stateToken := c.Query("state")
	if stateToken != "" {
		// Getting new token from Gdrive
		tokenCode := c.Query("code")
		if stateToken != "state-token" || tokenCode == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "auth_error", "description": "Invalid token"})
			return
		}
		Gds.CreateClient(c, tokenCode)
		c.JSON(http.StatusOK, gin.H{"success": "auth_success", "description": "Authentication success"})
		return
	} else {
		// First time auth. Provide auth URL to the user
		authURL, err := Gds.GetAuthURL()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "auth_error", "description": "Auth process error"})
			return
		}
		c.Redirect(http.StatusSeeOther, authURL)
		return
	}
}

// SearchInDoc ...
func SearchInDoc(c *gin.Context) {
	if Gds.HasClient() == true {

		fileID := strings.TrimSpace(c.Param("id"))
		if fileID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id_error"})
			return
		}

		word := strings.TrimSpace(c.Query("word"))
		if word == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "word_error"})
			return
		}

		Gds.SearchInDoc(fileID, word)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "auth_error", "description": "not implemented yet"})
		return

	}

	Auth(c)

}

// CreateFile ...
func CreateFile(c *gin.Context) {
	if Gds.HasClient() == true {
		file := &models.File{}
		if err := c.BindJSON(file); c.Request.ContentLength == 0 || err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bind_error", "description": err.Error()})
			return
		}
		driveFile, err := Gds.CreateFile(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "save_error", "description": err.Error(), "drivefile": driveFile})
			return
		}
		c.JSON(201, driveFile)
		return
	}

	Auth(c)

}
