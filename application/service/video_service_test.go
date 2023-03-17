package service_test

import (
	"encoder/application/repositories"
	"encoder/application/service"
	"encoder/domain"
	"encoder/framework/database"
	"log"
	"testing"
	"time"

	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func prepare() (*domain.Video, *repositories.VideoRepositoryDb) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.FilePath = "teste.mp4"
	video.CreatedAt = time.Now()

	repo := repositories.NewVideoRepositoryDb(db)

	repo.Insert(video)

	return video, repo
}

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error on loading .env file")
	}
}

func TestVideoService(t *testing.T) {
	video, repo := prepare()

	videoService := service.NewVideoService(video, repo)

	err := videoService.Download("go-encoder-videos")

	require.Nil(t, err)

	err = videoService.Fragment()
	require.Nil(t, err)
}
