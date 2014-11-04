![githook](https://cloud.githubusercontent.com/assets/2718133/4896295/bc003450-63f7-11e4-92fb-6875b5525e0c.png)


Githook is a simple "[Continuous Integration](http://en.wikipedia.org/wiki/Continuous_integration)"-like self-hosted service that can receive [GitHub webhooks](https://developer.github.com/webhooks/) and update local repositories accordingly. Githook is currently best for static sites, like my own [austindizzy.me](https://austindizzy.me), that have their entire static source [hosted on GitHub](https://github.com/AustinDizzy/austindizzy.me).

GitHook currently supports the following:

 - [x] [HMAC - SHA1 Secrets](https://developer.github.com/webhooks/securing/)
 - [x] Monitoring multiple repositories (theoretically, an unlimited amount)
 - [x] Easy configuration
 - [x] Easy service monitoring
 - [x] Easy deployment (with and without Go)
 - [ ] More features to come.
 
Currently, githook is as simple as can be and was done in a few hours as a proof of concept and time consumer.
 
 
##Installation
 
Installing and setting up githook is simple (hopefully).
 
__With Go:__
 
```bash
$ go get github.com/austindizzy/githook
$ cd $GOPATH/src/github.com/austindizzy/githook
$ make install
```
 
__Without Go:__
 
 * Find your appropriate binary at [GoBuild](http://gobuild.io/github.com/austindizzy/githook)
 * Copy `/init.d/githook` to `/etc/init.d/githook` on your system as root.
 * As root, chmod `/etc/init.d/githook` to global execute (e.g. `sudo chmod +X /etc/init.d/githook`)
 * Edit githook.json to your needs and copy it to the binary's working directory or to `/usr/local/etc/` as githook.json
 * Move below to "Starting githook"
 
####Starting githook
 
Once githook is installed, it can be managed as an upstart service with the following commands:
```bash
sudo service githook start #starts githook
sudo service githook stop # stops githook
sudo service githook restart # restarts githook
sudo service githook status # prints running status and PID of githook
--
sudo update-rc.d githook defaults # sets githook as a default service, runs on system boot
```
 
####Configuration
 
An example configuration file can be found as [githook.json](githook.json). For githook to load its configuration details properly, the file must be named "githook.json" as must be stored in one of the two following locations:
 1. `/usr/local/etc/`
 2. The working directory of the binary (e.g. if the binary is running from `/var/githook`, the githook.json file must be stored there as well in the top directory).
  
**Note**: If you really must house the githook.json file elsewhere on the system, the path to it can be edited in [main.go#L41](main.go#L41).
  
##License
  
**git**​hook is made available under [the MIT license](LICENSE).
  
Copyright 2014 (c) Austin Dizzy
