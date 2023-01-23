package com.xyh.collections;

import java.util.ArrayDeque;
import java.util.Queue;

public class QueueDemo {
    public static void show() {
        Queue<String> queue = new ArrayDeque<>();
        queue.add("a");
        queue.add("b");
        queue.add("c");
        queue.offer("d");
        var e1 = queue.peek();
        var e2 = queue.remove();
        var e3 = queue.poll();
        queue.element();

    }
}
