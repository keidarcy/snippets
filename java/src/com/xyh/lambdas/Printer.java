package com.xyh.lambdas;

public interface Printer {
    void print(String message);

    // it's possible to have implementation detail in interface
    // but should not use it
//    default void printTwice(String message) {
//        System.out.println(message);
//    }
}
