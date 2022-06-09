# SOBED

> I want embed, SOBED  
> Really, really BED  

![who's bad?](https://i.giphy.com/media/11BbSJZpwZAaXu/200.webp)

SOBED is a proof of concept of embedding a shared object (A.K.A. .so, A.K.A. library) in a Go binary.

From there, it loads it via `dlopen()` and get the right symbol via `dlsym`.
Then, with some extra wrapping, you can call it natively from Go (but it's still disgusting CGo behind the scene).

## Why?

It could be a way to really distribute your binary completely statically, embedding its external dependencies
in the right version for the right architecture. Of course, if the library uses functionality absent for the current host
you will still face suffering and failure.

For graphical libraries, there is no way to not pass by CGo, but let's avoid some pain and just `go get` this project and run it,
without previous dependencies install step.

## Security?

None, many steps in the process can fail, and I didn't find a way to have some standard API or wrapper.

Also, YOU are now responsible for providing the patched lib version if you don't want some security issues.

I guess, it makes more sense for easy to try lib/app, but then, really use the provided libs from your system and install + update
them accordingly.

## Try

```
go run github.com/dolanor/sobed@latest
```

If you want to play with the libgreet C lib, you can just modify it in a clone and rebuild it with a
```
go generate ./greet
```

