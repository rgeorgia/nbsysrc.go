# netbsd-sysrc
Very simple sysrc for NetBSD in rust.

### Why in Rust?

Because I am learning Rust and this is a self imposed, useful project

# nbsysrc
This started out to be a simple sysrc like application for NetBSD.

Written with go

I am tired of opening rc.conf with vi{m} in order to add a service, or a flag. I don’t want to have to “cat” the rc.conf to see the rc status of a or flag.

I need a single application that does it for me, something similar to sysrc used by FreeBSD and it’s derivatives.
Basically netsysrc will take an argument from the command line and insert or append it to the rc.conf file.

#### Example:
```netsysrc dbus=YES```

## Work Flow

1. Read cli argument, only one. service=value
2. Check for $HOME/.nbsysrc.toml file
3. If $HOME/.nbsysrc.toml is present, read the file for path names.
4. If $HOME/.nbsysrc.toml is **not** present, or the keys have no value, then use default values.

$HOME/.nbsysrc.toml

| key         | value | default value               |
|-------------|-------|-----------------------------|
| ect_rcd     | value | /etc/rc.d                   |
| example_rcd | value | /usr/pkg/local/example/rc.d |
| rc_conf     | value | rc.conf                     |


Some Details
------------

Check out our [wiki](https://github.com/rgeorgia/nbsysrc/wiki) for more details.

Check out our [Roadmap wiki](https://github.com/rgeorgia/nbsysrc/wiki/Roadmap) for more details.
