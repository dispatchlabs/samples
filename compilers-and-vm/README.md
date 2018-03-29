# GCC fronend tooling/frameowrk to compile different languages
- Supported languages https://en.wikipedia.org/wiki/GNU_Compiler_Collection#Languages
- To dump opcodes run
```shell
gcc main.c -fdump-tree-all-graph
or
gcc main.c -fdump-tree-original-raw
```

# LLVM - C
- Same as GCC but a bit more modern
- Supported languages https://en.wikipedia.org/wiki/LLVM
- To dump AST run
```shell
clang -Xclang -ast-dump main.c
```
- To dump opcodes run
```shell
clang -cc1 -triple x86_64-pc-linux-gnu -emit-llvm -mrelax-all -disable-free -disable-llvm-verifier -discard-value-names -main-file-name main.c -mrelocation-model pic -pic-level 2 -pic-is-pie -mthread-model posix -mdisable-fp-elim -fmath-errno -masm-verbose -mconstructor-aliases -munwind-tables -fuse-init-array -target-cpu x86-64 -dwarf-column-info -debugger-tuning=gdb -v -resource-dir /usr/lib/clang/6.0.0 -internal-isystem /usr/local/include -internal-isystem /usr/lib/clang/6.0.0/include -internal-externc-isystem /include -internal-externc-isystem /usr/include -fdebug-compilation-dir /home/nicu/Temp/evm -ferror-limit 19 -fmessage-length 150 -stack-protector 2 -fobjc-runtime=gcc -fdiagnostics-show-option -fcolor-diagnostics -x c main.c
```

# LLVM - Swift
```shell
swiftc main.swift -emit-ir > main.ll
```


# Swift Crash Course
https://gist.github.com/nic0lae/3fd8079719b2339fe79a