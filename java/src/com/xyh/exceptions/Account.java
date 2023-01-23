package com.xyh.exceptions;

import java.io.IOException;

public class Account {
    private float balance;
    public void deposit(float value) throws IOException {
        if (value <= 0) {
            throw new IOException("BAD VALUE");
        }
    }

//    public void withdraw(float value) throws InsufficientFundsException {
//        if (value > balance) {
//            throw new InsufficientFundsException("bad money");
//        }
//    }

    public void withdraw(float value) throws AccountException {
        balance = 0;
        if (value > balance) {
//            var fundsException = new InsufficientFundsException("bad money");
//            var accountException = new AccountException();
//            accountException.initCause(fundsException);
//            throw accountException;
            throw new AccountException(new InsufficientFundsException("bad money"));
        }
    }
}
