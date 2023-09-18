package s3org

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetS3() *s3.S3 {
	// AWS認証情報をセットアップ
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Credentials: credentials.NewEnvCredentials(), // 環境変数から認証情報を読み込む
	}))

	// S3クライアントを作成
	s3Client := s3.New(sess)
	return s3Client
}
