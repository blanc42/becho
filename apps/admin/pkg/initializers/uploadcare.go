package initializers

import (
	"log"
	"os"

	"github.com/uploadcare/uploadcare-go/ucare"
)

func GetUploadcareClient() ucare.Client {
	creds := ucare.APICreds{
		SecretKey: os.Getenv("UPLOADCARE_PRIVATE"),
		PublicKey: os.Getenv("UPLOADCARE_PUBLIC"),
	}

	conf := &ucare.Config{
		SignBasedAuthentication: true,
		APIVersion:              ucare.APIv06,
	}

	client, err := ucare.NewClient(creds, conf)
	if err != nil {
		log.Printf("=>>>>>>>>>>>>>>>>> creating uploadcare API client: %s", err)
	}

	return client
}
