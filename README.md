# synx

A small Go-package that extends the standard sync package. WaitGroup with timeout.
`synx` is a package with an extra `WaitGroup.WaitWithTimeout()` function.

## Usage

```go

wg := &synx.WaitGroup{}
wg.Add(1)
wg.WaitWithTimeout(10 * time.Second)
```



## Example

```go

// WaitWithTimeout returns the value "true" when  the [WaitGroup] counter is zero.
// And returns the value "false" when the wait is completed by timeout.
//
func shutdownServices() {
	wg := &synx.WaitGroup{}
	wg.Add(1)
	shutdownService1(wg)

	wg.Add(1)
	shutdownService2(wg)

	if ok := wg.WaitWithTimeout(10 * time.Second); !ok {
		log.Println("Force shutdown by timeout")
	}
	log.Println("Shutdown finished")
	os.Exit(0)
}

```
