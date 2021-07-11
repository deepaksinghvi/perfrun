package cmd

import (
	"fmt"
	"github.com/deepaksinghvi/perfrun/pkg/api"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// perfrunwsCmd represents the perfrunws command
var perfrunwsCmd = &cobra.Command{
	Use:   "perfrunws",
	Short: "A brief description of your command",
	Long: `A longer description .`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("perfrunws API Service Start-up")
		router := gin.Default()
		initAPI(router)
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		router.Run(fmt.Sprintf(":%d", 8080))
	},
}

func init() {
	rootCmd.AddCommand(perfrunwsCmd)
}


func initAPI(router *gin.Engine) {
	apis := router.Group("/api/v1")
	perfrunws := apis.Group("/perfrunws")
	{
		perfrunws.POST("", api.CreatePerfRun)
		perfrunws.GET(":id", api.GetPerfRun)
	}
	health := apis.Group("/health")
	{
		health.GET("", HealthCheck)
	}
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	c.JSON(http.StatusOK, res)
}
