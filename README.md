# cobra-fork
Fork/Detach a Cobra subcommand

An example of how to detach a process, defined as a cobra command, and keep it running separate from the parent process.

* create a child process
* record its process id
* exit the parent
* kill the child process using its recorded pid
