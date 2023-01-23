package com.xyh.basic;

public class Main {

    public void doSomething(TaxCalculator calculator) {

    }
    public static void main(String[] args) {

        var calculator = new TaxCalculator2022(100_000);
        var report = new TaxReport(calculator);
        report.show();


        report.setCalculator(new TaxCalculator2023());
        report.show();

//        UIControl[] controls = { new TextBox(), new Checkbox() };
//        for (var control : controls) {
//            control.render();
//        }

//        var point1 = new Point(1, 2);
//        var point2= new Point(1, 2);
//        System.out.println(point1.hashCode());
//        System.out.println(point2.hashCode());

//        var control = new UIControl(true);
//        var textBox = new TextBox();
//        show(textBox);

//        var control = new TextBox();
//        control.disable();
//        System.out.println(control.isEnable());

//        int principal = (int) Console.readNumber("Principal: ", 1000, 1_000_000);
//        float annualInterest = (float) Console.readNumber("Annual Interest Rate: ", 1, 30);
//        byte years = (byte) Console.readNumber("Period (years): ", 1, 30);
//
//        var calculator = new MortgageCalculator(principal, annualInterest, years);
//        var report = new MortgageReport(calculator);
//        report.printMortgage();
//        report.printPaymentSchedule();
    }

    public static void show(UIControl control) {
        if (control instanceof TextBox) {
            var textBox = (TextBox)control;
            textBox.setText("FFF");
        }
        System.out.println(control);
    }

}