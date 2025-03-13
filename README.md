# Chapter 4: Resource limits(cgroups)

 ## Overview
This chapter applies **CPU & memory constraints** using cgroups.
 ## Usage

 ### Using Golang

```sh
 cat  /sys/fs/cgroup/container/memory.max 
 cat  /sys/fs/cgroup/container/cpu.max
 ```

 ## Expected Output
 104857600
 50000 100000