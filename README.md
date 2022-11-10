# The Go Programming Language

This repository acts as a reference for my learning from the book "The Go Programming Language".

I have read the book through and am now attempting to do a fast run through the book building out some of the
examples so that I may understand better some of the concepts.

I have been working on various go codebases - mainly related to hyperledger fabric. This repo has been started 
to assist in improving my overall knowledge in relation to the Go programming language.

### Lissajous
This example outputs a gif file which can be written/redirected from Stdout to a file:

<img src="https://github.com/robevansuk/theGoProgrammingLanguage/blob/main/src/main/chapter1/Lissajous.gif">

The following command:
```
go build Lissajous_v1.go 
```
will create `Lissajous_v1` as a binary that can be executed using slash-dot notation

```
./Lissajous_vs > Lissajous.gif

ls -al
...
-rwxr-xr-x   1 user  owner  2131506 10 Nov 22:15 Lissajous_v1
...
```

Lissajous_v2.go changes the palette for black and green for a "more authentic" LissajousV2.gif

<img src="https://github.com/robevansuk/theGoProgrammingLanguage/blob/main/src/main/chapter1/LissajousV2.gif">