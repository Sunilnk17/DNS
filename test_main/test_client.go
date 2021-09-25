package main

import (
	"drone-navigation-service/app/config"
	"drone-navigation-service/client"
	"drone-navigation-service/client/operations"
	"fmt"
)

//swagger generate spec -b github.com/cloudqwest/oas-service -o ./swagger.json --scan-models
//swagger generate client -f swagger.json
//use to generate client
//install swagger-go from https://goswagger.io/install.html

func main() {

	config.Initialize()
	param := operations.NewGetLocationParams()
	param.SectorID = "12345"
	param.CompanyID = "atlas"
	param.X = "12.2"
	param.Y = "1.4"
	param.Z = "1.4"
	param.Vel = "1.5"

	client := client.NewHTTPClientWithConfig(nil, client.DefaultTransportConfig().WithHost("localhost:5010"))
	client.Operations.SetTransport(client.Transport)

	location, err := client.Operations.GetLocation(param)
	if err == nil {
		fmt.Print(location.Payload.Data)
	} else {
		data, ok := err.(*operations.GetLocationDefault)
		if ok {
			fmt.Print(data.Payload.Error)
		}
	}
}
