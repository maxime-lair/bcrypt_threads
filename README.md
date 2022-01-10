# bcrypt_threads
Test of subroutines implementation in go to calculate list of bcrypt hash

Requires `go 1.17`

It highlights the difference between an iterative program and a subroutines-driven one.

Without subroutines, It takes 50s to complete, with threads 100ms

## Usage

Download the repository then

`make all`

## Sample

With process usage on the side (accelerated)

![threads_go](https://user-images.githubusercontent.com/72258375/148839199-2484d111-74e1-43d7-a0b4-d9c6073f748d.gif)



