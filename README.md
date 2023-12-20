# Linux-reverse-shell
A simple and stable linux reverse shell written in Golang

## To build statically linked linux execuatble :
> <br>$ export CGO_ENABLED=0 </br>
> $ env GOOS=linux GOARCH=amd64 go build


## Start the listener
> $ nc -lvnp <your_port>
