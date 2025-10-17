package main

import (
	"controlF_back/internal/domain/auth"
	"controlF_back/internal/domain/user"
	"controlF_back/internal/models"
	"controlF_back/internal/utils"
	"controlF_back/internal/version"
	"net/http"
	"os"
	"strings"

	_ "controlF_back/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Debug().Msg(".env file not find")
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	models.ConnectDataBase()
}

// @title           controlF API
// @version         1.0
// @description     Esta é a documentação da API controlF.
// @host      localhost:7002
// @BasePath  /api
// @securityDefinitions.apikey BearerAuth
// @in                       header
// @name                     Authorization
func main() {
	// messagebroker.Start()

	r := gin.New()
	r.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/healthcheck"),
		gin.Recovery(),
	)

	if value, ok := os.LookupEnv("CORS_ORIGIN"); ok {
		config := cors.DefaultConfig()
		if value == "*" {
			config.AllowOriginFunc = func(_ string) bool { return true }
		} else {
			config.AllowOrigins = strings.Split(value, ",")
		}

		config.AddAllowHeaders("Authorization")

		r.Use(cors.New(config))
	}

	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"service":    "go-template",
			"git_commit": version.GitCommit,
			"build_os":   version.BuildOS,
			"build_date": version.BuildDate,
			"start_time": version.StartTime,
			"up_time":    version.GetUptime(),
			"version":    version.Version,
		})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	setupRoutes(r)

	host := utils.GetEnv("HOST", "")
	port := utils.GetEnv("PORT", "8999")
	r.Run(host + ":" + port)
}

func setupRoutes(r *gin.Engine) {
	authController := auth.InitAuthService()
	auth.RegisterRoutes(r, *authController)

	userController := user.InitUserService()
	user.RegisterRoutes(r, *userController)
}
