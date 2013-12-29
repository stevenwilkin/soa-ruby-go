# soa-ruby-go

A simple example of developing an API and providing a library to access that API.

## Starting the API 

Ensure [Go](http://golang.org/) is installed and your ``$GOPATH`` is set, then run the following:

	go get .
	go run ./api.go

## Run the demo script

Ensure [Bundler](http://bundler.io/) and an appropriate Ruby interpreter are available and run:

	bundle install
	./demo.rb

The API stores items in-memory and the demo script uses the library in ``item.rb`` to add and then manipulate some data. Output should be like the following:

	$ ./demo.rb
	> Create items:
	first
	second
	third
	forth
	
	> List items:
	1 - first
	2 - second
	3 - third
	4 - forth
	
	> Retrieve item 2:
	2 - second
	
	> Delete item 3:
	true
	
	> Delete item 3 again:
	false
	
	> List items:
	1 - first
	2 - second
	4 - forth
	
	> Update item 2:
	true
	
	> List items:
	1 - first
	2 - very second
	4 - forth
