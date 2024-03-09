package supabase

import (
	"mime/multipart"
	"os"

	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
)

type SupabaseInterface interface {
	UploadFile(file *multipart.FileHeader) (string, error)
}

type Supabase struct {
	client *supabasestorageuploader.Client
}

func Init() SupabaseInterface {
	supClient := supabasestorageuploader.New(
		os.Getenv("SUPABASE_URL"),
		os.Getenv("SUPABASE_TOKEN"),
		os.Getenv("SUPABASE_BUCKET_NAME"),
	)

	return &Supabase{
		client: supClient,
	}
}

func (s *Supabase) UploadFile(file *multipart.FileHeader) (string, error) {
	link, err := s.client.Upload(file)

	if err != nil {
		return link, err
	}

	return link, nil
}
