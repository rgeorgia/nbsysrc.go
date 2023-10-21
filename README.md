# netsysrc
This started out to be a simple sysrc like application for NetBSD.

Written with go

I am tired of opening rc.conf with vi{m} in order to add a service, or a flag. I don’t want to have to “cat” the rc.conf to see the rc status of a or flag.

I need a single application that does it for me, something similar to sysrc used by FreeBSD and it’s derivatives. 
Basically netsysrc will take an argument from the command line and insert or append it to the rc.conf file. 

  #### Example:
  ```netsysrc dbus=YES```

## Work Flow

1. Read rc.conf file
2. Read cli arguments
3. Check rc.conf if the service is already there.
4. If the service is present in the rc.conf file and the value is the same as the cli argument, do nothing.
5. If the service is present in the rc.conf and the value is **not** the same as the cli argument, change the value.
6. If the service is **not** present, append service=value to the rc.conf file
   1. Check etc/rc.d to see if the service is present, if not, alert the user the service is not available in etc/rc.d.
   2. Check example/rc.d to see if the service is present, if it is alert the user that the service is in the example/rc.d directory but not in /etc/rc.d.
   3. If the service is not in /etc/rc.d or example/rc.d, alert the user that the service has not been installed.
