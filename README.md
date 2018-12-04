# IASIP Title Card Generator

![](https://i.imgur.com/BcPsMe1.png)

This is a simple [Go](https://golang.org) program that generates title cards similar to those from
[It's Always Sunny in Philadelphia](https://en.wikipedia.org/wiki/It%27s_Always_Sunny_in_Philadelphia).
I wrote this in a few hours to learn Go a little better (and to make IASIP title cards super easy to
generate on-the-fly!).

## Usage

Once compiled to an executable file, place a copy of the Textile font (in TrueType format) into the
same folder as the executable, named "textile.ttf".
The font has not been included with this repository for licensing reasons.

Once you're done with that, open a terminal window. go to the directory of the files, and type:

```
./iasip-generator "\"The Gang Creates an IASIP Title Card\""
```

The output will be saved to `out.png` in the same folder.
