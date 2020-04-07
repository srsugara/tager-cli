# Tager CLI
> A simple CLI Application to manage your <strong>task</strong> using Golang and CouchDB.

## Installation

``` bash
# clone project
git clone git@github.com:srsugara/tager-cli.git

# go to root project
cd tager-cli

# install dependencies
go mod download

# install bin file
go install tager-cli

# check list command
tager-cli

# check list task
tager-cli list-task

# add task
tager-cli add-task -t "title task" -d "description task" -g "tags task"
or
tager-cli add-task --title "title task" --description "description task" --tags "tags task"

# update status task
tager-cli update-task -i "id task" -r "revision task" -s "status task"
or
tager-cli update-task --id "id task" --rev "revision task" --status "status task"

# delete task
tager-cli delete-task -i "id task" -r "revision task"
or
tager-cli delete-task --id "id task" --rev "revision task"
```

### GOPATH
**Notes** : Make sure your `PATH` includes the `$GOPATH/bin` directory so your commands can
be easily used:
```
export PATH=$PATH:$GOPATH/bin
```

## Installation using Docker

``` bash
# pull image from docker hub
docker pull srsugara/tager-cli

# run list command
docker run --rm -ti tager-cli

# run list task
docker run --rm -ti tager-cli list-task

# run add task
docker run --rm -ti tager-cli add-task -t "title task" -d "description task" -g "tags task"

# run update task
docker run --rm -ti tager-cli update-task -i "id task" -r "revision task" -s "status task"

# run delete task
docker run --rm -ti tager-cli delete-task -i "id task" -r "revision task"
```

## Features
- Show existing task list
- Create task
- Completing/Done a task
- Edit/Delete a task
- Could work offline / low bandwidth / flaky connection **[soon]**
 
## Additional Library
- [cobra](https://github.com/spf13/cobra) (Go CLI interactions)
- [kivic](https://github.com/go-kivik/kivik) (Common interface to CouchDB)
- [table](https://github.com/rodaine/table) (Go CLI Table Generator)
- [color](https://github.com/fatih/color) (Color package for Go)
 
## Example

``` bash
1. check list command
tager-cli
```
| <img src="https://imgur.com/T1vhlBL.png"> |
|:---:|

``` bash
2. check list task
tager-cli list-task
```
| <img src="https://imgur.com/iYaFrHy.png"> |
|:---:|

``` bash
3. add task
tager-cli add-task -t "title task" -d "description task" -g "tags task"
```
| <img src="https://imgur.com/0Fx4UJa.png"> |
|:---:|

``` bash
4. update status task
tager-cli update-task -i "id task" -r "revision task" -s "status task"
```
| <img src="https://imgur.com/8EgIrXK.png"> |
|:---:|

``` bash
5. delete task
tager-cli delete-task -i "id task" -r "revision task"
```
| <img src="https://imgur.com/gL2Buaq.png"> |
|:---:|

## License
MIT