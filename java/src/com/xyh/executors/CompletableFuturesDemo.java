package com.xyh.executors;

import java.time.Duration;
import java.time.LocalTime;
import java.util.concurrent.CompletableFuture;
import java.util.stream.Collectors;

public class CompletableFuturesDemo {
    public static void show() {
//        Runnable task = () -> System.out.println("a");
//        var future = CompletableFuture.runAsync(task);

//        Supplier<Integer> task = () -> 1;
//        var future = CompletableFuture.supplyAsync(task);
//        try {
//            var result = future.get();
//            System.out.println(result);
//        } catch (InterruptedException e) {
//            throw new RuntimeException(e);
//        } catch (ExecutionException e) {
//            throw new RuntimeException(e);
//        }

//        var service = new MailService();
//        service.sendAsync();
//        System.out.println("HELLO");
//
//        try {
//            Thread.sleep(5000);
//        } catch (InterruptedException e) {
//            throw new RuntimeException(e);
//        }

//        var future = CompletableFuture.supplyAsync(() -> 1);
////        future.thenRun(() -> {
////            System.out.println(Thread.currentThread().getName());
////            System.out.println("DONE");
////        });
////        future.thenAccept(result -> {
//        future.thenAcceptAsync(result -> {
//            System.out.println(Thread.currentThread().getName());
//            System.out.println(result);
//        });
//
//        try {
//            Thread.sleep(2000);
//        } catch (InterruptedException e) {
//            throw new RuntimeException(e);
//        }
//        var future = CompletableFuture.supplyAsync(() -> {
//            System.out.println("Getting the current weather");
//            throw new IllegalStateException();
//        });
//        try {
//            var temperature = future.exceptionally(ex -> 1).get();
//            System.out.println(temperature);
//        } catch (InterruptedException e) {
//            throw new RuntimeException(e);
//        } catch (ExecutionException e) {
//            throw new RuntimeException(e);
//        }

        // Transforming a Completable Future
//        var future = CompletableFuture.supplyAsync(() -> 20);
////        var result = future
////                .thenApply(CompletableFuturesDemo::toFahrenheit)
////                .get();
////        System.out.println(result);
//         future
//                .thenApply(CompletableFuturesDemo::toFahrenheit)
//                .thenAccept(System.out::println);

//        getUserEmailAsync()
//                .thenCompose(CompletableFuturesDemo::getPlaylistAsync)
//                .thenAccept(System.out::println);

        // Combing Completable Futures
        // 20 USD
        // 0.9
//        var first = CompletableFuture
//                .supplyAsync(() -> "20")
//                .thenApply(str -> {
//                    var price = str.replace("USD", "");
//                    return Integer.parseInt(price);
//                });
//        var second = CompletableFuture.supplyAsync(() -> 0.9);
//        first
//                .thenCombine(second, (price, exchangeRate) -> price * exchangeRate)
//                .thenAccept(System.out::println);

//        // Waiting for Many tasks to Complete
//        var first = CompletableFuture.supplyAsync(() -> 1);
//        var second = CompletableFuture.supplyAsync(() -> 2);
//        var third = CompletableFuture.supplyAsync(() -> 3);
//        var all = CompletableFuture.allOf(first, second, third);
//        all.thenRun(() -> {
//            try {
//                System.out.println(first.get() + second.get() + third.get());
//            } catch (InterruptedException e) {
//                throw new RuntimeException(e);
//            } catch (ExecutionException e) {
//                throw new RuntimeException(e);
//            }
//            System.out.println("ALL task success");
//        });


//        // Waiting for the first Task
//        var first = CompletableFuture.supplyAsync(() -> {
//            LongTask.simulate();
//            return 20;
//        });
//
//        var second = CompletableFuture.supplyAsync(() -> 20);
//
//        CompletableFuture
//                .anyOf(first, second)
//                .thenAccept(System.out::println);

//        // Handling timeouts
//        var future = CompletableFuture.supplyAsync(() -> {
//            LongTask.simulate();
//            return 1;
//        });
//
//        try {
//            var res = future
////                    .orTimeout(1, TimeUnit.SECONDS)
//                    .completeOnTimeout(200, 1, TimeUnit.SECONDS)
//                    .get();
//            System.out.println(res);
//        } catch (InterruptedException e) {
//            throw new RuntimeException(e);
//        } catch (ExecutionException e) {
//            throw new RuntimeException(e);
//        }

        var service = new FlightService();
//        service
//                .getQuote("site1")
//                .thenAccept(System.out::println);
        var futures = service.getQuotes()
                .map(future -> future.thenAccept((System.out::println)))
                .collect(Collectors.toList());
        var start = LocalTime.now();
        CompletableFuture
                .allOf(futures.toArray(new CompletableFuture[0]))
                .thenRun(() -> {
                    var end = LocalTime.now();
                    var duration = Duration.between(start, end);
                    System.out.println("Retrieved all quotes in " + duration.toMillis() + "msec.");
                });

        try {
            Thread.sleep(10_000);
        } catch (InterruptedException e) {
            throw new RuntimeException(e);
        }


    }

    public static CompletableFuture<String> getUserEmailAsync() {
        return CompletableFuture.supplyAsync(() -> "email");
    }

    public static CompletableFuture<String> getPlaylistAsync(String email) {
        return CompletableFuture.supplyAsync(() -> "playlist");
    }

    public static Integer toFahrenheit(int celsius) {
        return (int) (celsius * 1.8) + 32;
    }

}
