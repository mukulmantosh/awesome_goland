Goroutines are integral to nearly all Go programs. However, employing numerous goroutines can complicate debugging efforts.

Starting from Go 1.9, you have the ability to capture extra information to understand the execution flow. This includes recording any labels you choose as part of profiling data, which can be later be utilized to analyze profiler outputs effectively.


Here are some good blog posts on this topic:

- [How to Find Goroutines During Debugging](https://blog.jetbrains.com/go/2020/03/03/how-to-find-goroutines-during-debugging/)
- [Using profiler labels](https://www.jetbrains.com/help/go/using-profiler-labels.html)

![demo](./label_goroutines.gif)