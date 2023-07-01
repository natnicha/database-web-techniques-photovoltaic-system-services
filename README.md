# Photovoltaic System Services

## Installing the project on your machine

1 - Install Go by following this instruction
~~~
https://go.dev/doc/install
~~~

2 - Install dependencies using this command
~~~
go build
~~~

3 - Tidy your project by this command. It will remove dependencies and tidy your project
~~~
go mod tidy
~~~
or tidy using Makefile on Windows
~~~
MinGW32-make tidy 
~~~

## Setting up the database destination
1 - Set a database destination in a .env file 


## Setting up a Python destination
1 - Set a Python destination in a .env file. 

For the best practise, you should install Photovoltaic System Batch first. Install all dependencies and go back to this project to set Python path in a .env file. Otherwise, making a request through `/api/v1/project/generate-report/{id}` and `/api/v1/product/generate-report/{id}` would fail.

## Running the project on your machine
1 - Run this command
~~~
go run main.go
~~~
or run using Makefile on OSX
~~~
make run 
~~~
or run using Makefile on Windows
~~~
MinGW32-make run 
~~~
2 - Enjoy the service on specified port in a .env file, default at http://localhost:8000 


