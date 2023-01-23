package com.xyh.basic;

public class TaxCalculator2022 implements TaxCalculator {
    private double taxableIncome;

    public TaxCalculator2022(double taxableIncome) {
        this.taxableIncome = taxableIncome;
    }

    @Override
    public double calculateTax() {
        return taxableIncome * .4;
    }
}
