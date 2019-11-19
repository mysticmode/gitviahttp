# gitviahttp
Git HTTP backend in Go. Available in both CLI and library flavors.

## Installation
Example folder structure for the above installation
```
gitviahttp
|
-- repositories
   |
   -- mysticmode
      |
      -- ohlovely.git
```

For CLI mode on Linux
```
git clone https://github.com/mysticmode/gitviahttp
cd gitviahttp
go get
go build ./cmd/gitviahttp
./gitviahttp -port=1937 -directory=./repositories
```

For CLI mode on Windows. You can use Powershell for example as shown below
```
git clone https://github.com/mysticmode/gitviahttp
cd .\gitviahttp\
go get
go build .\cmd\gitviahttp
.\gitviahttp.exe -directory=.\repositories\
```

Now clone the repository
```
git clone http://localhost:1937/mysticmode/ohlovely.git
Cloning into 'ohlovely'...
warning: You appear to have cloned an empty repository.
```
