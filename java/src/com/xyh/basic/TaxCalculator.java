package com.xyh.basic;

public interface TaxCalculator {
    float minimumTax = 100_000;

    static double getTaxableIncome(double income, double expense) {
        return income - expense;
    }
    double calculateTax();
}
