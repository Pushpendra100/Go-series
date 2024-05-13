# Complete Go series
This is the complete go series by Hitesh Choudhary on YouTube. All chapters are covered with comments for more explanation of the code.

I am listing down some extra points and some commands from the series :-

1. ```go mod init _package_name_```   
 provides the project with go.mod file
2. ```go run main.go```    
is used to run the program
3. ```go help```
4. ```go help _command_```         
 it gives more info about that particular command
5. ```go env```     
    You can use the go env command to portably set the default value for an environment variable for future go commands

6. ```go build```    
    The go build command compiles the packages, along with their dependencies, but it doesn't install the results. The go install command compiles and installs the packages.

7. ```GOOS="windows" go build```     
build for windows - gives .exe file
similarly can do for linux, darwin(mac)

8. ```go get```     
The go get command is primarily used for retrieving remote packages from version control repositories and making them available for use in your projects

9. ```go mod tidy```  
Tidy makes sure go. mod matches the source code in the module. It adds any missing modules necessary to build the current module's packages and dependencies, and it removes unused modules that don't provide any relevant packages. It also adds any missing entries to go. sum and removes any unnecessary ones.

10. ```go mod verify```    
You may want to add “go mod verify” after download, so you can be sure your go.sum file matches what you downloaded and someone didn’t “git push — force” over a tag you’re using into something hacked.


11. ```go list -m all```   
give the list of all the dependencies

12. ```go list -m -versions github.com/gorilla/mux```   
list down all the versions available of the package

13. ```go mod why github.com/gorilla/mux```    
gives us the list of packages that are dependent on the package github.com/gorilla/mux

14. ```go mod graph```   
gives us the dependency of each other 

15. ```go mod edit```    
    to change the go.mod file from command line
        -> go mod edit -go 1.16 
        to change the go version
        -> 

16. ```go mod vendor```  
    just like node_modules, instead of storing in cache of the go locally, you are getting packages directly in working folder
    To first see packages in vendor instead of cache or web, we have to run following command
    -> go run -mod=vendor main.go

    --------

- **LEXER**
    - make sure that you are following the grammar of the language
    - go has lexer, and it handles semicolons itself

- **TYPES**
    - case sensitive; almost
    - Variable type should be known in advance
    - Everything is a type

    - String, Bool, Integer, Floating, Complex
    - Array, Slices, Maps, Structs, Pointers
    - Functions, Channels

- **GOOD RESOURCE** - https://go.dev/ref/spec

- **GO PACKAGES** - pkg.go.dev



- **MEMORY MANAGEMENT**
    - memory allocation and deallocation happens automatically
    - new()
        - allocate memory but no INIT
        - you will get a memory address
        - zeroed storage
    - make()
        - allocate memory and INIT
        - you will get a memory address
        - non-zeroed storage
    - GC(garbage collection) happens automatically
    - Out of Scope or nil
    - https://pkg.go.dev/runtime

- **GO PLAYGROUND** - https://go.dev/play/
(there are fixed value in Go Playground regarding the seed value.)

- **EXPORT** - if anything that starts with a capital letter simply means that is exportable

- **DEFER** - A "defer" statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because the surrounding function executed a return statement, reached the end of its function body, or because the corresponding goroutine is panicking

- **GO MONGODB DRIVERS** - https://go.mongodb.org/

- **CONTEXT PACKAGE** - Package context defines the context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes


> There are many more topics covered, take a look at 28 directories ...

**Happy coding :)**