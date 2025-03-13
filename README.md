# Chapter 3: FileSystem Isolation
 
 ## Overview
This chapter introduces **filesystem isolation** using `chroot`.  

 ## Usage

 ### Using Golang
 ```sh
 go run main.go run /bin/sh
 ```

 ### Using Docker
 ```sh
 docker run <image_name> /bin/sh
 ```

 ### Using OContainer
 ```sh
 OContainer run /bin/sh
 ```