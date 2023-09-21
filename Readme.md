# Inuminati
犬の画像を共有するインスタグラムのようなサービス。

## 技術スタック
- Go / Echo
- TypeScript / React
- Tailwind
- Docker MySQL
- AWS S3

## インフラ
AWSに期間限定で公開(コストがかかる割にユーザーが少ないので現在は閉鎖)。
GoのapiはECSに、ReactはCloudFrontとS3へデプロイ。

- VPC
- ECS
- ECR
- ALB
- CloudFront
- S3
- Route53
- ACM

## アプリ画像
- トップ画面
<img width="1104" alt="スクリーンショット 2023-09-18 11 20 31" src="https://github.com/Kenny-Chrysostomus/Inuminati/assets/104039651/1a7a1364-0a81-4d17-a4bc-07c3dd95405d">
- 投稿詳細
<img width="1095" alt="スクリーンショット 2023-09-18 11 21 23" src="https://github.com/Kenny-Chrysostomus/Inuminati/assets/104039651/16ee0270-2a44-4712-acde-d86d0d5d21ae">
- スマホ版
<img width="1095" alt="スクリーンショット 2023-09-18 11 21 23" src="https://github.com/Kenny-Chrysostomus/Inuminati/assets/104039651/11f6a067-bc48-4e84-b188-ddc08fbe2645">



