# Demo Walkthrough

### Support for go:embed Directive

![demo](./go_embed.gif)

Go 1.16 introduced a new feature called _go:embed_.

This works using a compiler directive, _//go:embed_, and a variable, or more, of type _string_, _[]byte_, or _embed.FS_.

If you want to embed a directory, then you must use the _embed.FS_ type for the variable in which you need to embed the content.

You don't need to do anything special to get the IDE to support this feature.

Open any Go 1.16, or newer, based project, and embed what you need.
