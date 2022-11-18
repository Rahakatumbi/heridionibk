package depensedocs

import (
	"github.com/Raha2071/heridionibd/config"
	"github.com/Raha2071/heridionibd/models"
	"github.com/gin-gonic/gin"
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
	var help HelperFrom
	c.ShouldBind(&help)
	// if help.Document == "" {
	// 	file, handle, err := c.Request.FormFile("document")
	// 	// The file cannot be received.
	// 	if err != nil {
	// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 			"message": "Mauvaise rqequette",
	// 			"error":   err.Error(),
	// 		})
	// 		return
	// 	}
	// 	defer file.Close()
	// 	filename := path.Base(handle.Filename)
	// 	// destination, err := os.Create(filename)

	// 	newFileName := uuid.New().String() + filename

	// 	if err != nil {
	// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
	// 			"message": "Crash du System",
	// 			"error":   err.Error(),
	// 		})
	// 		return
	// 	}

	// 	// The file is received, so let's save it
	// 	if err := c.SaveUploadedFile(handle, "./public/docs/"+newFileName); err != nil {
	// 		c.AbortWithStatusJSON(200, gin.H{
	// 			"message": "Unable to save the file",
	// 			"error":   err,
	// 		})
	// 		return
	// 	}
	doc := models.Documents{Object: help.Object, Intitule: help.Intitule, Description: help.Description, Createdby: help.CreatedBy}
	config.DB.Create(&doc)
	// config.DB.Updates(&models.Documents{Id: doc.Id, Document: newFileName})
	// }
	c.JSON(200, doc)
}
func Documents(c *gin.Context) {
	var document []models.Documents
	config.DB.Find(&document)
	c.JSON(200, document)
}
