# ascii-art-reverse

Ascii-art is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII

### Features

- The project is written in Go.
- The code respects the good practices.

Ascii-reverse consists on reversing the process, converting the graphic representation into a text. You will have to create a text file containing a graphic representation of a random string given as an argument.

The argument will be a flag, --reverse=<fileName>, in which --reverse is the flag and <fileName> is the file name. The program must then print this string in normal text.

- The flag must have exactly the same format as above, any other formats must return the following usage message:

```
Usage: go run . [OPTION]

EX: go run . --reverse=<fileName>
```

If there are other ascii-art optional projects implemented, the program should accept other correctly formatted [OPTION] and/or [BANNER].
Additionally, the program must still be able to run with a single [STRING] argument.



### Examples

```bash
$ cat -e file.txt
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
$
$ go run . --reverse=file.txt
hello$
$


```
## Authors

- [@mkassymk](https://01.alem.school/git/mkassymk)
- [@tlsh0](https://www.github.com/tlsh0)

