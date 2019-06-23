#!/bin/sh
nginx
goss autoadd package curl busybox
goss autoadd port 80
goss validate
