from queue import Queue

queue = Queue()

queue.put(1)
queue.put(2)
queue.put(3)

print("Size of stack is %d." % queue.qsize())
print("Leftmost element of the stack is %d." % queue.get())

queue.put(4)
queue.put(5)

print("FIFO:", end="")
while not queue.empty():
    print(queue.get(), end=" ")



