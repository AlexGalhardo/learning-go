<div align="center">
 <h1 align="center"><a href="https://go.dev/" target="_blank">Learning Go</a></h1>
</div>

## Introduction
- Repository to save my Go programming language studies notes and references.

## Install
- https://go.dev/doc/install

## Tools
- Playground: <https://go.dev/play/>
- Tour: <https://go.dev/tour/>
- <https://gobyexample.com/>
- <https://go-proverbs.github.io/>
- <https://www.asemanago.dev/>
- <https://golangweekly.com/>

## Libs and Frameworks

- HTTP
   - Gin: <https://gin-gonic.com/>
   - Fiber: <https://docs.gofiber.io/>
      - Examples: <https://github.com/gofiber/recipes>
   - Iris: <https://www.iris-go.com/>
- ORM
   - GoRM: <https://gorm.io/>
- Dotenv
   - GoDotEnv: <https://github.com/joho/godotenv>
- Live Reloading
   - Air: <https://github.com/air-verse/air>
- Admin Dashboard
   - GoAdmin: <https://www.go-admin.com/>
- Tests
   - Testify: <https://github.com/stretchr/testify>
- Utility
   - GoFunk: <https://github.com/thoas/go-funk>
   - GoUtil : <https://github.com/gookit/goutil>

## Online Courses

- YouTube
   - Aprenda Go PT-BR: <https://www.youtube.com/playlist?list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg>
- Udemy
   - Learn How To Code: Google's Go (golang) Programming Language: <https://www.udemy.com/course/learn-how-to-code>

## Articles
- Frequently Asked Questions (FAQ): <https://go.dev/doc/faq>
- Why did you create a new language? <https://go.dev/doc/faq#creating_a_new_language>

## Books
- [A Linguagem de Programação Go](https://www.amazon.com.br/Linguagem-Programa%C3%A7%C3%A3o-Go-Alan-Donovan-ebook/dp/B07KDDJQXS/ref=tmm_kin_swatch_0?_encoding=UTF8&qid=&sr=#customerReviews)

## Github
- <https://github.com/vkorbes/aprendago>
- <https://github.com/geiltonxavier/aprenda-go>
- <https://github.com/mmcgrana/gobyexample>
- <https://github.com/adonovan/gopl.io>
- <https://github.com/golang-standards/project-layout>

## Cheat Sheet CLI commands

### Module Management

- **Initialize a new module**
```sh
go mod init <module-name>
```

- **Add dependencies**
```sh
go get <package>
```

- **Update dependencies**
```sh
go get -u <package>
```

- **Update dependencies and show details**
```sh
go get -u -v <package>
```

- **Organize the go.mod file**
```sh
go mod tidy
```

- **Check module consistency**
```sh
go mod verify
```

- **Show module dependencies**
```sh
go list -m all
```

### Compiling and Running

- **Compile code**
```sh
go build
```

- **Compile for a specific file**
```sh
go build -o <filename>
```

- **Run code**
```sh
go run <file>.go
```

- **Install executable**
```sh
go install
```

### Tests

- **Run tests**
```sh
go test
```

- **Run tests with details**
```sh
go test -v
```

- **Run tests on all packages**
```sh
go test ./...
```

### Documentation

- **Show documentation for a package**
```sh
go doc <package>
```

- **Show documentation and examples for a package**
```sh
go doc -all <package>
```

### Miscellaneous Tools

- **Format code**
```sh
go fmt <file>.go
```

- **Format all project files**
```sh
go fmt ./...
```

- **Check code style**
```sh
go vet
```

- **Download dependencies**
```sh
go mod download
```

- **Check dependency versions**
```sh
go list -m -u all
```

### Binary Generation

- **Compile for multiple platforms**
```sh
GOOS=<operating-system> GOARCH=<architecture> go build -o <filename>
```

Example for compiling for 64-bit Linux:
```sh
GOOS=linux GOARCH=amd64 go build -o meu_programa_linux
```

### Tool Specific Commands

- **Install a specific tool**
```sh
go install <package>
```

Example for installing `golint`:
```sh
go install golang.org/x/lint/golint@latest
```

### Package Management

- **Check for outdated packages**
```sh
go list -u -m all
```

### Debugging

- **Run the debugger**

Note: `dlv` is part of the Delve package for debugging.
```sh
dlv debug
```

