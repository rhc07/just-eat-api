# Just Eat Restaurant Finder

A simple application that generates local restaurants when the user inputs a postcode.

### Stack
- Go on the backend/server
- Vanilla Javascript on the Front-end (Still in progress ⌛️)
- HTML static page

### Requirements
- [Go](https://golang.org/doc/install)
- [Just Eat API](https://uk.api.just-eat.io/restaurants/bypostcode/)

### Setting up your Environment
#### Clone:
- `git clone git@github.com:rhc07/just-eat-api.git`

**OR**

- `git clone https://github.com/rhc07/just-eat-api.git`

#### Download Dependencies:
- `go mod tidy`
- `go mod vendor`

### Build and Run

#### Build:
- `go build`

#### Run:
- `./main`
- Type in a postcode after seeing this command in the termainal: `Enter your postcode:`

#### Visit:
- [Local Host](http://localhost:8080/)
- Your local host link should look something like this:
![Local host Screenshot](./images/local-host.png "Local Host Screenshot")

##### UI (Not finished⌛):
- You can also see these restaurants on the UI after re-submitting your postcode on the form. Form can't send POST requests yet, so you cannot use it dynamically:




https://user-images.githubusercontent.com/76108704/159246302-9b166710-bf5a-44bf-b774-f2e5fe6f76f0.mov



#### Tests:
- `go test`
