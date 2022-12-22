package controllers

import (
	"context"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Eiliv17/FileVaultWebApp/initializers"
	"github.com/Eiliv17/FileVaultWebApp/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func Upload(c *gin.Context) {
	// database setup
	dbname := os.Getenv("DB_NAME")
	coll := initializers.DB.Database(dbname).Collection("files")

	// get file and password from req body
	file, err := c.FormFile("file")
	if err != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"error": "You must upload a file",
		})
		return
	}

	rawpassword := c.PostForm("password")
	if rawpassword == "" {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"error": "You must write a password",
		})
		return
	}

	// hash the password using bcrypt
	hashPass, err := bcrypt.GenerateFromPassword([]byte(rawpassword), 10)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{
			"error": "Internal server error",
		})
		return
	}

	// separate file name from file extension
	filextension := ""
	filename := file.Filename
	fileslice := strings.Split(file.Filename, ".")
	if len(fileslice) > 1 {
		filextension = fileslice[len(fileslice)-1]
		filename = strings.Join(fileslice[:len(fileslice)-1], "")
	}

	// create obj id
	timenow := time.Now()
	objID := primitive.NewObjectIDFromTimestamp(timenow)

	// save the file in a directory
	dirname := os.Getenv("FILE_DIR")
	filesavepath := dirname + "/" + objID.Hex() + "." + filextension

	err = c.SaveUploadedFile(file, filesavepath)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{
			"error": "Internal server error",
		})
		return
	}

	// create file model
	filem := models.File{
		FileID:        objID,
		FileName:      filename,
		FileExtension: filextension,
		FileLocation:  "/" + filesavepath,
		FilePassword:  string(hashPass),
		UploadedAt:    timenow,
	}

	_, err = coll.InsertOne(context.TODO(), filem)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{
			"error": "Internal server error",
		})
		return
	}

	// redirects to file page
	c.Redirect(http.StatusSeeOther, "/file/"+objID.Hex())
}

func Download(c *gin.Context) {
	// database setup
	dbname := os.Getenv("DB_NAME")
	coll := initializers.DB.Database(dbname).Collection("files")

	// get file id parameter from
	fileID := c.PostForm("fileid")
	password := c.PostForm("password")

	if fileID == "" || password == "" {
		c.HTML(http.StatusInternalServerError, "protectedfile.html", gin.H{
			"error":  "FileID or Password missing",
			"fileid": fileID,
		})
		return
	}

	objID, err := primitive.ObjectIDFromHex(fileID)
	if err != nil {
		c.HTML(http.StatusOK, "protectedfile.html", gin.H{
			"error":  "Internal Server Error",
			"fileid": fileID,
		})
		return
	}

	// retrieve file from database
	filter := bson.D{primitive.E{Key: "_id", Value: objID}}
	result := coll.FindOne(context.TODO(), filter)

	var file models.File
	err = result.Decode(&file)
	if err != nil {
		c.HTML(http.StatusOK, "protectedfile.html", gin.H{
			"error":  "File not found",
			"fileid": fileID,
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(file.FilePassword), []byte(password))
	if err != nil {
		c.HTML(http.StatusOK, "protectedfile.html", gin.H{
			"error":  "Wrong password",
			"fileid": fileID,
		})
		return
	}

	c.FileAttachment("."+file.FileLocation, file.FileName+"."+file.FileExtension)
}

func GetDownloadPage(c *gin.Context) {
	fileID := c.Param("id")

	c.HTML(http.StatusOK, "protectedfile.html", gin.H{
		"fileid": fileID,
	})
}
