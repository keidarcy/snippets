package com.xyh.basic;

public class Browser {
    public void navigate(String address) {
        String ip = findIpAddress(address);
        String html = sendHttpRequest(ip);
    }

    private String sendHttpRequest(String ip) {
        return "<html></html>";
    }

    private String findIpAddress(String address) {
        return "127.0.0.1";
    }
}
