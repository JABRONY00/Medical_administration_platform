package clients

import (
	"context"

	"github.com/gin-gonic/gin"
)

func CreateClient(c *gin.Context) {

}
func GetClients(c *gin.Context) {
	err := DB.Ping(context.Background())
	if err != nil {

		return
	}
}
func GetClient(c *gin.Context) {

}
func DeleteClient(c *gin.Context) {

}
func UpdateClient(c *gin.Context) {

}
