package migration

import (
	"context"
	"log"
	"user-manages/config"
	"user-manages/modules/user"
	"user-manages/pkg/database"
	"user-manages/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func userDbConn(pctx context.Context, cfg *config.Config) *mongo.Database {
	return database.DbConn(pctx, cfg).Database("user_db")
}

func UserMigrate(pctx context.Context, cfg *config.Config) {
	db := userDbConn(pctx, cfg)
	defer db.Client().Disconnect(pctx)

	col := db.Collection("users")

	// indexs
	indexs, _ := col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{"_id", 1}}},
		{Keys: bson.D{{"email", 1}}},
	})
	log.Println(indexs)

	documents := func() []any {
		roles := []*user.User{
			{
				Email: "player001@sekai.com",
				Password: func() string {
					// Hashing password
					hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
					return string(hashedPassword)
				}(),
				Name:      "Player001",
				CreatedAt: utils.LocalTime(),
			},
			{
				Email: "player002@sekai.com",
				Password: func() string {
					// Hashing password
					hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
					return string(hashedPassword)
				}(),
				Name:      "Player002",
				CreatedAt: utils.LocalTime(),
			},
			{
				Email: "player003@sekai.com",
				Password: func() string {
					// Hashing password
					hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
					return string(hashedPassword)
				}(),
				Name:      "Player003",
				CreatedAt: utils.LocalTime(),
			},
			{
				Email: "admin001@sekai.com",
				Password: func() string {
					// Hashing password
					hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
					return string(hashedPassword)
				}(),
				Name:      "Player003",
				CreatedAt: utils.LocalTime(),
			},
		}

		docs := make([]any, 0)
		for _, r := range roles {
			docs = append(docs, r)
		}
		return docs
	}()

	results, err := col.InsertMany(pctx, documents, nil)
	if err != nil {
		panic(err)
	}
	log.Println("Migrate auth completed: ", results)
}
