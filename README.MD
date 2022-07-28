# Installation (Linux)
For download go to this [site](https://go.dev/doc/install) and download the tar.gz file

Go to your download folder and extract the zip file in `usr/local` folder

You can also use this command from your download directory

``` sh
sudo tar -C /usr/local -xvzf go1.15.2.linux-amd64.tar.gz
```

Now open the bash profile from your terminal and run the command. Be ensure you are on home/${USER} directory
```sh
nano .bashrc
```
and add the follwings at the end of the file and save it

```sh
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH

```

# RUN

To build and run use this command

```sh
 go run hello-world.go
```
You can also use build and run separately by using this two command

```sh
go build hello-world.go
./hello-world
```