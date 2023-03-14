package repositories_test

import (
	"encoder/application/repositories"
	"encoder/domain"
	"encoder/framework/database"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestJobRepositoryInsertAndFind(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "/test"
	video.CreatedAt = time.Now()

	repoVideo := repositories.NewVideoRepositoryDb(db)
	repoVideo.Insert(video)

	job, err := domain.NewJob("output_path", "Pending", video)
	require.Nil(t, err)

	repoJob := repositories.NewjobRepositoryDb(db)

	repoJob.Insert(job)

	j, err := repoJob.Find(job.ID)

	require.Nil(t, err)
	require.NotEmpty(t, j)
	require.Equal(t, j.ID, job.ID)
	require.Equal(t, j.VideoID, video.ID)
}

func TestJobRepositoryUpdate(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "/test"
	video.CreatedAt = time.Now()

	repoVideo := repositories.NewVideoRepositoryDb(db)
	repoVideo.Insert(video)

	job, err := domain.NewJob("output_path", "Pending", video)
	require.Nil(t, err)

	repoJob := repositories.NewjobRepositoryDb(db)
	repoJob.Insert(job)

	job.Status = "Complete"

	repoJob.Update(job)

	j, err := repoJob.Find(job.ID)

	require.Nil(t, err)
	require.NotEmpty(t, j)
	require.Equal(t, j.Status, job.Status)
}
