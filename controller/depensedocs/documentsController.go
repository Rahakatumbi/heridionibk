package depensedocs

import (
	"net/http"
	"path"

	"github.com/Raha2071/heridionibd/config"
	"github.com/Raha2071/heridionibd/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HelperFrom struct {
	Id          int    `form:"id" json:"id"`
	Document    string `form:"document" json:"document"`
	Intitule    string `form:"intitule" json:"intitule"`
	Object      string `form:"object" json:"object"`
	Description string `form:"description" json:"description"`
	CreatedBy   int    `form:"createdby" json:"createdby"`
}

func Document(c *gin.Context) {
	c.MultipartForm()
	file, handle, err := c.Request.FormFile("document")
	// The file cannot be received.
	help := HelperFrom{}
	c.ShouldBind(&help)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Mauvaise rqequette",
			"error":   err.Error(),
			"data":    help,
		})
		return
	}
	defer file.Close()
	filename := path.Base(handle.Filename)
	// destination, err := os.Create(filename)

	newFileName := uuid.New().String() + filename

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Crash du System",
			"error":   err.Error(),
		})
		return
	}
	// The file is received, so let's save it
	if err := c.SaveUploadedFile(handle, "./public/docs/"+newFileName); err != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"message": "Unable to save the file",
			"error":   err,
		})
		return
	}
	config.DB.Save(&models.Documents{Id: uint(help.Id), Object: help.Object, Intitule: help.Intitule, Document: newFileName, Description: help.Description, Createdby: help.CreatedBy})
}
func Documents(c *gin.Context) {
	var document []models.Documents
	config.DB.Find(&document)
	c.JSON(200, document)
}
