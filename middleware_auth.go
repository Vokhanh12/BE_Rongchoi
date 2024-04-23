package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	firebase "firebase.google.com/go"
	"github.com/vokhanh12/RongchoiApp/backend/internal/auth"
	"github.com/vokhanh12/RongchoiApp/backend/internal/database"
	"google.golang.org/api/option"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

// Hàm callback trả về một slice byte chứa khoá bí mật
func (apiCfg *apiConfig) middlewareAuthAPIKey(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			log.Fatal("error can't get user by api key: %v", err)
		}

		handler(w, r, user)

	}
}
func (apiCfg *apiConfig) middlewareAuthBearer(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		opt := option.WithCredentialsFile("json")
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			log.Fatalf("error initializing app: %v", err)
		}

		authFirebase, err := app.Auth(context.Background())
		if err != nil {
			log.Fatalf("error auth app : %v", err)
		}

		token, err := auth.GetBearer(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		// Authenticate users using tokens sent from the client.
		userToken, err := authFirebase.VerifyIDTokenAndCheckRevoked(context.Background(), token)
		if err != nil {

			// Kiểm tra xem lỗi có phải là "ID token has been revoked" hay không
			if err.Error() == "ID token has been revoked" {

				log.Fatalf("ID token has been revoked: %v", err)
			} else {

				// Tách chuỗi bằng dấu cách
				parts := strings.Split(err.Error(), " ")

				// Lấy phần số ở phần cuối cùng của chuỗi
				expiredAtStr := parts[len(parts)-1]

				// Chuyển đổi chuỗi số thành số nguyên
				expiredAt, err := strconv.ParseInt(expiredAtStr, 10, 64)
				if err != nil {
					fmt.Println("Failed to parse expired time:", err)
					respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid token"))
					return
				}

				// Giả sử exp là thời điểm hết hạn của token
				exp := expiredAt

				// Chuyển đổi thời điểm Unix thành đối tượng thời gian
				expirationTime := time.Unix(exp, 0)

				// Xử lý các trường hợp lỗi khác
				log.Printf("The expiration time of the token: %v, [%s]", err, expirationTime)
				respondWithError(w, http.StatusBadRequest, fmt.Sprintf("The expiration time of the token"))
				return
			}

			respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid token"))
			return

		}

		// This email is used to compare data in the database to retrieve user information.
		emailIdentity, ok := userToken.Firebase.Identities["email"]
		if !ok {
			fmt.Println("Error: Email identity not found")
			return
		}

		// Get email
		_email := strings.Trim(fmt.Sprintf("%v", emailIdentity), "[]")

		// Get the API key of a user by email.
		apiKey, err := apiCfg.DB.GetAPIKeyByEmail(r.Context(), _email)

		// Check error for login
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't login: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to get user: %v", err))
			return
		}

		handler(w, r, database.User{
			ID:            user.ID,
			CreateAt:      user.CreateAt,
			UpdateAt:      user.UpdateAt,
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			Email:         user.Email,
			NickName:      user.NickName,
			NumberPhone:   user.NumberPhone,
			DayOfBirth:    user.DayOfBirth,
			Address:       user.Address,
			Role:          user.Role,
			ApiKey:        user.ApiKey,
			ApiIat:        user.ApiIat,
			ApiExp:        user.ApiExp,
			RefreshApiKey: user.RefreshApiKey,
			RefApiIat:     user.RefApiIat,
			RefApiExp:     user.RefApiExp,
		})
	}
}
