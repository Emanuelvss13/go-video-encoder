package service

import (
	"context"
	"encoder/application/repositories"
	"encoder/domain"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"cloud.google.com/go/storage"
)

type VideoService struct {
	Video           *domain.Video
	VideoRepository repositories.VideoRepository
}

func NewVideoService(video *domain.Video, repo repositories.VideoRepository) VideoService {
	return VideoService{
		Video:           video,
		VideoRepository: repo,
	}
}

func (v *VideoService) Download(bucketName string) error {

	ctx := context.Background()

	client, err := storage.NewClient(ctx)

	if err != nil {
		return err
	}

	bucket := client.Bucket(bucketName)
	obj := bucket.Object(v.Video.FilePath)

	r, err := obj.NewReader(ctx)

	if err != nil {
		return err
	}

	defer r.Close()

	body, err := ioutil.ReadAll(r)

	if err != nil {
		return err
	}

	file, err := os.Create("/tmp" + "/" + v.Video.ID + ".mp4")

	if err != nil {
		return err
	}

	_, err = file.Write(body)

	fmt.Println(file.Name())

	if err != nil {
		return err
	}

	defer file.Close()

	log.Printf("video %v has been stored", v.Video.ID)

	return nil
}

func (v *VideoService) Fragment() error {
	err := os.Mkdir(os.Getenv("LOCAL_STORAGE_PATH")+v.Video.ID, os.ModePerm)

	if err != nil {
		return err
	}

	source := os.Getenv("LOCAL_STORAGE_PATH") + "/" + v.Video.ID + ".mp4"
	target := os.Getenv("LOCAL_STORAGE_PATH") + "/" + v.Video.ID + ".frag"

	cmd := exec.Command("mp4fragment", source, target)

	output, err := cmd.CombinedOutput()

	if err != nil {
		return err
	}

	PrintOutput(output)

	return nil
}

func PrintOutput(out []byte) {
	if len(out) > 0 {
		fmt.Printf("=> output: %s\n", string(out))
	}
}
