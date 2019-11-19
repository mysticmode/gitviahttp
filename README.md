# gitviahttp
Git HTTP backend in Go. Available in both CLI and library flavors.

## Installation & Usage
Example folder structure for the below installation
```
gitviahttp
|
-- repositories
   |
   -- mysticmode
      |
      -- oh-lovely.git
```
### CLI mode
On Linux
```
git clone https://github.com/mysticmode/gitviahttp
cd gitviahttp
go build ./cmd/gitviahttp
./gitviahttp -port=1937 -directory=.
```

On Windows
```
git clone https://github.com/mysticmode/gitviahttp
cd .\gitviahttp\
go build .\cmd\gitviahttp
.\gitviahttp.exe -port=1937 -directory=.
```

Now clone the repository
```
git clone http://localhost:1937/mysticmode/oh-lovely.git
Cloning into 'oh-lovely'...
warning: You appear to have cloned an empty repository.
```

If you have **repositories directory somewhere else**, specify the absolute path of that directory as shown below

For example on Linux, you have the repository at `/home/git/repositories`
```
./gitviahttp -port=1937 -directory=/home/git/repositories
```

For example on Windows, you have the repository at `D:\Git\repositories`
```
.\gitviahttp.exe -port=1937 -directory=D:\Git\repositories
```

And then clone the repository.

### Library mode
I'm using [Gorilla Mux](https://www.gorillatoolkit.org/pkg/mux) router below to show an example of how gitviahttp will work as a library.
```
package main

import (
    "github.com/gorilla/mux"
    "gopkg.in/mysticmode/gitviahttp.v1"
)

func main() {
    m := mux.NewRouter()
   
    repoDir := "/home/git/repositories"
   
    m.PathPrefix("/+{username}/{reponame[\\d\\w-_\\.]+\\.git$}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
       // Do your authentication here if you want
       // and then call gitviahttp.Context()
       gitviahttp.Context(w, r, repoDir)
    }).Methods("GET", "POST")
}
```
