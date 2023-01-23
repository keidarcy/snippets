package com.xyh.generics;

public class Utils {
    public static <T extends Comparable<T>> T max(T first, T second) {
        return first.compareTo(second) > 0 ? first : second;
    }

    public static <K, V> void print(K key, V value) {
        System.out.println(key + "=" + value);
    }

    public static void printUser(User user) {
        System.out.println(user);
    }

    // class CAP#1 extends User {}
    public static void printUsers(GenericList<? extends User> users) {
        // GenericList<? extends User> => read
        User u = users.get(0);

        // GenericList<? super User> => write
//        users.add(new User(2));

    }
}
