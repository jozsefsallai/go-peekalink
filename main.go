// Package peekalink is an API wrapper for Peekalink (peekalink.io) for the Go
// programming language. It provides interfaces and methods that make it easy
// to work with link previews through Peekalink.
//
// Usage
//
// It is easy to use the package:
//
// 	package main
//
// 	import (
// 		"github.com/jozsefsallai/go-peekalink"
// 	)
//
// 	func main() {
// 		client := peekalink.NewClient("your-api-key")
// 		preview, err := peekalink.Preview("https://okuna.io")
//
// 		// ... do something with `preview`
// 	}
package peekalink
