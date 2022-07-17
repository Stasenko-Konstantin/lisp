# lisp
desing is very human. very easy to use  
simple lisp in ± 800 lines of code (including some tests)
```bash
git clone https://github.com/Stasenko-Konstantin/lisp
cd lisp
go build lisp.go

usage: ./lisp [*.scm] [arg] where arg:
        --help   -- prints the help
        --past   -- prints intermediate 
        
```
```haskell
./lisp
<<< :h
>>> :l, :load *.scm      -- evaluate file
>>> :h, :help    -- prints the help
>>> :p, :past    -- prints intermediate ast
>>> :q, :quit    -- program exit
<<< :p
>>> OK
<<< (println 1)

type = LIST_O, content = 
        type = NAME_O, content = println, x = 1, y = 0;
        type = NUM_O, content = 1, x = 10, y = 0;, x = 0, y = 0;

1 
>>> ( 1 )
<<< 

```
