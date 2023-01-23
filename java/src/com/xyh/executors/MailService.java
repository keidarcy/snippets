package com.xyh.executors;

import java.util.concurrent.CompletableFuture;

public class MailService {
    public void send() {
        LongTask.simulate();
        System.out.println("Mail was sent.");
    }

    public CompletableFuture<Void> sendAsync() {
        return CompletableFuture.runAsync(() -> send());
    }
}
