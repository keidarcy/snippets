package com.xyh.threads;

public class DownloadFileTask implements Runnable{

    private final DownloadStatus status;

    public DownloadFileTask(
            DownloadStatus status
    ) {
//        this.status = new DownloadStatus();
        this.status = status;
    }

    @Override
    public void run() {
        System.out.println("Downloading a file: " + Thread.currentThread().getName());
        for (int i = 0; i < 10_000; i++) {
            if (Thread.currentThread().isInterrupted()) break;
//            System.out.println("Downloading byte " + i);
            status.incrementTotalBytes();
        }
        status.done();
        synchronized (status) {
            status.notifyAll();
        }
//        try {
//            Thread.sleep(5000);
//        } catch (InterruptedException e) {
//            e.printStackTrace();
//        }
        System.out.println("Download complete: " + Thread.currentThread().getName());

    }

    public DownloadStatus getStatus() {
        return status;
    }
}
