package com.xyh.collections;

import java.util.*;

public class MapDemo {
    public static void show() {
        var c1 = new Customer("a", "e1");
        var c2 = new Customer("b", "e2");

        Map<String, Customer> map = new HashMap<>();
        map.put(c1.getEmail(), c1);
        map.put(c2.getEmail(), c2);

        var customer = map.getOrDefault("e10", new Customer("b", "a"));
        var exists = map.containsKey("e10");
        map.replace("e1", new Customer("a++", "e1"));

        System.out.println(map);

//        for (var key: map.keySet()) {
//        for (var customer: map.values()) {
        for (var entry : map.entrySet()) {
            System.out.println(entry.getValue());
        }

    }
}
