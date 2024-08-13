package main

import (
	"context"
	"fmt"
	"os"

	"github.com/crowdstrike/gofalcon/falcon"
	"github.com/crowdstrike/gofalcon/falcon/client/cspm_registration"
)

func main() {
	falconClientId := os.Getenv("FALCON_CLIENT_ID")
	falconClientSecret := os.Getenv("FALCON_CLIENT_SECRET")
	clientCloud := os.Getenv("FALCON_CLOUD")

	client, err := falcon.NewClient(&falcon.ApiConfig{
		ClientId:     falconClientId,
		ClientSecret: falconClientSecret,
		Cloud:        falcon.Cloud(clientCloud),
		Context:      context.Background(),
	})
	if err != nil {
		panic(err)
	}

	filter := "use_current_scan_ids:true"
	limit := int64(10)

	res, err := client.CspmRegistration.GetConfigurationDetectionIDsV2(
		&cspm_registration.GetConfigurationDetectionIDsV2Params{
			Context: context.Background(),
			Filter:  &filter,
			Limit:   &limit,
		},
	)

	fmt.Printf("Response: %v\n", res)
	fmt.Printf("Error: %v\n", err)
}