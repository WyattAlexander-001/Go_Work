# Guide:
* Packages are simply, dir(s) in a project with source files
    * They begin with package dirName
* Module, usualy JUST one, has the dependencies with versions
    * Go sum holds info of the ^^Dependencies^^
    * go mod init mycompany.com/projectName
    * Use terminal to get depencies
    ```
    go get github.com/gorrila/mux
    ```
