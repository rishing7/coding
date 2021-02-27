class Node:
    def __init__(self, data):
        self.data = data
        self.next = None


head = Node(1)
head.next = Node(2)
head.next.next = Node(3)
head.next.next.next = Node(4)
head.next.next.next.next = Node(5)


def printLL(head):
    temp = head
    while temp:
        print(temp.data, end=" ")
        temp = temp.next


def sort(head, K):
    i = 0
    tail_head = None
    while i <= K:

        tail = tail_head
        temp = head

        while temp.next:
            if i == temp.data:
                if tail_head:
                    tail.next = temp
                else:
                    tail_head = temp
                    tail = tail_head
            temp = temp.next

        temp.next = tail_head
        i = i + 1
    return tail_head


head = sort(head, 6)
printLL(head)
