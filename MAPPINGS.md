| docker command          | dtools equivalent                   | Comments                                                                                 |
|-------------------------|-------------------------------------|------------------------------------------------------------------------------------------|
| login                   | login, auth login                   | Login to remote registry                                                                 |
| ps -a                   | ls                                  | List containers (*enhanced output*)                                                      |
| [un]pause               | [un]pause                           | [Un]pause container(s)                                                                   |
| rename                  | rename                              | Rename a container                                                                       |
| rm                      | rm                                  | Remove container(s)                                                                      |
| log                     | log                                 | Display the container's log                                                              |
| inspect                 | inspect                             | Inspect a container (JSON format)                                                        |
| exec                    | exec                                | **NOT IMPLEMENTED YET** Execute commands in container *(Oh man, so many issues here...)* |
| stop/kill/start/restart | stop/kill/start/restart             | Stop, kill, start, restart container(s)                                                  |
| ---                     | stopall/killall/startall/restartall | same as above, but for all containers                                                    |
| images, image ls        | image ls, lsi                       | List images (*enhanced output*)                                                          |
| image rm, rmi           | image rm, rmi                       | Remove image(s)                                                                          |
| pull                    | pull                                | Pull image (*with extra feature, relies on `repo ls`*)                                   |
| push                    | push                                | **BROKEN** Push image (*with extra feature, relies on `repo ls`*)                        |
| tag                     | tag                                 | Tag a docker image                                                                       |
| network ls              | network ls                          | List networks                                                                            |
| network rm              | network rm                          | Remove network(s)                                                                        |
| network [dis]connect    | network [dis]connect                | [Dis]connect a network from a container                                                  |
| network create          | network add                         | **NOT IMPLEMENTED YET** *multiple issues*                                                |
| info                    | system info, info                   | **NOT FULLY IMPLEMENTED, WILL NOT BE** show daemon info                                  |
| ---                     | repo add                            | Add a default docker registry config (for the -d flag)                                   |
| ---                     | repo rm                             | Remove the current default docker registry config                                        |
| ---                     | repo ls                             | Show the current default docker registry config                                          |
| ---                     | get catalog                         | List all images from a remote registry                                                   |
| ---                     | get tags                            | List all tags of from a remote registry docker image                                     |
