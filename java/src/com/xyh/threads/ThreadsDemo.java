package com.xyh.threads;

import java.util.*;
import java.util.concurrent.ConcurrentHashMap;

public class ThreadsDemo {
    public static void show() {

//        System.out.println(Thread.activeCount());
//        System.out.println(Runtime.getRuntime());
//        System.out.println(Thread.currentThread().getName());

//        Thread thread = new Thread(new DownloadFileTask());
//        thread.start();
//
//        try {
//            Thread.sleep(1000);
//        } catch (InterruptedException e) {
//            e.printStackTrace();
//        }
//        thread.interrupt(); // do not force interrupt thread, just send a interrupt request
////        try {
////            thread.join(); // block method
////        } catch (InterruptedException e) {
////            e.printStackTrace();
////        }
//
//        System.out.println("The file is ready to be scanned.");

//        var status = new DownloadStatus();
//
//        List<Thread> threads = new ArrayList<>();
//        List<DownloadFileTask> tasks = new ArrayList<>();
//        for (int i = 0; i < 10; i++) {
//            var task = new DownloadFileTask(status);
//            var thread = new Thread(task);
//            thread.start();
//            tasks.add(task);
//            threads.add(thread);
//        }
//
//        for (var thread : threads) {
//            try {
//                thread.join();
//            } catch (InterruptedException e) {
//                throw new RuntimeException(e);
//            }
//        }

        // Strategies for Thread Safety
        // 1. Confinement
        // 2. Immutability
        // 3. Synchronization
        // 4. Atomic

        // Confinement
//        var totalBytes = tasks.stream()
//                .map(t -> t.getStatus().getTotalBytes())
//                .reduce(0, Integer::sum);
//        System.out.println(totalBytes);

        // Immutability - lock
//        System.out.println(status.getTotalBytes());

        // Synchronization

        // volatile -> solve visibility, not solve race condition
//        var status = new DownloadStatus();
//        var thread1 = new Thread(new DownloadFileTask(status));
//        var thread2 = new Thread(() -> {
//            while (!status.isDone()) {
//                synchronized (status) {
//                    try {
//                        status.wait();
//                    } catch (InterruptedException e) {
//                        e.printStackTrace();
//                    }
//                }
//            }
//            System.out.println(status.getTotalBytes());
//        });
//        thread1.start();
//        thread2.start();

        // atomic
//        var status = new DownloadStatus();
//
//        List<Thread> threads = new ArrayList<>();
//        List<DownloadFileTask> tasks = new ArrayList<>();
//        for (int i = 0; i < 10; i++) {
//            var task = new DownloadFileTask(status);
//            var thread = new Thread(task);
//            thread.start();
//            tasks.add(task);
//            threads.add(thread);
//        }
//
//        for (var thread : threads) {
//            try {
//                thread.join();
//            } catch (InterruptedException e) {
//                throw new RuntimeException(e);
//            }
//        }
//
//        System.out.println(status.getTotalBytes());
//
//
        // Synchronized Collections
//        Collection<Integer> collection = new ArrayList<>(); // has race condition
//        Collection<Integer> collection = Collections.synchronizedCollection(new ArrayList<>());
//        var thread1 = new Thread(() -> {
//            collection.addAll(Arrays.asList(1, 2, 3));
//        });
//        var thread2 = new Thread(() -> {
//            collection.addAll(Arrays.asList(4, 5, 6));
//        });
//        thread1.start();
//        thread2.start();
//
//        try {
//            thread1.join();
//            thread2.join();
//        } catch (InterruptedException e) {
//            throw new RuntimeException(e);
//        }
//        System.out.println(collection);

        // Concurrent Collections
//        Map<Integer, String> map = new HashMap<>();
        Map<Integer, String> map = new ConcurrentHashMap<>();
        map.put(1, "a");
        map.get(1);
        map.remove(1);

    }
}
