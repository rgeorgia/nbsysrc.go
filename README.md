# netsysrc
This started out to be a simple sysrc for NetBSD; however there may be some room to allow for FreeBSD, HardenedBSD and DragonflyBSD.

Written with go

I am tired of opening rc.conf with vi in order to add a service, or a flag. I don’t want to have to “cat” the rc.conf to see the rc status of a or flag.

I need a single application that does it for me, something similar to sysrc used by FreeBSD and it’s derivatives. Basically netsysrc will take an argument from the command line and insert or append it to the rc.conf file. 

  #### Example:
  ```netsysrc dbus=YES```
