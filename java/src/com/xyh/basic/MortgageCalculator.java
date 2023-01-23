package com.xyh.basic;

public class MortgageCalculator {

    private final static byte MONTH_IN_YEAR = 12;
    private final static byte PERCENT = 100;

    private final int principal;
    private final float annualInterest;
    private final byte years;

    public MortgageCalculator(int principal, float annualInterest, byte years) {
        this.principal = principal;
        this.annualInterest = annualInterest;
        this.years = years;
    }

    public double calculateBalance(short numberOfPaymentsMade) {
        float monthlyInterest = getMonthlyInterest();
        float numberOfPayments = getNumberOfPayments();

        return principal
                * (Math.pow(1 + monthlyInterest, numberOfPayments) - Math.pow(1 + monthlyInterest, numberOfPaymentsMade))
                / (Math.pow(1 + monthlyInterest, numberOfPayments) - 1);
    }

    public double[] getRemainingBalances() {
        var balances = new double[getNumberOfPayments()];
        for (short month = 1; month <= balances.length; month++) {
            balances[month - 1] = calculateBalance(month);
        }
        return  balances;
    }

    public double calculateMortgage() {
        float monthlyInterest = getMonthlyInterest();
        float numberOfPayments = getNumberOfPayments();

        return principal * (monthlyInterest * Math.pow(1 + monthlyInterest, years))
                / (Math.pow(1 + monthlyInterest, numberOfPayments) - 1);
    }

    private int getNumberOfPayments() {
        return years * MONTH_IN_YEAR;
    }

    private float getMonthlyInterest() {
        return annualInterest / PERCENT / MONTH_IN_YEAR;
    }
}
