# gitviahttp
Git HTTP backend in Go. Available in both CLI and library flavors.

## Installation
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
.\gitviahttp.exe -port=1937 -directory=.\repositories\
```

Now clone the repository
```
git clone http://localhost:1937/mysticmode/oh-lovely.git
Cloning into 'oh-lovely'...
warning: You appear to have cloned an empty repository.
```

Now if you have repositories directory somewhere else, specify the absolute path of that directory as shown below

For example on Linux, you have the repository at `/home/git/repositories`
```
./gitviahttp -port=1937 -directory=/home/git/repositories
```

For example on Windows, you have the repository at `D:\Git\repositories`
```
.\gitviahttp.exe -port=1937 -directory=D:\Git\repositories
```

And then clone the repository.
