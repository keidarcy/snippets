package com.xyh.basic;

import java.text.NumberFormat;

public class MortgageReport {


    private final MortgageCalculator calculator;
    private final NumberFormat currency;

    public MortgageReport(MortgageCalculator calculator) {
        this.calculator = calculator;
        currency = NumberFormat.getCurrencyInstance();
    }

    public void printPaymentSchedule() {
        System.out.println();
        System.out.println("Payment Schedule");
        System.out.println("-----------------");
        for (double balance: calculator.getRemainingBalances()) {
            System.out.println(currency.format(balance));
        }
    }

    public void printMortgage() {
        double mortgage = calculator.calculateMortgage();
        String formattedMortgage = currency.format(mortgage);
        System.out.println("MORTGAGE");
        System.out.println("---------");
        System.out.println("Mortgage: " + formattedMortgage);
    }
}
