# mtBST

This was just a little challenge project for a job interview (a binary search tree that reports on max depth, including one or more values at that depth.  The challenge did not specify language so I coded it first in two languages: (1) Node.js and (2) Go.  Furthermore, I write a shell script to run both sample tests on each version:

```bash
run.sh
```

## Node.js

This version was very quick to write, becuase I code frequently in Node.js.  It is also short because Node.js is a higher level and very expressive language.  Mostly this came down to that versatile data structures don't require special definitions.  You can mix types and even dynamically modify the data structures and your code can look at it, reflectively.  This is a very useful language for "always on" network software because it JIT compiles and caches, making it easy to build a system whereby you add/modify code at run-time for upgrades/fixes.  It also has great support for sockets, encryption, various communication protocols, and access to the underlying system.

## Go

This version took me some time because this is literally the first program I have ever written in Go.  It feels a lot like an improved version of C, honestly.  I structured this implementation as closely as I could to the Node.js implementation.  However, it's considerably more code because it's a slightly lower level language and, as such, didn't have the various shortcuts available as did Node.js.  Compiling is noticeably slower but I imagine execution speeds are probably better.  Go feels to me like a language that could have been designed for building an operating system (like C was and Node.js is not).  

Another advantage of Go is that it has far broader architectural support than Node.js.  Node is built on the LLVM compiler backend, which is faster than GCC, but only supports x86, x86-64, ARM, and WebAssembly (plus somehow Z System but not sure how).  Go has versions available for GCC (supports just about anything that computes except for Web Assembly), for LLVM, and also it's own.  

There also appears to be good support for threads, sockets, and at least some limited support for encryption (need to research furhter).  It's a pretty cool language.
