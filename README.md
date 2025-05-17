# GoPlayground
Place to test &amp; learn Go lang

## Notes
Think of a *T as its own type
This means that if something is expecting an interface, the type implementing the methods must be equal to that of the receiver - Important for Stringer interface [[Ref]](https://www.reddit.com/r/golang/comments/10yrpic/why_not_mix_value_and_pointer_receivers/)  
* This is the reason why value and pointer receivers shouldn't be mixed

Malloc Consideration  
Go is garbage collected. Don't consider malloc, copy, & move semantics. If you really care look at go code on github

