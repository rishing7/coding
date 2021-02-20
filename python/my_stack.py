from queue import LifoQueue

stack = LifoQueue()
stack.put(1)
stack.put(2)
stack.put(3)
stack.put(3)

print("Size of stack is %d." % stack.qsize())
print("Top element of the stack is %d." % stack.get())

stack.put(4)
stack.put(5)
print("LIFO: ", end="")
while not stack.empty():
    print(stack.get(), end=" ")

