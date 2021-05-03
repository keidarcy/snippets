class Node {
  constructor(data, next = null) {
    this.data = data;
    this.next = next;
  }
}

class LinkedList {
  constructor() {
    this.head = null;
    this.size = 0;
  }

  // Insert first node
  insertFirst(data) {
    this.head = new Node(data, this.head);
    this.size++;
  }

  // Insert last node
  insertLast(data) {
    let node = new Node(data);
    let current;
    if (!this.head) {
      this.head = node;
    } else {
      current = this.head;
      while (current.next) {
        current = current.next;
      }
      current.next = node;
    }

    this.size++;
  }

  // Insert at index
  insertAt(data, index) {
    // If index is out of range
    if (index > 0 && index > this.size) return;

    if (index === 0) {
      this.insertFirst(data);
    }

    const node = new Node(data);
    let current, previous;
    current = this.head;

    let count = 0;
    while (count < index) {
      previous = current; // node before index
      current = current.next;
      count++;
    }
    node.next = current;
    previous.next = node;
    this.size++;
  }

  // Get at index

  getIndex(index) {
    let current = this.head;
    let count = 0;
    while (current) {
      if (count === index) {
        console.log({ index: `index ${index}: ${current.data}` });
      }
      current = current.next;
      count++;
    }
    return null;
  }

  // Remove at index
  removeAt(index) {
    if (index > this.size) return;

    let current = this.head;
    let previous;
    let count = 0;

    if (!index) {
      this.head = current;
    } else {
      while (count < index) {
        previous = current;
        current = current.next;
        count++;
      }
      previous.next = current.next;
    }
    this.size--;
  }

  // Clear list
  clearList() {
    this.head = null;
    this.size = 0;
  }

  // Print list data
  printListData() {
    let current = this.head;
    while (current) {
      console.log(current.data);
      current = current.next;
    }
    if (!this.size) console.log('nothing');
  }
}

const ll = new LinkedList();

ll.insertFirst(400);
ll.insertFirst(300);
ll.insertFirst(200);
ll.insertFirst(100);
ll.removeAt(2);
ll.printListData();
ll.clearList();
ll.printListData();
