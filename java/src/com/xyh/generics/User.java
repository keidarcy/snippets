package com.xyh.generics;

public class User implements Comparable<User> {
    private int points;

    public User(int points) {
        this.points = points;
    }

    @Override
    public int compareTo(User other) {
        // this > u => 1
        // this == u => 0
        // this < u => -1
        return points - other.points;
    }

    @Override
    public String toString() {
        return "Points=" + points;
    }
}
