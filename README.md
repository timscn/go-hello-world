# Hello World

Hello World is my first Project which is based on Go.

## Installation

Use the following command to run:

```shell
[vd@Fri May 05 go-hello-world]$ go run .
Hello, world :D
Don't communicate by sharing memory, share memory by communicating.
[vd@Fri May 05 go-hello-world]$ 
```
## Test

Use the following command to test it :

```shell
[vd@Fri May 05 go-hello-world]$ go clean -testcache && go test 
PASS
ok      go-hello-world/hello    0.138s
[vd@Fri May 05 go-hello-world]$ 
```

## Usage

```go

func main() {
	fmt.Println(Hello())
	fmt.Println(quote.Go())
}

```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
Inspired by: Brad Traversy

[MIT](https://choosealicense.com/licenses/mit/)