## Develop

#### ⚒️ Setup environment variables
```bach
cp .env.sample .env
```

#### 🏃 Run Locally without "Hot Reloading"
```bash
go run main.go
```

#### ♻️  Enable Hot Reloading
Install Air
```bash
go install github.com/cosmtrek/air@latest
```
Run air to start Echo with hot reloading
```bash
air 
```
More info at
[Air Documentation](https://github.com/cosmtrek/air)
