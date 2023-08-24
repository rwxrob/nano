![](naNO.gif)

# Just Say NO to nano

[![Go Report Card](https://goreportcard.com/badge/gitlab.com/skilstak/code/go/nano)](https://goreportcard.com/report/gitlab.com/skilstak/code/go/nano) 

## Installation

If you have `go` installed:

```
go get -u gitlab.com/skilstak/code/go/nano
```

## Motivation

Nano is a horrible editor, beyond horrible! Here is a list of the reasons why:

*	encourages beginners to use Ctrl-z, Ctrl-s, and such (which backgrounds and suspends stuff everywhere else),
*	makes beginners dependent on a silly, spammy editor no professional actually uses regularly.

The fact that some distributions have made `nano` the default editor on some Linux variants is just plain stupid. I mean really stupid. I seriously doubt any professional Linux educator was involved in that decision.

The argument goes that `nano` is there for the Linux beginner who just wants to edit a config file quickly but that logic falls on its face when you consider if the user is a "casual" Linux user they already have graphical editors they can use. If they are required to use the command line then they are no longer in the "casual" user category. As soon as you open a command line you are no longer casual, period. Anyone encouraging "casual" users to mess around at the command prompt without learning exactly what they are doing is being irresponsible.

Case in point: `nano` encourages the use of the very troublesome `ctrl-z` and `ctrl-s` combination. Casual users, being "casual" and being on the command line will find themselves stuck in something eventually and unable to get out. So what will they naturally do? Yep, `ctrl-z` and now that "casual" user is ready to shut the terminal app down just to get out of it potentially leaving all kinds of zombie processes around. That's right. This is `nano`'s fault for teaching these beginners that horrible habit. I have seen this behavior in hundreds of new Linux users. This is no fluke.

If you are using the command line then use the editor built from the very beginning for the command line: vi (not even vim, unless arrow keys are disabled).
