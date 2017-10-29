- Introduction
	- A Taste of Go
		- Go playground
		- "Yeah, it works!" (instead of hello-word)		
	- What Makes Go Special?
		- Simple
		- Opinionated
		- First Class Concurrency
	- Go Origins
		- C
		- Limbo
		- Google
	- Who Uses Go?
	- The Go community
	- Say Hi to Juvenile Delinkcious (the sample app the book follows)
		- A delicious-like link manager
	
- Get Up and Running
	- Installation
	- Go in a box (Complete Go environment in a Docker container)	
	- Development Environment Options
		- Gogland
		- Visual Studio Code
		- Your favorite text editor		
	- Tooling
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
	- Troubleshooting and debugging
		Delve

- The Go Type System	
	- Types	
		- Primitive types
		- Strings
		- Arrays and Slices
		- Maps
		- Interfaces
		- Structs
		- Functions
		- Pointers
	- Type declaration and Aliasing		
	- Type Assertions
	- Reflection
	
- Foundations
	- Naming and visibility
	- Variables and Assignment
	- Type Inference
	- Control Flow
	- Iterations

- Modularity
	- Packaging
	- Importing Packages
	- Package Hygiene
	
			
- Go as a Multi-Paradigm Language
	- Structured Programming
	- Object-Oriented Programming
	- Functional Programming
	- Generic Programming
	- Concurrent Programming
		- Parallel vs. Concurrent

- Structured Programming in Go
	- Functions

- Functional Programming in Go
	- Functions as first class citizens in Go
	
- Generic Programming
	- Built-in generic containers
	- The empty interface
		- An interface is a pointer
		- Objects that operate at the generic level
		- Using an empty interface
		- Interfaces with empty interface arguments and specialized implementations
	
	- Code Generation of specialized containers
	- Build your own set
		- Set of empty interfaces
		- Set with generic interface and specialized implementations
			- The Set interface accepts empty interface
			- No type-safety
			- Need to cast
		- Generating strongly-typed sets
			- All the benefits
			- Need to use a code generator
	 	
- Advanced Object-Oriented Design
    - Quick Introduction to Object-Oriented Programming in Go
    	- Structs and Methods
    	- Visibility
    	- The Receiver
    	
    - Struct Embedding
    	- Unnamed Struct Field
    	- Poorman's Inheritance
    	- Multiple inheritance
    	- Pointer Embedding
    	
    - SOLID design in Go
		- Single Responsibility Principle
		- Open / Closed Principle
		- Liskov Substitution Principle
		- Interface Segregation Principle
		- Dependency Inversion Principle
        
	- Interfaces FTW
		- Primary Abstraction Mechanism
		- Duck interfaces
	- The WINIWYG principle
		- Dependency injection a.k.a 3rd party binding
	- Factories
	    - Singletons
	- Data-only structs
		- No methods
		- Fields may contain any object
		- Prefer primitive types, standard library types and interfaces
	- Modeling your domain
	  - Primary Entities - User, Link and Tag
	  - Primary operations
	- Object-oriented Design in Go vs. Other Languages
		- Go vs. C
			- Go structs vs. C structs
		- Go vs. C++
			- Go interfaces vs. struct with pure virtual functions
		- Go vs. Java/C#
			- Go empty interface vs. Java/C# Object
		- Go vs. Python
			- Go packages vs. Python namespaces
	
- Industrial-Strength Testing in Go
	- Unit Testing with the testing package
	- Code Coverage
	- BDD
	- Ginkgo and Gomega
	- Mocking
		- With Interfaces
		- Replacing function variables
	- Integration Testing
	- System Testing
		- End to end testing
		- Performance Testing
		- Load Testing

- Developing Plugin-Based Systems
	- Dynamic Plugins
	- Writing Plugins
	- Loading Plugins
	- Using Plugins	
	
- Concurrent Programming in Go
	- CSP
	- Goroutines
	- Channels
	- Synchronization
	- Contexts
	- Race Detection
	- The Go memory object model

- Error Handling
	- Multiple return values
		- Named return values
	- The Error interface
	- No exceptions, except...
		- Panic
		- Defer
		- Recover
	
	- Concurrent Error Handling
		- Error channel
	- Checking for errors
		- Always assign the return values of a function
		- Use underscores to explicitly ignore irrelevant values
		
	- Advanced Error Handling with Merry
		- Stack traces FTW
		- Wrapping and Unwrapping
		- Using pre-defined errors
		
	- Is it possible to emulate exception handling with Go?
	
- Writing Command-Line Tools
	- Flags
- Database Programming
	- Working with Relational Databases
		- To ORM or not to ORM
		- sqlx
		- 
	- Working with NoSQL Databases
		- Etcd3

- Writing GRPC-Based Microservices
	- Quick Introduction to GRPC
	- Why GRPC?
	- GRPC Proto files 
	- Designing and implementing the GRPC Service
	- Consuming GRPC

- Exposing GraphQL-Based APIs 
	- Quick Introduction to GraphQL
	- Why GraphQL?
	- The GraphQL Schema
	- 

- Dockerizing a Go Application
	- Quick Introduction to Docker
	- Why Docker?
	- The Dockerfile
	- Using Docker compose
	- Deploying to the cloud
	- Considering Orchestration
		- Kubernetes
		- Docker Swarm
		- AWS ECS
	
- Welcome to the Future
	- Go Warts
	- Go 2.0
	- Plugins
	- Stability

