# Overview

Welcome to "Modern Go Programming"! In this short overview I'll tell what the book is about, 
who I am, who you are, why you should read this book and how the book is organized.

## What is this Book About?

This book is about programming Go. It covers everything you need to know if you're just starting 
with Go. But, it really shines compared to other Go books when it comes to more advanced topics.
The goal is to provide you a lot of value that you can't glean just from reading the documentation.  
For example, I cover topics like gRPC-based microservices and GraphQL APIs. In addition to the 
specifics of the Go language and its libraries I also weave in guidelines and advice on best 
practices that transcend any programming language like proper object-oriented design and testing
mindset. 

## Who am I?

I'm Gigi Sayfan. I've been developing software professionally for 22 years in domains as diverse as:
- Instant messaging
- Morphing
- Chip fabrication process control
- Embedded multimedia applications for game consoles
- Brain-inspired machine learning
- Custom browser development
- Web services for 3D distributed game platforms
- IoT sensors
- Virtual reality
- Genomics and bioinformatics

I wrote production code every day in many languages such as: 

- Go
- Python
- C#, 
- C/C++,
- Java
- Javascript
- Delphi
- Cobol(back in the day)

The code I wrote that runs on:
- Windows, 
- Linux
- macOS
- Android
- Lynx (embedded)
- Sony PlayStation and GearVR

My technical expertise includes:  
- Distributed systems
- Data stores
- Containers and orchestration
- Low-level networking
- Unorthodox user interfaces
- General software development life cycle.

For me programming and software development are much more than just a job. It is my passion. I'm 
constantly learning by playing with new technologies, writing articles and books. I started writing
technical articles in [2001](http://www.drdobbs.com/a-quick-test-window-for-com/184416298) when 
technical magazines were printed and I never stopped. You can find my articles here:

- [Dr. Dobbs](http://www.drdobbs.com/sitesearch?queryText=gigi+sayfan&type=site)
- [DevX.com](https://www.google.com/search?q=gigi+sayfan+site%3Adevx.com)
- [Compose.io](https://www.compose.com/articles/search/?s=gigi%20sayfan)
- [TutsPlus.com](https://tutsplus.com/authors/gigi-sayfan) 

I also recently wrote the book [Mastering Kubernetes](https://www.amazon.com/Mastering-Kubernetes-container-deployment-management/dp/1786461005). Check it out. 


I have life besides programming and writing (hard to believe I know). I live in Davis, California. I
have a georgeous wife and three lovely children. I tried to infect them with programming too with pretty good results. Saar, my oldest son, is now a CS student and had a couple of contruct job as a software engineers. Guy, my youngest son, is now getting into game development using Python (after trying Lua). I even managed to persuade Liat, my wife, to learn some Python and SQL abd write some programs to automate Excell and SPSS. My only utter failure is Ophir, my daughter, who prefers animals to computers and is an animal science student at Cal Poli. 

I used to play a lot of basketball and did Crossfit for several years. But, with great age come 
great injuries and in the last 5 years my athletic outlet is Brazilian Jiu Jitsu that I train with 
my sons Saar and Guy.

## Who Are You?

Enough about me. What about you? You can be anyone of course, but when I write this book I address 
in my mind a software developer who is enthusiastic about Go and has some programming experience 
beyond just school projects. You don't have to know anything about Go coming in (except that it's cool and you want to know more). If you have little experience you can still benefit, but you may need to refer to some external sources to explain basic concepts that I'll just mention without detailed explanations (variables, conditionals, loops, etc).

If you have a lot of professional exeprience programming in other mainstream languages like Python, 
C++, Java or C# then you'll appreciate the comparisons I make throuout the book to other languages.  

## Why Should you Read This Book?

There are already several other Go books out there. But, "Modern Go Programming" is different. It
is designed to be easy to learn, easy to use and provide value in real-world situations. I attack 
every subject from multiple angles: conceptual level, comparison to other languages, simple code
examples, in the contenxt of a large system. 

### Concepts

I start every topic at the conceptual level. This is very important for grounding

### Compare to other languages

If you have experience in another language then concepts like iteration or pointers can be made 
clear by comparing them to the concepts in a familiar language. Sometimes pointing out the 
differences in philosophy, design or implementation between languages is the key to understanding.

### Simple Examples

I will provide simple stand-alone code examples for every topic. Talk is cheap, but code is king.

### Full-Fledged Sample Application

The whole is often more than the sum of its parts. It is often frustrating when you try to take some
toy example code and incorporate it into a large system. Sometimes, standalone examples are too 
abstract to understand how to integrate a particular concept into a real system. Here, in addition
to stand-alone examples we will also build together a non-trivial system that consists of multiple
collaborating services and data stores. 

### Real-World, From the Trenches Author

Everything I write is based on direct experience, building large-scale systems. deployed in 
production, evolving over years, in multi-developer teams. I'm not an academic and I'm not a 
consultant brought in for a quick fix. I'm a full time software-architect working every day at the 
code level on enterprise-grde software systems.   

### Always Up to Date

I plan to keep this book up to date and extend it as the Go language evolves and new best practices 
or technologies become popular.

### Docker

To make sure, you can always try the code I will provide Docker images with all the necessary 
dependencies and instructions how to run all the code. I'll update the Docker image along with the
rest of the book.

### Github

I write this book on Github and you can participate opening issues, commenting and PRs. I welcome
discussion and contributions.

### No homework

You will not find the dreaded exercise toi the reader here! 

## How is the Book Organized?

The book is organized very simply. The first part up to and including chapter 6 introduces the Go 
language, its basic features, its tooling and sets you up write simple Go programs. The second part
is composed of chapters that cover some topic in depth. Some of the topics may have been covered 
briefly earlier. The order is pretty arbitrary. Throughout the second part we will build the Delinkcious sample application together. 

