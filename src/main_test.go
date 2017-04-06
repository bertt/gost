package main

import (
	"fmt"
	"strconv"
	"testing"

	net "net/http"

	"github.com/geodan/gost/src/configuration"
	"github.com/geodan/gost/src/database/postgis"
	"github.com/geodan/gost/src/http"
	"github.com/geodan/gost/src/mqtt"
	"github.com/geodan/gost/src/sensorthings/api"
	"github.com/stretchr/testify/assert"
)

func TestVersionHandler(t *testing.T) {
	//arrange
	server := "localhost"
	port := 8088

	cfg := configuration.Config{}
	mqttServer := mqtt.CreateMQTTClient(configuration.MQTTConfig{})
	database := postgis.NewDatabase("", 123, "", "", "", "", false, 50, 100, 200)
	api := api.NewAPI(database, cfg, mqttServer)
	api.Start()

	gostServer := http.CreateServer(server, port, &api, false, "", "")
	go gostServer.Start()
	versionURL := fmt.Sprintf("%s/Version", "http://"+server+":"+strconv.Itoa(port))

	fmt.Println(versionURL)
	// act
	request, _ := net.NewRequest("GET", versionURL, nil)
	res, _ := net.DefaultClient.Do(request)

	//assert
	assert.Equal(t, 200, res.StatusCode, "result should be http 200")

	// teardown
	gostServer.Stop()
}
