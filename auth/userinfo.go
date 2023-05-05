package auth

import (
	"context"

	firebase "firebase.google.com/go"
)

func GetUserInfo(ctx context.Context, token string) (string, error) {
    // Firebase Admin SDKの初期化
    app, err := firebase.NewApp(ctx, nil)
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