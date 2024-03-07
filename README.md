# simple-fibonacci-server

Write an HTTP service in Go, with one endpoint;

1. The endpoint should take a Fibonacci number as input and output the next Fibonacci number.
2. The user should be able to choose between JSON or MessagePack for both the input and the output.
3. For the sake of this exercise, the input and output both should look like this:

{ "n": THE_FIBONACCI_NUMBER }
