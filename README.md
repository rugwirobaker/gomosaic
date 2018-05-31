# Mosaic Make

[![CircleCI](https://circleci.com/gh/rugwirobaker/gomosaic.svg?style=svg)](https://circleci.com/gh/rugwirobaker/gomosaic)

Mosaic Make is a web application that allows users to upload pictures that are are converted into mosaic artifacts and made available through an online publicly available gallery for the world to admire.

## Motivation
What motivated me to start this project is to learn new programming concepts with the Go programming language including the following:

* Image processing
* Concurency with Go routines
* Basic algorithms and datastructures
* Cloud Object Storage
* Deploying with Google App engine

## Application architecture
I will be dividing the application into 3 micro-services namely:
* fronted:          | On Firebase
* image-processor:  | On Google App Engine

This is for the sake of efficient  scaling on top of the google cloud infrastracture.

## FACT
In this new project I will be building upon the knowledge acquired from past projects like [structure](http://www.github.com/structure) and goes even further. I hope I can share with those are far behind me in this journey but also learn from those better. So lets get to work and bring this to life. `"Happy Hacking"`