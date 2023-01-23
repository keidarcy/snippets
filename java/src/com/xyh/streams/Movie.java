package com.xyh.streams;

public class Movie {

    private String title;
    private int likes;


    private Genre genre;

    public Movie(String title, int likes, Genre genre) {
        this.title = title;
        this.likes = likes;
        this.genre = genre;
    }

    public int getLikes() {
        return likes;
    }

    @Override
    public String toString() {
        return "movie name: " + title;
    }

    public String getTitle() {
        return title;
    }

    public Genre getGenre() {
        return genre;
    }
}
