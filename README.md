# dotenv

Another **minimal** dotenv library, aim for **fast, small and instant break** if error occured!

### Available functions

- `dotenv.Load(filenames ...string) error`
- `dotenv.ForceLoad(filenames ...string) error` - Overwrite all exist variables!

### Future plan

- Support multi-line value
- Improve reader

## Examples and notes

You should use `Load` or `ForceLoad` as early as your code (ideally at top of main)

If your project structure like this:
```
│ .env
│ .env.test
│ go.mod
│ go.sum
│ README.md
│   
├───cmd
│   ├───api
│   │   ├───main.go
...
```

When you call function at main.go, it should be `dotenv.Load("../../.env", "../../.env.test")`. This should all variables, duplicate skipped. If you use `ForceLoad`, it will load `.env` then load `.env.test` which all duplicate variable will priority from `.env.test`