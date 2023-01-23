package com.xyh.threads;

import java.util.concurrent.atomic.AtomicInteger;
import java.util.concurrent.atomic.LongAdder;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

public class DownloadStatus {
    private volatile boolean isDone;
//    private int totalBytes;
//    private AtomicInteger totalBytes = new AtomicInteger();
    private LongAdder totalBytes = new LongAdder();

    private int totalFiles;

    private final Object totalBytesLock = new Object();
    private final Object totalFilesLock = new Object();
//    private final Lock lock = new ReentrantLock();

    public int getTotalBytes() {
//        return totalBytes;
//        return totalBytes.get();
        return totalBytes.intValue(); // sum()
    }

    public void incrementTotalBytes() {
//        lock.lock();
//        try {
//            totalBytes++;
//        } finally {
//            lock.unlock();
//        }

//        synchronized (totalBytesLock) {
//            totalBytes++;
//        }
//        totalBytes++;
//        totalBytes.incrementAndGet(); // ++a
//        totalBytes.getAndIncrement(); // a++
    totalBytes.increment();
    }

    public void incrementTotalFiles() {
//        synchronized (this) {
        synchronized (totalFilesLock) {
            totalFiles++;
        }
    }
//    public synchronized void incrementTotalFiles() {
//        totalFiles++;
//    }

    public int getTotalFiles() {
        return totalFiles;
    }

    public boolean isDone() {
        return isDone;
    }

    public void done() {
        isDone = true;
    }
}
