package com.xyh.streams;

import java.util.*;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

import static java.util.stream.Collectors.*;

public class StreamsDemo {
    public static void show () {
        List<Movie> movies = List.of(
                new Movie("a", 10, Genre.COMEDY),
                new Movie("b", 15, Genre.COMEDY),
                new Movie("c", 30, Genre.ACTION),
                new Movie("d", 30, Genre.THRILLER)
        );
//        System.out.println(movies);

//        // Imperative Programming
//        int count = 0;
//        for (var movie :
//                movies) {
//            if (movie.getLikes() > 10)
//                count++;
//        }

        // Declarative (Functional) Programming
//        var count2 = movies.stream()
//                .filter(movie -> movie.getLikes() > 10)
//                .count();
//
//        System.out.println(count2);
//
//        Collection<Integer> x;
//        var list = new ArrayList<Integer>();
////        list.stream();
//
//        int[] numbers = {1, 2, 3};
//        Arrays.stream(numbers).forEach(System.out::println);
//        var s1 = Stream.of(1, 2, 4);
//        var s2 = Stream.generate(Math::random);
//        s2.limit(3).forEach(System.out::println);
//
//        Stream.iterate(1, n -> n + 1)
//                .limit(10)
//                .forEach(System.out::println);
//
//        movies.stream()
//                .mapToInt(Movie::getLikes)
////                .map(movie -> movie.getTitle())
//                .forEach(name -> System.out.println(name));

//        var stream = Stream.of(List.of(1, 2, 3), List.of(4, 5, 6));
//        stream
//             .flatMap(Collection::stream)
//                .forEach(System.out::println);

//        Predicate<Movie> isPopular = m -> m.getLikes() > 10;
//        movies.stream()
//                .filter(isPopular)
//                .forEach(System.out::println);

//        // Slicing
//        movies.stream()
//                .skip(2)
//                .limit(3)
//                .forEach(m -> System.out.println(m.getTitle()));
//
//        movies.stream()
//                // .dropWhile() loop until return true
//                .takeWhile(m -> m.getLikes() < 15) // filter => loop, takeWhile => loop until return false
//                .forEach(System.out::println);

        // Sorting Stream
//        movies.stream()
//                .sorted(Comparator.comparing(Movie::getTitle).reversed())
////                .sorted((m1, m2) -> m2.getTitle().compareTo(m1.getTitle()))
//                .forEach(System.out::println);

//        // Getting Unique Elements
//        movies.stream()
//                .map(Movie::getLikes)
//                .distinct()
//                .forEach(System.out::println);

//        Predicate<Movie> isP = (m) -> m.getLikes() > 10;
//        // Peak
//        movies.stream()
//                .filter(isP)
//                .peek(m -> System.out.println("\nfiltered: " + m.getTitle()))
//                .map(Movie::getTitle)
//                .peek(t -> System.out.println("mapped: " + t))
//                .forEach(System.out::println);


        // Reducers
//        var result = movies.stream()
////                .anyMatch(m -> m.getLikes() > 20)
////        .allMatch(m -> m.getLikes() > 20)
////                .findAny()
////                .findFirst()
//        .max(Comparator.comparing(Movie::getLikes))
//                .get();
//        System.out.println(result);

//        Optional<Integer> sum = movies.stream()
//                .map(m -> m.getLikes())
////                .reduce((a, b) -> a + b);
//                .reduce(Integer::sum);
//        Integer sum = movies.stream()
//                .map(m -> m.getLikes())
////                .reduce((a, b) -> a + b);
//                .reduce(0, Integer::sum);
//        System.out.println(sum);


        // Collectors
//        var result = movies.stream()
//                .filter(m -> m.getLikes() > 10)
//                .collect(Collectors.toList());
//        var result = movies.stream()
//                .filter(m -> m.getLikes() > 10)
////                .collect(Collectors.toMap(Movie::getTitle, Function.identity()));
////        .collect(Collectors.summingDouble(Movie::getLikes));
//        .collect(Collectors.summarizingInt(Movie::getLikes));
//        System.out.println(result);
//        System.out.println(movies.stream().map(Movie::getTitle).collect(Collectors.joining(", ")));

//        // Grouping Elements
//        var result = movies.stream()
////                .collect(groupingBy(Movie::getGenre));
//        .collect(
//                Collectors.groupingBy(
//                        Movie::getGenre,
//                        Collectors.mapping(
//                                Movie::getTitle,
//                                Collectors.joining(" - ")
//                        )
//                )
//        );
//        System.out.println(result);


//        // Partitioning Elements
//        var result = movies.stream()
////                .collect(Collectors.partitioningBy(m -> m.getLikes() > 20));
//        .collect(Collectors.partitioningBy(
//                m -> m.getLikes() > 20,
//                Collectors.mapping(Movie::getTitle, Collectors.joining(" + "))
//                ));
//        System.out.println(result);

        // Primitive Type Streams
        IntStream.rangeClosed(1, 5)
                .forEach(System.out::println);
        IntStream.range(1, 5)
                .forEach(System.out::println);



    }
}
