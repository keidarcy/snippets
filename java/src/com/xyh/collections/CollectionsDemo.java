package com.xyh.collections;

import java.util.ArrayList;
import java.util.Collection;
import java.util.Collections;

public class CollectionsDemo {
    public static void show() {
        Collection<String> collection = new ArrayList<>();
        Collections.addAll(collection, "a", "b");
//
//        Object[] o = collection.toArray();
//        String[] s = collection.toArray(new String[0]);
//        System.out.println(s[0].);

        Collection<String> other = new ArrayList<String>();
        other.addAll(collection);

        System.out.println(other == collection);
        System.out.println(other.equals(collection));



    }
}
