package com.xyh.collections;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

public class ListDemo {
    public static void show() {
//        List<String> list = new ArrayList<>();
//        list.add("a");
//        list.add("b");
//        list.add(0, "8");
//        list.set(0, "a+");
//        list.remove("a");
//        Collections.addAll(list, "b", "c");
//        System.out.println(list);
//        System.out.println(list.get(0));
//        System.out.println(list.indexOf("b"));
//        System.out.println(list.lastIndexOf("b"));
//        System.out.println(list.subList(0,4));

        List<Customer> customers = new ArrayList<>();
        customers.add(new Customer("a", "e4"));
        customers.add(new Customer("c", "e2"));
        customers.add(new Customer("b", "e3"));

        Collections.sort(customers, new EmailComparator());
        System.out.println(customers);

    }
}
