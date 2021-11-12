#### What  
From the Golang site's create a Go module tutorial.  
[Golang tutorial](https://golang.org/doc/tutorial/create-module)  

[Golang tour](https://tour.golang.org/)  

Eventually more could be added.  

#### Misc  

`go build`  
Compiles into executable in PWD  

`go install`  
Compiles and installs packages into executable in go install path  

`go list -f '{{.Target}}'`  
Display install path  


#### Tests/Testing  

##### Go unit tests  

`go test`  
Naming files filename*_test.go* allows the `go test` command to test functions  
Functions to test need to start with *Test*functionName  
