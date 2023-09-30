# Domain To IP
![Static Badge](https://img.shields.io/badge/Go-100%25-brightgreen)
## Description

This tool checks whether the given domain is UP or not, and if it is UP, check that's IP in Active mode and show to you.

- Show IP
- Show IP And Domain


## Table of Contents 


- [Installation](#installation)
- [Usage](#usage)


## Installation

```
go install github.com/destan0098/ipshow/cmd/ipshow@latest
```

## Usage

show Just IP
```
    cat iplist.txt | ipshow -ip
or
    echo google.com| ipshow -ip
or
    ipshow -ipd -d google.com
```
show Domain and  IP
```
    cat iplist.txt | ipshow -ipd -pipe
or
    echo google.com| ipshow -ipd -pipe
or
    ipshow -ipd -d google.com
```
show Domain and  IP info
```
    cat iplist.txt | ipshow -ipin -pipe
or
    echo google.com| ipshow -ipin -pipe
or
    ipshow -ipin -d google.com
```

---

## ScreenShot

![IP Show](/ScreenShot.png?raw=true "IP Show")


## Features

Scan and show Fastly 

