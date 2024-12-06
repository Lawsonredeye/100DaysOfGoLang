# More on Docker!

So today, i'll be diving into some concepts of docker and how to use the docker command on the command line.

```Docker
docker ps # list all running containers

docker ps -a # both running and exited containers

docker run # starts a container in attach modes where you can see logs and others

docker run --help # see all help commands which can be used with docker run.

docker attach `NAME_OR_ID_OF_CONTAINER` # Attach local standard input, output, and error streams to a running container

docker logs NAME_OF_CONTAINER # Fetch the logs of a container

docker start NAME_OF_CONTAINER # Starts container in detach mode without fetching logs 
``` 
