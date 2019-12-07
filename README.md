# morse_thinklight

Morse code powered thinklight. Use your thinkpad to spread the word of Morse!

This is some simple script I finandled around during one beautiful morning.

Super simple, it takes your command line input and turns your led into morse encoder!

To run (you can also build it), just do 

``` sh
go run main.go a b c
```

which will give you `.- / -... / -.-.` and start blinking! Make sure you `chmod 666` you
led path. It's usually `/sys/class/leds/tpacpi::thinklight/brightness`

**BONUS**

You can loop your message by passing the `-l` flag, so running

``` sh
go run main.go -l hello, world
```

will blink "hello, world" forever
