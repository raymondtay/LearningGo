#!/bin/sh

gcc -c callClib/*.c
ar rs callC.a *.o

