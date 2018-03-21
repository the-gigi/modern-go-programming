# Get Up and Running

In this chapter you'll get ready to roll. I'll discuss local installation on macOS and Windows, using a ready to go Docker image sporting a Linux distro, some IDEs you may like, the awsome tooling that comes with Go, how to organize your code and finally some troubleshooting and debugging advice. At the end of this chapter, you'll understand the options for local development and you'll have picked and set up at least one environment. You'll be ready to go!

## Installation

Go is cross-platform and can be installed on macOS (that's what they call Mac OSX these days), Linux and Windows. Just follow the instructions and you'll be fine: https://golang.org/doc/install

On Mac you may prefer to use [homebrew](https://brew.sh/). I do. This option is not mentioned in the docs. Just type: `brew install go`

That's all there is to it. To find out where Go is installed on your system type (Mac and Linux):

```
> ls -la `which go`
lrwxr-xr-x  1 gigi.sayfan  admin  23 Sep 19 16:44 /usr/local/bin/go -> ../Cellar/go/1.9/bin/go
```

If you're using Windows. It's probably "C:\go". You can find out for sure by opening a PowerShell window and using the Get-Command command:

```powershell
﻿PS C:\Users\gigi> get-command go

CommandType     Name                                                Definition
-----------     ----                                                ----------
Application     go.exe                                              C:\Go\bin\go.exe
```

To verify everything is fine type `go` in a console window. You should see this:

```
> go
Go is a tool for managing Go source code.

Usage:

	go command [arguments]

The commands are:

	build       compile packages and dependencies
	clean       remove object files
	doc         show documentation for package or symbol
	env         print Go environment information
	bug         start a bug report
	fix         run go tool fix on packages
	fmt         run gofmt on package sources
	generate    generate Go files by processing source
	get         download and install packages and dependencies
	install     compile and install packages and dependencies
	list        list packages
	run         compile and run Go program
	test        test packages
	tool        run specified go tool
	version     print Go version
	vet         run go tool vet on packages

Use "go help [command]" for more information about a command.

Additional help topics:

	c           calling between Go and C
	buildmode   description of build modes
	filetype    file types
	gopath      GOPATH environment variable
	environment environment variables
	importpath  import path syntax
	packages    description of package lists
	testflag    description of testing flags
	testfunc    description of testing functions

Use "go help [topic]" for more information about that topic.
```

To make sure it really works let's run our little "Yeah, it works!" program. Save it to a file call `yeah.go`:

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Yeah, it works!")
}
```

Go can compile your programs with the `go build` command, but if you want to immediately run your program you can use the `go run` command that will compile and execute the program immediately.

```
> go run yeah.go
Yeah, it works!
```

## Go in a box (Complete Go environment in a Docker container)

If you're too lazy to click links, download files and type commands I got you. You can use a Docker image with Go pre-installed. Here is the fast way to go about it:

- Make sure [Docker](https://docs.docker.com/engine/installation/) is installed
- Use a pre-built Golang Docker image
- Launch the Docker image while mounting your source directory into the container
- Build or run your program.

Here is a one-liner that builds and runs the yeah.go program inside a docker container given you have Docker installed.

`docker run --rm -v "$PWD":/go/src/yeah -w /go/src/yeah golang:1.9 go run yeah.go`

Let's understand what's going on here by explaining each element:

- docker run - the docker command that executes commands inside docker containers
- -v "$PWD":/go/src/yeah - mounts the current directory on the host into the container at /go/src/yeah
- -w /go/src/yeah - set the working directory inside the container
- golang:1.9 - the Docker image to load (if not available locally it will be downloaded)
- go run yeah.go - the command to execute inside the container

The first time you run this command Docker will download the golang:1.9 image from the docker repository. The output should look like:

```
Unable to find image 'golang:1.9' locally
1.9: Pulling from library/golang
3e17c6eae66c: Pull complete
fdfb54153de7: Pull complete
a4ca6e73242a: Pull complete
93bd198d0a5f: Pull complete
2a43f474a764: Pull complete
e19893b2f35c: Pull complete
3b8a1a0cc426: Pull complete
85a9bedd68ab: Pull complete
Digest: sha256:57ef143775b37c6100d4be60ca0f9493e18c68cde3ea76c36a22f41857df11f4
Status: Downloaded newer image for golang:1.9
Yeah, it works!
```

But, running the command again will be much more concise:

```
> docker run --rm -v "$PWD":/go/src/yeah -w /go/src/yeah golang:1.9 go run yeah.go
Yeah, it works!
```

## Native Installation vs. Docker

There are pros and cons for each approach:

- Native installation is more streamlined
- Debugging native installation may be easier
- Running in disposable Docker containers makes it easy to run against multipel versions of Go
- Running in disposable Docker containers makes it easy to isoalte different environments without messing with environment variables
- Running in disposable Docker containers doesn't clutter your OS, file syste and environment variables
- Docker allows building a full-fledged environment for your application ready to be deployed

You may even use a hybrid approach where you run some programs natively and other in Docker containers.

Check out https://github.com/docker-library/docs/tree/master/golang for more options. 	

## Integrated Development Environments

I was always a big fan of IDEs. Early in my career, I used Microsoft's Visual Studio for C/C++ development as well as C#. I've used NetBeans for a short stint of Java development. I've used JetBrains' PyCharm for Python development for many years, and nowadays I use JetBrains Goland for Go development. But, Goland is not the only game in town. Visual Studio Code is a great Go IDE and has many followers. A lot of people just use text editors with various plugins. There are also several other IDEs that are less popular.

### Goland

Goland is the JetBrains Go IDE. It doesn't get any better IMO. Unfortunately, there is no community edition. So, if you want to use it you have to pay for it (or have your company pay for it). Goland has all the bells and whistles: 

- Project management
- Super-powerful code editor
- Lots of refactoring
- Go tools integrations
- Integrated testing including coverage
- Integrated debugging

But, it also has tons of other auxiliary features such as:

- Strong source control integration
- Built-in terminal
- Docker and vagrant integration
- Web development (HTML, CSS, Javascript)
- Database tools 
- Powerful git support
- Markdown support
- Lots of other plugins

![Here is a screenshot of Goland](images/chapter-2/goland.png)


### Visual Studio Code

I don't use [Visual Studio Code](https://code.visualstudio.com/) myself, but I know many people who swear by it. Since, Goland will cost you $$$ you may prefer Visual Studio Code as your Go IDE. What makes Visual Studio Code special? Here are a few tidbits:
- It is based on [Electron](https://en.wikipedia.org/wiki/Electron_(software_framework)), same as [Atom](https://en.wikipedia.org/wiki/Atom_(text_editor)), but it provides its own editor - [Monaco]()
- It was voted developer environment tool in the 2018 StackOverflow survey with 34% of the votes
- It was developed by Microsoft and released under the MIT license in 2015. The code is on [Github](https://github.com/Microsoft/vscode)



![Here is a screenshot of Goland](images/chapter-2/visual-studio-code.png)


### Your favorite text editor	
### Other Go IDEs

## Tooling

    The go command		
        build       compile packages and dependencies
        clean       remove object files
        doc         show documentation for package or symbol
        env         print Go environment information
        bug         start a bug report
        fix         run go tool fix on packages
        fmt         run gofmt on package sources
        generate    generate Go files by processing source
        get         download and install packages and dependencies
        install     compile and install packages and dependencies
        list        list packages
        run         compile and run Go program
        test        test packages
        tool        run specified go tool
        version     print Go version
        vet         run go tool vet on packages	
    x/tools
	
https://godoc.org/golang.org/x/tools	
	
## Organizing your Go code
		
## Troubleshooting and debugging
 
### Gogland Debugger 
### Delve

## Conclusion
