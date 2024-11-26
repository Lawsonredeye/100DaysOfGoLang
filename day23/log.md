# Packages and Mocking

Packages are just like your regular folders which stores files or source codes containing variables, functions, structs and so on.

In `Go` you can have your very own package/folder where you would store function which can be reuseable at any given time.

## Creating a Package
To create a package you just need to do the following:
1. create a folder for your project e.g `my-package-project`
2. Initialize the folder using `go mod init MODULE_NAME_FOR_YOUR_PROJECT`
3. Create another folder/package to store your source code e.g `math` 
4. Create a source code with your code logic inside
5. Save your code and run `go install` on your terminal to install the package on the project and go environment.
6. Create a `main.go` file to test your package.
7. Within your main.go file import your file using the name of the module with the folder name e.g `MODULE_NAME_FOR_YOUR_PROJECT/math`.
8. Call your method in your main() using the keyword `math._name_of_func_in_package`.

