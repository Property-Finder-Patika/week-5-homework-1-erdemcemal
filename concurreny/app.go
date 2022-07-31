package main

import "fmt"

// App is the interface that defines the Run function
type App interface {
	Run() error
}

// runAdobe is the function that runs the application in a goroutine
func runAdobe(a *AdobeProxy, c chan int) {
	if err := a.Run(); err != nil {
		fmt.Println(err)
	}
	c <- 1
}

// License struct to hold the license key
type License struct {
	Key string
}

// Adobe is the application that we want to run
type Adobe struct{}

// Run is the function that runs the application
func (a *Adobe) Run() error {
	fmt.Println("Running Adobe")
	return nil
}

// AdobeProxy is the proxy that runs the application
type AdobeProxy struct {
	adobe   Adobe
	license *License
}

// NewAdobeProxy is the function that creates a new AdobeProxy
func NewAdobeProxy(license *License) *AdobeProxy {
	return &AdobeProxy{
		license: license,
	}
}

// Run is the function that runs the application
func (a *AdobeProxy) Run() error {
	// check license is in list of valid licenses
	if a.license.Key == "valid" {
		return a.adobe.Run()
	}
	return fmt.Errorf("license %s is not valid", a.license.Key)
}
