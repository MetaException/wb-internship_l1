package main

import "fmt"

type APIInfoGetter interface {
	getData()
}

type SomeApi1 struct {
	ApiConf string
	Logger  interface{}
}

func (api *SomeApi1) fetchInformation() string {
	//...
	fmt.Println("Using api1")
	return "some_data"
}

type SomeApi2 struct {
	Logger    interface{}
	AuthCreds string
}

func (api *SomeApi2) getData() string {
	//...
	fmt.Println("Using api2")
	return "some_data"
}

type API2Adapter struct {
	Adaptee *SomeApi1
}

func (api API2Adapter) getData() {
	fmt.Println("Using converted api1")
	api.Adaptee.fetchInformation()
}

func main() {
	api1 := &SomeApi1{
		ApiConf: "123",
		Logger:  "iamlogger",
	}

	var api2 APIInfoGetter = &API2Adapter{
		Adaptee: api1,
	}

	api2.getData()
}
