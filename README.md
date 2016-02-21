# Lietu's Go Game Engine

This engine gives you stuff. You can totally wotally write games on Go now. Yay!

It's got a license. Read it. It's in `LICENSE.md`.

## Getting SDL stuff compiled

First time running anything the SDL libraries will have to be compiled, and they need to locate the SDL include and lib directories.

### On *nix

If using \*nix you'll probably want to install SDL development packages using whatever package manager you have, and then you'll need to set up the `CGO_*` environment variables to point to the SDL directories.
 
```
CGO_CFLAGS="-I/path/to/sdl/include"
CGO_LDFLAGS="-L/path/to/sdl/lib"
```

You can figure out what these are with e.g.:

```
find / -wholename "*/SDL2/SDL.h"
```

### On Windows

You'll likely want to get [msys2](https://msys2.github.io/), install GCC and the SDL libraries on that.

Start up the MSYS2 Shell, and run:

```
pacman -S mingw-w64-x86_64-gcc
pacman -S mingw-w64-x86_64-SDL2 mingw-w64-x86_64-SDL2_image mingw-w64-x86_64-SDL2_mixer mingw-w64-x86_64-SDL2_ttf
```

Then set the environment variables for Go to find the SDL stuff.

```
set CGO_CFLAGS=-Ic:/msys64/mingw64/include
set CGO_LDFLAGS=-Lc:/msys64/mingw64/lib
```



pacman -S mingw-w64-x86_64-openal