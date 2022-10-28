# foo

## 1. create `.env` file 

```bash
cd $PROJECT_ROOT
cp ./env-template.env ./.env 
```

## 2. modify `.env` file 

## 3. set go proxy 

```bash 
go env -w GOPROXY=https://goproxy.cn,direct
```

## 4. start server

```bash
go mod tidy 
go run main.go 
```