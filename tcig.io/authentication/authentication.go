package authentication

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

func initClientStore() *store.ClientStore {
	clientStore := store.NewClientStore()

	clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost",
	})

	clientStore.Set("!UHY_VBRF_xdCC_pmOP_&lk8", &models.Client{
		ID:     "!UHY_VBRF_xdCC_pmOP_&lk8",
		Secret: "gFtv_9&dr_=GSx_WS98_$Stg",
		Domain: "http://localhost",
	})

	clientStore.Set("Hubert", &models.Client{
		ID:     "Hubert",
		Secret: "17Elvis17",
		Domain: "http://localhost:8080",
	})

	return clientStore
}

var _srv *server.Server

func Init(router *gin.Engine) {
	manager := manage.NewDefaultManager()

	manager.MustTokenStorage(store.NewFileTokenStore(os.Getenv("ROOT_DIR") + "/_tokens/store"))
	//manager.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte("00000000"), jwt.SigningMethodHS512))

	manager.MapClientStorage(initClientStore())

	// refresh_token wanted
	manager.SetClientTokenCfg(&manage.Config{
		AccessTokenExp:    time.Hour * 1,
		RefreshTokenExp:   time.Hour * 2,
		IsGenerateRefresh: true,
	})

	_srv = server.NewDefaultServer(manager)
	_srv.SetAllowGetAccessRequest(true)
	_srv.SetClientInfoHandler(server.ClientFormHandler)

	_srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	_srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	router.POST("/o/token", func(c *gin.Context) {
		_srv.HandleTokenRequest(c.Writer, c.Request)
	})

	router.POST("/o/refresh", func(c *gin.Context) {
		_srv.HandleTokenRequest(c.Writer, c.Request)
	})
}

/*
func CheckAccess(c *gin.Context, proceed func(string)) {
	token, err := _srv.ValidationBearerToken(c.Request)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	proceed(token.GetScope())
}
*/

func CheckAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := _srv.ValidationBearerToken(c.Request)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid API token"})
		}
		c.Set("scope", token.GetScope())
		c.Next()
	}
}
