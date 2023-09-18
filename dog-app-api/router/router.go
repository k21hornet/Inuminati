package router

import (
	"dog-app-api/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// ログイン中のユーザーがいない場合のUnauthorizedに対するエラーハンドリング
func CustomJWTErrorHandler(c echo.Context, err error) error {
	// わざとステータスコード200を返し、ログイン中のユーザーがいないことを示す
	return c.JSON(http.StatusOK, echo.Map{
		"user_id": nil,
	})
}

// Router
func NewRouter(uc controller.IUserController, dc controller.IDogController) *echo.Echo {
	e := echo.New()
	//reactとの通信用
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	//
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		// CookieSameSite: http.SameSiteDefaultMode,
		//CookieMaxAge:   60,
	}))

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)
	e.GET("/:userId", uc.GetUserById)
	e.POST("/:userId/upicon", uc.UpdateIcon)

	// ログイン中のuser_idを取得するエンドポイント
	u := e.Group("/user")
	u.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:   []byte(os.Getenv("SECRET")),
		TokenLookup:  "cookie:token",
		ErrorHandler: CustomJWTErrorHandler, // カスタムエラーハンドラを指定
	}))
	u.GET("", uc.GetLoggedInUserID)

	d := e.Group("/dogs")
	d.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	d.GET("/allusers", dc.GetAllUsersDogs)
	d.GET("/:userId/user", dc.GetAllDogs)
	d.GET("/:dogId", dc.GetDogById)
	d.POST("", dc.CreateDog)
	d.PUT("/:dogId", dc.UpdateDog)
	d.DELETE("/:dogId", dc.DeleteDog)
	d.POST("/upload", dc.UploadFile)
	d.POST("/upload_s3", dc.UploadToS3)

	// aws ヘルスチェック用
	e.GET("/health_check", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Healty")
	})

	return e
}
