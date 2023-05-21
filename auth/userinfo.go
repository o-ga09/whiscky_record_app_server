package auth

import (
	"context"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func GetUserInfo(ctx context.Context, token string) (string, error) {
    // Firebase Admin SDKの初期化
    opt := option.WithCredentialsJSON([]byte(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")))

    app, err := firebase.NewApp(ctx, nil, opt)
    if err != nil {
        return "", err
    }
    client, err := app.Auth(ctx)
    if err != nil {
        return "", err
    }

    // アクセストークンの検証
    tokenInfo, err := client.VerifyIDToken(ctx, token)
    if err != nil {
        return "", err
    }

    // ユーザー情報の取得
    userInfo, err := client.GetUser(ctx, tokenInfo.UID)
    if err != nil {
        return "", err
    }
    return userInfo.UID, nil
}