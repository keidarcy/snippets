package com.xyh.generics;

import java.util.Iterator;
import java.util.function.Consumer;

//public class GenericList<T extends Comparable & Cloneable> {
public class GenericList<T> implements Iterable<T> {
    private T[] items = (T[])new Object[10];
    private int count;

    public void add(T item)  {
        items[count++] = item;
    }

    public T get(int i) {
        return items[i];
    }

    @Override
    public Iterator<T> iterator() {
        return new ListIterator(this);
    }

    private class ListIterator implements Iterator<T> {
        private GenericList<T> list;
        private int index;
        public ListIterator(GenericList<T> list) {
            this.list = list;
        }

        @Override
        public boolean hasNext() {
            return index < list.count;
        }

        @Override
        public T next() {
            return list.items[index++];
        }
    }
}
