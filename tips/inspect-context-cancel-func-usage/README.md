# Demo Walkthrough

### Check context.CancelFunc Usage

![demo](./context_cancel.gif)

In code that uses a _context.WithCancel_ call, check for any warning messages. If one shows up, you can use the _Show Context Action_ menu, <kbd>⌥⏎</kbd> (macOS) / <kbd>Alt+Enter</kbd> (Windows/Linux), and use either the **Navigate to 'cancel'** function or **Navigate to the end of the problematic execution path** options from it. By using these feature, you can then gain the insight into what the IDE sees as a problematic execution path, and place the correct call to the cancel function.

<hr>

<em>The following content is directly taken from the JetBrains Guide.</em>
