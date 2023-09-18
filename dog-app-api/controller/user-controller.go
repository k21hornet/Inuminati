package controller

import (
	"dog-app-api/model"
	"dog-app-api/s3org"
	"dog-app-api/usecase"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// インターフェース
type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
	CsrfToken(c echo.Context) error
	GetLoggedInUserID(c echo.Context) error
	GetUserById(c echo.Context) error
	UpdateIcon(c echo.Context) error
}

// 構造体
type userController struct {
	uu usecase.IUserUsecase
}

// コンストラクタ、usecaseへ
func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

// signupコントローラー
func (uc *userController) SignUp(c echo.Context) error {
	//クライアントから受け取るリクエストをUserの構造体にバインド
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// usecaseへ委譲
	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, userRes)
}

// loginコントローラー
func (uc *userController) LogIn(c echo.Context) error {
	user := model.User{}
	//クライアントから受け取るリクエストをUserの構造体にバインド
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// usecaseへ委譲
	tokenString, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// 取得したjwtをcookieに保存
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	// 有効期限は24h
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	//クライアントのjsからtokenの値が読み取れないように設定
	cookie.HttpOnly = true
	//クロスドメイン間でのcookieの送受信をするために設定
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

// logoutコントローラー
func (uc *userController) LogOut(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	//有効期限がすぐ切れるように
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

// csrfToken取得
func (uc *userController) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}

// ログイン中のユーザーid
func (uc *userController) GetLoggedInUserID(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	return c.JSON(http.StatusOK, echo.Map{
		"user_id": userId,
	})
}

func (uc *userController) GetUserById(c echo.Context) error {
	id := c.Param("userId")
	userId, _ := strconv.Atoi(id)
	dogRes, err := uc.uu.GetDogById(uint(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, dogRes)
}

func (uc *userController) UpdateIcon(c echo.Context) error {
	id := c.Param("userId")
	userId, _ := strconv.Atoi(id)

	// s3にアップロード
	s3Client := s3org.GetS3()
	// Reactから送信されたファイルを取得
	file, err := c.FormFile("icon")
	if err != nil {
		log.Println("Error copying file:", err)
		return err
	}

	originalFileName := file.Filename // 元のファイル名

	// ファイル名を'.'で分割
	parts := strings.Split(originalFileName, ".")
	if len(parts) < 2 {
		fmt.Println("Invalid file name")
	}

	// ランダムな文字列を生成して新しいファイル名を作成
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var builder strings.Builder
	for i := 0; i < 16; i++ {
		builder.WriteByte(charset[rand.Intn(len(charset))])
	}
	newFileName := fmt.Sprintf("%s.png", &builder)

	// S3のバケット名とオブジェクトキーを設定
	bucketName := "dog-app-upload-dev-1"
	objectKey := "icons/" + newFileName // オブジェクトキーを適切に設定

	// ファイルをオープン
	src, err := file.Open()
	if err != nil {
		log.Println("Error copying file2:", err)
		return err
	}
	defer src.Close()

	// S3にファイルをアップロード
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   src,
	})
	if err != nil {
		log.Println("Error copying s3:", err)
		return err
	}

	// アップロードされたファイルのS3 URLを生成
	s3URL := "https://" + bucketName + ".s3.amazonaws.com/" + objectKey

	user := model.User{}
	userRes, err := uc.uu.UpdateIcon(user, uint(userId), s3URL)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, userRes)
}
