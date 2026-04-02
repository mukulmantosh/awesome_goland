# Demo Walkthrough

### Detect Incorrect Usages of t/b.Fatal* Calls in Goroutines

![demo](./detect_tfatal.gif)

Calling \*t/b.Fatal\*\* methods from a goroutine are not advised as it may produce unexpected behavior.

You don't need to do anything to run this inspection as it's enabled by default.

Open your tests or benchmarks and see if it catches any issues.
