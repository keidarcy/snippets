package com.xyh.lambdas;

import java.util.List;
import java.util.function.*;

public class LambdasDemo {
    public static String prefix = "-";

//    public LambdasDemo(String message) {}

    public static void print(String message) {}

    public void print2(String message) {}

    public static void show() {
//        greet(new ConsolePrinter());

        // anonymous inner class
//        greet(new Printer() {
//            @Override
//            public void print(String message) {
//                System.out.println(message);
//            }
//        });

        // lambda expression
//        greet((message) -> {
//            System.out.println(message);
//        });
//        greet((message) -> System.out.println(message));

//        Printer printer = new ConsolePrinter();
//        Printer printer = message -> System.out.println(message);

//        Printer printer = message -> System.out.println(prefix + message);
//        greet(printer);
//
//        // Method Reference: Class/Object::method
//        greet(System.out::println);
//        greet(message -> print(message));
//        greet(LambdasDemo::print);

//        var demo = new LambdasDemo("hello");
//        greet(demo::print2);
//
//        greet(message -> new LambdasDemo(message));
//        greet(LambdasDemo::new);

        // Consumer
        List<String> list = List.of("a", "b");

//        // Imperative Programming
//        for (var item: list)
//            System.out.println(item);
//
//        // Declarative Programming
//        list.forEach((item) -> System.out.println(item));
        Consumer<String> print = System.out::println;
        Consumer<String> printUpperCase = item -> System.out.println(item.toUpperCase());
//        list.forEach(print.andThen(printUpperCase).andThen(print));

        // Supplier
        Supplier<Double> getRandom = Math::random;
        var random = getRandom.get();
        System.out.println(random);

        DoubleSupplier sup = () -> Math.random();
        System.out.println(sup.getAsDouble());

        // Function
        Function<String, Integer> map = str -> str.length();
        var length = map.apply("SKY");
        System.out.println(length);

        // "key:value"
        // first: "key=value"
        // second: "{key=value}"
        Function<String, String> replaceColon = str -> str.replace(":", "=");
        Function<String, String> addBraces = str -> "{" + str + "}";

        System.out.println(replaceColon.andThen(addBraces).apply("key:value"));
        System.out.println(addBraces.compose(replaceColon).apply("key:value"));

        // Predicate
        Predicate<String> isLongerThan5 = str -> str.length() > 5;
        var result = isLongerThan5.test("sky");

        Predicate<String> hasLeftBrace = str -> str.startsWith("{");
        Predicate<String> hasRightBrace = str -> str.endsWith("}");

        Predicate<String> hasLeftAndRightBraces = hasLeftBrace.and(hasRightBrace);
        System.out.println(hasLeftAndRightBraces.test("{s}"));
        Predicate<String> or = hasLeftBrace.or(hasRightBrace);
        Predicate<String> noLeftBrace = hasLeftBrace.negate();
        System.out.println(noLeftBrace.test("s"));


        // BinaryOperator
        // a, b -> a + b -> square
        BinaryOperator<Integer> add = (a, b) -> a + b;
        Function<Integer, Integer> square = a -> a * a;
        System.out.println(add.andThen(square).apply(1 , 2));

        // UnaryOperator
        UnaryOperator<Integer> square2 = n -> n * n;
        UnaryOperator<Integer> increment = n -> n + 1;
        System.out.println(increment.andThen(square2).apply(1));






    }

    public static void greet(Printer printer) {
        printer.print("HELLO WORLD");

    }
}
