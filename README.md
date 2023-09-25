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
go install github.com/destan0098/ipshow@latest
```

## Usage

show Just IP
```
cat iplist.txt | ipshow -ip
 echo google.com| ipshow -ip
```
show Domain and  IP
```
    cat iplist.txt | ipshow -ip -pipe
or
    echo google.com| ipshow -ip -pipe
or
    ipshow -ip -d google.com
```


---

## ScreenShot

![IP Show](/ScreenShot.png?raw=true "IP Show")


## Features

Scan and show Fastly 

