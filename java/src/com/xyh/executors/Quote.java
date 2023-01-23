package com.xyh.executors;

public class Quote {
    private final String site;


    private final int price;

    public Quote(String site, int price) {
        this.site = site;
        this.price = price;
    }

    public String getSite() {
        return site;
    }

    public int getPrice() {
        return price;
    }

    @Override
    public String toString() {
        return site + " " + price;
    }

}
