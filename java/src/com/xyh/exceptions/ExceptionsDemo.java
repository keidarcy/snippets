package com.xyh.exceptions;

import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.FileWriter;
import java.io.IOException;
import java.text.ParseException;
import java.text.SimpleDateFormat;

public class ExceptionsDemo {
    public static void show() throws IOException {
//        FileReader reader = null;
//        try {
//            reader = new FileReader("file.txt");
//            var value = reader.read();
//            reader.close();
//            new SimpleDateFormat().parse("");
//        } catch (IOException | ParseException  e) {
//            e.printStackTrace();
//        } finally {
//            if (reader != null) {
//                try {
//                    reader.close();
//                } catch (IOException e) {
//                    throw new RuntimeException(e);
//                }
//            }
//        }

//        try (var reader = new FileReader("file.txt");
//             var writer = new FileWriter("some.txt");
//        ) {
//            var value = reader.read();
//            reader.close();
//            new SimpleDateFormat().parse("");
//        } catch (IOException | ParseException  e) {
//            e.printStackTrace();
//        }

        var account = new Account();
//        try {
//            account.deposit(-1);
//        } catch (IOException e) {
//            System.out.println("Logging");
//            throw e;
//        }
        try {
            account.withdraw(10);
        } catch (AccountException e) {
            var cause= e.getCause();
            System.out.println(cause.getMessage());
            e.printStackTrace();
        }
    }
}
