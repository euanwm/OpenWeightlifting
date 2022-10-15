<h1>Backend Local Testing</h1>
Ok, so basically CORS is a massive issue. When starting the serverMain.go file, simypl add "local" as an additional argument and it <i>should</i> disable the CORS middleware that is in place.

```
go build serverMain.go
./serverMain.go local
```

I'm not writing much more, someone else can write documentation. I can't be arsed.