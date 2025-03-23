# Golang

## Initializing a Module

Initialize the Module: Use the go mod init command to create a new module. Replace my-go-project with your desired module name. <br>

```
go mod init my-go-project
```

This command generates a go.mod file, which is crucial for managing your project’s dependencies. <br>

## Understanding go.mod File

The go.mod file is the heart of a Go Module. It contains essential information about your module, including: <br>

Module Path: The unique identifier for your module, typically a URL or a local path.
Go Version: Specifies the version of Go used to build the module.
Dependencies: Lists all external packages required by your module, along with their versions. <br>
Here’s a simple example of a go.mod file: <br>

```
module my-go-project
go 1.16
require (
github.com/go-sql-driver/mysql v1.5.0
)
```

In this example, the module depends on the mysql driver, specifying the exact version needed. This ensures that anyone building your project will use the same library version, maintaining consistency. <br>

By following these steps, you’ll have a fully initialized Go Module, ready to support your development efforts. For more detailed guidance, consider exploring resources like the Go Modules Reference or Francesc Campoy’s insightful video series on Go modules. <br>

## Managing Dependencies in Golang Module

Navigating the world of dependencies in a golang module can seem daunting at first, but with Go Modules, the process becomes both intuitive and efficient. This section will guide you through adding, upgrading, and removing dependencies, ensuring your projects remain robust and up-to-date.

## Adding Dependencies

When working with a golang module, adding dependencies is a straightforward task that enhances your project’s functionality by incorporating external packages.

## Using go get Command

The go get command is your primary tool for fetching new dependencies. It downloads the specified package and updates your go.mod file automatically. Here’s how you can use it:

```
go get github.com/some/package
```

This command fetches the latest version of the package, making it available for use in your project. The seamless integration of go get with Go Modules simplifies the process, allowing you to focus on development rather than dependency management. <br>

## Specifying Versions

Version control is a critical aspect of managing dependencies in a golang module. By specifying versions, you ensure that your project uses the exact library versions you intend, avoiding unexpected behavior due to changes in newer releases. To specify a version, simply append it to the package path: <br>

```
go get github.com/some/package@v1.2.3
```

This command fetches version v1.2.3 of the package, updating your go.mod file accordingly. Such precision in versioning is one of the key benefits of using Go Modules, providing stability and predictability in your projects. <br>

## Upgrading Dependencies

Keeping your dependencies up-to-date is essential for leveraging the latest features and security patches. Go Modules make this process efficient and manageable.

## Checking for Updates

Before upgrading, it’s wise to check for available updates. You can do this by running:

```
go list -u -m all
```

This command lists all modules in your project, highlighting those with newer versions available. Regularly checking for updates ensures your project remains current and secure. <br>

## Updating to a New Version

Once you’ve identified outdated dependencies, upgrading them is simple. Use the go get command with the desired version: <br>

```
go get github.com/some/package@latest
```

This command updates the package to its latest version, reflecting the change in your go.mod file. By maintaining updated dependencies, you enhance your project’s performance and security. <br>

## Removing Dependencies

Over time, some dependencies may become obsolete or unused. Efficiently managing these is crucial for keeping your golang module lean and efficient. <br>

## Identifying Unused Dependencies

To identify unused dependencies, Go provides the go mod tidy command. This tool cleans up your go.mod and go.sum files by removing unnecessary entries: <br>

```
go mod tidy
```

Running this command helps streamline your project, ensuring only essential dependencies are retained. <br>

## Cleaning Up go.mod

After identifying unused dependencies, it’s important to clean up your go.mod file. The go mod tidy command not only identifies but also removes these dependencies, keeping your project organized and efficient.

# Command

## GO RUN FILE_NAME.GO

- It compiles and executes the main package which contains the .go file specified on the terminal. This command compiles the statement and stores it in a temporary file. The syntax for the command is - `go run filename. go` <br>

- If you run this in a command line you’ll get the output from the main function’s print statement but the compilation won’t be stored anywhere. As the code complies from high-level language to low language and stores in a temporary file,

- Using go run is useful and highly recommended when you’re learning the Go or working/testing a small programme or on the testing phase of your application and don’t want to take the storage. <br>

## GO BUILD

go build scrutinizes the files in the directory to see which file (from the list) or part of the file is actually included in the main package. It is usually used to compile the packages and dependencies that are user-defined i.e. defined by you or pre-defined ones but you used in your file. When you have the binary file you made for the application or part of the bigger project to use later or it is a file that can be used in other applications or want to excess this file remotely then go build is the command for you. Using this command helps you build a different and permanent binary file in the current directory of your project/application.
