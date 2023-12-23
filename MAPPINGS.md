| section / verb     | docker command          | dtools equivalent                   | Comments                                                           |
|--------------------|-------------------------|-------------------------------------|--------------------------------------------------------------------|
| **authentication** | login                   | login, auth login                   | Login to remote registry                                           |
| **containers**     | ps -a                   | ls                                  | List containers (*enhanced output*)                                |
|                    | pause/unpause           | pause/unpause                       | Pause/unpause container(s)                                         |
|                    | rename                  | rename                              | Rename a container                                                 |
|                    | rm                      | rm                                  | Remove container(s)                                                |
|                    | log                     | log                                 | Display the container's log                                        |
|                    | inspect                 | inspect                             | Inspect a container (JSON format)                                  |
|                    | exec                    | exec                                | Execute commands in container                                      |
|                    | stop/kill/start/restart | stop/start/restart                  | Stop, kill, start, restart container(s)                            |
|                    | n/a                     | stopall/killall/startall/restartall | same as above, but for all containers                              |
|                    | run                     | run                                 | **NOT YET IMPLEMENTED**                                            |
| **images**         | images, image ls        | lsi                                 | List images (*enhanced output*)                                    |
|                    | image rm, rmi           | rmi                                 | Remove docker image(s)                                             |
|                    | pull                    | pull                                | Pull image (*with extra feature, relies on `repo ls`*)             |
|                    | push                    | push                                | **BROKEN** Push image (*with extra features, relies on `repo ls`*) |
|                    | tag                     | tag                                 | Tag a docker image                                                 |
| **network**        | ls                      | ls                                  | List networks                                                      |
|                    | rm                      | rm                                  | Remove network(s)                                                  |
|                    | connect, disconnect     | connect, disconnect                 | Connect or disconnect a network from a container                   | 
|                    | create                  | add                                 | **NOT IMPLEMENTED YET** *multiple issues*                          |
| **info**           | info                    | system info                         | **NOT FULLY IMPLEMENTED, WILL NOT BE** show daemon info            |
| **NEW FEATURES**   | n/a                     | repo add                            | Add a default docker registry config (for the -d flag)             |
|                    | n/a                     | repo rm                             | Remove the current default docker registry config                  |
|                    | n/a                     | repo ls                             | Show the current default docker registry config                    |
|                    | n/a                     | get catalog                         | List all images from a remote registry                             |
|                    | n/a                     | get tags                            | List all tags of from a remote registry docker image               |
| **volumes**        |                         | volume                              | **NOT YET IMPLEMENTED**                                            | 

<br>Other commands will be forthcoming, especially equivalents to `docker system` (all the prune commands, for instance)