class Node:
    def __init__(self, data):
        self.data = data
        self.left = None
        self.right = None


def printTree(root):
    if root:
        print(root.data, end=" ")
        printTree(root.left)
        printTree(root.right)
    return None


"""
      6
    /   \
   10   12
  /  \   / \
13   14  15  16
"""

root = Node(6)
root.left = Node(10)
root.right = Node(12)
root.left.left = Node(13)
root.left.right = Node(14)
root.right.left = Node(15)
root.right.right = Node(16)
printTree(root)
