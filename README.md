# substrate

> the base on which an organism lives
> “substrate,” Merriam-Webster.com Dictionary, https://www.merriam-webster.com/dictionary/substrate. Accessed 10/18/2023.

A collection of odds and ends, sometimes with a cutesy name.

* buffers: Support for rotating buffers
* dontio: Models std{in,out,err} as a struct, also ansi escape codes?
* files: augments fs.FS to include WRITING TO A FILE WHICH THE STDLIB DOESNT SUPPORT OMG WHY
* mirrors: Reflection stuff
* peruse: lexer and Pratt parser implementations
* rng: Additional rand/v2 source implementations
* skelly: Some data structures
    * bitset{32,64}: uint{32/64} based bitset
    * graph{32,64}: map[uint{32,64}]bitset{32,64} graph
    * hashset: what it says on the tin
    * queue: slice based queue
    * shufflequeue: growable queue that dequeues elements randomly
    * stack: slice based stacks, fixed size and growable implementations
* slipup: error and panic helpers
* stageleft: primarily concerned with modeling exit codes (as uint8, _NOT_ uint32)
