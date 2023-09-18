package controller

import (
	"dog-app-api/model"
	"dog-app-api/s3org"
	"dog-app-api/usecase"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// インターフェース
type IDogController interface {
	GetAllUsersDogs(c echo.Context) error
	GetAllDogs(c echo.Context) error
	GetDogById(c echo.Context) error
	CreateDog(c echo.Context) error
	UpdateDog(c echo.Context) error
	DeleteDog(c echo.Context) error
	UploadFile(c echo.Context) error
	UploadToS3(c echo.Context) error
}

// 構造体
type dogController struct {
	du usecase.IDogUsecase
}

// コンストラクタ、usecaseへ
func NewDogController(du usecase.IDogUsecase) IDogController {
	return &dogController{du}
}

// 全てのイヌ取得
func (dc *dogController) GetAllUsersDogs(c echo.Context) error {
	dogsRes, err := dc.du.GetAllUsersDogs()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, dogsRes)
}

// ユーザーの全てのイヌ
func (dc *dogController) GetAllDogs(c echo.Context) error {
	// ユーザーを取得
	id := c.Param("userId")
	userId, _ := strconv.Atoi(id)

	dogsRes, err := dc.du.GetAllDogs(uint(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, dogsRes)
}

// イヌ取得
func (dc *dogController) GetDogById(c echo.Context) error {
	id := c.Param("dogId")
	dogId, _ := strconv.Atoi(id)
	dogRes, err := dc.du.GetDogById(uint(dogId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, dogRes)
}

// 犬を作る
func (dc *dogController) CreateDog(c echo.Context) error {
	// ユーザー情報を取得
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	dog := model.Dog{}
	if err := c.Bind(&dog); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	dog.UserId = uint(userId.(float64))
	dogRes, err := dc.du.CreateDog(dog)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, dogRes)
}

func (dc *dogController) UpdateDog(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("dogId")
	dogId, _ := strconv.Atoi(id)

	dog := model.Dog{}
	if err := c.Bind(&dog); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	dogRes, err := dc.du.UpdateDog(dog, uint(userId.(float64)), uint(dogId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, dogRes)
}

func (dc *dogController) DeleteDog(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("dogId")
	dogId, _ := strconv.Atoi(id)

	err := dc.du.DeleteDog(uint(userId.(float64)), uint(dogId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

// ローカルに画像をupload
func (dc *dogController) UploadFile(c echo.Context) error {
	// フォームデータからファイルを取得
	file, err := c.FormFile("image")
	if err != nil {
		log.Println("Error copying file1:", err)
		return err
	}

	originalFileName := file.Filename // 元のファイル名

	// ファイル名を'.'で分割
	parts := strings.Split(originalFileName, ".")
	if len(parts) < 2 {
		fmt.Println("Invalid file name")
	}

	// 拡張子を取得
	extension := parts[len(parts)-1]

	// ランダムな文字列を生成して新しいファイル名を作成
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var builder strings.Builder
	for i := 0; i < 16; i++ {
		builder.WriteByte(charset[rand.Intn(len(charset))])
	}
	newFileName := fmt.Sprintf("%s.%s", &builder, extension)

	// アップロード先ディレクトリを指定
	uploadDir := "../dog-app-front/public/uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.MkdirAll(uploadDir, 0755)
	}

	// ファイルの保存パスを生成
	destPath := filepath.Join(uploadDir, newFileName)

	// ファイルをオープンして保存
	src, err := file.Open()
	if err != nil {
		log.Println("Error copying file2:", err)
		return err
	}
	defer src.Close()

	dst, err := os.Create(destPath)
	if err != nil {
		log.Println("Error copying file3:", err)
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		log.Println("Error copying file4:", err)
		return err
	}

	// 保存したファイルのURLをクライアントに返す
	fileURL := "uploads/" + file.Filename
	return c.JSON(http.StatusOK, map[string]string{"url": fileURL})
}

// S3に画像をアップロード
func (dc *dogController) UploadToS3(c echo.Context) error {
	s3Client := s3org.GetS3()
	// Reactから送信されたファイルを取得
	file, err := c.FormFile("image")
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
	objectKey := "images/" + newFileName // オブジェクトキーを適切に設定

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
	return c.JSON(http.StatusOK, map[string]string{"url": s3URL})
}
