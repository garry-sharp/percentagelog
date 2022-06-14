# Percentage Log

A very small and relatively stupid and pointless application to print updated status to the screen, it's only really tested on UNIX based terminals 

## Overview

So there's a interface of type `Printable` which exports 2 methods.
- `String() string` a standard `toString()` style method
- `Percentage() float32` which should return a float between 0 and 100.

You can then pass a spread of any structs which implements this interface and call `PrintUnitilFinished`. The application will output continuously until all percentages hit 100% Below is sample output from a demo dummy web downloader app

```
hello_world.txt                                   | ==================================================> 100.00 %
my_cat.jpg                                        | ===================================>                 71.22 %
amazing_go_code.go - v.1.2.2                      | ==================================================> 100.00 %
MyPasswords.kdbx                                  | ==================================================> 100.00 %
Docker.dmg                                        | >                                                     0.55 %
Application.zip                                   | ===================>                                 39.76 %
mysql-workbench-community-8.0.29-macos-x86_64.dmg | =>                                                    2.23 %
offsetexplorer.dmg                                | =>                                                    3.88 %
```

You can see a demo application in the demo folder. You can run this from the folder by using `go run demo.go`

## Usecases

Mainly where logging requires something to continually say what percentage of completion has been done. Examples could include

- Long running go routines
- Downloading multiple files from the internet

## Import

you can import it to your application using `go get -u "github.com/garry-sharp/percentagelog"`