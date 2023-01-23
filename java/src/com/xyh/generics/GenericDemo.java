package com.xyh.generics;

public class GenericDemo {
    public static void show() {
//        var u1 = new User(1);
//        var u2 = new User(2);
//        if (u1.compareTo(u2) > 0) {
//            System.out.println("u1 > u2");
//        } else if (u1.compareTo(u2) == 0) {
//            System.out.println("u1 == u2");
//        } else {
//            System.out.println("u1 < u2");
//        }

//        var u = Utils.max(new User(1), new User(2));
//        System.out.println(u);

//        Utils.print("1", "2");


        User u = new Instructor(2);
        Utils.printUser(u);
        var users = new GenericList<User>();
//        Utils.printUsers(new GenericList<Instructor>());


        var list = new GenericList<Integer>();
        list.add(8);
        list.add(99);
        for (var item: list) {
            System.out.println(item);
        }
//        var iterator = list.iterator();
//
//        while(iterator.hasNext()) {
//            var current = iterator.next();
//            System.out.println(current);
//        }

//        for (var item: list) {
//            System.out.println(item);
//        }
    }
}
