#include <stdio.h>
#include <sys/ioctl.h>
#include "msg.h"

int main(void) {
  struct winsize sz;
  ioctl( 0, TIOCGWINSZ, &sz );
  if (sz.ws_col >= 88 && sz.ws_col >= 22)
    puts(nano1);
  else
    puts(nano1s);
}