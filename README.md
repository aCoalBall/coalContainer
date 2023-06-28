# coalContainer

A simple and self-designed container tool for Linux. It creates user-level visualized containers similar to Docker.

## Namespace

In the coalContainer, you have an independent user-level run-time environment, including :

Unix Timesharing System 

Process Space 

File System (Ubuntu 20.04) 

Users 

IPC 

## Usage

Install:

git clone https://github.com/aCoalBall/coalContainer

Add to PATH:

export PATH="$coalContainer/bin:$PATH"

Run:

coalContainer run /bin/bash




