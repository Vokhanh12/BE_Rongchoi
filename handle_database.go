package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func handlerUpdateAPIUsers(db *sql.DB) {
	// Thiết lập lập lịch để cập nhật API key mỗi giờ
	tickerAPIKey := time.NewTicker(time.Hour * 1)
	defer tickerAPIKey.Stop()

	// Thiết lập lập lịch để cập nhật Refresh API key mỗi hai giờ
	tickerRefAPIKey := time.NewTicker(time.Hour * 2)
	defer tickerRefAPIKey.Stop()

	// Khởi chạy goroutine để cập nhật API key
	go func() {
		for {
			select {
			case <-tickerAPIKey.C:
				if err := updateAPIKey(db); err != nil {
					log.Println("Failed to update API key:", err)
				}
			}
		}
	}()

	// Khởi chạy goroutine để cập nhật Refresh API key
	go func() {
		for {
			select {
			case <-tickerRefAPIKey.C:
				if err := updateRefAPIKey(db); err != nil {
					log.Println("Failed to update Refresh API key:", err)
				}
			}
		}
	}()

	// Để main goroutine không kết thúc ngay lập tức
	select {}
}

func updateAPIKey(db *sql.DB) error {
	_, err := db.Exec("UPDATE users SET api_key = encode(sha256(random()::text::bytea), 'hex')")
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE users SET api_iat = NOW()")
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE users SET api_exp = NOW() + INTERVAL '1 hour'")
	if err != nil {
		return err
	}

	fmt.Println("API key updated successfully")
	return nil
}

func updateRefAPIKey(db *sql.DB) error {
	_, err := db.Exec("UPDATE users SET refresh_api_key = encode(sha256(random()::text::bytea), 'hex')")
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE users SET ref_api_iat = NOW()")
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE users SET ref_api_exp = NOW() + INTERVAL '1 hour'")
	if err != nil {
		return err
	}

	fmt.Println("Refresh API key updated successfully")
	return nil
}
