# Neopixel Display
## A 2d Go drawing library for Rapsberry Pi driven Neopixels

## Install

`go get github.com/fresh4less/neopixel-display`

## Build

`go install`

To build with Raspberry Pi neopixel support using the [rpi_ws281x](https://github.com/jgarff/rpi_ws281x) library:
   1. Ensure the ws281x header files `pwm.h`, `rpihw.h`, and `ws2811.h`, are in `/usr/include`
   2. Ensure the ws281x library file `ws2811.a` is in `/usr/lib`
   3. Run `go install`

The library also provides a command-line visualization to help develop on a machine that isn't connected to LEDs.
Build with `go install -tags noled`.

## Test
To test with neopixels:
`./neopixel-display`

To test in 256 color terminal:
`./neopixel-display -m console`

You should see an animation transitioning between three different frames.

