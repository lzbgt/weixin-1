## Summary

This is a demo using github.com/wizjin/weixin

## Start
* docker build -t img-wxapp
* docker run -it -d -p 80:80 --name wxapp img-wxapp
or

* docker run -v "$(pwd)":/myapp /myapp img-wxapp go build -v; ./myapp


