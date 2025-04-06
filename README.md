# a06-go-basics

Yes, I found your code showinng race conditions. Because the results printed from 'Prog Two' process is a negative number. Because the go power_h processes are asynchronous. They also have a shared global state 'res'. One way to fix it is to use synchronous recursion. It would work because it creates a process once the previous one has successfully modify the 'res' value.